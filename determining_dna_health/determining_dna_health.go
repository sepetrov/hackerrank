// Package determining_dna_health contains the solution for HackerRank challenge
// https://www.hackerrank.com/challenges/determining-dna-health/problem.
package determining_dna_health

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"sync"
)

// Run reads the input from reader and returns the respective total healths of
// the unhealthiest (minimum total health) and healthiest (maximum total health)
// strands of DNA.
func Run(reader *bufio.Reader) (min int, max int) {
	min, max = -1, -1

	// Read the total number of genes and healths.
	n, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	// Read the genes.
	genes := strings.Split(readLine(reader), " ")
	if len(genes) != int(n) {
		panic(fmt.Sprintf("got %d genes, want %d", len(genes), n))
	}

	// Read the healths.
	healths := func() []int {
		var out []int
		for _, s := range strings.Split(readLine(reader), " ") {
			i, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				panic(err)
			}
			out = append(out, int(i))
		}
		return out
	}()
	if len(healths) != int(n) {
		panic(fmt.Sprintf("got %d healths, want %d", len(healths), n))
	}

	// Read the total number of strands.
	n, err = strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	ch := make(chan int)
	wg := new(sync.WaitGroup)
	wg.Add(int(n))
	go func() {
		for i := 0; i < int(n); i++ {
			strand := strings.Split(readLine(reader), " ")

			start, err := strconv.ParseInt(strand[0], 10, 64)
			checkError(err)

			end, err := strconv.ParseInt(strand[1], 10, 64)
			checkError(err)

			go func(genes []string, healths []int, dna string) {
				defer wg.Done()
				ch <- calculateHealth(genes, healths, dna)
			}(genes[start:end+1], healths[start:end+1], strand[2])
		}

		wg.Wait()
		close(ch)
	}()

	for health := range ch {
		if min == -1 || max == -1 {
			min, max = health, health
			continue
		}

		if health > max {
			max = health
		}
		if health < min {
			min = health
		}
	}

	return
}

// node is Aho–Corasick trie node.
type node struct {
	char        rune // The char this node corresponds to.
	parent      *node
	isEndOfWord bool  // True if the node represents the end of a word.
	matches     []int // The indexes of the found words. len(matches) > 0 only if isEndOfWord == true.
	children    map[rune]*node

	strictSuffix *node

	dictSuffix *node
}

func buildTrie(dictionary []string) *node {
	root := &node{
		char:     '_',
		matches:  []int{},
		children: make(map[rune]*node),
	}

	// Contains pointers to the nodes for each level of the tree.
	levelNodes := make(
		[]map[*node]struct{},
		func() int {
			var levels int
			for _, w := range dictionary {
				if l := len(w); l > levels {
					levels = l
				}
			}
			return levels
		}(),
	)

	// Build the tree.
	for wordIndex, word := range dictionary {
		n := root
		for charIndex, char := range word {
			c, ok := n.children[char]
			if !ok {
				c = &node{
					char:        char,
					isEndOfWord: charIndex == len(word)-1,
					parent:      n,
					matches:     []int{},
					children:    make(map[rune]*node),
				}
				n.children[char] = c
			}
			if charIndex == len(word)-1 {
				c.isEndOfWord = true
				c.matches = append(c.matches, wordIndex)
			}

			if levelNodes[charIndex] == nil {
				levelNodes[charIndex] = make(map[*node]struct{})
			}
			levelNodes[charIndex][c] = struct{}{}

			// c becomes n, so that on the next iteration the node for char is a child of c.
			n = c
		}
	}

	// Update the strict suffixes.
	for _, nodes := range levelNodes {
		for n := range nodes {
			if n.parent == root {
				n.strictSuffix = root
				continue
			}
			for p := n.parent.strictSuffix; p != nil; p = p.strictSuffix {
				if c, ok := p.children[n.char]; ok {
					n.strictSuffix = c
				}
			}
		}
	}

	// Update dictionary suffixes.
	for _, nodes := range levelNodes {
		for n := range nodes {
			for p := n.strictSuffix; p != nil; p = p.strictSuffix {
				if p.isEndOfWord {
					n.dictSuffix = p
					break
				}
			}
		}
	}

	// Update matches of the nodes where node.isEndOfWord = true.
	for _, nodes := range levelNodes {
		for n := range nodes {
			if n.isEndOfWord && n.dictSuffix != nil {
				n.matches = append(
					n.matches,
					n.dictSuffix.matches...,
				)
			}
		}
	}

	return root
}

func calculateHealth(genes []string, healths []int, dna string) int {
	var score int
	var matches []int // Contains indexes of genes found in dna.

	root := buildTrie(genes)

	n := root

NextChar:
	for _, char := range dna {
		if c, ok := n.children[char]; ok {
			matches = append(matches, c.matches...)
			n = c
			continue
		}

		for n.strictSuffix != nil {
			n = n.strictSuffix
			if c, ok := n.children[char]; ok {
				matches = append(matches, c.matches...)
				n = c
				continue NextChar
			}
		}
	}

	for _, i := range matches {
		score += healths[i]
	}

	return score
}

func readLine(reader *bufio.Reader) string {
	s, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	return strings.TrimRight(s, "\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
