package attrdet

import "golang.org/x/net/html"

type AttributeDetector interface {
	DetectAttributes(*html.Node, map[string]int)
}
