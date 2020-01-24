package parser

import (
	"fmt"
	"testing"
)

func TestParseToTree(t *testing.T) {
	query := "select users[id,email,password,firstname,lastname]"
	action, tree, err := parseToTree(query)
	if err != nil {
		t.Fatal(err)
	}

	if action != "select" {
		t.Fatalf("expected select to be the action of : %s", query)
	}

	if tree == nil {
		t.Fatalf("expected tree to be not nil from %s", query)
	} else {
		if tree.name != "users" {
			t.Fatalf("expected root table name to be users but got %s from %s", tree.name, query)
		}

		if tree.childs != nil {
			t.Fatalf("expected root childs to be nil from %s", query)
		}

		if tree.columns == nil {
			t.Fatalf("expected root columns to be not nil from %s", query)
		}
		expectedColumns := []string{"id", "email", "password", "firstname", "lastname"}
		if len(expectedColumns) != len(tree.columns) {
			t.Fatalf("expected root columns length to be %d but got %d", len(expectedColumns), len(tree.columns))
		}

		for i := range tree.columns {
			if expectedColumns[i] != tree.columns[i] {
				t.Fatalf("expected root to have column %s but got %s", expectedColumns[i], tree.columns[i])
				break
			}
		}
	}
}

func TestParseToTreeTableBetweenColumns(t *testing.T) {
	query := "select users[id,email,password,firstname,products[id, email],lastname]"
	action, tree, err := parseToTree(query)
	if err != nil {
		t.Fatal(err)
	}

	if action != "select" {
		t.Fatalf("expected select to be the action of : %s", query)
	}

	if tree == nil {
		t.Fatalf("expected tree to be not nil from %s", query)
	} else {
		if tree.name != "users" {
			t.Fatalf("expected root table name to be users but got %s from %s", tree.name, query)
		}

		if tree.childs == nil {
			t.Fatalf("expected root childs to be not nil %s", query)
		}

		if tree.columns == nil {
			t.Fatalf("expected root columns to be not nil from %s", query)
		}
		expectedColumns := []string{"id", "email", "password", "firstname", "lastname"}
		if len(expectedColumns) != len(tree.columns) {
			fmt.Println(tree.columns)
			t.Fatalf("expected root columns length to be %d but got %d", len(expectedColumns), len(tree.columns))
		}

		for i := range tree.columns {
			if expectedColumns[i] != tree.columns[i] {
				t.Fatalf("expected root to have column %s but got %s", expectedColumns[i], tree.columns[i])
				break
			}
		}
	}
}

func TestParseToTreeNestedTables(t *testing.T) {
	query := "select users[id, email,password,firstname,lastname,products[id,title,description,created_at,updated_at]]"
	action, tree, err := parseToTree(query)
	if err != nil {
		t.Fatal(err)
	}

	if action != "select" {
		t.Fatalf("expected select to be the action of : %s", query)
	}

	if tree == nil {
		t.Fatalf("expected tree to be not nil from %s", query)
	} else {
		if len(tree.childs) == 0 {
			t.Fatalf("expected root table name to have childs")
		}
		if tree.childs[0].name != "products" {
			t.Fatalf("expected child table name to be users but got %s from %s", tree.childs[0].name, query)
		}

		if tree.childs[0].childs != nil {
			t.Fatalf("expected child childs to be nil from %s", query)
		}

		if tree.childs[0].columns == nil {
			t.Fatalf("expected child columns to be not nil from %s", query)
		}
		expectedColumns := []string{"id", "title", "description", "created_at", "updated_at"}
		if len(expectedColumns) != len(tree.childs[0].columns) {
			t.Fatalf("expected child columns length to be %d but got %d", len(expectedColumns), len(tree.childs[0].columns))
		}

		for i := range tree.childs[0].columns {
			if expectedColumns[i] != tree.childs[0].columns[i] {
				t.Fatalf("expected child to have column %s but got %s", expectedColumns[i], tree.childs[0].columns[i])
				break
			}
		}
	}
}

func TestParseToTreeInvalidQuery(t *testing.T) {
	query := "select users[id,email,,password,firstname,lastname]"
	_, _, err := parseToTree(query)
	if err == nil {
		t.Fatal("expected err to be not nil but got nil")
	}
}

func BenchmarkParseToTree51QueryLength(b *testing.B) {
	for n := 0; n < b.N; n++ {
		parseToTree("select users[id,email,password,firstname,lastname]")
	}
}

func BenchmarkParseToTree103QueryLength(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		parseToTree("select users[id, email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]")
	}
}

func BenchmarkParseToTree200QueryLength(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		parseToTree("select users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]]]")
	}
}

func BenchmarkParseToTree400QueryLength(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		parseToTree("select users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]]]]")
	}
}

func BenchmarkParseToTree5400QueryLength(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		parseToTree("select users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]]]]]")
	}
}
