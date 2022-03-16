package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"os"
	"strings"
)

type myHTMLBean struct {
	href, text string
}

type myHTMLBeans []myHTMLBean

var htmlBeans myHTMLBeans

func displayHTMLNode(node *html.Node) {
	for node := node.FirstChild; node != nil; node = node.NextSibling {
		////fmt.Println(node)
		if node.Type == html.ElementNode && node.Data == "a" {
			var tempHTMLBean myHTMLBean
			for _, attr := range node.Attr {
				if attr.Key == "href" {
					tempHTMLBean.href = attr.Val
					if node.FirstChild.Type == html.TextNode {
						tempHTMLBean.text = node.FirstChild.Data
					} else if node.LastChild.Type == html.TextNode {
						tempHTMLBean.text = node.LastChild.Data
					}
				}
			}
			htmlBeans = append(htmlBeans, tempHTMLBean)
			//fmt.Printf("Type:%v\n", node.Type)
			//fmt.Printf("Attribute:%v\n", node.Attr)
			//fmt.Printf("Data:%v\n", node.Data)
			//fmt.Printf("Child Att:%v\n", node.FirstChild.Attr)
			//fmt.Printf("Child Data:%v\n", node.FirstChild.Data)
		}
		displayHTMLNode(node)

	}
}

func main() {
	htmlReader, readError := os.Open("sampleFile4.html")
	if readError != nil {
		log.Fatal(readError.Error())
	}
	htmlDoc, parseError := html.Parse(htmlReader)
	if parseError != nil {
		log.Fatal(parseError.Error())
	}
	displayHTMLNode(htmlDoc)
	for _, htmlBean := range htmlBeans {
		fmt.Printf("{")
		fmt.Printf("\nhref::%v,", htmlBean.href)
		fmt.Printf("\ntext::%v\n", strings.TrimSpace(htmlBean.text))
		fmt.Printf("},\n")
	}
}
