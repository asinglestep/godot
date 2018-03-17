package godot

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"text/template"
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

// dotTemplate dotTemplate
const dotTemplate = `
digraph g {
	{{ .Attr.String }}

	{{ range .NodeGlobalAttr }}
	{{ if .Len }}
	{{ printf "node [ %s ];" .String }}
	{{ end }}
	{{ end }}

	{{ if .EdgeGlobalAttr.Len }}
	{{ printf "edge [ %s ];" .EdgeGlobalAttr.String }}
	{{ end }}

	{{ range .Nodes.Sort }}
	{{ if .Attr.Len }}
	{{ printf "%s [ %s ];" .Name .Attr.String }}
	{{ else }}
	{{ printf "%s;" .Name }}
	{{ end }}
	{{ end }}

	{{ range .Edges.Sort }}
	{{ printf "%s%s -> %s%s;" .Src .SrcPort .Dst .DstPort }}
	{{ end }}
}
`

// GenerateDotFile 生成dot文件
func GenerateDotFile(fileName string, g *Graph) error {
	var buf bytes.Buffer

	if err := template.Must(template.New("dot").Parse(dotTemplate)).Execute(&buf, g); err != nil {
		return err
	}

	return ioutil.WriteFile(fileName, buf.Bytes(), 0666)
}
