package testdata

import (
	"fmt"
	"gvb_server/utils"
	"testing"
)

func TestHashPwd(t *testing.T) {
	pwd := "12312314"

	pwdHash, err := utils.PasswordHash(pwd)

	fmt.Println(pwdHash, err)
}
func TestCheckPwd(t *testing.T) {
	pwd := "12312314"

	pwdHash, _ := utils.PasswordHash(pwd)
	ok := utils.PasswordVerify(pwd, pwdHash)

	fmt.Println(pwdHash, ok)
}
