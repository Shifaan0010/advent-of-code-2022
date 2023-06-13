package trie

type Trie[T comparable] struct {
	IsLeaf   bool
	Children map[T]*Trie[T]
}

func (trie *Trie[T]) Add(word []T) {
	for _, letter := range word {
		if trie.Children == nil {
			trie.Children = map[T]*Trie[T]{}
		}

		if trie.Children[letter] == nil {
			trie.Children[letter] = &Trie[T]{}
		}

		trie = trie.Children[letter]
	}

	trie.IsLeaf = true
}
