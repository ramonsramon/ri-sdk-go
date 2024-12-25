package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/hatch-ed-com/ri-sdk-go/pkg/rapididentity"
)

func main() {
	baseUrl, err := url.Parse(os.Getenv("RI_URL"))
	if err != nil {
		log.Fatal(err)
	}
	options := rapididentity.Options{
		HTTPClient:      &http.Client{},
		BaseUrl:         baseUrl,
		ServiceIdentity: os.Getenv("RI_KEY"),
	}

	client := rapididentity.New(options)

	input := rapididentity.GetConnectFilesInput{
		Path:    "/",
		Project: "sec_mgr",
	}

	output, err := client.GetConnectFiles(input)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", output)
}
