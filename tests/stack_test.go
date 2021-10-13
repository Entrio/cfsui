package tests

import (
	"github.com/Entrio/cfsui/internal"
	"testing"
)

func BenchmarkStack(b *testing.B) {
	s := internal.NewStack()
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}

	for j := 0; j < b.N; j++ {
		s.Pop()
	}
}

func TestStackPush(t *testing.T) {
	s := internal.NewStack()
	s.Push(1)

	if s.Count() != 1 {
		t.Fatalf("Stack count was not 1 afre a single push")
	}
}

func TestStackPop(t *testing.T) {
	s := internal.NewStack()
	s.Push(1)
	s.Pop()

	if got, want := s.Count(), 0; got != want {
		t.Fatalf("wanted %d, got %d after pop", want, got)
	}
}

func TestStackPushMulti(t *testing.T) {
	s := internal.NewStack()

	cnt := 200
	for i := 0; i < cnt; i++ {
		s.Push(i)
	}

	if got, want := s.Count(), cnt; got != want {
		t.Fatalf("wanted %d, got %d after push", want, got)
	}
}

func TestStackPopMulti(t *testing.T) {
	s := internal.NewStack()
	cnt := 200
	for i := 0; i < cnt; i++ {
		s.Push(i)
	}

	if got, want := s.Count(), cnt; got != want {
		t.Fatalf("wanted %d, got %d after push", want, got)
	}

	for j := cnt; j > 0; j-- {
		s.Pop()
	}

	if got, want := s.Count(), 0; got != want {
		t.Fatalf("wanted %d, got %d after pop", want, got)
	}
}
