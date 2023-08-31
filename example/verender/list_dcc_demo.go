package main

import (
	"fmt"
)

func ListDccDemo() {
	v := getVerenderInstance()
	dccs, err := v.ListDccs()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(dccs)
}
