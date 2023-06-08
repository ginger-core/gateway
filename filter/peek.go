package filter

import (
	"regexp"
	"strings"
)

var reservedWords = []string{
	"(", ")",
	">=", "<=", "!=", "=", ">", "<",
	"&=", "&!=", "&~", "&!~", " IS ", " IN ",
}

func (p *parser) peek() string {
	peeked, _ := p.peekWithLength()
	return peeked
}

func (p *parser) peekWithLength() (string, int) {
	if p.i >= len(p.queryString) {
		return "", 0
	}
	for _, rWord := range reservedWords {
		token := strings.ToUpper(p.queryString[p.i:min(len(p.queryString), p.i+len(rWord))])
		if token == rWord {
			return token, len(token)
		}
	}
	if p.queryString[p.i] == '\'' { // Quoted string
		return p.peekQuotedStringWithLength()
	}
	return p.peekIdentifierWithLength()
}

func (p *parser) peekQuotedStringWithLength() (string, int) {
	if len(p.queryString) < p.i || p.queryString[p.i] != '\'' {
		return "", 0
	}
	for i := p.i + 1; i < len(p.queryString); i++ {
		if p.queryString[i] == '\'' && p.queryString[i-1] != '\\' {
			return p.queryString[p.i+1 : i], len(p.queryString[p.i+1:i]) + 2 // +2 for the two quotes
		}
	}
	return "", 0
}

func (p *parser) peekIdentifierWithLength() (string, int) {
	for i := p.i; i < len(p.queryString); i++ {
		if matched, _ := regexp.MatchString(`[a-zA-Z0-9_*]`, string(p.queryString[i])); !matched {
			return p.queryString[p.i:i], len(p.queryString[p.i:i])
		}
	}
	return p.queryString[p.i:], len(p.queryString[p.i:])
}
