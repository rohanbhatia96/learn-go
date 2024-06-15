package main

import (
	"fmt"
	"net/url"
)

const exampleUrl = "https://therohanbhatia.com/my-projects?param1=test&param2=here"

func main() {
	fmt.Println("Hello")
	fmt.Println(exampleUrl)

	result, err := url.Parse(exampleUrl)

	checkNilError(err)

	fmt.Println(result.Scheme)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Host)

	qparams := result.Query()

	fmt.Printf("Params: %T\n", qparams)

	fmt.Printf("Param1: %v\n", qparams["param1"])

	for key, val := range qparams {
		fmt.Printf("%v is %v\n", key, val)
	}

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
