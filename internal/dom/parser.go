package dom

import (
	"strings"
	"fmt"

	"golang.org/x/net/html"
)

func parseDocument(htmlDocument string) {

	doc, err := html.Parse(strings.NewReader(htmlDocument))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", doc)
	fmt.Println(doc)

}
