package mapusage

import (
	"fmt"
	"sort"
)

func UsingMapOfSlices() {
	fmt.Println("Usage : Map of Slices.")

	type Person struct {
		Name string
		Likes []string
	}
	var people []*Person
	people = append(people, &Person{"jay",
		[]string{"apple","banana","cheese"},
	})
	people = append(people, &Person{"kim",
		[]string{"apple","photo","cheese"},
	})
	people = append(people, &Person{"park",
		[]string{"apple","banana","music"},
	})
	likes := make(map[string][]*Person)
	for _, p := range people {
		for _, l := range p.Likes {
			likes[l] = append(likes[l], p)
		}
	}
	// To print a list who likes cheese :
	for _, p := range likes["cheese"] {
		fmt.Println(p.Name, "likes cheese.")
	}
	// To print the number of people who like apple :
	fmt.Println(len(likes["apple"]), "people like apple.")
}

func UsingLikeSet(){
	fmt.Println("Usage : Using Like Set.")

	type Node struct {
		Next  *Node
		Value interface{}
	}
	first := &Node{nil,0}
	second := &Node{
		&Node{
			&Node{
				&Node{
					&Node{
						first,0}, 4}, 3}, 2}, 1}
	first.Next = second

	visited := make(map[*Node]bool)
	for n := first; n != nil; n = n.Next {
		if visited[n] {
			fmt.Println("cycle detected")
			break
		}
		visited[n] = true
		fmt.Println(n.Value)
	}

}

func IterationOrder() {
	m := make(map[int]string)
	m[2] = "2"
	m[1] = "1"
	m[3] = "3"
	keys := make([]int,0)
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Println("key: ", k, "Value :", m[k])
	}
}
