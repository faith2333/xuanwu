package dag

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type MyNode struct {
	name     string
	runAfter []string
}

func (my *MyNode) NodeName() string {
	return my.name
}

func (my *MyNode) PrevNodeNames() []string {
	return my.runAfter
}

func TestNew_Parallel(t *testing.T) {

	// parallel
	//    a   b   c

	a := &MyNode{name: "a"}
	b := &MyNode{name: "b"}
	c := &MyNode{name: "c"}

	g, err := NewDAG([]NamedNode{a, b, c})
	assert.NoError(t, err)
	assertNode(t, g.Nodes["a"], "a", nil, nil)
	assertNode(t, g.Nodes["b"], "b", nil, nil)
	assertNode(t, g.Nodes["c"], "c", nil, nil)
}

func TestNew_JoinMultipleRoots(t *testing.T) {

	//   a    b   c
	//   | \ /
	//   x  y
	//   |
	//   z

	a := &MyNode{name: "a"}
	b := &MyNode{name: "b"}
	c := &MyNode{name: "c"}
	xRunAfterA := &MyNode{name: "x", runAfter: []string{"a"}}
	yRunAfterAAndB := &MyNode{name: "y", runAfter: []string{"a", "b"}}
	zRunAfterX := &MyNode{name: "z", runAfter: []string{"x"}}

	g, err := NewDAG([]NamedNode{a, b, c, xRunAfterA, yRunAfterAAndB, zRunAfterX})
	assert.NoError(t, err)
	assertNode(t, g.Nodes["a"], "a", nil, []string{"x", "y"})
	assertNode(t, g.Nodes["b"], "b", nil, []string{"y"})
	assertNode(t, g.Nodes["c"], "c", nil, nil)
	assertNode(t, g.Nodes["x"], "x", []string{"a"}, []string{"z"})
	assertNode(t, g.Nodes["y"], "y", []string{"a", "b"}, nil)
	assertNode(t, g.Nodes["z"], "z", []string{"x"}, nil)
}

func TestBuild_FanInFanOut(t *testing.T) {

	//   a
	//  / \
	// d   e
	//  \ /
	//   f
	//   |
	//   g

	a := &MyNode{name: "a"}
	dRunAfterA := &MyNode{name: "d", runAfter: []string{"a"}}
	eRunAfterA := &MyNode{name: "e", runAfter: []string{"a"}}
	fRunAfterDAndE := &MyNode{name: "f", runAfter: []string{"d", "e"}}
	gRunAfterF := &MyNode{name: "g", runAfter: []string{"f"}}

	g, err := NewDAG([]NamedNode{a, dRunAfterA, eRunAfterA, fRunAfterDAndE, gRunAfterF})
	assert.NoError(t, err)
	assertNode(t, g.Nodes["a"], "a", nil, []string{"d", "e"})
	assertNode(t, g.Nodes["d"], "d", []string{"a"}, []string{"f"})
	assertNode(t, g.Nodes["e"], "e", []string{"a"}, []string{"f"})
	assertNode(t, g.Nodes["f"], "f", []string{"d", "e"}, []string{"g"})
	assertNode(t, g.Nodes["g"], "g", []string{"f"}, nil)
}

func TestNew_Invalid(t *testing.T) {
	a := &MyNode{name: "a"}
	xRunAfterA := &MyNode{name: "x", runAfter: []string{"a"}}
	zRunAfterX := &MyNode{name: "z", runAfter: []string{"x"}}
	aRunAfterZ := &MyNode{name: "a", runAfter: []string{"z"}}
	aSelfLinkAfterA := &MyNode{name: "a", runAfter: []string{"a"}}
	aInvalidTaskAfterNone := &MyNode{name: "a", runAfter: []string{"none"}}

	tcs := []struct {
		caseName string
		nodes    []NamedNode
	}{{
		caseName: "self-link-after",
		nodes:    []NamedNode{aSelfLinkAfterA},
	}, {
		caseName: "cycle-runAfter",
		nodes:    []NamedNode{xRunAfterA, zRunAfterX, aRunAfterZ},
	}, {
		caseName: "duplicate-node",
		nodes:    []NamedNode{a, a},
	}, {
		caseName: "invalid-task-name-after",
		nodes:    []NamedNode{aInvalidTaskAfterNone},
	},
	}
	for _, tc := range tcs {
		t.Run(tc.caseName, func(t *testing.T) {
			if _, err := NewDAG(tc.nodes); err == nil {
				t.Errorf("expected to see an error for invalid DAG %v but had none", tc.nodes)
			} else {
				t.Log(err)
			}
		})
	}
}

func assertNode(t *testing.T, n Node, expectedName string, expectedPrev []string, expectedNext []string) {
	assertSameNodeName(t, n, expectedName)
	assertSameNodeDepends(t, n.PrevNodes(), expectedPrev)
	assertSameNodeDepends(t, n.NextNodes(), expectedNext)
}

func assertSameNodeName(t *testing.T, n Node, expectedName string) {
	assert.Equal(t, expectedName, n.NodeName())
}

func assertSameNodeDepends(t *testing.T, depends []Node, expected []string) {
	assert.Equal(t, len(expected), len(depends))
	prevNames := make(map[string]struct{})
	for _, prev := range depends {
		prevNames[prev.NodeName()] = struct{}{}
	}
	for _, ep := range expected {
		_, ok := prevNames[ep]
		assert.True(t, ok)
	}
}
