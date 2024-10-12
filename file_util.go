package file

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"path"
)

// IsExist 返回文件是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// Md5Sum 生成path指定文件的Md5
func Md5Sum(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

// Create 创建文件
func Create(name string) (*os.File, error) {
	// 如果文件夹不存在则创建文件夹后在创建文件
	dirPath := path.Dir(name)
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err = os.MkdirAll(dirPath, 0755)
		if err != nil {
			return nil, err
		}
	}
	// 返回的文件句柄不使用时需要关闭
	return os.Create(name)
}

// Remove 移除文件
func Remove(name string) error {
	if IsExist(name) {
		return os.Remove(name)
	}
	return nil
}
