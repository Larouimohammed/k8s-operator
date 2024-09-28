package cmd

import "fmt"
func Hello( s string) error{
	
	_,err:=fmt.Printf("Hello: %s \n",s)
	return  err

}					

