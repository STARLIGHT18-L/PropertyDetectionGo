package tool

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
)

// Decrypt 使用 AES-ECB 模式和 PKCS5 填充进行解密
func Decrypt(password string) (string, error) {
	// 解密密钥，必须与前端加密时使用的密钥相同
	decryptionKey := []byte("Intellectual_Property_Detection_")
	// 对 Base64 编码的加密数据进行解码
	decodedData, err := base64.StdEncoding.DecodeString(password)
	if err != nil {
		return "", err
	}
	// 创建 AES 块
	block, err := aes.NewCipher(decryptionKey)
	if err != nil {
		return "", err
	}
	// 进行解密操作
	decryptedBytes := make([]byte, len(decodedData))
	ecb := NewECBDecrypter(block)
	ecb.CryptBlocks(decryptedBytes, decodedData)
	// 去除 PKCS5 填充
	decryptedBytes = pkcs5Unpad(decryptedBytes)
	// 将解密后的字节数组转换为字符串
	return string(decryptedBytes), nil
}

// NewECBDecrypter 创建 ECB 解密器
func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
	return ecbDecrypter{block: b}
}

type ecbDecrypter struct {
	block cipher.Block
}

func (x ecbDecrypter) BlockSize() int { return x.block.BlockSize() }

func (x ecbDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.block.BlockSize() != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.block.Decrypt(dst, src[:x.block.BlockSize()])
		src = src[x.block.BlockSize():]
		dst = dst[x.block.BlockSize():]
	}
}

// pkcs5Unpad 去除 PKCS5 填充
func pkcs5Unpad(data []byte) []byte {
	padding := int(data[len(data)-1])
	return data[:len(data)-padding]
}

// Sha256Decrypt 比较 SHA256 加密结果
func Sha256Decrypt(words []int, plaintext string) bool {
	// 将 int 数组转换为字节数组
	encryptedBytes := intArrayToByteArray(words)
	// 使用 SHA256 进行加密
	hashedBytes := sha256.Sum256([]byte(plaintext))
	// 比较加密结果
	return bytes.Equal(encryptedBytes, hashedBytes[:])
}

// IntArrayToByteArray 将 int 数组转换为字节数组
func intArrayToByteArray(intArray []int) []byte {
	var buf bytes.Buffer
	for _, num := range intArray {
		err := binary.Write(&buf, binary.BigEndian, int32(num))
		if err != nil {
			panic(err)
		}
	}
	return buf.Bytes()
}

// 辅助函数，用于将 JSON 数组解析为 int 切片
func ParseJSONArrayToIntSlice(jsonStr string) ([]int, error) {
	var words []int
	err := json.Unmarshal([]byte(jsonStr), &words)
	if err != nil {
		return nil, err
	}
	return words, nil
}
