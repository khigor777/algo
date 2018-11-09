package algo

import (
	"testing"
)

func TestList(t *testing.T) {
	n := new(Node)

	n.Append(&List{
		data: 1,
	})
	n.Append(&List{
		data: 1,
	})

	n.Append(&List{
		data: 1,
	})

	n.Append(&List{
		data: 1,
	})

	n.Append(&List{
		data: 2,
	})
	n.Append(&List{
		data: 2,
	})

	n.RemoveDuplicateCounted(3)
	if n.length != 3 {
		t.Error("dosen't remove", n.length)
	}

	//if you want to see a result
	//currentPost := n.start
	//for currentPost.data != 0 {
	//	fmt.Println(currentPost.data)
	//	if currentPost.next == nil {
	//		break
	//	}
	//	currentPost = currentPost.next
	//}
}
