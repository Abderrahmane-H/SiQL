package parser

import "testing"

func TestParseToTree(t *testing.T) {
	query := "select users[id,email,password,firstname,lastname]"
	action, tree, err := ParseToTree(query)
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
				t.Fatalf("expected root child %s but got %s", expectedColumns[i], tree.columns[i])
				break
			}
		}
	}
}

func TestParseToTreeInvalidQuery(t *testing.T) {
	query := "select users[id,email,,password,firstname,lastname]"
	_, _, err := ParseToTree(query)
	if err == nil {
		t.Fatal("expected err to be not nil but got nil")
	}
}

func BenchmarkParseToTree51QueryLength(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ParseToTree("select users[id,email,password,firstname,lastname]")
	}
}

func BenchmarkParseToTree103QueryLength(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		ParseToTree("select users[id, email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]")
	}
}

func BenchmarkParseToTree200QueryLength(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		ParseToTree("select users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]]]")
	}
}

func BenchmarkParseToTree400QueryLength(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		ParseToTree("select users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]]]]")
	}
}

func BenchmarkParseToTree5400QueryLength(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		ParseToTree("select users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]],users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at,users[id,email,password,firstname,lastname,products[id,title,descritpion,created_at,updated_at]]]]]]")
	}
}
