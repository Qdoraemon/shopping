package utils

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/nfnt/resize"
)

type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
}

func Success(Data interface{}, Message string) Response {
	return Response{
		Data:    Data,
		Message: Message,
		Code:    0,
	}

}

func Error(Code int, Message string) Response {
	return Response{
		// Data:    Data,
		Message: Message,
		Code:    Code,
	}

}

func GenFileNameByUUID(prefix string) string {
	return prefix + "_" + uuid.New().String()
}

func GetFileNameByTime() string {
	return time.Now().Format("20060102150405")
}

// resize image
func ResizeImage(file *multipart.FileHeader, dir string, targetHeight int) (string, error) {
	// 检查文件类型，这里简单判断文件后缀是否为图片
	if file.Header.Get("Content-Type") != "image/jpeg" && file.Header.Get("Content-Type") != "image/png" {
		return "", fmt.Errorf("文件类型错误")
	}

	// // 检查文件大小，这里简单判断文件大小不超过 5MB
	if file.Size > 5*1024*1024 {
		return "", fmt.Errorf("文件大小超过限制")
	}

	// 打开上传的文件
	srcFile, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("無法打開上傳文件")
	}
	defer srcFile.Close()

	// 通过 image.Decode 解码文件
	img, _, err := image.Decode(srcFile)
	if err != nil {
		// http.Error(w, "Unable to decode image", http.StatusInternalServerError)
		return "", fmt.Errorf("image 解碼失敗")
	}

	// 原始图像的宽高
	originalWidth := img.Bounds().Dx()
	originalHeight := img.Bounds().Dy()

	// 指定的目标高度
	// targetHeight := 800

	// 计算缩放后的宽度，保持原始宽高比
	scaleFactor := float64(targetHeight) / float64(originalHeight)
	targetWidth := int(float64(originalWidth) * scaleFactor)

	// 使用 resize.Resize 来进行缩放
	resizedImg := resize.Resize(uint(targetWidth), uint(targetHeight), img, resize.Lanczos3)

	// 创建一个 800x800 的黑色背景图
	finalImg := image.NewRGBA(image.Rect(0, 0, targetHeight, targetHeight))

	// 填充背景为黑色
	black := color.RGBA{0, 0, 0, 255}
	for y := 0; y < targetHeight; y++ {
		for x := 0; x < targetHeight; x++ {
			finalImg.Set(x, y, black)
		}
	}

	// 将缩放后的图片粘贴到黑色背景上，居中显示
	offsetX := (targetHeight - targetWidth) / 2
	offsetY := (targetHeight - targetHeight) / 2

	// 将缩放后的图像绘制到新图像的合适位置
	for y := 0; y < targetHeight; y++ {
		for x := 0; x < targetWidth; x++ {
			finalImg.Set(x+offsetX, y+offsetY, resizedImg.At(x, y))
		}
	}

	// 生成文件名
	fileName := GenFileNameByUUID("Image") + ".jpg"

	// 生成存儲路徑
	dest := filepath.Join(dir, fileName)

	// 打开一个文件来保存处理后的图像
	outFile, err := os.Create(dest)
	if err != nil {
		return "", fmt.Errorf("無法創建文件")
	}
	defer outFile.Close()

	// 将处理后的图像保存为 JPEG 格式
	err = jpeg.Encode(outFile, finalImg, nil)
	if err != nil {

		return "", fmt.Errorf("無法保存文件")
	}
	return fileName, nil

}
