package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

//AES加密,CBC
func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

//AES解密
func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS7UnPadding(origData)
	return origData, nil
}

func AesGCMEncrypt(data []byte, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

func AesGCMDecrypt(ciphertext []byte, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, content := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, content, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}

type Aes struct {
	securityKey []byte
	iv          []byte
}

/**
 * constructor
 */
func AesTool(securityKey []byte, iv string) *Aes {
	return &Aes{securityKey, []byte(iv)}
}

/**
 * 加密
 * @param string $plainText 明文
 * @return bool|string
 */
func (a Aes) Encrypt(plainText string) (string, error) {
	block, err := aes.NewCipher(a.securityKey)
	if err != nil {
		return "", err
	}
	plainTextByte := []byte(plainText)
	blockSize := block.BlockSize()
	plainTextByte = addPKCS7Padding(plainTextByte, blockSize)
	cipherText := make([]byte, len(plainTextByte))
	mode := cipher.NewCBCEncrypter(block, a.iv)
	mode.CryptBlocks(cipherText, plainTextByte)
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

/**
 * 解密
 * @param string $cipherText 密文
 * @return bool|string
 */
func (a Aes) Decrypt(cipherText string) (string, error) {
	block, err := aes.NewCipher(a.securityKey)
	if err != nil {
		return "", err
	}
	cipherDecodeText, decodeErr := base64.StdEncoding.DecodeString(cipherText)
	if decodeErr != nil {
		return "", decodeErr
	}
	mode := cipher.NewCBCDecrypter(block, a.iv)
	originCipherText := make([]byte, len(cipherDecodeText))
	mode.CryptBlocks(originCipherText, cipherDecodeText)
	originCipherText = stripPKSC7Padding(originCipherText)
	return string(originCipherText), nil
}

/**
 * 填充算法
 * @param string $source
 * @return string
 */
func addPKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	paddingText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, paddingText...)
}

/**
 * 移去填充算法
 * @param string $source
 * @return string
 */
func stripPKSC7Padding(cipherText []byte) []byte {
	length := len(cipherText)
	unpadding := int(cipherText[length-1])
	return cipherText[:(length - unpadding)]
}
