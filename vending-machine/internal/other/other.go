package other

import "fmt"

func Catch() {
	rec := recover()
	if rec != nil {
		fmt.Println("Error :", rec)
	}
}
