package util

import (
	"bytes"
	"errors"
	"io"
	"regexp"
	"strconv"

	"golang.org/x/net/html"
)

type ParamFindNode struct {
	NodeName   string
	Attributes map[string]string
}

func verifyParam(node *html.Node, param ParamFindNode) bool {
	if param.NodeName != node.Data {
		return false
	}

	nodeAttr := make(map[string]string)
	for _, attr := range node.Attr {
		nodeAttr[attr.Key] = attr.Val
	}

	for key, value := range param.Attributes {
		if nodeAttr[key] != value {
			return false
		}
	}

	return true
}

func FindNode(doc *html.Node, param ParamFindNode) ([]*html.Node, error) {
	if param.NodeName == "" && param.Attributes == nil {
		return nil, errors.New("ERROR: Node uninformed")
	}

	nodes := []*html.Node{}
	var crawler func(*html.Node)
	crawler = func(node *html.Node) {
		if node.Type == html.ElementNode && verifyParam(node, param) {
			nodes = append(nodes, node)
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			crawler(child)
		}
	}

	crawler(doc)

	if len(nodes) != 0 {
		return nodes, nil
	}

	return nil, errors.New("ERROR: Missing the node tree")
}

func RenderNode(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, n)
	return buf.String()
}

func GetContent(node *html.Node) string {
	content := RenderNode(node)
	re := regexp.MustCompile(">(.+)<")
	return re.FindStringSubmatch(content)[1]
}

func convertCurrencyBrStringToFloat(str string) (float64, error) {
	re := regexp.MustCompile(`^(\d{0,3})\.?(\d{0,3})\.?(\d{0,3})\.?(\d{1,3}),(\d{2})$`)
	currencyValue := re.ReplaceAllString(str, "$1$2$3$4.$5")

	totalValue, err := strconv.ParseFloat(currencyValue, 64)
	if err != nil {
		return 0, err
	}

	return totalValue, nil
}
