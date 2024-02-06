package passwrod

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"go-gin-example/internal/model/sys_user_model"
	"io"
)

func generateRandomBytes(length int) ([]byte, error) {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}
	return randomBytes, nil
}

func GenerateRandomSalt() (string, error) {
	saltLength := 32 // 你可以根据需要选择适当的 salt 长度
	randomBytes, err := generateRandomBytes(saltLength)
	if err != nil {
		return "", err
	}
	salt := base64.StdEncoding.EncodeToString(randomBytes)
	return salt, nil
}

func GeneratePassword(password string, user *sys_user_model.SysUser) (string, error) {
	passwordWithSalt := password + user.Salt
	fmt.Printf("passwordWith: %s\n", passwordWithSalt)
	hash := md5.New()
	_, err := io.WriteString(hash, passwordWithSalt)
	if err != nil {
		return "", err
	}
	hashedPassword := fmt.Sprintf("%x", hash.Sum(nil))
	return hashedPassword, nil
}
