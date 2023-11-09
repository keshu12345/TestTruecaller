// create an autocomplete
// injest a list of words
// build the trie
// then take an input as prefix
// then print all the possible words

package main

import "fmt"

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func Newtrie() *Trie {
	return &Trie{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
		},
	}

}

func (t *Trie) Insert(word string) {
	node := t.root

	for _, char := range word {
		if _, exits := node.children[char]; !exits {
			node.children[char] = &TrieNode{children: make(map[rune]*TrieNode)}

		}
		node = node.children[char]

	}
	node.isEnd = true
}

func (t *Trie) Search(prefix string) []string {
	node := t.root
	for _, char := range prefix {

		if _, exits := node.children[char]; !exits {
			return []string{}
		}
		node = node.children[char]
	}

	return t.CollectionWords(node, prefix)

}

func (t *Trie) CollectionWords(node *TrieNode, prefix string) []string {
	words := []string{}
	if node.isEnd {
		words = append(words, prefix)
	}

	for char, child := range node.children {
		words = append(words, t.CollectionWords(child, prefix+string(char))...)
	}
	return words

}

func main() {
	trie := Newtrie()

	words := []string{"apple", "app", "banana", "bat", "ball", "cat", "application", "apprant"}

	for _, word := range words {
		trie.Insert(word)
	}

	prefix := "app"
	suggestion := trie.Search(prefix)
	fmt.Println(suggestion)

}
