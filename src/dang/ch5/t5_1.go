// 5_1
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)

	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}

}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	return get_links(n.FirstChild, links)
	//	for c := n.FirstChild; c != nil; c = c.NextSibling {
	//		links = visit(links, c)
	//	}
	//	return links
}

func get_links(n *html.Node, links []string) []string {
	if n == nil {
		return links
	}
	links = visit(links, n)
	return get_links(n.NextSibling, links)
}
