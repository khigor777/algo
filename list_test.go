package algo

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	n := new(Node)

	n.Append(&List{
		data: 1,
	})
	n.Append(&List{
		data: 2,
	})

	n.Append(&List{
		data: 3,
	})

	n.Append(&List{
		data: 2,
	})

	n.Append(&List{
		data: 2,
	})

	//O(n)
	n.RemoveDuplicate()

	if n.length != 3 {
		t.Error("dosen't remove duplicates")
	}

	currentPost := n.start
	for currentPost.data != 0 {
		fmt.Println(currentPost.data)
		if currentPost.next == nil {
			break
		}
		currentPost = currentPost.next
	}
}
