package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	Children []*Node
	// feel free to add fields as you see fit
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	nodes := make(map[int]*Node)
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})
	if records[0].ID != records[0].Parent {
		return nil, errors.New("root node has parent")
	}
	nodes[0] = &Node{0, nil}

	lastNodeId := 0
	for _, record := range records[1:] {
		if lastNodeId+1 != record.ID {
			return nil, errors.New("non-continuous")
		}
		if record.ID == record.Parent {
			return nil, errors.New("cycle_directly")
		}
		if record.ID < record.Parent {
			return nil, errors.New("higher id parent of lower id")
		}

		_, nodeExists := nodes[record.ID]
		if nodeExists {
			return nil, errors.New("duplicate node")
		}
		parentNode, parentExists := nodes[record.Parent]
		newNode := Node{record.ID, nil}
		if parentExists {
			parentNode.Children = append(parentNode.Children, &newNode)
		}
		nodes[record.ID] = &newNode

		lastNodeId++
	}
	return nodes[0], nil
}
