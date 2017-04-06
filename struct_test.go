package triplestore

import (
	"testing"
	"time"
)

type TestStruct struct {
	Name  string    `predicate:"name"`
	Age   int       `predicate:"age"`
	Size  int64     `predicate:"size"`
	Male  bool      `predicate:"male"`
	Birth time.Time `predicate:"birth"`

	// special cases that should be ignored
	NoTag       string
	Unsupported complex64 `predicate:"complex"`
	Pointer     *string   `predicate:"ptr"`
}

func TestStructToTriple(t *testing.T) {
	now := time.Now()
	s := TestStruct{
		Name:  "donald",
		Age:   32,
		Size:  186,
		Male:  true,
		Birth: now,
	}

	exp := []Triple{
		SubjPred("me", "name").StringLiteral("donald"),
		SubjPred("me", "age").IntegerLiteral(32),
		SubjPred("me", "size").IntegerLiteral(186),
		SubjPred("me", "male").BooleanLiteral(true),
		SubjPred("me", "birth").DateTimeLiteral(now),
	}

	tris := TriplesFromStruct("me", s)
	if got, want := Triples(tris), Triples(exp); !got.Equal(want) {
		t.Fatalf("got %s\n\n want %s", got, want)
	}

	tris = TriplesFromStruct("me", &s)
	if got, want := Triples(tris), Triples(exp); !got.Equal(want) {
		t.Fatalf("got %s\n\n want %s", got, want)
	}
}

func TestReturnEmptyTriplesOnNonStructElem(t *testing.T) {
	var ptr *string
	tcases := []struct {
		val interface{}
	}{
		{true}, {"any"}, {ptr},
	}

	for i, tc := range tcases {
		tris := TriplesFromStruct("", tc.val)
		if len(tris) != 0 {
			t.Fatalf("case %d: expected no triples", i+1)
		}
	}
}
