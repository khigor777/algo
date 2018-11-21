package algo

import (
	"testing"
)

func TestList(t *testing.T) {
	n := new(Nodes)

	n.Append(&Item{
		data: 1,
	})
	n.Append(&Item{
		data: 2,
	})

	n.Append(&Item{
		data: 3,
	})

	n.Append(&Item{
		data: 2,
	})

	n.Append(&Item{
		data: 2,
	})
	n.Append(&Item{
		data: 2,
	})


	n.RemoveDuplicateCounted(3)
	if n.length != 2 {
		t.Error("dosen't remove")
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
