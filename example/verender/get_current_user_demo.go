package main

import (
	"encoding/json"
	"fmt"
)

func GetCurrentUserDemo() {
	v := getVerenderInstance()
	resp, err := v.GetCurrentUser()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	data, _ := json.Marshal(resp)
	fmt.Println(string(data))
}
