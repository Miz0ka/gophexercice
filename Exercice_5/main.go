package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	htmlparser "gophexercice/Exercice_4/lib"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type SiteMap struct {
	Url      string
	Sitemaps []SiteMap
}

func main() {
	fmt.Println("Exercice 5")

	var urlBase string

	flag.StringVar(&urlBase, "url", "https://www.calhoun.io/", "Url to the site to parse")
	flag.Parse()

	//var sitemap SiteMap
	linkvisited := make([]string, 0)
	domaine := domainExtract(urlBase)

	linkvisited = append(linkvisited, "")

	sitemap := exploreSite(urlBase, &linkvisited, domaine)

	file, err := xml.MarshalIndent(sitemap, "", " ")
	file = []byte(xml.Header + string(file))
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	_ = ioutil.WriteFile("sitemap.xml", file, 0644)
}

func exploreSite(current string, linkvisited *[]string, domaine string) SiteMap {
	sitemap := &SiteMap{
		Url:      current,
		Sitemaps: make([]SiteMap, 0),
	}
	*linkvisited = append(*linkvisited, current)
	res, err := http.Get(current)
	if err != nil {
		fmt.Println("Can't access to :", current)
	}
	if res.StatusCode == 200 {
		newurl := htmlparser.SearchALink(res.Body)
		for _, val := range newurl {
			nexturl := checkUrl(val.Href, domaine)
			if !contains(*linkvisited, nexturl) {
				undersite := exploreSite(nexturl, linkvisited, domaine)
				sitemap.Sitemaps = append(sitemap.Sitemaps, undersite)
				fmt.Println(val.Href)
			}
		}
		//Check URL(Domain + already visited)
	} else {
		log.Println(current, " failed to acces : ", res.StatusCode)
	}
	return *sitemap
}

func domainExtract(url string) string {
	re := regexp.MustCompile(`(https?:\/\/.*\/)`)
	return re.FindStringSubmatch(url)[0]
}

func checkUrl(url string, domaine string) string {
	if strings.Contains(url, domaine) {
		return url
	}
	re := regexp.MustCompile(`$\/`)
	if re.Match([]byte(url)) {
		return domaine + url
	}
	return ""
}

func contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

//TODO correctly manage domain
