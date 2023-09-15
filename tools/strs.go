package tools

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"regexp"
)

func NormalCodeVerify(str string) error {
	matched, err := regexp.MatchString("^[a-zA-Z0-9_]+$", str)
	if err != nil {
		return err
	}
	if !matched {
		return errors.New("invalid str: ^[a-zA-Z0-9_]+$")
	}
	return nil
}

func NormalIDVerify(str string) error {
	return NormalCodeVerify(str)
}

func NormalKeyVerify(str string) error {
	return NormalCodeVerify(str)
}

func Md5(str string) string {
	hash := md5.Sum([]byte(str))
	hashString := hex.EncodeToString(hash[:])
	return hashString
}
