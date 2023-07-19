package iter_json

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type JsonIterator struct {
}

func NewIterator() *JsonIterator {
	return &JsonIterator{}
}

type PathKind int

const (
	_ PathKind = iota
	OBJECT
	ARRAY
)

type JsonPathKind struct {
	Kind PathKind
	Key  interface{}
}

type JsonEntry struct {
	Path []JsonPathKind
	Val  interface{}
}

func (e *JsonEntry) PathString() string {
	var path []string
	for _, p := range e.Path {
		switch p.Kind {
		case OBJECT:
			path = append(path, fmt.Sprintf("%s", p.Key))
		case ARRAY:
			path = append(path, fmt.Sprintf("%d", p.Key))
		}
	}
	return strings.Join(path, ".")
}

func walkArray(dec *json.Decoder, out chan<- JsonEntry, path []JsonPathKind) {
	for i := 0; ; i++ {
		t, err := dec.Token()
		if err == io.EOF {
			return
		}
		if err != nil {
			return
		}
		switch t.(type) {
		case json.Delim:
			switch t.(json.Delim) {
			case '{':
				walkObject(dec, out, append(path, JsonPathKind{ARRAY, i}))
			case '[':
				walkArray(dec, out, append(path, JsonPathKind{ARRAY, i}))
			case '}':
				return
			case ']':
				return
			}
		case string, bool, nil, float64, json.Number:
			out <- JsonEntry{append(path, JsonPathKind{ARRAY, i}), t}
		}
	}
}

func walkObject(dec *json.Decoder, out chan<- JsonEntry, path []JsonPathKind) {
	for {
		key, err := dec.Token()
		if err == io.EOF {
			return
		}
		if err != nil {
			return
		}
		switch key.(type) {
		case json.Delim:
			switch key.(json.Delim) {
			case '{':
				walkObject(dec, out, path)
			case '[':
				walkArray(dec, out, path)
			case '}':
				return
			case ']':
				return
			}
		}
		val, err := dec.Token()
		if err == io.EOF {
			return
		}
		if err != nil {
			return
		}
		switch val.(type) {
		case json.Delim:
			switch val.(json.Delim) {
			case '{':
				walkObject(dec, out, append(path, JsonPathKind{OBJECT, key}))
			case '[':
				walkArray(dec, out, append(path, JsonPathKind{OBJECT, key}))
			case '}':
				return
			case ']':
				return
			}
		case string, bool, nil, float64, json.Number:
			out <- JsonEntry{append(path, JsonPathKind{OBJECT, key}), val}
		}
	}
}

func (i *JsonIterator) Iterate(val string) (<-chan JsonEntry, error) {
	ch := make(chan JsonEntry)

	go func(out chan<- JsonEntry) {
		defer close(out)
		decoder := json.NewDecoder(strings.NewReader(val))
		t, err := decoder.Token()
		if err == io.EOF {
			return
		}
		if err != nil {
			// TODO: maybe add an error channel
			return
		}
		switch t.(type) {
		case json.Delim:
			switch t.(json.Delim) {
			case '{':
				walkObject(decoder, out, []JsonPathKind{})
			case '[':
				walkArray(decoder, out, []JsonPathKind{})
			}
		}
	}(ch)
	return ch, nil
}
