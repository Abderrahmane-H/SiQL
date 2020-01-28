package parser

import "fmt"

type queryBuilder struct {
	action  string
	tables  []string // tables names
	columns []string // columns names
	where   []string // where clause
	options []string // order by, limit offset ...
}

func (q *queryBuilder) AddTable(table string) {
	q.tables = append(q.tables, table)
}

func (q *queryBuilder) AddColumn(column string) {
	q.columns = append(q.columns, column)
}

func (q *queryBuilder) Where(condition string) {
	q.where = append(q.where, condition)
}

func (q *queryBuilder) Option(op string) {
	q.options = append(q.options, op)
}

func (q *queryBuilder) SetAction(action string) {
	q.action = action
}

func (q *queryBuilder) ToSQL() (string, error) {
	switch q.action {
	case "SELECT":
		query := q.action + " "
		for i := range q.columns { // columns
			if i == (len(q.columns) - 1) {
				query += q.columns[i] + " FROM "
			} else {
				query += q.columns[i] + ", "
			}
		}
		for i := range q.tables { // columns
			if i == (len(q.tables) - 1) {
				query += q.tables[i] + " "
			} else {
				query += q.tables[i] + ", "
			}
		}
		return query, nil
	case "UPDATE":
		return "", nil
	case "DELETE":
		return "", nil
	case "INSERT":
		return "", nil
	default:
		return "", fmt.Errorf("unsupported action %s", q.action)
	}
}
