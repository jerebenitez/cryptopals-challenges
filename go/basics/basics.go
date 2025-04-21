package basics

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
)

func Hex2Base64(str string) (string, error) {
	hex, err := hex.DecodeString(str)

	if err != nil {
		return "", errors.New(fmt.Sprintf("hex decode string: %v", err.Error()))
	}

	return base64.RawStdEncoding.EncodeToString(hex), nil
}
