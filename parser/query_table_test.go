package parser

import "testing"

func TestQueryTableAddcolumn(t *testing.T) {
	qt := &queryTable{}

	qt.Addcolumn("id")
	if len(qt.columns) != 1 {
		t.Fatalf("expected columns length to be 1 but got %d", len(qt.columns))
	}
}

func TestQueryTableAddChild(t *testing.T) {
	qt := &queryTable{}

	qt.AddChild(&queryTable{})
	if len(qt.childs) != 1 {
		t.Fatalf("expected childs length to be 1 but got %d", len(qt.childs))
	}
}

func BenchmarkQueryTableAddcolumn(b *testing.B) {
	qt := &queryTable{}
	for i := 0; i < b.N; i++ {
		qt.Addcolumn("id")
	}
}

func BenchmarkQueryTableAddChild(b *testing.B) {
	qt := &queryTable{}
	for i := 0; i < b.N; i++ {
		qt.AddChild(&queryTable{})
	}
}
