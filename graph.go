package godot

import (
	"fmt"
	"strings"
)

// Graph Graph
type Graph struct {
	NodeGlobalAttr []NodeAttr // node全局属性
	Nodes          Nodes

	EdgeGlobalAttr EdgeAttr // edge全局属性
	Edges          Edges

	Attr GraphAttr
}

// NewGraph NewGraph
func NewGraph() *Graph {
	g := &Graph{}
	g.Attr = GraphAttr{}
	return g
}

// AddNode AddNode
func (g *Graph) AddNode(node *Node) {
	g.Nodes.Nodes = append(g.Nodes.Nodes, node)
}

// AddEdge AddEdge
func (g *Graph) AddEdge(edge *Edge) {
	g.Edges.Edges = append(g.Edges.Edges, edge)
}

// SetNodeGlobalAttr 设置node全局属性
func (g *Graph) SetNodeGlobalAttr(attr NodeAttr) {
	g.NodeGlobalAttr = append(g.NodeGlobalAttr, attr)
}

// SetEdgeGlobalAttr 设置edge全局属性
func (g *Graph) SetEdgeGlobalAttr(attr EdgeAttr) {
	g.EdgeGlobalAttr = attr
}

// GraphAttr GraphAttr
type GraphAttr map[string]string

// Add Add
func (attr GraphAttr) Add(key string, val string) {
	attr[key] = val
}

// String String
func (attr GraphAttr) String() string {
	list := []string{}
	for k, v := range attr {
		list = append(list, fmt.Sprintf("%s=%s", k, v))
	}

	return strings.Join(list, ";\n")
}
