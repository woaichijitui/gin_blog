package utils

import (
	"fmt"
	"testing"
)

func TestHashPwd(t *testing.T) {
	pwd := "12312314"

	pwdHash, err := PasswordHash(pwd)

	fmt.Println(pwdHash, err)
}
func TestCheckPwd(t *testing.T) {
	pwd := "12312314"

	pwdHash, err := PasswordHash(pwd)
	err = PasswordVerify(pwd, pwdHash)

	fmt.Println(pwdHash, err)
}
