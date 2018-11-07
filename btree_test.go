package algo

import "testing"

func TestBtree(t *testing.T) {
	b := &Btree{}
	b.Insert(1)
	b.Insert(14)
	b.Insert(6)
	b.Insert(2)
	h := Height(b)
	if h != 4 {
		t.Errorf("wrong height, h == %d", h)
	}
}
