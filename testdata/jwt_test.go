package testdata

import (
	"fmt"
	"gvb_server/config"
	"gvb_server/core"
	"gvb_server/utils"
	"testing"
)

func TestGenerateTokenUsingRS256(t *testing.T) {
	core.InitConfig()

	tokenStr, err := utils.GenerateTokenUsingRS256(1, "htt")

	fmt.Println(tokenStr, err)
}
func TestParseTokenRs256(t *testing.T) {
	core.InitConfig()

	tokenStr, err := utils.GenerateTokenUsingRS256(1, "htt")

	fmt.Println(tokenStr, err)

	claimsRs256, err := utils.ParseTokenRs256(tokenStr)
	fmt.Println(claimsRs256, err)

}

func TestParsePubKeyBytes(t *testing.T) {
	core.InitConfig()

	pub, err := utils.ParsePubKeyBytes([]byte(config.PUB_KEY))
	if err != nil {
		fmt.Println(pub)
		fmt.Println(err)
	}
	fmt.Println(pub)
	fmt.Println(err)

}
