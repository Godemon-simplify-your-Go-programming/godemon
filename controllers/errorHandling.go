package controllers

import "fmt"

func ErrorHandle(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
