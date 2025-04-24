package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword 对密码进行哈希处理
func HashPass(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPasswordHash 验证密码是否匹配
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	fmt.Println("CheckPasswordHash err", err)
	return err == nil
}
