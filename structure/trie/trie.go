// Package trie provides Trie data structures in golang.
//
// Wikipedia: https://en.wikipedia.org/wiki/Trie
package trie

import "fmt"

// Node represents each node in Trie.
type Node struct {
	children map[rune]*Node // map children nodes
	isLeaf   bool           // current node value
}

// NewNode creates a new Trie node with initialized
// children map.
func NewNode() *Node {
	n := &Node{}
	n.children = make(map[rune]*Node)
	n.isLeaf = false
	return n
}

// Insert inserts words at a Trie node.
func (n *Node) Insert(s string) {
	curr := n
	for _, c := range s {
		next, ok := curr.children[c]
		if !ok {
			next = NewNode()
			curr.children[c] = next
		}
		curr = next
	}
	curr.isLeaf = true
}

// Find finds words at a Trie node.
func (n *Node) Find(s string) bool {
	curr := n
	for _, c := range s {
		next, ok := curr.children[c]
		if !ok {
			return false
		}
		curr = next
	}
	return true
}

// Logical removal of a word
func (n *Node) Remove(s string) error {
	curr := n
	for _, c := range s {
		if curr.children[c] == nil {
			return fmt.Errorf("word not found")
		}
		curr = curr.children[c]
	}
	if curr.isLeaf {
		curr.isLeaf = false
	}
	return nil
}
