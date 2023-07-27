# 📦 Attrdet

Welcome to the **Attrdet** 🚀 project! This package is all about empowering you to effortlessly detect special attributes in HTML files. With the help of two charming 🌟 detectives, `AngularDetector` and `ThymeleafDetector`, you'll uncover the secrets of Angular (🅰️) and Thymeleaf (🌿) attributes hidden within your web pages.

## What's Inside?

🕵️‍♂️ **AngularDetector**: 🅰️

Discover the magic of Angular as the `AngularDetector` takes you on a thrilling journey through the DOM, revealing any attributes with the distinct "ng-" prefix. You'll be amazed at how simple and efficient it is to find Angular's hidden treasures!

🕵️‍♀️ **ThymeleafDetector**: 🌿

Embrace the tranquility of Thymeleaf with the `ThymeleafDetector`. Watch in awe as it gracefully glides through the DOM, skillfully detecting any attributes adorned with the "th:" prefix. Let the enchanting beauty of Thymeleaf leave you spellbound!

## How to Use?

It's as easy as 1-2-3! Just import the package and let the detectives do their magic:

```
package main

import (
	"fmt"
	"strings"

	"github.com/your_username/attribute_detector"
)

func main() {
	content := `<html><body><div ng-click="doSomething()" th:if="isVisible">Content</div></body></html>`

	doc, _ := html.Parse(strings.NewReader(content))

	// Angular Detector
	angularAttrs := make(map[string]int)
	angularDetector := &attribute_detector.AngularDetector{}
	angularDetector.DetectAttributes(doc, angularAttrs)
	printAttributes("Angular (ng) attributes:", angularAttrs)

	// Thymeleaf Detector
	thymeleafAttrs := make(map[string]int)
	thymeleafDetector := &attribute_detector.ThymeleafDetector{}
	thymeleafDetector.DetectAttributes(doc, thymeleafAttrs)
	printAttributes("Thymeleaf attributes:", thymeleafAttrs)
}

func printAttributes(title string, attrs map[string]int) {
	fmt.Println(title)
	for attr, count := range attrs {
		fmt.Printf("%s: %d occurrences\n", attr, count)
	}
}
```

## Contributions 🤝

Contributions, bug reports, and suggestions are more than welcome! If you spot any issues or have ideas to make this package even better, please create an "issue" in the repository.

## License 📜

This package is open-source and available under the MIT License. See the LICENSE file for more details.

Start the investigation now! 🕵️‍♂️🕵️‍♀️ Unveil the secrets of Angular (🅰️) and Thymeleaf (🌿) attributes with the Attribute Detector and enjoy the magic 🎉 in your HTML files!
