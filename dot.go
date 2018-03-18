package godot

import (
	"bytes"
	"io/ioutil"
	"text/template"
)

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
