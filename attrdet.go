package attrdet

import "golang.org/x/net/html"

type AttributeDetector interface {
	DetectAttributes(*html.Node, map[string]int)
	DetectVariables(*html.Node, map[string]int)
	DetectScripts(*html.Node, int) int
}
