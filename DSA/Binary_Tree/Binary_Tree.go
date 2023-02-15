package main

type node struct {
	data  int
	left  *node
	right *node
}

func new(val int) *node {
	return &node{
		data:  val,
		left:  nil,
		right: nil,
	}

}

func (t *node) insert(val int) *node {
	if val < t.data {
		if t.left != nil {
			t.left.insert(val)
		} else {
			t.left = new(val)
		}
	} else {
		if t.right != nil {
			t.right.insert(val)
		} else {
			t.right = new(val)
		}
	}
	return t
}

func main() {

}
