package parser

import (
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	query := "select users[id,email,password,firstname,lastname]"
	eqSQL := "SELECT users.id, users.email, users.password, users.firstname, users.lastname FROM users"
	sql, err := parse(query)
	if err != nil {
		t.Fatal(err)
	}

	if strings.TrimSpace(sql) != strings.TrimSpace(eqSQL) {
		t.Fatalf("wrong sql %s expected %s", sql, eqSQL)
	}
}

func TestParseNestedTables(t *testing.T) {
	query := "select users[id,email,password,products[id, title]]"
	eqSQL := "SELECT users.id, users.email, users.password, products.id, products.title FROM users, products"
	sql, err := parse(query)
	if err != nil {
		t.Fatal(err)
	}

	if strings.TrimSpace(sql) != strings.TrimSpace(eqSQL) {
		t.Fatalf("wrong sql %s expected %s", sql, eqSQL)
	}
}

func TestParse3NestedTables(t *testing.T) {
	query := "select users[id,email,products[id, company[id]]]"
	eqSQL := "SELECT users.id, users.email, products.id, company.id FROM users, products, company"
	sql, err := parse(query)
	if err != nil {
		t.Fatal(err)
	}

	if strings.TrimSpace(sql) != strings.TrimSpace(eqSQL) {
		t.Fatalf("wrong sql %s expected %s", sql, eqSQL)
	}
}
