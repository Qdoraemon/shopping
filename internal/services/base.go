package services

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

func SaveFile(file *multipart.FileHeader) (string, error) {
	// 確保 uploads 資料夾存在
	saveDir := "uploads"
	if _, err := os.Stat(saveDir); os.IsNotExist(err) {
		os.Mkdir(saveDir, os.ModePerm)
	}

	// 避免檔名衝突，使用時間戳
	filename := fmt.Sprintf("%d_%s", time.Now().Unix(), filepath.Base(file.Filename))
	savePath := filepath.Join(saveDir, filename)

	// 保存文件
	err := saveUploadedFile(file, savePath)
	if err != nil {
		return "", err
	}

	return savePath, nil
}

// 封裝 SaveUploadedFile 以避免循環引用 gin.Context
func saveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = out.ReadFrom(src)
	return err
}
