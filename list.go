package algo

import (
	"errors"
)

type List struct {
	data int64
	next *List
}

type Node struct {
	length int
	start  *List
}

//add to the end
func (n *Node) Append(list *List) {
	if n.length == 0 {
		n.start = list
	} else {
		current := n.start
		for current.next != nil {
			current = current.next
		}
		current.next = list
	}
	n.length++
}

//remove one
func (n *Node) Remove(data int64) {
	if n.length == 0 {
		panic(errors.New("node is empty"))
	}

	var previous *List
	current := n.start

	for current.data != data {
		if current.next == nil {
			panic(errors.New("no such list found"))
		}

		previous = current
		current = current.next
	}
	previous.next = current.next

	n.length--
}

//remove all duplicates
func (n *Node) RemoveDuplicate() {
	var previous *List
	current := n.start
	hash := make(map[int64]int)

	for current != nil {
		if current == nil {
			break
		}
		if _, ok := hash[current.data]; ok {
			previous.next = current.next
			n.length--
		} else {
			hash[current.data]++
			previous = current
		}
		current = current.next
	}
}

//removes founded num count duplicates
func (n *Node) RemoveDuplicateCounted(num int) {
	for k := range n.getCountedDuplicates(num) {
		var previous *List
		current := n.start
		hash := make(map[int64]int)

		for current != nil {
			if current == nil {
				break
			}
			if current.data == k {
				if previous == nil {
					previous = current
					n.length++
					continue
				}
				previous.next = current.next
				n.length--
			} else {
				hash[current.data]++
				previous = current
			}
			current = current.next
		}
	}

}

func (n *Node) getCountedDuplicates(num int) map[int64]bool {
	hash := make(map[int64]int)
	res := make(map[int64]bool)
	current := n.start
	for current.data != 0 {
		hash[current.data]++
		if val, _ := hash[current.data]; val >= num {
			res[current.data] = true
		}
		if current.next == nil {
			break
		}
		current = current.next
	}
	return res
}
