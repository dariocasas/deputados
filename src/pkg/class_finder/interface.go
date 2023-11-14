package class_finder

import "golang.org/x/net/html"

type ClassFinderInterface interface {
	SearchClass(tokeniker *html.Tokenizer, class, stopClass string) *TagFound
	SearchNextToken(tokeniker *html.Tokenizer, tokenName string) *TagFound
	ScanToEndToken(tokeniker *html.Tokenizer, endToken string) bool
	ScanToNextTextToken(tokenizer *html.Tokenizer,
		onText func(text string, index int, stop *bool),
	)
}
