package parser

import (
	"fmt"
	"io"
	"strings"
)

const endOfQuery string = "golang.io.EOF"

type queryTable struct {
	// the name of the table
	name string
	// childs of this table
	// example:
	// users [id, email, products[id, title]]
	// products title is child of users table
	childs []*queryTable
	// the parent of this child
	// used so we can go back in the tree easily
	parent *queryTable
	//columns of this table
	// example: users [id, email, password, products[id, title]]
	// users.columns : id, email, password
	// products.columns: id, title
	columns []string
}

func (c *queryTable) addColumn(column string) {
	c.columns = append(c.columns, column)
}

func (c *queryTable) addChild(child *queryTable) {
	c.childs = append(c.childs, child)
}

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

func ParseToTree(query string) (string, *queryTable, error) {
	queryReader := strings.NewReader(query)
	action := getAction(queryReader)

	var treeRoot *queryTable
	var currentTable *queryTable
	var lastToken *token = &token{}

	for t, err := readToken(queryReader); err != io.EOF; t, err = readToken(queryReader) {
		switch t.Type {
		case tableToken:
			lastToken.Name = "["
			lastToken.Type = separatorToken
			if currentTable == nil { // first Table
				treeRoot = &queryTable{
					name:    t.Name,
					childs:  nil,
					parent:  nil,
					columns: nil,
				}
				currentTable = treeRoot
			} else { // child table
				table := &queryTable{
					name:    t.Name,
					childs:  nil,
					parent:  currentTable,
					columns: nil,
				}
				currentTable.addChild(table)
				currentTable = table
			}
		case columnToken:
			lastToken.Name = ","
			lastToken.Type = separatorToken
			if currentTable == nil {
				return "", nil, fmt.Errorf("unexpected token %s at positon %d", t.Name, getCurrentPosition(queryReader))
			}
			currentTable.addColumn(t.Name)
			if t.isLast == true {
				currentTable = currentTable.parent
				lastToken.Name = "]"
			}
		case separatorToken:
			if lastToken.Type == separatorToken && ((lastToken.Name == "]" && t.Name != "]" && t.Name != ",") || lastToken.Name == "," && t.Type == separatorToken) {
				return "", nil, fmt.Errorf("Unexpected token %s at %d", string(t.Name), getCurrentPosition(queryReader))
			}
		default:
		}
	}

	return action, treeRoot, nil
}

// select users(id)[id, email, password, products(id)[id, title]]

// returns the first token of the query
// which is an action to be sent to the database
// supported actions are : select, insert, update and delete
func getAction(reader *strings.Reader) string {
	action := ""
	for {
		b, err := reader.ReadByte()
		if err == io.EOF {
			reader.UnreadByte()
			break
		}
		if b == 32 {
			break
		}
		action += string(b)
	}
	return action
}

// readToken -- reads and returns a token
func readToken(reader *strings.Reader) (*token, error) {
	var word []byte = []byte{}
	var separator byte = 0
	for { // read until we find a separator or EOF
		b, err := readByte(reader)
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
			return nil, fmt.Errorf("Unexpected token %s at %d", string(b), getCurrentPosition(reader))
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
			return nil, fmt.Errorf("Unexpected token %s at %d, expected [ or ] or ,", string(separator), getCurrentPosition(reader))
		}
	} else {
		return &token{Name: string(separator), Type: separatorToken}, nil
	}
}

func readByte(reader *strings.Reader) (byte, error) {
	b, err := reader.ReadByte()
	// ignore whitespaces
	for isWhiteSpace(b) {
		b, err = reader.ReadByte()
	}
	return b, err
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

func getCurrentPosition(reader *strings.Reader) int64 {
	return reader.Size() - int64(reader.Len())
}
