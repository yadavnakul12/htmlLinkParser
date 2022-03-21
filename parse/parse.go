package parse

import (
	"golang.org/x/net/html"
	"log"
	"net/http"
	"strings"
)

type MyHTMLLinkBean struct {
	Href, Text string
}

type MyHTMLLinkBeans []MyHTMLLinkBean

var HtmlLinkBeans MyHTMLLinkBeans

func isDuplicate(htmlLinkBean MyHTMLLinkBean) bool {
	duplicate := false
	for _, value := range HtmlLinkBeans {
		if value.Href == htmlLinkBean.Href {
			duplicate = true
		}
	}
	return duplicate
}

func displayHTMLNode(node *html.Node) {
	for node := node.FirstChild; node != nil; node = node.NextSibling {
		////fmt.Println(node)
		if node.Type == html.ElementNode && node.Data == "a" {
			var tempHTMLBean MyHTMLLinkBean
			for _, attr := range node.Attr {
				if attr.Key == "href" {
					tempHTMLBean.Href = attr.Val
					if node.FirstChild.Type == html.TextNode {
						tempHTMLBean.Text = node.FirstChild.Data
					} else if node.LastChild.Type == html.TextNode {
						tempHTMLBean.Text = node.LastChild.Data
					}
				}
			}
			if strings.TrimSpace(tempHTMLBean.Href) != "" && !isDuplicate(tempHTMLBean) {
				HtmlLinkBeans = append(HtmlLinkBeans, tempHTMLBean)
			}
			//fmt.Printf("Type:%v\n", node.Type)
			//fmt.Printf("Attribute:%v\n", node.Attr)
			//fmt.Printf("Data:%v\n", node.Data)
			//fmt.Printf("Child Att:%v\n", node.FirstChild.Attr)
			//fmt.Printf("Child Data:%v\n", node.FirstChild.Data)
		}
		displayHTMLNode(node)

	}
}

func GetHTMLLinkBeansFromUrl(url string) (MyHTMLLinkBeans, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	htmlDoc, parseError := html.Parse(resp.Body)
	if parseError != nil {
		log.Fatal(parseError.Error())
	}
	displayHTMLNode(htmlDoc)
	return HtmlLinkBeans, nil
}

//
//func main() {
//	htmlReader, readError := os.Open("sampleFile4.html")
//	if readError != nil {
//		log.Fatal(readError.Error())
//	}
//	htmlDoc, parseError := html.Parse(htmlReader)
//	if parseError != nil {
//		log.Fatal(parseError.Error())
//	}
//	displayHTMLNode(htmlDoc)
//	for _, htmlBean := range HtmlLinkBeans {
//		fmt.Printf("{")
//		fmt.Printf("\nHref::%v,", htmlBean.Href)
//		fmt.Printf("\nText::%v\n", strings.TrimSpace(htmlBean.Text))
//		fmt.Printf("},\n")
//	}
//}
