package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l2 := &ListNode{
		Val:  2,
		Next: nil,
	}

	l1 := &ListNode{
		Val:  1,
		Next: l2,
	}

	removeNthFromEnd3(l1, 1)

}

func getLength(head *ListNode) (length int) {
	for ; head != nil; head = head.Next {
		length++
	}
	return
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := getLength(head)
	dummy := &ListNode{0, head}
	cur := dummy
	for i := 0; i < length-n; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dummy.Next
}

func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	nodes := []*ListNode{}
	dummy := &ListNode{0, head}
	for node := dummy; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	prev := nodes[len(nodes)-1-n]
	prev.Next = prev.Next.Next
	return dummy.Next
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	first, second := head, dummy
	for i := 0; i < n; i++ {
		first = first.Next
	}
	for ; first != nil; first = first.Next {
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
}

// 递归实现
func removeNthFromEnd3(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	dfs(dummy, &n)
	return dummy.Next
}

func dfs(head *ListNode, n *int) {
	if head.Next != nil {
		dfs(head.Next, n)
	}
	*n--
	if *n == -1 {
		head.Next = head.Next.Next
	}
}
