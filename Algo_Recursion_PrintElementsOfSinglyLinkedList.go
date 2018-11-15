package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type SinglyLinkedListNode struct {
	data int32
    next *SinglyLinkedListNode
}

type SinglyLinkedList struct {
	head *SinglyLinkedListNode
	tail *SinglyLinkedListNode
}

func (SinglyLinkedList *SinglyLinkedList) insertNodeIntoSinglyLinkedList (nodeData int32) {
	node = &SinglyLinkedListNode {
		next: nil
		data: nodeData
	}
	if SinglyLinkedList.head == nil {
		SinglyLinkedList.head = node
	} else {
		SinglyLinkedList.tail.next = node
	}
	SinglyLinkedList.tail = node
}

func printLinkedList(head *SinglyLinkedListNode) {
    if head != nil {
        fmt.Printf("%v\n", head.data)
        printLinkedList(head.next)
    }
}

//TODO
funcMain(){
	fmt.Println("TODO")
}