// Package cssrefs returns a slice of `Reference{URI, Token string}`s from an `io.Reader`.
package cssrefs

import (
	"io"
	"io/ioutil"
	"regexp"

	"github.com/weisjohn/css/scanner"
)

// `Reference` are simply two strings: a `URI` and a `Token`
type Reference struct{ URI, Token string }

// a map of which tokens :: attr names to examine
var identTerminators = map[string]string{
	"@import":          ";",
	"@font-face":       "}",
	"background":       ";",
	"background-image": ";",
}

// a small map of the CSS-spec identifiers to token type
var identTokens = map[string]string{
	"@import":          "css",
	"@font-face":       "font",
	"background":       "img",
	"background-image": "img",
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

	// the current identifier matching against
	ident := ""
	_ = ident

	// the regex for `url("[match]")`, and for "[match]"
	urireg, _ := regexp.Compile(`url\([\'\"]([^)]+)[\'\"]\)`)
	quotes, _ := regexp.Compile(`[\'\"]([^)]+)[\'\"]`)

	for {

		// find the next token in the document
		token := doc.Next()

		// shorter access
		Type, Value := token.Type, token.Value

		// exit condition
		if Type == scanner.TokenEOF || Type == scanner.TokenError {
			break
		}

		// continue condition for Types
		switch Type {
		case scanner.TokenAtKeyword, scanner.TokenIdent, scanner.TokenURI, scanner.TokenString, scanner.TokenChar:
		default:
			continue
		}

		// continue conditions for Values
		switch Value {
		case "", ":", ",", "{":
			continue
		}

		// find identifiers
		if Type == scanner.TokenAtKeyword {
			if Value == "@import" || Value == "@font-face" {
				ident = Value
			}
			continue
		} else if Type == scanner.TokenIdent {
			if Value == "background" || Value == "background-image" {
				ident = Value
			}
			continue
		}

		// if we've found an identifier, find URIs based on uriregex
		if Type == scanner.TokenURI && ident != "" {
			urimatches := urireg.FindAllStringSubmatch(Value, -1)
			// if there's no url reference and ident != @import, bolt
			if len(urimatches) >= 1 && len(urimatches[0]) >= 2 {
				refs = append(refs, Reference{URI: urimatches[0][1], Token: identTokens[ident]})
			}
			continue
		}

		// look for non URI sources for @import statements
		if Type == scanner.TokenString && ident == "@import" {
			quotematches := quotes.FindAllStringSubmatch(Value, -1)
			if len(quotematches) >= 1 && len(quotematches[0]) >= 2 {
				refs = append(refs, Reference{URI: quotematches[0][1], Token: identTokens[ident]})
			}
			continue
		}

		// terminate finding an identifier
		if Type == scanner.TokenChar {
			if ident != "" && Value == identTerminators[ident] {
				ident = ""
			}
		}
	}

	return refs
}
