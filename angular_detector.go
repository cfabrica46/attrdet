package attrdet

import (
	"golang.org/x/net/html"
)

type AngularDetector struct {
	BaseDetector
}

func (ad *AngularDetector) DetectAttributes(node *html.Node, attrs map[string]int) {
	ad.BaseDetector.DetectAttributes(node, attrs, "ng-")
}

func (ad *AngularDetector) DetectVariables(node *html.Node, variables map[string]int) {
	ad.BaseDetector.DetectVariables(node, variables, "{{", "}}")
}
