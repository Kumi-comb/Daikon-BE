package libs

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

// CreateLinkCode
func CreateLinkCode(UniqueCode string) string {
	return generateLinkCode(UniqueCode, time.Now().Format("20060102150405"))
}

func generateLinkCode(UniqueCode string, Timestamp string) string {

	b := sha256.Sum256([]byte(UniqueCode + Timestamp))
	s := hex.EncodeToString(b[:])[:4]

	return s
}
