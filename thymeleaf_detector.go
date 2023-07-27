package attrdet

import (
	"strings"

	"golang.org/x/net/html"
)

type ThymeleafDetector struct{}

func (td *ThymeleafDetector) DetectAttributes(node *html.Node, attrs map[string]int) {
	if node.Type == html.ElementNode {
		for _, attr := range node.Attr {
			if strings.HasPrefix(attr.Key, "th:") {
				attrs[attr.Key]++
			}
		}
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		td.DetectAttributes(c, attrs)
	}
}
