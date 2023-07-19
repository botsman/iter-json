package iter_json

import (
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

	ch, err := i.Iterate(`{ "a": 1 }`)
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

	ch, err := i.Iterate(`[2]`)
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

	ch, err := i.Iterate(`{"a": {"b": 3}}`)
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

	ch, err := i.Iterate(`{"a": [4]}`)
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

	ch, err := i.Iterate(`{"a": [{"b": 5}]}`)
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

	ch, err := i.Iterate(`{}`)
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

	ch, err := i.Iterate(`[]`)
	if err != nil {
		t.Error("Iterate() should not return error")
	}
	for range ch {
		t.Error("Should not iterate")
	}
}
