package tree

import (
	"fmt"
	"sort"
)

// Record represents an input to the tree.
type Record struct {
	ID, Parent int
}

func (record *Record) isRoot() bool {
	return record.ID == 0
}
func (record *Record) isChild() bool {
	return !record.isRoot()
}
func (record *Record) isChildOf(node *Node) bool {
	return record.isChild() && record.Parent == node.ID
}

// Node represents a component of a tree.
type Node struct {
	ID       int
	Children []*Node
}

func (node *Node) build(records []Record, size int) {
	if size == len(records) {
		return
	}
	size++

	for _, record := range records {
		if record.isChildOf(node) {
			child := &Node{ID: record.ID}
			child.build(records, size)

			node.Children = append(node.Children, child)
		}
	}
}

func reorder(records []Record) {
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})
}

func validate(records []Record) (err error) {
	for index, record := range records {
		if index >= 1 && record.ID == records[index-1].ID {
			return fmt.Errorf("duplicate record found")
		}
		if record.ID >= len(records) {
			return fmt.Errorf("records not contiguous")
		}
		if record.ID < record.Parent {
			return fmt.Errorf("indirect cycle detected")
		}
		if record.isChild() && record.ID == record.Parent {
			return fmt.Errorf("cycle detected")
		}
		if record.isRoot() && record.Parent != 0 {
			return fmt.Errorf("root has a parent")
		}
	}
	return
}

// Build contructs a valid tree from a slice of records.
func Build(records []Record) (root *Node, err error) {
	if len(records) == 0 {
		return
	}

	err = validate(records)
	if err != nil {
		return nil, err
	}

	reorder(records)

	root = &Node{}
	root.build(records, 1)

	return
}
