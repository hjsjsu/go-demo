package utils

import (
	"demo/entity"
	"encoding/json"
	"github.com/tidwall/gjson"
)

func InitSqlConnect() {

}

func RToJson(data entity.R) []byte {
	dataMap := map[string]any{}
	bytesData, _ := json.Marshal(data.GetData())
	dataMap["code"] = data.GetCode()
	dataMap["data"] = gjson.Parse(string(bytesData)).Value()
	dataMap["message"] = data.GetMessage()
	bytes, _ := json.Marshal(dataMap)
	return bytes
}

func JsonToMap(str string) map[string]any {
	var tempMap map[string]interface{}
	err := json.Unmarshal([]byte(str), &tempMap)
	if err != nil {
		panic(err)
	}
	return tempMap
}
