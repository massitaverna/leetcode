package addtwonumbers

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	cur := result
	prev := result

	var carry int

	for l1 != nil && l2 != nil {
		cur.Val = l1.Val + l2.Val + carry
		carry = 0
		if cur.Val >= 10 {
			cur.Val -= 10
			carry = 1
		}

		cur.Next = &ListNode{}
		prev = cur
		cur = cur.Next

		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		cur.Val = l1.Val + carry
		carry = 0
		if cur.Val >= 10 {
			cur.Val -= 10
			carry = 1
		}

		cur.Next = &ListNode{}
		prev = cur
		cur = cur.Next

		l1 = l1.Next
	}

	for l2 != nil {
		cur.Val = l2.Val + carry
		carry = 0
		if cur.Val >= 10 {
			cur.Val -= 10
			carry = 1
		}

		cur.Next = &ListNode{}
		prev = cur
		cur = cur.Next

		l2 = l2.Next
	}

	if carry == 1 {
		cur.Val = 1
	} else {
		prev.Next = nil
	}

	return result
}
