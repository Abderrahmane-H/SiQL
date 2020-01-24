package parser

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
