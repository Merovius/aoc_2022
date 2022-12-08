package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	m, err := ReadInput(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d trees are visible\n", CountVisible(m))
}

func ReadInput(r io.Reader) (*Map, error) {
	var out [][]int
	s := bufio.NewScanner(r)
	for s.Scan() {
		var row []int
		l := s.Text()
		for i := 0; i < len(l); i++ {
			if l[i] < '0' || l[i] > '9' {
				return nil, fmt.Errorf("invalid line %q", l)
			}
			row = append(row, int(l[i]-'0'))
		}
		out = append(out, row)
	}
	if len(out) == 0 {
		return nil, errors.New("empty input")
	}
	m := &Map{
		Width:  len(out[0]),
		Height: len(out),
		Cells:  out,
	}
	for _, r := range m.Cells[1:] {
		if len(r) != m.Width {
			return nil, errors.New("number of columns is inconsistent")
		}
	}
	return m, nil
}

type Map struct {
	Width  int
	Height int
	Cells  [][]int
}

func (m *Map)

func CountVisible(m *Map[int]) int {
	var n int
	for i := 0; i < m.Height; i++ {
		for j := 0; j < m.Width; j++ {
			visible := true
			for ii := 0; ii < i; ii++ {
				if m.Cells[ii][j] >= m.Cells[i][j] {
					visible = false
					break
				}
			}
			if visible {
				n++
				continue
			}
			visible = true
			for ii := i + 1; ii < m.Height; ii++ {
				if m.Cells[ii][j] >= m.Cells[i][j] {
					visible = false
					break
				}
			}
			if visible {
				n++
				continue
			}
			visible = true
			for jj := 0; jj < j; jj++ {
				if m.Cells[i][jj] >= m.Cells[i][j] {
					visible = false
					break
				}
			}
			if visible {
				n++
				continue
			}
			visible = true
			for jj := j + 1; jj < m.Width; jj++ {
				if m.Cells[i][jj] >= m.Cells[i][j] {
					visible = false
					break
				}
			}
			if visible {
				n++
				continue
			}
		}
	}
	return n
}

func ViewingDistances(m *Map[int])
