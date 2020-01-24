package parser

import (
	"fmt"
	"io"
	"strings"
)

const endOfQuery string = "golang.io.EOF"

type tokenType byte

const (
	tableToken     tokenType = 1
	columnToken    tokenType = 2
	separatorToken tokenType = 3
)

type token struct {
	Name   string
	Type   tokenType
	isLast bool
}

type queryReader struct {
	*strings.Reader
	position int64
	line     int64 // todo
}

func (r *queryReader) readNext() (byte, error) {
	b, err := r.ReadByte()
	if err == nil {
		r.position++
	}
	// ignore whitespaces
	for isWhiteSpace(b) {
		b, err = r.ReadByte()
		if err == nil {
			r.position++
		}
	}
	return b, err
}

func (r *queryReader) getPosition() int64 {
	return r.position
}

// readToken -- reads and returns a token
func (r *queryReader) readToken() (*token, error) {
	var word []byte = []byte{}
	var separator byte = 0
	for { // read until we find a separator or EOF
		b, err := r.readNext()
		if err != nil { // could be io.EOF
			return nil, err
		}
		// we found a separator, break the loop
		if isSeparator(b) && separator == 0 {
			separator = b
			break
		}
		// not a separator
		if isValidByte(b) == false { // make sure the byte is valid character
			return nil, fmt.Errorf("Unexpected token %s at %d", string(b), r.getPosition())
		}

		word = append(word, b)
	}

	if len(word) > 0 {
		switch separator {
		case 91: // [
			return &token{Name: string(word), Type: tableToken}, nil
		case 44: // ,
			return &token{Name: string(word), isLast: false, Type: columnToken}, nil
		case 93: // ]
			return &token{Name: string(word), isLast: true, Type: columnToken}, nil
		default:
			return nil, fmt.Errorf("Unexpected token %s at %d, expected [ or ] or ,", string(separator), r.getPosition())
		}
	} else {
		return &token{Name: string(separator), Type: separatorToken}, nil
	}
}

// returns the first token of the query
// which is an action to be sent to the database
// supported actions are : select, insert, update and delete
func (r *queryReader) getAction() string {
	action := ""
	for {
		b, err := r.ReadByte()
		if err == io.EOF {
			r.UnreadByte()
			break
		}
		if b == 32 {
			break
		}
		action += string(b)
	}
	return action
}

func isWhiteSpace(b byte) bool {
	return (b == 32 || b == 10 || b == 13 || b == 9)
}

func isSeparator(b byte) bool {
	return (b == 91 || b == 93 || b == 44 || b == 123 || b == 125)
}

func isValidByte(b byte) bool {
	return (b >= 97 && b <= 122) || (b >= 65 && b <= 90) || (b == 42) || (b == 95)
}

func newReader(query string) *queryReader {
	return &queryReader{Reader: strings.NewReader(query), position: 0, line: 0}
}
