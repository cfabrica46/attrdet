package attrdet

import (
	"golang.org/x/net/html"
)

type ThymeleafDetector struct {
	BaseDetector
}

func (td *ThymeleafDetector) DetectAttributes(node *html.Node, attrs map[string]int) {
	td.BaseDetector.DetectAttributes(node, attrs, "th:")
}

func (td *ThymeleafDetector) DetectVariables(node *html.Node, variables map[string]int) {
	td.BaseDetector.DetectVariables(node, variables, "${", "}")
}
