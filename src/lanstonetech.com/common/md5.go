package common

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func MakeMD5(arg ...string) string {
	info := strings.Join(arg, "-")
	m := md5.New()
	m.Write([]byte(info))
	return hex.EncodeToString(m.Sum(nil))
}

func MakeByteMd5(arg []byte) string {
	m := md5.New()
	m.Write(arg)
	return hex.EncodeToString(m.Sum(nil))
}
