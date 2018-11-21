package algo

import (
	"errors"
)

type Item struct {
	data int64
	next *Item
}

type Nodes struct {
	length int
	start  *Item
}

//add to the end
func (n *Nodes) Append(list *Item) {
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
func (n *Nodes) Remove(data int64) {
	if n.length == 0 {
		panic(errors.New("node is empty"))
	}

	var previous *Item
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
func (n *Nodes) RemoveDuplicate() {
	var previous *Item
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
func (n *Nodes) RemoveDuplicateCounted(num int) {
	for k := range n.getCountedDuplicates(num) {
		var previous *Item
		current := n.start
		hash := make(map[int64]int)

		for current != nil {
			if current == nil {
				break
			}
			if current.data == k {
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

func (n *Nodes) getCountedDuplicates(num int)map[int64]bool {
	hash := make(map[int64]int)
	res := make(map[int64]bool)
	current := n.start
	for current.data != 0 {
		if current.next == nil {
			break
		}
		hash[current.data]++
		if val, _ := hash[current.data] ; val >= num {
			res[current.data] = true
		}
		current = current.next
	}
	return res
}
