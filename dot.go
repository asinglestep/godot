package godot

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
	"text/template"
)

type Graph struct {
	Nodes []*Node
	Edges []*Edge
}

type Node struct {
	Name string
	Attr Attr
}

type Attr map[string]string

func (attr Attr) String() string {
	list := []string{}
	for k, v := range attr {
		list = append(list, fmt.Sprintf("%s=%s", k, v))
	}

	return strings.Join(list, ",")
}

type Edge struct {
	Src     string
	SrcPort string
	Dst     string
	DstPort string
}

const dotTemplate = `
digraph g {
	node [ height=.1, shape=record, width=.1 ];

	{{ range .Nodes }}
	{{ printf "\"%s\" [ %s ];" .Name .Attr.String }}
	{{ end }}

	{{ range .Edges }}
	{{ printf "\"%s\"%s -> \"%s\"%s;" .Src .SrcPort .Dst .DstPort }}
	{{ end }}
}
`

func GenerateDotFile(fileName string, g *Graph) error {
	var buf bytes.Buffer

	if err := template.Must(template.New("dot").Parse(dotTemplate)).Execute(&buf, g); err != nil {
		return err
	}

	return ioutil.WriteFile(fileName, buf.Bytes(), 0666)
}
