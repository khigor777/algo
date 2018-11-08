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

func (f *Node) Append(list *List) {
	if f.length == 0 {
		f.start = list
	} else {
		currentPost := f.start
		for currentPost.next != nil {
			currentPost = currentPost.next
		}
		currentPost.next = list
	}
	f.length++
}

func (f *Node) Remove(data int64) {
	if f.length == 0 {
		panic(errors.New("node is empty"))
	}

	var previousPost *List
	currentPost := f.start

	for currentPost.data != data {
		if currentPost.next == nil {
			panic(errors.New("no such list found"))
		}

		previousPost = currentPost
		currentPost = currentPost.next
	}
	previousPost.next = currentPost.next

	f.length--
}

func (f *Node) RemoveDuplicate() {
	var previousPost *List
	currentPost := f.start
	hash := make(map[int64]int)

	for currentPost != nil {
		if currentPost == nil {
			break
		}
		if val, ok := hash[currentPost.data]; ok && val >= 2 {
			previousPost.next = currentPost.next
			f.length--
		} else {
			hash[currentPost.data]++
			previousPost = currentPost
		}
		currentPost = currentPost.next
	}
}
