package main

import (
	"client-api/entity"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetProductFromAPI() entity.Result {
	request, err := http.NewRequest("POST", "http://localhost:8080/api/v1/products", nil)
	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	responseInByte, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	var result entity.Result
	err = json.Unmarshal(responseInByte, &result)
	if err != nil {
		fmt.Println(err)
	}

	return result
}

func main() {
	message := GetProductFromAPI().Message
	result := GetProductFromAPI().Result

	fmt.Println("message:", message)
	if len(result) != 0 {
		fmt.Println("products: ")
		for i := 0; i < len(result); i++ {
			fmt.Println(result[i])
		}
		return
	}
}
