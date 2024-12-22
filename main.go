package main

import (
	"cmp"
	"fmt"
	"math/rand"
	"slices"
)

type Member struct {
	Name   string
	Number int
}

func divideInto4and2(members []Member) ([]Member, []Member) {
	if len(members) != 6 {
		panic("members must be 6")
	}
	return members[:2], members[2:]
}
func divideInto4and1(members []Member) ([]Member, []Member) {
	if len(members) != 6 {
		panic("members must be 6")
	}
	return members[:1], members[1:]
}

func main() {
	m := make([]Member, 0, len(members))
	for _, v := range members {
		m = append(m, Member{Name: v, Number: rand.Intn(100)})
	}
	slices.SortFunc(m, func(i, j Member) int {
		return cmp.Compare(i.Number, j.Number)
	})

	var alpha, beta []Member
	if len(members) == 6 {
		alpha, beta = divideInto4and2(m)
	} else if len(members) == 5 {
		alpha, beta = divideInto4and1(m)
	} else {
		panic("members must be 5 or 6")
	}

	fmt.Println("-----alpha team-----")
	for _, v := range alpha {
		println(v.Name)
	}
	fmt.Println("-----beta team-----")
	for _, v := range beta {
		println(v.Name)
	}

}
