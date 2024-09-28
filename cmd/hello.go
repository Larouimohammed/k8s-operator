package cmd

import ("fmt")

// "github.com/kubernetes/client-go
func Hello( s string) error{
	
	_,err:=fmt.Printf("Hello: %s \n",s)
	return  err

}