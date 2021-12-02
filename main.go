package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/flrnprz/color"
)

const api = "https://larder.io/api/1/@me/"
const searchQ = "search/?q="

// Body data request
type Body []byte

func main() {
	t := Token{}
	c := Config{}

	if len(os.Args[1:]) < 1 {
		clr := color.New(color.Bold)
		clr.Println("Usage: ")
		fmt.Println("$ larder-cli <search>")
		os.Exit(1)
	}

	c.fileName = "config.json"
	c.directory = join(usrHomeDir(), ".config/lader-cli/")
	c.path = join(c.directory, c.fileName)

	query := os.Args[1]

	apiSearch := join(api, searchQ)
	url := join(apiSearch, query)

	handle, err := ioutil.ReadFile(c.path)

	if os.IsNotExist(err) {
		err := firstTimeRun(c)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Config file created.")
		os.Exit(1)
	}

	if os.IsPermission(err) {
		log.Fatal(err)
	}

	stringToken, err := t.Parse(handle)
	if err != nil {
		log.Fatal("Parse Error: ", err)
	}

	token := fmt.Sprintf("Token %s", stringToken)
	body, err := getBody(request(token, url))
	if err != nil {
		log.Fatal("getRequest: ", err)
	}

	searchData, err := body.search(query)
	if err != nil {
		log.Fatal("Search ", err)
	}
	searchData.parseResult()
}

// efficient way to concatenate strings
func join(strs ...string) string {
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(str)
	}
	return sb.String()
}

func (s SearchResult) parseResult() {
	green := color.New(color.FgGreen).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	white := color.New(color.FgWhite).SprintFunc()
	yellow := color.New(color.FgHiYellow).SprintFunc()

	for n, res := range s.Results {
		fmt.Printf("%s. %s\n", cyan(n+1), green(res.Title))
		fmt.Printf("    %s %s\n", red(">"), cyan(res.URL))
		if len(res.Description) > 1 {
			fmt.Printf("    %s %s\n", red("+"), white(res.Description))
		}
		if len(res.Tags) > 1 {
			fmt.Printf("    %s ", red("tags: "))
			for _, tag := range res.Tags {
				fmt.Printf("#%s ", yellow(tag.Name))
			}
			fmt.Println()
		}
		fmt.Println()
	}
	fmt.Printf("Total found: %s", green(s.Count))
}

func request(token string, url string) *http.Request {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", token)
	return req
}

func getBody(req *http.Request) (Body, error) {
	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func (b Body) search(query string) (SearchResult, error) {
	var s SearchResult
	return s, json.Unmarshal(b, &s)
}
