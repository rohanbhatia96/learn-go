package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://randomuser.me/api"

func main() {
	fmt.Println("Making a get request")

	response, err := http.Get(url)
	checkNilError(err)
	defer response.Body.Close()
	fmt.Printf("Response is of type %T\n", response)

	data, err := io.ReadAll(response.Body)

	checkNilError(err)

	content := string(data)

	fmt.Println(content)
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
