package elrond

import "github.com/pkg/errors"

type Node struct {
	url  string
	Name string
}

func NewNode(url string) (*Node, error) {
	node := Node{
		url: url,
	}

	status, err := node.Status()

	if err != nil {
		return nil, errors.Wrapf(err, "get status of node '%s'", url)
	}

	name, ok := status.Details["erd_node_display_name"]

	if !ok {
		return nil, errors.New("erd_node_display_name not available")
	}

	node.Name, ok = name.(string)

	if !ok {
		return nil, errors.New("erd_node_display_name is not a string")
	}

	return &node, nil
}
