package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	html "golang.org/x/net/html"
)

type Links struct {
	Href string
	Text string
}

func main() {
	var htmlPath string
	fmt.Println("Exercice 4")

	flag.StringVar(&htmlPath, "html", ".\\ex4.html", "Link to the html file")
	flag.Parse()
	f, e := os.Open(htmlPath)
	if e != nil {
		fmt.Println("Can't open the html file")
		panic(e)
	}
	res := SearchALink(f)
	fmt.Printf("%+v", res)
}

func SearchALink(r io.Reader) []Links {
	links := make([]Links, 0)
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Println("Error in the HTML Page")
	}
	links = explorenode(doc, links)
	return links
}

func explorenode(n *html.Node, links []Links) []Links {
	if n.Type == html.ElementNode && n.Data == "a" {
		var link Links
		link.Href = gethref(n)
		link.Text = strings.TrimSpace(getdata(n.FirstChild))
		links = append(links, link)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = explorenode(c, links)
	}
	return links
}

func gethref(n *html.Node) string {
	for _, val := range n.Attr {
		if strings.ToLower(val.Key) == "href" {
			return val.Val
		}
	}
	return ""
}

func getdata(n *html.Node) string {
	var str string = ""
	for s := n.NextSibling; s != nil; s = s.NextSibling {
		str += getdata(s)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		str += getdata(c)
	}
	fmt.Printf("%+v\n", n)
	if n.Type != html.TextNode {
		return str
	}
	return strings.TrimSpace(n.Data) + " " + str
}
