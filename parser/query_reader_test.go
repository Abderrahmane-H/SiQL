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

func BenchmarkQueryReaderReadNext(b *testing.B) {
	for i := 0; i < b.N; i++ {
		qr := newReader("select users[id, email]")
		qr.readNext()
	}
}
