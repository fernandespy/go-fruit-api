package crawler

import (
	"fmt"
	"go-fruit-api/services"
)

func RunCrawler() {
	fmt.Println("Executing automatic crawler...")
	err := services.LoadFruitsFromAPI()
	if err != nil {
		fmt.Println("Erro uploading fruits using crawler", err)
	} else {
		fmt.Println("Crawler executed successfuly")
	}
}
