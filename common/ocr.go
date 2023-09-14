package common

import (
	"fmt"
	"github.com/otiai10/gosseract/v2"
	"log"
)

func Ocr(path string) string {
	client := gosseract.NewClient()
	defer func(client *gosseract.Client) {
		if err := client.Close(); err != nil {
			log.Println("ocr client close failed!")
		}
	}(client)

	if err := client.SetImage(path); err != nil {
		log.Println("image ocr failed!")
		return ""
	}
	text, _ := client.Text()
	fmt.Println(text)
	return text
}
