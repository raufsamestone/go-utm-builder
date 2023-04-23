package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/common-nighthawk/go-figure"
)

func main() {
	// Generate the ASCII art banner
	myFigure := figure.NewFigure("Go-UTM Builder", "", true)
	myFigure.Print()

	// Prompt the user for the URL
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the URL: ")
	urlStr, _ := reader.ReadString('\n')
	urlStr = strings.TrimSpace(urlStr)

	// Prompt the user for the source parameter
	fmt.Print("Enter the utm_source parameter: ")
	source, _ := reader.ReadString('\n')
	source = strings.TrimSpace(source)

	// Prompt the user for the medium parameter
	fmt.Print("Enter the utm_medium parameter: ")
	medium, _ := reader.ReadString('\n')
	medium = strings.TrimSpace(medium)

	// Prompt the user for the campaign parameter
	fmt.Print("Enter the utm_campaign parameter: ")
	campaign, _ := reader.ReadString('\n')
	campaign = strings.TrimSpace(campaign)

	// Generate the UTM parameters
	values := url.Values{}
	values.Set("utm_source", source)
	values.Set("utm_medium", medium)
	values.Set("utm_campaign", campaign)

	// Add the UTM parameters to the URL
	u, err := url.Parse(urlStr)
	if err != nil {
		fmt.Printf("Error parsing URL: %v\n", err)
		return
	}

	u.RawQuery = values.Encode()

	fmt.Printf("Done! 🎊 ")

	// Print the UTM link
	fmt.Printf("Your UTM link 🔗: https://www.%v\n", u.String())
}
