package algo

import (
	"testing"
)

func TestList(t *testing.T) {
	n := new(Nodes)

	n.Append(&Item{
		data: 1,
	})
<<<<<<< HEAD
	n.Append(&Item{
		data: 2,
	})

	n.Append(&Item{
		data: 3,
	})

	n.Append(&Item{
		data: 2,
=======
	n.Append(&List{
		data: 1,
	})

	n.Append(&List{
		data: 1,
	})

	n.Append(&List{
		data: 1,
>>>>>>> 69f03f274b861b5b28c01f1868f410f822bb346d
	})

	n.Append(&Item{
		data: 2,
	})
	n.Append(&Item{
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
