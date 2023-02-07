package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/antchfx/htmlquery"
)

func FetchGoogleDriveDirectURL(fileId string) string {

	// Create HTTP Client
	client := &http.Client{}
	// Create Google Drive Direct Download Link
	url := fmt.Sprintf("https://drive.google.com/uc?id=%s&export=download", fileId)
	// Create HTTP GET Request
	req, _ := http.NewRequest("GET", url, nil)
	// Add Headers
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_6_6 rv:4.0; en-US) AppleWebKit/531.16.5 (KHTML, like Gecko) Version/4.1 Safari/531.16.5")

	// Send Request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close()

	// Read Response Body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return ""
	}

	// Parse HTML
	doc, err := htmlquery.Parse(strings.NewReader(string(body)))
	if err != nil {
		log.Println(err)
		return ""
	}

	list := htmlquery.Find(doc, "//form")

	if list[0].Attr[0].Val == "download-form" {
		return list[0].Attr[1].Val
	}

	return ""
}

func main() {
	id := ""
	ref := FetchGoogleDriveDirectURL(id)
	fmt.Println(ref)
}
