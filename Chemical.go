package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/TwiN/go-color"
)

func main() {
	// Parse command-line flags
	updateFlag := flag.Bool("u", false, "Update list of chemicals")
	wordFlag := flag.String("w", "nil", "Chemical for Norway government list")
	fileFlag := flag.String("f", "", "File with chemicals to search for")
	flag.Parse()

	fmt.Println("     --- LovData Narkotika Søk 0.15 ---")

	// If the update flag is set, update the list of chemicals
	if *updateFlag {
		fmt.Println(color.Ize(color.Green, "[ + ] Update list of chemicals"))
		updateChemicalList()
		os.Exit(0)
	}

	// Check if the list of chemicals exists
	if _, err := os.Stat("tables.txt"); err == nil {
		if *wordFlag != "nil" {
			searchTable(*wordFlag)
		} else if *fileFlag != "" {
			// Open the file specified by the -f flag
			file, err := os.Open(*fileFlag)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			defer file.Close()

			// Read the contents of the file
			contents, err := ioutil.ReadAll(file)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			// Split the contents of the file into lines
			lines := strings.Split(string(contents), "\n")

			// Search for each line in the "tables.txt" file
			for _, line := range lines {
				searchTable(line)
			}
		} else {
			fmt.Println("[ ! ] You have NOT put a Chemicals in -w")
			os.Exit(1)
		}
	} else {
		fmt.Println(color.Ize(color.Red, "[ ! ] No list of chemicals"))
		os.Exit(0)
	}
}

func updateChemicalList() {
	// Set the URL for the website
	url := "https://lovdata.no/dokument/SF/forskrift/2013-02-14-199"

	// Download the website by making an HTTP GET request to the URL
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()

	// Read the response body and convert it to a string
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	bodyStr := string(body)

	// Parse the HTML from the response body
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(bodyStr))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create a new file to store the tables
	file, err := os.Create("tables.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Find all the tables in the HTML
	doc.Find("table").Each(func(i int, table *goquery.Selection) {
		// Write the table to the file
		tableHtml, _ := table.Html()
		file.WriteString(tableHtml + "\n")
	})
}

func searchTable(word string) {
	// Convert the input word to two variations: one with the first letter in uppercase and the rest in lowercase, and the other with the first letter in lowercase and the rest in lowercase.
	if word == "" {
		os.Exit(0)
	}
	lenString := len(word)
	upperVariation := strings.ToUpper(word[0:1]) + word[1:lenString]
	lowerVariation := strings.ToLower(word[0:1]) + word[1:lenString]

	// Read the contents of the "tables.txt" file.
	html, err := ioutil.ReadFile("tables.txt")
	if err != nil {
		// If an error occurred while reading the file, print the error and return.
		fmt.Println(err)
		return
	}

	// Check if either variation of the input word is contained in the file contents.
	if strings.Contains(string(html), upperVariation) || strings.Contains(string(html), lowerVariation) {
		// If either variation is found, print a message indicating that the chemical is banned.
		fmt.Println(color.Ize(color.Red, "[ ! ] "+word+": This chemical does have a ban on it"))
	} else {
		// If neither variation is found, print a message indicating that the chemical is not banned.
		fmt.Println(color.Ize(color.Green, "[ + ] "+word+": This chemical does NOT have a ban on it"))
	}
}
