package server

import (
	"encoding/base64"

	"github.com/tidwall/gjson"
)

func ParseUserInfo(str string) (name string, id int, err error) {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return
	}
	// fmt.Println(string(data))
	id = int(gjson.GetBytes(data, "id").Int())
	name = gjson.GetBytes(data, "username").String()
	return
}
