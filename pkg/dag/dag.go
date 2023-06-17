package dag

import (
	"github.com/pkg/errors"
	"strings"
)

type DAG struct {
	Nodes                        map[string]*internalNode
	allowMarkArbitraryNodeAsDone bool
	allowNotCheckCycle           bool
}

type NamedNode interface {
	NodeName() string
	PrevNodeNames() []string
}

type Node interface {
	NamedNode
	PrevNodes() []Node
	NextNodes() []Node
	NextNodeNames() []string
}

type internalNode struct {
	name          string
	prevNodeNames []string
	prevNodes     []*internalNode
	nextNodes     []*internalNode
}

type OperationFunc func(dag *DAG)

func WithAllowMarkArbitraryNodeAsDone(allow bool) OperationFunc {
	return func(dag *DAG) {
		dag.allowMarkArbitraryNodeAsDone = allow
	}
}

func WithAllowNotCheckCycle(allow bool) OperationFunc {
	return func(dag *DAG) {
		dag.allowNotCheckCycle = allow
	}
}

func NewDAG(nodes []NamedNode, ops ...OperationFunc) (*DAG, error) {
	dag := &DAG{
		Nodes:                        make(map[string]*internalNode),
		allowNotCheckCycle:           false,
		allowMarkArbitraryNodeAsDone: false,
	}

	for _, op := range ops {
		op(dag)
	}

	// init DAG
	for _, node := range nodes {
		if err := dag.addNode(node); err != nil {
			return nil, errors.Errorf("failed add node %q to DAG %v", node.NodeName(), err)
		}
	}

	// link node in DAG
	for _, node := range dag.Nodes {
		for _, prevNode := range node.PrevNodeNames() {
			if err := dag.addLink(node, prevNode); err != nil {
				return nil, errors.Wrap(err, "add link failed")
			}
		}
	}

	return dag, nil
}

func (dag *DAG) addNode(n NamedNode) error {
	if _, ok := dag.Nodes[n.NodeName()]; ok {
		return errors.Errorf("duplicate node found: %s", n.NodeName())
	}

	dag.Nodes[n.NodeName()] = &internalNode{name: n.NodeName(), prevNodeNames: n.PrevNodeNames()}
	return nil
}

func (dag *DAG) addLink(n Node, prevNodeName string) error {
	prevNode, ok := dag.Nodes[prevNodeName]
	if !ok {
		return errors.Errorf("node %s depend on non-existent node %s", n.NodeName(), prevNodeName)
	}

	if err := dag.linkTwoNodes(prevNode, n); err != nil {
		return errors.Errorf("create link form %q to %q failed: %v", n.NodeName(), prevNodeName, err)
	}
	return nil
}

func (dag *DAG) linkTwoNodes(from, to Node) error {
	if !dag.allowNotCheckCycle {
		if err := validateNodes(from, to); err != nil {
			return err
		}
	}

	// link node
	to.(*internalNode).prevNodes = append(to.(*internalNode).prevNodes, from.(*internalNode))
	from.(*internalNode).nextNodes = append(from.(*internalNode).nextNodes, to.(*internalNode))
	return nil
}

func validateNodes(from, to Node) error {
	// self cycle check
	if from.NodeName() == to.NodeName() {
		return errors.Errorf("self cycle deteced, node %q depends on itself", from.NodeName())
	}

	// check cycle
	path := []string{to.NodeName(), from.NodeName()}
	if err := visit(to, to.PrevNodes(), path); err != nil {
		return errors.Wrap(err, "cycle detected")
	}
	return nil
}

func visit(startNode Node, prevNodes []Node, visitedPath []string) error {
	for _, n := range prevNodes {
		visitedPath = append(visitedPath, n.NodeName())
		if n.NodeName() == startNode.NodeName() {
			return errors.Errorf("%s", getVisitedPath(visitedPath))
		}

		if err := visit(startNode, n.PrevNodes(), visitedPath); err != nil {
			return err
		}
	}
	return nil
}

// generate visited path like a->b->c->a
func getVisitedPath(path []string) string {
	// reverse the path since we traversed the DAG using prev pointers.
	for i := len(path)/2 - 1; i >= 0; i-- {
		opp := len(path) - 1 - i
		path[i], path[opp] = path[opp], path[i]
	}
	return strings.Join(path, " -> ")
}

func (inter *internalNode) NodeName() string {
	return inter.name
}

func (inter *internalNode) PrevNodeNames() []string {
	return inter.prevNodeNames
}

func (inter *internalNode) PrevNodes() []Node {
	prevNodes := make([]Node, len(inter.prevNodes))
	for index, prevNode := range inter.prevNodes {
		prevNodes[index] = prevNode
	}
	return prevNodes
}

func (inter *internalNode) NextNodes() []Node {
	nextNodes := make([]Node, len(inter.nextNodes))
	for index, nextNode := range inter.nextNodes {
		nextNodes[index] = nextNode
	}
	return nextNodes
}

func (inter *internalNode) NextNodeNames() []string {
	names := make([]string, len(inter.nextNodes))
	for index, node := range inter.nextNodes {
		names[index] = node.name
	}
	return names
}
