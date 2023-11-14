package class_finder

import (
	"strings"

	"golang.org/x/net/html"
)

type TagFound struct {
	tokenType html.TokenType
	Token     html.Token
	class     string
}

type ClassFinder struct {
}

func NewClassFinder() *ClassFinder {
	return &ClassFinder{}
}

func (c ClassFinder) SearchClass(tokeniker *html.Tokenizer, class, stopClass string) *TagFound {

	result := &TagFound{}

	for {
		result.tokenType = tokeniker.Next()
		switch {
		case result.tokenType == html.ErrorToken:
			return nil

		case result.tokenType == html.StartTagToken:
			result.Token = tokeniker.Token()

			for _, attrib := range result.Token.Attr {
				if attrib.Key == "class" {
					result.class = strings.Split(attrib.Val, " ")[0]
					if result.class > "" && result.class == stopClass {
						return nil
					}
				}
				if result.class == class {
					return result
				}
			}
		}
	}
}

func (c ClassFinder) SearchNextToken(tokeniker *html.Tokenizer, tokenName string) *TagFound {

	result := &TagFound{}

	for {
		result.tokenType = tokeniker.Next()
		switch {
		case result.tokenType == html.ErrorToken:
			return nil

		case result.tokenType == html.StartTagToken:
			result.Token = tokeniker.Token()
			if result.Token.Data == tokenName {
				return result
			}
			return nil

		}
	}
}

func (c ClassFinder) ScanToEndToken(tokeniker *html.Tokenizer, endToken string) bool {

	for {
		tokenType := tokeniker.Next()
		switch {
		case tokenType == html.ErrorToken:
			return false

		case tokenType == html.EndTagToken:
			token := tokeniker.Token()
			tokenName := token.Data
			if tokenName == "" || tokenName == endToken {
				return true
			}
		}
	}
}

func (c ClassFinder) ScanToNextTextToken(tokenizer *html.Tokenizer,
	onText func(text string, index int, stop *bool),
) {

	var index = 0
	var stop = false

	for !stop {
		tokenType := tokenizer.Next()

		switch {
		case tokenType == html.ErrorToken:
			return

		case tokenType == html.TextToken:
			token := tokenizer.Token()
			text := strings.Trim(
				strings.TrimLeft(
					strings.TrimRight(
						strings.Trim(token.Data, " "),
						"\n"),
					"\n"),
				" ")
			onText(text, index, &stop)
			index++
		}
	}
}
