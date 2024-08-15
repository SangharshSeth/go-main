package main

import (
	"fmt"
	"runtime"
)

func main(){
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
	fmt.Println("Hello World!")
	fmt.Printf("Docker Rocks!!!,\t Coming for you Kubernetes!")
}