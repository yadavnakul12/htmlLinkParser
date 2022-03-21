package main

import (
	"fmt"
	"htmlLinkParser/parse"
	"log"
	"strings"
)

func main() {
	htmlLinks, err := parse.GetHTMLLinkBeansFromUrl("https://www.calhoun.io/")
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, htmlBean := range htmlLinks {
		fmt.Printf("{")
		fmt.Printf("\nhref::%v,", htmlBean.Href)
		fmt.Printf("\ntext::%v\n", strings.TrimSpace(htmlBean.Text))
		fmt.Printf("},\n")
	}
}
