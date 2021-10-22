package main

import "fmt"

type LinkNode struct {
	Val  int
	Next *LinkNode
}

func main() {
	var items = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	link := buildLinkList(items)
	printLinkList(link)
	reversedlink := reverLinkList(link)
	printLinkList(reversedlink)
}

func reverLinkList(head *LinkNode) *LinkNode {
	p := head.Next
	if p == nil {
		return head
	}
	q := p.Next
	p.Next = nil
	for q != nil {
		r := q.Next
		q.Next = p
		p = q
		q = r
	}
	head.Next = p
	return head
}

func printLinkList(head *LinkNode) {
	node := head.Next
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

func buildLinkList(items []int) *LinkNode {
	hair := &LinkNode{}
	p := hair
	for _, v := range items {
		node := &LinkNode{Val: v}
		p.Next = node
		p = p.Next
	}
	p.Next = nil
	return hair
}
