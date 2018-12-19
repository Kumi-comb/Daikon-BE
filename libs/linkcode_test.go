package libs

import (
	"log"
	"testing"
)

func TestGetUniqueCode(t *testing.T) {
	if GetUniqueCode("2f3d") != "1a2f3c4d5e6f7a8b1a2f3c4d5e6f7a8b1a2f3c4d5e6f7a8b" {
		log.Panic("リンクコードとユニークコードの不一致")
	}
	if GetUniqueCode("") != "" {
		log.Panic("エラー処理の不具合")
	}
}
