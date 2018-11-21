package linked_list

type Item struct {
	next *Item
	Value string
}


func reverseList(ItemList *Item) *Item {
	current := ItemList
	var top *Item = nil
	for {
		if current == nil {
			break
		}
		temp := current.next
		current.next = top
		top = current
		current = temp
	}
	return top
}