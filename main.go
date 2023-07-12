package main

import (
	"fmt"
	"log"
	"os"
	"p2w/service"
)

func main() {
	cache1 := service.ToPDF("https://www.baidu.com")
	if err := os.WriteFile("sample.pdf", cache1, 0o644); err != nil {
		log.Fatal(err)
	}
	fmt.Println("wrote sample.pdf")
	cache2 := service.ToFullImage("https://www.baidu.com", 90)
	if err := os.WriteFile("sample.png", cache2, 0o644); err != nil {
		log.Fatal(err)
	}
	fmt.Println("wrote sample.png")
}
