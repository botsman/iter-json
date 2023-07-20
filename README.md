# iter-json

This tool is designed for iterating over a large (and probably unstructured) JSON files.  
It is useful for processing JSON files that are too large to fit into memory.

## Example

```go
package main

import (
    "fmt"
    "github.com/botsman/iter-json"
)

func main() {
    iter := NewIterator()
	ch, err := iter.Iterate(`{"a": 1, "b": 2, "c": 3, "d": [1, 2, 3], "e": {"f": {"g": 4}}}`)
	if err != nil {
		panic(err)
	}
	for e := range ch {
		fmt.Println(e.PathString(), e.Val)
	}
}
```

Output:

```
a 1
b 2
c 3
d.0 1
d.1 2
d.2 3
e.f.g 4
```

Interface:  
After creating an iterator (`NewIterator()`), you can iterate over a JSON string by calling `Iterate()` method.   
It returns a channel of `Entry` objects:

```go
type PathKind int

const (
    _ PathKind = iota
    OBJECT
    ARRAY
)

type JsonPathElement struct {
    Kind PathKind
    Key  interface{}
}

type JsonEntry struct {
    Path []JsonPathElement
    Val  interface{}
}
```
