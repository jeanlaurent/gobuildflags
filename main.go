package main


import (
	"fmt"
	"log"
)

func main() {
	version, err := GetWindowsVersion()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Windows system root is %s\n", version)
}

