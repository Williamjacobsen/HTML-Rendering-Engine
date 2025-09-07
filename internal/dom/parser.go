package dom

import (
	"strings"

	"golang.org/x/net/html"
)

func ParseDocument(htmlDocument []byte) (*html.Node, error) {
	reader := strings.NewReader(string(htmlDocument))
	doc, err := html.Parse(reader)
	if err != nil {
		return nil, err
	}
	return doc, nil
}
