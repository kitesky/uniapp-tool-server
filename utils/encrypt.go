package utils

import (
	"crypto/hmac"
	"crypto/md5"
	random "crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"os"

	"golang.org/x/crypto/bcrypt"
)

// bcrypt 加密密码
func GenerateFromPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// bcrypt 校验密码
func CompareHashAndPassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

// MD5加密
func MD5(data []byte) string {
	new := md5.New()
	new.Write(data)
	return hex.EncodeToString(new.Sum(nil))
}

func MD5File(path string) (string, error) {
	// 打开文件
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}

	defer file.Close()

	// 创建 MD5 Hash 对象
	hash := md5.New()

	// 从文件中读取内容并写入到 Hash 对象中
	_, err = io.Copy(hash, file)
	if err != nil {
		return "", err
	}

	// 获取最终的哈希值
	hashBytes := hash.Sum(nil)

	return fmt.Sprintf("%x", hashBytes), nil
}

// HmacSha256加密
func HmacSha256(data []byte, secretKey []byte) string {
	new := hmac.New(sha256.New, secretKey)
	new.Write(data)
	return hex.EncodeToString(new.Sum(nil))
}

func Sha256(data []byte) string {
	new := sha256.New()
	new.Write(data)
	return hex.EncodeToString(new.Sum(nil))
}

// 加密函数
func RSA_Encrypt(text string, publicKey []byte) ([]byte, error) {
	// 解析公钥PEM数据
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("公钥解析失败")
	}

	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	pub := pubInterface.(*rsa.PublicKey)

	// 对原文进行加密
	encryptedBytes, err := rsa.EncryptPKCS1v15(random.Reader, pub, []byte(text))
	if err != nil {
		return nil, err
	}

	return encryptedBytes, nil
}

// RSA解密
func RSA_Decrypt(cipherText []byte, privateKey []byte) (string, error) {
	//pem解码
	block, _ := pem.Decode(privateKey)
	//X509解码
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	//对密文进行解密
	dncryptedBytes, error := rsa.DecryptPKCS1v15(random.Reader, key, cipherText)
	//返回明文
	return string(dncryptedBytes), error
}
