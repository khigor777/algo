package linked_list

import (
	"testing"
)

func TestItem_Append(t *testing.T) {
	item := Item{Value:"1"}
	item.next = &Item{Value: "2"}
	reverted := reverseList(&item)

	if reverted.Value != "2" {
		t.Error(item.Value)
	}
}
