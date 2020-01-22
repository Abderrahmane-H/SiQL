package parser

import (
	"io"
	"strings"
)

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

func ParseToTree(query string) (string, *queryTable, error) {
	queryReader := strings.NewReader(query)
	action := getAction(queryReader)

	var treeRoot *queryTable
	var currentTable *queryTable

	token, next, err := readToken(queryReader)
	if err != nil {
		return "", nil, err
	}
	for token != "golang.io.EOF" {
		switch token {
		case "[":
			// validate next token
			// should be a character different than [, ], "," (start of a new identifier)
		case ",":
			// validate next token
			// should be a character different than [, ], "," (start of a new identifier)
		case "]":
			if currentTable.parent != nil {
				currentTable = currentTable.parent
			}
		case "golang.io.EOF": // end of string
		default: // an identifier either table or column name or options
			switch next {
			case "[": // table, next elements are either columns or tables
				// tree not initialized yet (we read the first table)
				if currentTable == nil {
					treeRoot = &queryTable{
						name:    token,
						childs:  nil,
						columns: nil,
						parent:  nil,
					}
					currentTable = treeRoot
				} else {
					// initialize a new table node
					node := &queryTable{
						name:    token,
						childs:  nil,
						columns: nil,
						parent:  currentTable,
					}
					// add the table node to the childs of the current table
					currentTable.childs = append(currentTable.childs, node)
					// set the current table to be the new one
					currentTable = node
				}
			case "]": // column and the parent table is closed
				currentTable.columns = append(currentTable.columns, token)
			case ",": // column
				currentTable.columns = append(currentTable.columns, token)
			}
		}
		token, next, err = readToken(queryReader)
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

// readToken -- reads current token and returns the current token and the next one
// the read cursor is set back to the end of the current token
// returned values :
// (a_word, golang.io.EOF) end of string
// (a_word, [) the read token is a table
// (a_word, "") the read token is either [, "," or ]
// (a_word, ",") the read token was a column and the we expect a column or a table after
// (a_word, "]") end of current table expect query options if current table is parent,
//    otherwise expect other columns or other tables
func readToken(reader *strings.Reader) (string, string, error) {
	word := ""
	nextToken := ""
	for {
		b, err := reader.ReadByte()
		if b == 32 { // ignore spaces
			continue
		}
		if err == io.EOF {
			if len(word) > 0 {
				break
			} else {
				nextToken = "golang.io.EOF"
				word = "golang.io.EOF"
				break
			}
		}
		if b == 91 && len(word) > 0 { // we read a [ after reading a word (Table)
			nextToken = string(b)
			reader.UnreadByte()
			break
		} else if b == 91 { // we read [, after this all we will have are probably columns
			nextToken = ""
			break
		}

		if b == 44 && len(word) > 0 { // we read a "," after reading a word (Column)
			nextToken = string(b)
			reader.UnreadByte()
			break
		} else if b == 44 { // we read , then a after this we will have probably columns or another table
			nextToken = ""
			break
		}

		if b == 93 && len(word) > 0 { // we read ] after reading a word (table end)
			nextToken = string(b)
			reader.UnreadByte()
			break
		} else if b == 93 { // we read ] table end
			nextToken = ""
			break
		}
		word += string(b)
	}
	return word, nextToken, nil
}

func isValidQueryCharacter(identifier string) bool {
	return true
}
