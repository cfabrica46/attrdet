package attrdet

import (
	"strings"

	"golang.org/x/net/html"
)

type BaseDetector struct{}

func (bd *BaseDetector) DetectAttributes(node *html.Node, attrs map[string]int, prefix string) {
	if node.Type == html.ElementNode {
		for _, attr := range node.Attr {
			if strings.HasPrefix(attr.Key, prefix) {
				attrs[attr.Key]++
			}
		}
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		bd.DetectAttributes(c, attrs, prefix)
	}
}

func (bd *BaseDetector) DetectVariables(node *html.Node, variables map[string]int, startTag, endTag string) {
	if node.Type == html.TextNode {
		text := node.Data
		startIdx := strings.Index(text, startTag)
		endIdx := strings.Index(text, endTag)

		if startIdx != -1 && endIdx != -1 && startIdx < endIdx {
			variable := strings.TrimSpace(text[startIdx+len(startTag) : endIdx])
			variables[variable]++
		}
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		bd.DetectVariables(c, variables, startTag, endTag)
	}
}
