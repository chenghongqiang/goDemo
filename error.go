package main

import (
	"fmt"
)

type DivideError struct {
	dividee int 
	divider int
}

func (de *DivideError) Error() string {

	strFormat := `
		dividee:%d
		divider:0
	`

	return fmt.Sprintf(strFormat, de.dividee)
}

func Divided(dividee int, divider int)(result int, errorMsg string){

	if divider ==0 {
		dData := DivideError{
			dividee: dividee,
			divider: divider,
		}
		errorMsg = dData.Error()
		return 

	} else {
		return dividee / divider, ""
	}

}

func main() {
	if result, errorMsg := Divided(10,2);errorMsg == "" {
		fmt.Println("10/2=", result)
	}

	if _, errorMsg := Divided(10,0);errorMsg != "" {
		fmt.Println("errorMsg is", errorMsg)
	}

}