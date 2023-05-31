package dag

import "github.com/pkg/errors"

func (dag *DAG) GetSchedulableNodes(finishes ...string) (map[string]Node, error) {
	roots := dag.getRoots()
	finishNodesMap, err := dag.toMap(finishes...)
	if err != nil {
		return nil, err
	}

	result := make(map[string]Node)
	visited := make(map[string]Node)

	for _, root := range roots {
		schedulable := findSchedulableNode(root, visited, finishNodesMap)
		for _, n := range schedulable {
			result[n] = dag.Nodes[n]
		}
	}

	return result, nil
}

func (dag *DAG) GetSchedulableNodeNames(finishes ...string) ([]string, error) {
	schedulableNodes, err := dag.GetSchedulableNodes(finishes...)
	if err != nil {
		return nil, err
	}

	result := make([]string, 0)
	for name := range schedulableNodes {
		result = append(result, name)
	}

	return result, nil
}

func (dag *DAG) getRoots() []Node {
	roots := make([]Node, 0)
	for _, node := range dag.Nodes {
		if len(node.PrevNodeNames()) == 0 {
			roots = append(roots, node)
		}
	}

	return roots
}

func (dag *DAG) toMap(nodeNames ...string) (map[string]Node, error) {
	nodesMap := make(map[string]Node)
	for _, nodeName := range nodeNames {
		if node, ok := dag.Nodes[nodeName]; ok {
			nodesMap[nodeName] = node
		} else {
			return nil, errors.Errorf("node %s not found", nodeName)
		}
	}

	return nodesMap, nil
}

func findSchedulableNode(n Node, visited, doneNodes map[string]Node) []string {
	// if the node n is visited, returned directly
	if _, ok := visited[n.Name()]; ok {
		return nil
	}

	visited[n.Name()] = n

	// the node n is finished, traverse the next node of node n, and returned
	if _, ok := doneNodes[n.Name()]; ok {
		result := make([]string, 0)
		for _, nextNode := range n.NextNodes() {
			result = append(result, findSchedulableNode(nextNode, visited, doneNodes)...)
		}
		return result
	}

	// the node n is not finished, check it is schedulable
	if isSchedulable(n, doneNodes) {
		return []string{n.Name()}
	}

	// the node is not finished and can not be scheduled
	return nil
}

func isSchedulable(n Node, doneNodes map[string]Node) bool {
	// the root node, can be scheduled
	if len(n.PrevNodes()) == 0 {
		return true
	}

	// if this node all previous node is finished, it can be scheduled
	var collected []string
	for _, prev := range n.PrevNodes() {
		if _, ok := doneNodes[prev.Name()]; ok {
			collected = append(collected, prev.Name())
		}
	}
	return len(collected) == len(n.PrevNodeNames())
}

func checkNotVisited(doneTaskMap, visited map[string]Node) []string {
	var notVisited []string
	for done := range doneTaskMap {
		if _, ok := visited[done]; !ok {
			notVisited = append(notVisited, done)
		}
	}
	return notVisited
}
