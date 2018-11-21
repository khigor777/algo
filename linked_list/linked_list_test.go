package linked_list

import (
	"testing"
)

func TestItem_Append(t *testing.T) {
	reverted := reverseList(nil)
	if reverted != nil{
		t.Error("must be nill")
	}
	item := Item{Value:"1"}
	reverted = reverseList(&item)

	if reverted.Value != "1"{
		t.Error("must be 1")
		}

	item.next = &Item{Value: "2"}
	reverted = reverseList(&item)

	if reverted.Value != "2" {
		t.Error("must be 2")
	}
}


