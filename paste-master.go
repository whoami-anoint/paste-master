package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	// Set your API key and URL
	apiKey := "your-api-key"
	url := "https://pastebin.com/api/api_post.php"

	// Set up the parameters for the POST request
	data := url.Values{}
	data.Set("api_dev_key", apiKey)
	data.Set("api_option", "paste")
	data.Set("api_paste_code", "This is the text that will be saved to the file.")
	data.Set("api_paste_name", "example-file")

	// Create an HTTP client and send the POST request to the pastebin API
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(data.Encode()))
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
