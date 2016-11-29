package main

import "fmt"


type tree struct {
	value int
	left, right *tree
}

func Sort(values []int) []int {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
	return values
}

func appendValues(values []int, t *tree) []int  {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	fmt.Println(values)
	return values
}

func add(t *tree, value int) *tree  {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}


func main()  {
	values := [] int {2, 3,11,9,2,1,13,6,7,5,44,34,22}
	//values =
	//fmt.Println(values[:0])
	fmt.Println(cap(values), len(values))
	fmt.Println(Sort(values[:]))
}

