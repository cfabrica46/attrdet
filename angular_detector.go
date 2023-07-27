package attrdet

import (
	"strings"

	"golang.org/x/net/html"
)

type AngularDetector struct{}

func (ad *AngularDetector) DetectAttributes(node *html.Node, attrs map[string]int) {
	if node.Type == html.ElementNode {
		for _, attr := range node.Attr {
			if strings.HasPrefix(attr.Key, "ng-") {
				attrs[attr.Key]++
			}
		}
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		ad.DetectAttributes(c, attrs)
	}
}
