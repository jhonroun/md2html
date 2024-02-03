// This initial created by inspiration of https://github.com/sindresorhus/github-markdown-css/
// and https://github.com/gilliek/ghmd/tree/master
// Then license as "THE BEER-WARE LICENSE" (Revision 42):
// <kevin.gillieron@gw-computing.net> wrote this file. As long as you retain
// this notice you can do whatever you want with this stuff. If we meet some
// day, and you think this stuff is worth it, you can buy me a beer in return
// Kevin Gillieron
// guys muy him beer and maybe little cup for me :)
// LemTech

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

const (
	doctype    = "<!DOCTYPE html>"
	headFormat = "<head><meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\">%s</head>"
)

type markdown struct {
	Text    string `json:"text"`
	Mode    string `json:"mode"`
	Context string `json:"context"`
}

func fatal(a ...interface{}) {
	fmt.Fprintln(os.Stderr, a...)
	os.Exit(1)
}

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func findH(doc *goquery.Document, pattern string) map[string]string {
	var temp map[string]string = make(map[string]string)
	doc.Find(pattern).Each(func(i int, s *goquery.Selection) {
		element := s.Text()
		_, exists := temp[element]
		if !exists {
			rndId := RandStringBytes(16)
			temp[element] = rndId
			s.SetAttr("id", rndId)
		}
	})
	return temp
}

func render(path string, out *os.File) {
	md, err := os.ReadFile(path)
	if err != nil {
		fatal(err)
	}

	title := strings.Split(filepath.Base(path), ".")[0]

	buf, err := json.Marshal(markdown{
		Text: string(md),
		Mode: "gfm",
	})
	if err != nil {
		fatal(err)
	}

	resp, err := http.Post("https://api.github.com/markdown",
		"application/json", bytes.NewBuffer(buf))
	if err != nil {
		fatal(err)
		return
	}
	defer resp.Body.Close()

	err = out.Truncate(0)
	if err != nil {
		fatal(err)
	}

	_, err = out.Seek(0, 0)
	if err != nil {
		fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fatal(err)
	}

	var linksHone map[string]string = findH(doc, "h1")
	var linksHtwo map[string]string = findH(doc, "h2")
	var linksHthree map[string]string = findH(doc, "h3")
	var linksHfour map[string]string = findH(doc, "h4")
	var linksHfive map[string]string = findH(doc, "h5")
	var linksHsix map[string]string = findH(doc, "h6")

	doc.Find("ul li").Each(func(i int, s *goquery.Selection) {
		element := s.Find("a")
		links := element.Text()
		_, exists := element.Attr("href")
		if exists && len(links) > 1 {
			newHref, exist := linksHone[links]
			if exist {
				element.SetAttr("href", "#"+newHref)
			}
			newHref, exist = linksHtwo[links]
			if exist {
				element.SetAttr("href", "#"+newHref)
			}
			newHref, exist = linksHthree[links]
			if exist {
				element.SetAttr("href", "#"+newHref)
			}
			newHref, exist = linksHfour[links]
			if exist {
				element.SetAttr("href", "#"+newHref)
			}
			newHref, exist = linksHfive[links]
			if exist {
				element.SetAttr("href", "#"+newHref)
			}
			newHref, exist = linksHsix[links]
			if exist {
				element.SetAttr("href", "#"+newHref)
			}
		}
	})

	doc.Find("p a").Each(func(i int, s *goquery.Selection) {
		hasLink, ok := s.Attr("href")
		decodedValue, err := url.QueryUnescape(hasLink)
		if err != nil {
			fatal(err)
			return
		}
		if ok {
			if s.Text() == "^" {
				decodedValue = strings.ReplaceAll(decodedValue, "#", "")
				s.SetAttr("href", "#"+linksHone[decodedValue])
				s.SetText("к " + decodedValue)
			}
		}
	})

	Body, err := doc.Html()
	if err != nil {
		fatal(err)
	}

	head := fmt.Sprintf(headFormat, "<title>"+title+"</title><link id=\"favicon\" rel=\"shortcut icon\" type=\"image/png\" href=\""+favicon+"\">")

	fmt.Fprintln(out, doctype, head, `<style>`, CSS, `</style><body><div class="readme"><article class="markdown-body">`)
	fmt.Fprintln(out, Body, "</article></div></body>")
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Simplest usage run md2html filename.md")
		os.Exit(1)
	}

	path := os.Args[1]
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fatal("input markdown file does not exist")
	}

	outputFile := ""
	if strings.Contains(path, ".md") {
		outputFile = strings.Split(filepath.Base(path), ".")[0] + ".html"
	} else {
		outputFile = filepath.Base(path) + ".html"
	}

	var out *os.File
	out, err := os.Create(outputFile)
	if err != nil {
		fatal(err)
	}
	defer out.Close()

	render(path, out)

	fmt.Println("Output file is: " + outputFile)

}
