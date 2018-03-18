package godot

import (
	"fmt"
	"sort"
	"strings"
)

// Edge Edge
type Edge struct {
	Src     string
	SrcPort string
	Dst     string
	DstPort string
}

// Edges Edges
type Edges struct {
	Edges []*Edge
}

// Len Len
func (e *Edges) Len() int {
	return len(e.Edges)
}

// Swap Swap
func (e *Edges) Swap(i, j int) {
	e.Edges[i], e.Edges[j] = e.Edges[j], e.Edges[i]
}

// Less Less
func (e *Edges) Less(i, j int) bool {
	return e.Edges[i].Src < e.Edges[j].Src
}

// Sort Sort
func (e *Edges) Sort() []*Edge {
	sort.Sort(e)
	return e.Edges
}

// EdgeAttr EdgeAttr
type EdgeAttr map[string]string

// String String
func (attr EdgeAttr) String() string {
	list := []string{}
	for k, v := range attr {
		list = append(list, fmt.Sprintf("%s=%s", k, v))
	}

	return strings.Join(list, ",")
}

// Len Len
func (attr EdgeAttr) Len() int {
	return len(attr)
}
