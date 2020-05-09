package link

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

//Link represents link in html anchor tag
type Link struct {
	Href string
	Text string
}

//Parse will take HTML document and
//return slices of anchor links
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)

	if err != nil {
		return nil, fmt.Errorf("Unable to parse html %v", err)
	}

	nodes := linkNodes(doc)
	var links []Link

	for _, node := range nodes {
		links = append(links, buildLink(node))
	}

	return links, nil
}

func buildLink(n *html.Node) Link {
	var ret Link

	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
			break
		}
	}
	//trim space
	t := strings.Join(strings.Fields(string(text(n))), " ")
	ret.Text = t
	return ret
}

func text(n *html.Node) []byte {
	if n.Type == html.TextNode {
		return []byte(n.Data)
	}

	if n.Type != html.ElementNode {
		return []byte("")
	}

	ret := []byte("")
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, text(c)...)
	}
	return ret
}

func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}

	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}
	return ret
}

func dfs(n *html.Node, padding string) {
	msg := n.Data

	if n.Type == html.ElementNode {
		msg = "<" + msg + ">"
	}

	fmt.Println(padding, msg)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, padding+" ")
	}
}
