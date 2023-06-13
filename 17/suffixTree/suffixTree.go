package suffixTree

import (
	"fmt"
	"strings"

	"example.com/prob17/trie"
)

type SuffixTree[T comparable] struct {
	// IsLeaf   bool
	Substr   []T
	Children []SuffixTree[T]
}

func New[T comparable](s []T) SuffixTree[T] {
	suffixTrie := trie.Trie[T]{}

	for i := 0; i <= len(s); i += 1 {
		suffixTrie.Add(s[i:])
	}

	tree := SuffixTree[T]{}

	tree.addTrie(&suffixTrie)

	return tree
}

func (tree *SuffixTree[T]) addTrie(suffixTrie *trie.Trie[T]) {
	for !suffixTrie.IsLeaf && len(suffixTrie.Children) == 1 {
		// get the single value in map
		for letter, child := range suffixTrie.Children {
			tree.Substr = append(tree.Substr, letter)
			suffixTrie = child
			break
		}
	}

	// if suffixTrie.IsLeaf {
	// 	tree.Children = append(tree.Children, SuffixTree[T]{})
	// }

	for letter, child := range suffixTrie.Children {
		node := SuffixTree[T]{}

		node.Substr = append(node.Substr, letter)

		// if tree.Children == nil {
		// 	tree.Children = []SuffixTree[T]{}
		// }

		node.addTrie(child)

		tree.Children = append(tree.Children, node)
	}
}

func buildString[T comparable](builder *strings.Builder, tree *SuffixTree[T], depth int) {
	for i := 0; i < depth-1; i += 1 {
		builder.WriteString("|   ")
	}

	if depth > 0 {
		builder.WriteString("|---")
	}

	builder.WriteString(fmt.Sprintf("%s\n", tree.Substr))

	for _, child := range tree.Children {
		buildString(builder, &child, depth+1)
	}
}

func (tree SuffixTree[byte]) String() string {
	builder := strings.Builder{}

	buildString[byte](&builder, &tree, 0)

	return builder.String()
}
