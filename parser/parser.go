package parser

import (
	"fmt"
	"io"
)

func ParseToTree(query string) (string, *queryTable, error) {
	queryReader := newReader(query)
	action := queryReader.getAction()

	var treeRoot *queryTable
	var currentTable *queryTable
	var lastToken *token = &token{}

	for t, err := queryReader.readToken(); err != io.EOF; t, err = queryReader.readToken() {
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
				return "", nil, fmt.Errorf("unexpected token %s at positon %d", t.Name, queryReader.getPosition())
			}
			currentTable.addColumn(t.Name)
			if t.isLast == true {
				currentTable = currentTable.parent
				lastToken.Name = "]"
			}
		case separatorToken:
			if lastToken.Type == separatorToken && ((lastToken.Name == "]" && t.Name != "]" && t.Name != ",") || lastToken.Name == "," && t.Type == separatorToken) {
				return "", nil, fmt.Errorf("Unexpected token %s at %d", string(t.Name), queryReader.getPosition())
			}
		default:
		}
	}

	return action, treeRoot, nil
}
