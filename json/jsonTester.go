package json

import (
	"encoding/json"
	"fmt"
)

type MyJson struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestEncodJson() {
	myJsonStr := MyJson{"mliu", 23}
	b, err := json.Marshal(&myJsonStr)
	if err != nil {
		fmt.Printf("encoding json error")
		return
	}
	var myJsObj MyJson
	json.Unmarshal(b, &myJsObj)
	fmt.Printf("反序列化json:%v", myJsObj.Name)
	fmt.Printf("序列化json:%v", string(b))
}
