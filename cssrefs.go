// Package cssrefs returns a slice of `Reference{URI, Token string}`s from an `io.Reader`.
package cssrefs

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/weisjohn/css/scanner"
)

// `Reference` are simply two strings: a `URI` and a `Token`
type Reference struct{ URI, Token string }

// a map of which tokens :: attr names to examine
var identifierTerminators = map[string]bool{
	"@import":          ";",
	"@font-face":       "}",
	"background":       ";",
	"background-image": ";",
}

// `All` takes a reader object (like the one returned from http.Get())
// It returns a slice of struct Reference{uri, nodetype}
// It does not close the reader passed to it.
func All(httpBody io.Reader) []Reference {
	refs := make([]Reference, 0)

	// copy the reader into a new buffer
	b, _ := ioutil.ReadAll(httpBody)

	// create a new document
	doc := scanner.New(string(b))

	// the current identifier, if any
	ident := ""

	for {

		// find the next token in the document
		token := doc.Next()

		// shorter access
		Type := token.Type

		// exit condition
		if Type == scanner.TokenEOF || Type == scanner.TokenError {
			break
		}

		// continue condition
		switch Type {
		case scanner.TokenAtKeyword, scanner.TokenIdent, scanner.TokenChar, scanner.TokenURI:
		default:
			continue
		}

		Value := token.Value

		fmt.Println(Type, Value)

		// find identifiers
		if Type == scanner.TokenAtKeyword {

		} else if Type == scanner.TokenIdent {
			if Value == "background" || Value == "background-image" {

			}
			continue
		}

	}

	return refs
}
