package parser

type queryOptions struct {
	name       string // the name that should be used in sorm query language
	equivalent string // the equivalent of the current option in sql (orderby => order by)
}

var supported_options []queryOptions = []queryOptions{
	{name: "orderby", equivalent: "ORDER BY"},
	{name: "asc", equivalent: "ASC"},
	{name: "desc", equivalent: "DESC"},
	{name: "limit", equivalent: "LIMIT"},
	{name: "offset", equivalent: "OFFSET"},
}
