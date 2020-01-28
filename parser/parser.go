package parser

import (
	"fmt"
	"io"
	"strings"
)

func parse(query string) (string, error) {
	queryReader := newReader(query)
	action := queryReader.getAction()

	currentTable := ""
	var lastToken *token = &token{}

	qb := &queryBuilder{}
	qb.SetAction(strings.ToUpper(action))

	for t, err := queryReader.readToken(); err != io.EOF; t, err = queryReader.readToken() {
		switch t.Type {
		case tableToken:
			lastToken.Name = "["            // keep track of last read token
			lastToken.Type = separatorToken // keep track of last read token
			// add table to query builder
			currentTable = t.Name
			qb.AddTable(t.Name)
		case columnToken:
			lastToken.Name = ","
			lastToken.Type = separatorToken
			qb.AddColumn(currentTable + "." + t.Name)
			if t.isLast == true {
				lastToken.Name = "]"
			}
		case separatorToken:
			if lastToken.Type == separatorToken && ((lastToken.Name == "]" && t.Name != "]" && t.Name != ",") || lastToken.Name == "," && t.Type == separatorToken) {
				return "", fmt.Errorf("Unexpected token %s at %d", string(t.Name), queryReader.getPosition())
			}
		default:
		}
	}
	sql, err := qb.ToSQL()
	if err != nil {
		return "", err
	}
	return sql, nil
}
