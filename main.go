package main

import (
	"fmt"
	"log"

	"github.com/otiai10/gosseract/v2"
)

func main() {
	client := gosseract.NewClient()
	defer client.Close()
	err := client.SetImage("./test.png")
	if err != nil {
		log.Fatal(err)
	}
	text, err := client.Text()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(text)

}
