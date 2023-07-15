package common

import (
	"golang.org/x/crypto/bcrypt"
)

// NewPassword 通过bcrypt算法来实现密码加密，自动生成solt，具体实现类似于php的password_hash
func NewPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// CheckPassword 验证hash与password是否匹配
func CheckPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
