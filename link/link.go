package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"os"
	"strings"
)

type Link struct {
	Href string
	Text string
}

func findLinks(n *html.Node) []Link {
	l := make([]Link, 0)

	if n.Type == html.ElementNode && n.Data == "a" {
		var href string

		for _, a := range n.Attr {
			if a.Key == "href" {
				href = a.Val
				break
			}
		}
		l = append(l, Link{href, findText(n)})
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		l = append(l, findLinks(c)...)
	}
	return l
}

func findText(n *html.Node) string {
	s := make([]string, 0)

	if n.Type == html.TextNode {
		s = append(s, n.Data)
	}

	var childText string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		childText = findText(c)

		if childText != "" {
			s = append(s, childText)
		}
	}
	return strings.TrimSpace(strings.Trim(strings.Join(s, " "), "\n"))
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No input file specified")
	}

	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal("Unable to read input file")
	}
	defer f.Close()

	doc, err := html.Parse(f)
	if err != nil {
		log.Fatal("Unable to parse HTML file")
	}

	links := findLinks(doc)
	fmt.Println(links)
}
