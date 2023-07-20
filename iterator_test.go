package iter_json

import (
	"strings"
	"testing"
)

func TestNewIterator(t *testing.T) {
	i := NewIterator()
	if i == nil {
		t.Error("NewIterator() should not return nil")
	}
}

func TestIterateObject(t *testing.T) {
	i := NewIterator()
	if i == nil {
		t.Error("NewIterator() should not return nil")
	}

	val := `{"a": 1}`
	ch, err := i.Iterate(strings.NewReader(val))
	if err != nil {
		t.Error("Iterate() should not return error")
	}
	for e := range ch {
		if e.Val.(float64) != 1 {
			t.Error("Wrong value")
		}
		if e.PathString() != "a" {
			t.Error("Wrong path")
		}
	}
}

func TestIterateArray(t *testing.T) {
	i := NewIterator()
	if i == nil {
		t.Error("NewIterator() should not return nil")
	}

	val := `[2]`
	ch, err := i.Iterate(strings.NewReader(val))
	if err != nil {
		t.Error("Iterate() should not return error")
	}
	for e := range ch {
		if e.Val.(float64) != 2 {
			t.Error("Wrong value")
		}
		if e.PathString() != "0" {
			t.Error("Wrong path")
		}
	}
}

func TestIterateNested(t *testing.T) {
	i := NewIterator()
	if i == nil {
		t.Error("NewIterator() should not return nil")
	}

	val := `{"a": {"b": 3}}`
	ch, err := i.Iterate(strings.NewReader(val))
	if err != nil {
		t.Error("Iterate() should not return error")
	}
	for e := range ch {
		if e.Val.(float64) != 3 {
			t.Error("Wrong value")
		}
		if e.PathString() != "a.b" {
			t.Error("Wrong path")
		}
	}
}

func TestIterateArrayNested(t *testing.T) {
	i := NewIterator()
	if i == nil {
		t.Error("NewIterator() should not return nil")
	}

	val := `{"a": [4]}`
	ch, err := i.Iterate(strings.NewReader(val))
	if err != nil {
		t.Error("Iterate() should not return error")
	}
	for e := range ch {
		if e.Val.(float64) != 4 {
			t.Error("Wrong value")
		}
		if e.PathString() != "a.0" {
			t.Error("Wrong path")
		}
	}
}

func TestIterateArrayNested2(t *testing.T) {
	i := NewIterator()
	if i == nil {
		t.Error("NewIterator() should not return nil")
	}

	val := `{"a": [{"b": 5}]}`
	ch, err := i.Iterate(strings.NewReader(val))
	if err != nil {
		t.Error("Iterate() should not return error")
	}
	for e := range ch {
		if e.Val.(float64) != 5 {
			t.Error("Wrong value")
		}
		if e.PathString() != "a.0.b" {
			t.Error("Wrong path")
		}
	}
}

func TestIterateEmptyObject(t *testing.T) {
	i := NewIterator()
	if i == nil {
		t.Error("NewIterator() should not return nil")
	}

	val := `{}`
	ch, err := i.Iterate(strings.NewReader(val))
	if err != nil {
		t.Error("Iterate() should not return error")
	}
	for range ch {
		t.Error("Should not iterate")
	}
}

func TestIterateEmptyArray(t *testing.T) {
	i := NewIterator()
	if i == nil {
		t.Error("NewIterator() should not return nil")
	}

	val := `[]`
	ch, err := i.Iterate(strings.NewReader(val))
	if err != nil {
		t.Error("Iterate() should not return error")
	}
	for range ch {
		t.Error("Should not iterate")
	}
}
