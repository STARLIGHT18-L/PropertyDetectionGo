package tool

import (
	"PropertyDetection/config"
	"fmt"
	"github.com/disintegration/imaging"
	"image"
	"image/color"
	"io"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// DownloadImage 下载网络图片到本地
func DownloadImage(host, url string, id int) *image.NRGBA {
	if config.Boot.Config.App.Env != "dev" {
		url = strings.Replace(url, host, "minio", 1)
	}
	resp, err := http.Get(url) // 发送 HTTP GET 请求
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK { // 检查响应状态码
		fmt.Println("HTTP request failed:", resp.Status)
		return nil
	}
	suffix := filepath.Ext(url)
	fileName := fmt.Sprintf("./image/%d%s", id, suffix)
	dir := filepath.Dir(fileName)                   // 获取文件所在目录
	if _, err := os.Stat(dir); os.IsNotExist(err) { // 检查目录是否存在，如果不存在则创建
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return nil
		}
	}
	file, err := os.Create(fileName) // 创建本地文件
	if err != nil {
		return nil
	}
	defer file.Close()
	_, err = io.Copy(file, resp.Body) // 将响应体中的数据复制到本地文件
	if err != nil {
		return nil
	}
	img, _ := imaging.Open(fileName)
	grayImg := imaging.Grayscale(img)                       //转换为灰度图像
	return imaging.Resize(grayImg, 64, 64, imaging.Lanczos) // 缩放图片
}

// ConvertToGray 将 *image.NRGBA 转换为 *image.Gray
func ConvertToGray(img *image.NRGBA) *image.Gray {
	bounds := img.Bounds()
	gray := image.NewGray(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			gray.SetGray(x, y, color.GrayModel.Convert(img.At(x, y)).(color.Gray))
		}
	}
	return gray
}

// ExtractFeatureVector 提取图片的特征向量
func ExtractFeatureVector(img *image.Gray) []float64 {
	bounds := img.Bounds()
	rows := bounds.Dy()
	cols := bounds.Dx()
	featureVector := make([]float64, 0)
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x += 2 {
			if x+1 < cols {
				featureVector = append(featureVector, float64(img.GrayAt(x, y).Y))
				featureVector = append(featureVector, float64(img.GrayAt(x+1, y).Y))
			} else {
				// 处理最后一个像素（如果列数为奇数）
				featureVector = append(featureVector, float64(img.GrayAt(x, y).Y))
				featureVector = append(featureVector, 0) // 用 0 填充
			}
		}
	}
	return featureVector
}

// SelectFeatures 根据方差选择特征
func SelectFeatures(featureVector []float64, numFeatures int) []float64 {
	numSamples := len(featureVector)
	// 计算每个特征的均值
	mean := 0.0
	for _, val := range featureVector {
		mean += val
	}
	mean /= float64(numSamples)
	// 计算每个特征的方差
	variances := make([]float64, numSamples)
	for i, val := range featureVector {
		diff := val - mean
		variances[i] = diff * diff
	}
	// 选择方差最大的前 numFeatures 个特征
	selectedIndices := make([]int, 0, numFeatures)
	for k := 0; k < numFeatures; k++ {
		maxVar := -math.MaxFloat64
		maxIndex := -1
		for i := 0; i < numSamples; i++ {
			if variances[i] > maxVar {
				maxVar = variances[i]
				maxIndex = i
			}
		}
		// 检查 maxIndex 是否为 -1
		if maxIndex == -1 {
			break // 没有可选择的特征了，提前结束循环
		}
		selectedIndices = append(selectedIndices, maxIndex)
		variances[maxIndex] = -math.MaxFloat64 // 标记为已选择
	}
	// 提取选择的特征
	selectedFeatures := make([]float64, len(selectedIndices))
	for j, index := range selectedIndices {
		selectedFeatures[j] = featureVector[index]
	}
	return selectedFeatures
}

// 假设这里有一个更全面的停用词列表
var stopwords = map[string]bool{
	"的": true,
	"是": true,
	"在": true,
	"和": true,
	"与": true,
	"或": true,
}

// 简单的分词函数，按单个字符分词
func simpleTokenize(text string) []string {
	var tokens []string
	for _, r := range text {
		tokens = append(tokens, string(r))
	}
	return tokens
}

// 文本预处理
func preprocess(text string) string {
	// 分词
	tokens := simpleTokenize(text)
	var filteredTerms []string
	// 遍历分词结果
	for _, word := range tokens {
		word = strings.ToLower(word)
		// 检查是否为停用词
		if !stopwords[word] {
			filteredTerms = append(filteredTerms, word)
		}
	}
	// 将过滤后的词语用空格连接成字符串
	return strings.Join(filteredTerms, " ")
}

// 计算词频
func getTermFrequency(text string) map[string]int {
	termFrequency := make(map[string]int)
	terms := strings.Fields(text)
	for _, term := range terms {
		termFrequency[term]++
	}
	return termFrequency
}

// 计算 TF-IDF
func getTFIDF(tf map[string]int, docCount int, df map[string]int) map[string]float64 {
	tfidf := make(map[string]float64)
	for term, termFreq := range tf {
		docFreq := df[term]
		idf := math.Log(float64(docCount) / float64(docFreq+1))
		tfidf[term] = float64(termFreq) * idf
	}
	return tfidf
}

// 计算点积
func dotProduct(vec1, vec2 map[string]float64) float64 {
	dotProduct := 0.0
	for term, value1 := range vec1 {
		if value2, ok := vec2[term]; ok {
			dotProduct += value1 * value2
		}
	}
	return dotProduct
}

// 计算向量的模
func magnitude(vector map[string]float64) float64 {
	magnitude := 0.0
	for _, value := range vector {
		magnitude += math.Pow(value, 2)
	}
	return math.Sqrt(magnitude)
}

// 计算余弦相似度
func CosineSimilarity(text1, text2 string, documents []string) float64 {
	// 文本预处理
	text1 = preprocess(text1)
	text2 = preprocess(text2)
	// 计算词频
	tf1 := getTermFrequency(text1)
	tf2 := getTermFrequency(text2)
	// 计算文档频率
	df := make(map[string]int)
	for _, doc := range documents {
		doc = preprocess(doc)
		uniqueTerms := make(map[string]bool)
		terms := strings.Fields(doc)
		for _, term := range terms {
			uniqueTerms[term] = true
		}
		for term := range uniqueTerms {
			df[term]++
		}
	}
	// 计算 TF-IDF
	docCount := len(documents)
	tfidf1 := getTFIDF(tf1, docCount, df)
	tfidf2 := getTFIDF(tf2, docCount, df)
	// 计算点积和模
	dotProduct := dotProduct(tfidf1, tfidf2)
	magnitude1 := magnitude(tfidf1)
	magnitude2 := magnitude(tfidf2)
	if magnitude1 == 0 || magnitude2 == 0 {
		return 0
	}
	return dotProduct / (magnitude1 * magnitude2)
}
