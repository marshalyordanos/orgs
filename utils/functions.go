package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func Contains(arr []string, val string) bool {
	for _, item := range arr {
		if item == val {
			return true
		}
	}
	return false
}

func GetFileType(ext string) string {
	imageExts := []string{".jpg", ".jpeg", ".png", ".gif", ".bmp"}
	videoExts := []string{".mp4", ".avi", ".mkv", ".mov"}
	documentExts := []string{".pdf", ".doc", ".docx", ".xls", ".xlsx"}

	if Contains(imageExts, ext) {
		return "image"
	} else if Contains(videoExts, ext) {
		return "video"
	} else if Contains(documentExts, ext) {
		return "document"
	} else {
		return "other"
	}
}
func GenerateUniqueFilename(fileExt string) string {
	timestamp := time.Now().UnixNano()
	randomString := GenerateRandomString(8) // Specify the length of the random string here
	return fmt.Sprintf("%d-%s%s", timestamp, randomString, fileExt)
}
func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	randomString := make([]byte, length)
	for i := range randomString {
		randomString[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(randomString)
}
