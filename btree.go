package algo

import (
	"fmt"
)

type Btree struct {
	Data  int
	Left  *Btree
	Right *Btree
}

func (b *Btree) Insert(data int) error {
	if b == nil {
		return fmt.Errorf("error, tree is nil")
	}
	switch {
	case data == b.Data:
		return nil
	case data < b.Data:
		if b.Left == nil {
			b.Left = &Btree{Data: data}
			return nil
		}
		return b.Left.Insert(data)
	case data > b.Data:
		if b.Right == nil {
			b.Right = &Btree{Data: data}
			return nil
		}
		return b.Right.Insert(data)
	}
	return nil
}

func Height(b *Btree) int64 {

	if b == nil {
		return -1
	}

	left := Height(b.Left)
	right := Height(b.Right)

	if left > right {
		return left + 1
	}
	return right + 1

}
