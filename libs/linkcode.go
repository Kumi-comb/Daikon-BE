package libs

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"time"
)

// CreateLinkCode returns onetime password of uniquecode(LinkCode).
func CreateLinkCode(UniqueCode string) string {
	return generateLinkCode(UniqueCode, time.Now().Format("20060102150405"))
}

// GetUniqueCode returns uniquecode with LinkCode.
func GetUniqueCode(LinkCode string) string {
	uniqueCode, err := searchUniqueCode(LinkCode)
	if err != nil {
		log.Printf("\"%s\" is invaild linkcode.", LinkCode)
		return ""
	}

	return uniqueCode
}

func generateLinkCode(UniqueCode string, Timestamp string) string {

	b := sha256.Sum256([]byte(UniqueCode + Timestamp))
	s := hex.EncodeToString(b[:])[:4]

	return s
}

func searchUniqueCode(LinkCode string) (string, error) {
	if LinkCode == "2f3d" {
		return "1a2f3c4d5e6f7a8b1a2f3c4d5e6f7a8b1a2f3c4d5e6f7a8b", nil
	}

	return "", nil
}
