package utils

import (
	"bytes"
	"time"
)

const TIME_FORMAT = "2006-01-02 15:04:05"

func TimeNowStr() string {
	return time.Now().Format(TIME_FORMAT)
}

func UnionStr(strList ...string) string {
	var buff bytes.Buffer

	for _, str := range strList {
		buff.WriteString(str)
	}

	return buff.String()
}