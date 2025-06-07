package api

import (
	"PropertyDetection/config"
	"PropertyDetection/model/entity"
	"PropertyDetection/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
)

type SystemApi struct {
}

func (api SystemApi) Upload(c *gin.Context) {
	file, err := c.FormFile("file") // 从请求中获取上传的文件
	if err != nil {
		c.JSON(400, tool.Fail())
		return
	}
	src, err := file.Open() // 打开上传的文件
	if err != nil {
		c.JSON(400, tool.Fail())
		return
	}
	defer src.Close()
	objectSize := file.Size               // 获取文件大小
	_, err = config.Boot.Minio.PutObject( // 调用 MinIO 上传文件的方法
		config.Boot.Config.Minio.Bucket,
		file.Filename,
		src,
		objectSize,
		minio.PutObjectOptions{},
	)
	if err != nil {
		c.JSON(400, tool.Fail())
		return
	}
	url := "http://" + config.Boot.Config.Minio.Host + "/" + config.Boot.Config.Minio.Bucket + "/" + file.Filename
	if strings.Contains(url, "minio") {
		url = strings.Replace(url, "minio", c.Request.Host, 1)
	}
	c.JSON(200, tool.SuccessData(
		map[string]interface{}{
			"url": url,
		},
	))
}
func (api SystemApi) Gc(c *gin.Context) {
	runtime.GC()
	c.JSON(200, tool.Success())
}
func (api SystemApi) CleanFile(c *gin.Context) {
	var wg sync.WaitGroup
	wg.Add(2)
	ossCleanCountCh := make(chan int) // 创建通道用于接收 OSS 和本地文件清理数量
	localCleanCountCh := make(chan int)
	go func() { // 清理 OSS 文件
		defer wg.Done()
		defer close(ossCleanCountCh)
		index := 0
		imgs := map[string]bool{}
		for info := range config.Boot.Minio.ListObjects(config.Boot.Config.Minio.Bucket, "", true, nil) {
			imgs[info.Key] = true
		}
		var images []entity.Image
		var copyrights []entity.Copyright
		var patents []entity.Patent
		var trademarks []entity.Trademark
		var drs []entity.DetectionRecords
		config.Boot.Mysql.Select("id, url").Where("deleted = ?", 1).Find(&images)
		config.Boot.Mysql.Select("id, url").Where("deleted = ?", 1).Find(&copyrights)
		config.Boot.Mysql.Select("id, url").Where("deleted = ?", 1).Find(&patents)
		config.Boot.Mysql.Select("id, url").Where("deleted = ?", 1).Find(&trademarks)
		config.Boot.Mysql.Select("id, url").Where("deleted = ?", 1).Find(&drs)
		list := []string{}
		for _, image := range images {
			lastSlashIndex := strings.LastIndex(image.Url, "/")
			if lastSlashIndex != -1 {
				imageName := image.Url[lastSlashIndex+1:]
				list = append(list, imageName)
			}
		}
		for _, copyright := range copyrights {
			lastSlashIndex := strings.LastIndex(copyright.Url, "/")
			if lastSlashIndex != -1 {
				copyrightName := copyright.Url[lastSlashIndex+1:]
				list = append(list, copyrightName)
			}
		}
		for _, patent := range patents {
			lastSlashIndex := strings.LastIndex(patent.Url, "/")
			if lastSlashIndex != -1 {
				patentName := patent.Url[lastSlashIndex+1:]
				list = append(list, patentName)
			}
		}
		for _, trademark := range trademarks {
			lastSlashIndex := strings.LastIndex(trademark.Url, "/")
			if lastSlashIndex != -1 {
				trademarkName := trademark.Url[lastSlashIndex+1:]
				list = append(list, trademarkName)
			}
		}
		for _, dr := range drs {
			lastSlashIndex := strings.LastIndex(dr.Url, "/")
			if lastSlashIndex != -1 {
				drName := dr.Url[lastSlashIndex+1:]
				list = append(list, drName)
			}
		}
		listCh := make(chan string)
		go func() {
			defer close(listCh)
			for _, s := range list {
				if imgs[s] {
					listCh <- s
					index++
				}
			}
		}()
		config.Boot.Minio.RemoveObjects(config.Boot.Config.Minio.Bucket, listCh)
		fmt.Println("All files in the OSS have been deleted.")
		ossCleanCountCh <- index // 将 OSS 清理数量发送到通道
	}()
	go func() { // 清理本地文件
		defer wg.Done()
		defer close(localCleanCountCh)
		folderPath := "./image"
		count := 0
		err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				err := os.Remove(path)
				if err != nil {
					return fmt.Errorf("failed to delete file %s: %w", path, err)
				}
				count++
			}
			return nil
		})
		if err != nil {
			fmt.Printf("Error walking the path %s: %v\n", folderPath, err)
			return
		}
		fmt.Println("All files in the folder have been deleted.")
		localCleanCountCh <- count // 将本地清理数量发送到通道
	}()
	ossCleanCount := <-ossCleanCountCh // 从通道接收 OSS 和本地文件清理数量
	localCleanCount := <-localCleanCountCh
	wg.Wait() // 等待两个协程完成
	c.JSON(200, tool.SuccessMsgData(fmt.Sprintf("清理OSS文件%v个，本地文件%v个", ossCleanCount, localCleanCount), nil))
}
