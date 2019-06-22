package disjointSet

import (
	"testing"
)

func TestElement_Join(t *testing.T) {
	set := DisjointSet{}
	// set initialization
	set.Initialize(10)
	// testing join
	for i := 1; i < 9; i *= 2 {
		set.Data[i].Join(set.Data[i+1])
	}
	// checking the roots of the elements
	countRoots := make(map[*Element]int, 2)
	for i := 0; i < 10; i++ {
		countRoots[set.Data[i].up]++
	}
	exp:=4
	got:=len(countRoots)
	if got!=exp {
		t.Errorf("wrong number of roots exp=%d, got=%d",exp,got)
	}
}


func TestElement_SameClass(t *testing.T) {
	set := DisjointSet{}
	// set initialization
	for i := 0; i < 10; i++ {
		set.Data = append(set.Data, &Element{0, i, nil})
	}
	// testing join
	for i := 1; i < 9; i *= 2 {
		set.Data[i].Join(set.Data[i+1])
	}

	if set.Data[0].SameClass(set.Data[1]) ||
	set.Data[1].SameClass(set.Data[4]) {
		t.Error("error with the same class")
	}


}
