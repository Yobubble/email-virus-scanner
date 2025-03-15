package utils

import (
	"encoding/base64"
	"os"
)

func GetFile(filepath string) []byte {
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	return bytes
}

func Base64Encode(file []byte) string {
	return base64.StdEncoding.EncodeToString(file)
}
