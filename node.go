package godot

import (
	"fmt"
	"sort"
	"strings"
)

// Node Node
type Node struct {
	Name string
	Attr NodeAttr
}

// Nodes Nodes
type Nodes struct {
	Nodes []*Node
}

// Len Len
func (n *Nodes) Len() int {
	return len(n.Nodes)
}

// Swap Swap
func (n *Nodes) Swap(i, j int) {
	n.Nodes[i], n.Nodes[j] = n.Nodes[j], n.Nodes[i]
}

// Less Less
func (n *Nodes) Less(i, j int) bool {
	return n.Nodes[i].Name < n.Nodes[j].Name
}

// Sort Sort
func (n *Nodes) Sort() []*Node {
	sort.Sort(n)
	return n.Nodes
}

// NodeAttr NodeAttr
type NodeAttr map[string]string

// String String
func (attr NodeAttr) String() string {
	list := []string{}
	for k, v := range attr {
		list = append(list, fmt.Sprintf("%s=%s", k, v))
	}

	return strings.Join(list, ",")
}

// Len Len
func (attr NodeAttr) Len() int {
	return len(attr)
}
