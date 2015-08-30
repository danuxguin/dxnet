package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	//"regexp"
	"text/template"
)

// func IsMobileNumber(str string) bool {
// 	reg := regexp.MustCompile(`^\+86(1[3-578])\d{9}$`)
// 	return reg.MatchString(str)
// }

func ConverterToJson(t interface{}) (string, error) {
	x, err := json.Marshal(t)
	return string(x), err
}

func ConvertJsonToMap(j string) (r map[string]interface{}, err error) {
	err = json.Unmarshal([]byte(j), &r)
	return
}

func ParseTemplate(text string, data interface{}) (string, error) {

	t := template.New("")
	t, err := t.Parse(text)
	if err != nil {
		return "", err
	}

	buff := bytes.NewBufferString("")
	if err := t.Execute(buff, data); err != nil {
		return "", err
	}

	return buff.String(), nil
}

func StringCombine(str1, str2 string) string {
	if str1 > str2 {
		return fmt.Sprintf("%s-%s", str2, str1)
	}

	return fmt.Sprintf("%s-%s", str1, str2)
}

func StringCutRune(str string, n int) string {

	b := []byte(str)
	i := 0
	index := bytes.IndexFunc(b, func(r rune) bool {
		i++

		if i > n {
			return true
		}

		return false
	})

	if index < 0 {
		return str
	}
	return string(b[:index])
}

func GetStringRuneNum(str string) int {

	b := []byte(str)
	return len(bytes.Runes(b))
}
