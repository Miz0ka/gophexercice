package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	htmlparser "gophexercice/Exercice_4/lib"
	"io/ioutil"
	"log"
	"net/http"
)

type SiteMap struct {
	Url      string
	Sitemaps []SiteMap
}

func main() {
	fmt.Println("Exercice 5")

	var urlBase, url string

	flag.StringVar(&urlBase, "url", "https://www.calhoun.io/", "Url to the site to parse")
	flag.Parse()

	//var sitemap SiteMap
	linktoexplore := make([]string, 0)
	linkvisited := make([]string, 0)

	linktoexplore = append(linktoexplore, urlBase)

	for len(linktoexplore) != 0 {
		url, linktoexplore = linktoexplore[len(linktoexplore)-1], linktoexplore[:len(linktoexplore)-1]
		linkvisited = append(linkvisited, url)
		res, err := http.Get(url)
		if err != nil {
			fmt.Println("Can't access to :", url)
		}
		if res.StatusCode == 200 {
			newurl := htmlparser.SearchALink(res.Body)
			for _, val := range newurl {
				fmt.Println(val.Href)
			}
			//Check URL(Domain + already visited)
		} else {
			log.Println(url, " failed to acces : ", res.StatusCode)
		}
	}

	sitemap := &SiteMap{
		Url: urlBase,
	}

	file, err := xml.MarshalIndent(sitemap, "", " ")
	file = []byte(xml.Header + string(file))
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	_ = ioutil.WriteFile("sitemap.xml", file, 0644)
}
