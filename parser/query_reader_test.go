package parser

import "testing"

func TestQueryReaderGetPosition(t *testing.T) {
	qr := newReader("select users[id, email]")
	if qr.getPosition() > 0 {
		t.Fatal("query position after initialization should be 0")
	}
	qr.readNext()
	if qr.getPosition() != 1 {
		t.Fatal("query position after initialization should be 0")
	}
	qr.readNext()
	if qr.getPosition() != 2 {
		t.Fatal("query position after initialization should be 0")
	}
}

func TestQueryReaderReadToken(t *testing.T) {
	qr := newReader("select users[id, email]")

	qr.getAction()

	v, err := qr.readToken()
	if err != nil {
		t.Fatalf("expected error to be not nil but got %s", err.Error())
	}
	if v.Name != "users" || v.Type != tableToken {
		t.Fatalf("expected token users to be tableToken")
	}
}

func TestQueryReaderGetAction(t *testing.T) {
	qr := newReader("select users[id, email]")

	a := qr.getAction()
	if a != "select" {
		t.Fatalf("expected action to be select but got %s", a)
	}
}

func BenchmarkQueryReaderReadNext(b *testing.B) {
	for i := 0; i < b.N; i++ {
		qr := newReader("select users[id, email]")
		qr.readNext()
	}
}
