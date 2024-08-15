package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/joho/godotenv"
)

func init(){
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error Loading .env file")
	}
}

func main(){
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
	fmt.Printf("Port is %s", os.Getenv("PORT"))

}