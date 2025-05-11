package common

import "fmt"

type Trie[T comparable] struct {
	root *node[T]
}

func NewTrie[T comparable](value T) *Trie[T] {
	return &Trie[T]{
		root: &node[T]{value: value, children: map[T]*node[T]{}},
	}
}

func (t *Trie[T]) Add(word []T) {
	if len(word) == 0 {
		// Ignore empty strings
		return
	}

	t.root.add(word)
}

func (t *Trie[T]) Search(prefix []T) ([][]T, error) {
	if len(prefix) == 0 {
		return [][]T{}, nil
	}

	node, err := t.root.getNode(prefix)

	if err != nil {
		return nil, err
	}

	if node == nil {
		return [][]T{}, nil
	}

	searchResults := [][]T{}

	collectedValues := node.collect()

	for _, values := range collectedValues {
		searchResults = append(searchResults, append(prefix[:len(prefix)-1], values...))
	}

	return searchResults, nil
}

type node[T comparable] struct {
	value    T
	children map[T]*node[T]
	isEnd    bool
}

func (n *node[T]) add(word []T) {
	if len(word) == 0 {
		return
	}

	first, word := word[0], word[1:]

	isEnd := len(word) == 0

	_, ok := n.children[first]

	if !ok {
		n.children[first] = &node[T]{
			value:    first,
			children: map[T]*node[T]{},
		}
	}

	if isEnd {
		n.children[first].isEnd = isEnd
	}

	n.children[first].add(word)
}

func (n *node[T]) getNode(prefix []T) (*node[T], error) {
	if len(prefix) == 0 {
		return nil, fmt.Errorf("cannot find any node with prefix len = 0")
	}

	curNode := n
	for _, ch := range prefix {

		nextNode, isEdgePresent := curNode.children[ch]

		if !isEdgePresent {
			return nil, nil
		}

		curNode = nextNode
	}

	return curNode, nil
}

func (n *node[T]) collect() [][]T {

	collectedValues := [][]T{}

	if n.isEnd {
		collectedValues = append(collectedValues, []T{n.value})
	}

	if len(n.children) == 0 && !n.isEnd {
		collectedValues = append(collectedValues, []T{n.value})

	}

	for _, children := range n.children {
		childrenCollectedValues := children.collect()

		for _, values := range childrenCollectedValues {
			collectedValues = append(collectedValues, append([]T{n.value}, values...))
		}
	}

	return collectedValues
}
