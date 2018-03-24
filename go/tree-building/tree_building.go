package tree

import (
	"fmt"
	"sort"
)

// Record represents an input to the tree.
type Record struct {
	ID, Parent int
}

// Node represents a component of a tree.
type Node struct {
	ID       int
	Children []*Node
}

func reorder(records []Record) []Record {
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	return records
}

func validate(record Record, nodes []*Node) (err error) {
	switch {
	case record.ID == 0 && record.Parent != 0:
		return fmt.Errorf("root has a parent")
	case record.ID < record.Parent:
		return fmt.Errorf("indirect cycle detected")
	case record.ID != 0 && record.ID == record.Parent:
		return fmt.Errorf("cycle detected")
	case record.ID >= len(nodes):
		return fmt.Errorf("records not contiguous")
	case nodes[record.ID] != nil:
		return fmt.Errorf("duplicate record found")
	}

	return
}

// Build contructs a valid tree from a slice of records.
func Build(records []Record) (root *Node, err error) {
	if len(records) == 0 {
		return
	}

	nodes := make([]*Node, len(records))
	for index, record := range reorder(records) {
		if err := validate(record, nodes); err != nil {
			return nil, err
		}

		nodes[index] = &Node{ID: record.ID}
		if index > 0 {
			parent := nodes[record.Parent]
			parent.Children = append(parent.Children, nodes[index])
		}
	}
	root = nodes[0]

	return
}
