package godot

import (
	"os/exec"
	"runtime"
	"testing"
)

func Test_dot(t *testing.T) {
	dGraph := NewGraph()

	dGraph.SetNodeGlobalAttr(map[string]string{
		"height":  ".1",
		"shape":   "record",
		"width":   ".1",
		"rankdir": "LR",
		"rotate":  "90",
	})

	attr0 := "a0"
	dNode0 := &Node{}
	dNode0.Name = "node0"
	dNode0.Attr = map[string]string{
		"label": attr0,
	}

	dGraph.AddNode(dNode0)

	attr1 := "a1"
	dNode1 := &Node{}
	dNode1.Name = "node1"
	dNode1.Attr = map[string]string{
		"label": attr1,
	}

	dGraph.AddNode(dNode1)

	dEdge := &Edge{}
	dEdge.Src = "node0"
	dEdge.Dst = "node1"

	dGraph.AddEdge(dEdge)

	if err := GenerateDotFile("test.dot", dGraph); err != nil {
		t.Fatal("GenerateDotFile err: ", err)
		return
	}

	if err := exec.Command("dot", "-Tpng", "test.dot", "-o", "test.png").Run(); err != nil {
		t.Fatal("exec.Command err: ", err)
		return
	}

	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", "test.png")
	}

	if err := cmd.Run(); err != nil {
		t.Fatal("cmd.Run err: ", err)
		return
	}
}
