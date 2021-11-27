package util

import (
	"crypto/sha1"
	"fmt"
	"os"
	"path"
)

func hashRequestParams(url string, body []byte) string {
	hash := sha1.New()
	hash.Write([]byte(url))

	if body != nil {
		hash.Write(body)
	}

	return fmt.Sprintf("%x", hash.Sum(nil))
}

func getFromCache(url string, body []byte) []byte {
	_ = os.MkdirAll(".cache", 0755)

	hash := hashRequestParams(url, body)
	fileName := path.Join(".cache", hash)

	content, err := os.ReadFile(fileName)
	if err != nil {
		return nil
	}

	return content
}

func saveToCache(url string, body []byte, content []byte) error {
	_ = os.MkdirAll(".cache", 0755)

	hash := hashRequestParams(url, body)
	fileName := path.Join(".cache", hash)

	err := os.WriteFile(fileName, content, 0644)
	if err != nil {
		return err
	}

	return nil
}

