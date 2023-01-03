package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	// Prompt the user for their pastebin API key and the name of the file they want to create
	var apiKey, fileName string
	fmt.Print("Enter your pastebin API key: ")
	fmt.Scanln(&apiKey)
	fmt.Print("Enter the name of the file you want to create: ")
	fmt.Scanln(&fileName)

	// Prompt the user for the text they want to save to the file
	var text string
	fmt.Print("Enter the text you want to save to the file: ")
	fmt.Scanln(&text)

	// Set up the parameters for the POST request
	data := url.Values{}
	data.Set("api_dev_key", apiKey)
	data.Set("api_option", "paste")
	data.Set("api_paste_code", text)
	data.Set("api_paste_name", fileName)

	// Create an HTTP client and send the POST request to the pastebin API
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://pastebin.com/api/api_post.php", bytes.NewBufferString(data.Encode()))
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// Print a message based on the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	if resp.StatusCode == 200 {
		location := resp.Header.Get("Location")
		fmt.Println("File created successfully at:", location)
	} else {
		fmt.Println("An error occurred: ", string(body))
	}
}
