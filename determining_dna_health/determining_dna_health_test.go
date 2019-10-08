// Package determining_dna_health contains the solution for HackerRank problem
// https://www.hackerrank.com/challenges/determining-dna-health/problem.
package determining_dna_health_test

import (
	"bufio"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/sepetrov/hackerrank/determining_dna_health"
)

func TestRun(t *testing.T) {
	t.Parallel()
	type out struct {
		min, max int
	}
	parseOutput := func(filename string) out {
		body, err := ioutil.ReadFile(filename)
		if err != nil {
			panic(err)
		}
		mm := strings.Fields(string(body))
		min, err := strconv.Atoi(mm[0])
		if err != nil {
			panic(err)
		}
		max, err := strconv.Atoi(mm[1])
		if err != nil {
			panic(err)
		}
		return out{min, max}
	}
	tests := []struct {
		in  string
		out out
	}{
		{
			"testdata/input.txt",
			parseOutput("testdata/output.txt"),
		},
		{
			"testdata/input02.txt",
			parseOutput("testdata/output02.txt"),
		},
		{
			"testdata/input07.txt",
			parseOutput("testdata/output07.txt"),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()
			f, err := os.Open(tt.in)
			if err != nil {
				t.Fatal(err)
			}
			min, max := determining_dna_health.Run(bufio.NewReaderSize(f, 1024*1024))

			if min != tt.out.min {
				t.Errorf("Run() min = %d, want %d", min, tt.out.min)
			}
			if max != tt.out.max {
				t.Errorf("Run() max = %d, want %d", max, tt.out.max)
			}
		})
	}
}
