package errors

import "fmt"

func errorHandle(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
