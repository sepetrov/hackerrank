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

	n, err := strconv.ParseInt(readLine(reader), 10, 64)
	genes := newGHSlice(
		strings.Split(readLine(reader), " "),
		strings.Split(readLine(reader), " "),
	)
	if int64(len(genes)) != n {
		panic(fmt.Sprintf("got %d genes and health values, want %d", len(genes), n))
	}

	n, err = strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	ch := make(chan int)
	wg := new(sync.WaitGroup)
	wg.Add(int(n))
	go func() {
		for i := int64(0); i < n; i++ {
			// first last d
			line := strings.Split(readLine(reader), " ")

			first, err := strconv.ParseInt(line[0], 10, 64)
			checkError(err)

			last, err := strconv.ParseInt(line[1], 10, 64)
			checkError(err)

			go func(genes []gh, strand string) {
				health := calculateHealth(genes, strand)
				// fmt.Printf("\n\ngenes: %#v\nstrand: %s\nhealth: %d\n", genes, strand, health)
				ch <- health
				wg.Done()
			}(genes[first:last+1], line[2])
		}

		wg.Wait()
		close(ch)
	}()

	for health := range ch {
		if min == -1 || max == -1 {
			min = health
			max = health
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

func newGHSlice(genes, health []string) []gh {
	var s []gh
	if len(genes) != len(health) {
		panic(fmt.Sprintf("got len(genes) = %d, len(health) = %d, want len(genes) == len(health)", len(genes), len(health)))
	}
	for i, n := 0, len(genes); i < n; i++ {
		healthItem, err := strconv.ParseInt(health[i], 10, 64)
		checkError(err)

		s = append(s, gh{
			gene:   genes[i],
			health: int(healthItem),
		})
	}
	return s
}

type gh struct {
	gene   string
	health int
}

func calculateHealth(genes []gh, strand string) int {
	var score int
	for _, g := range genes {
		score += countOccurrences(g.gene, strand) * g.health
	}
	return score
}

func countOccurrences(gene, strand string) int {
	n := 0
	for {
		i := strings.Index(strand, gene)
		if i == -1 {
			return n
		}
		n++
		strand = strand[i+1:]
	}
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
