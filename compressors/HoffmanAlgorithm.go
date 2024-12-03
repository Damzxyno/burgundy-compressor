package compressors

import (
	"container/heap"
)

type HuffmanNode struct {
	char        rune
	freq        int
	left, right *HuffmanNode
}

type PriorityQueue []*HuffmanNode

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].freq < pq[j].freq
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*HuffmanNode))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

func BuildHuffmanTree(freqMap map[rune]int) *HuffmanNode {
	pq := &PriorityQueue{}
	heap.Init(pq)

	for char, freq := range freqMap {
		heap.Push(pq, &HuffmanNode{char: char, freq: freq})
	}

	for pq.Len() > 1 {
		left := heap.Pop(pq).(*HuffmanNode)
		right := heap.Pop(pq).(*HuffmanNode)

		merged := &HuffmanNode{
			char:  0,
			freq:  left.freq + right.freq,
			left:  left,
			right: right,
		}

		heap.Push(pq, merged)
	}

	return heap.Pop(pq).(*HuffmanNode)
}

func GenerateCodes(root *HuffmanNode, code string, codes map[rune]string) {
	if root == nil {
		return
	}

	if root.char != 0 {
		codes[root.char] = code
	}

	GenerateCodes(root.left, code+"0", codes)
	GenerateCodes(root.right, code+"1", codes)
}

func HuffmanEncode(data string, codes map[rune]string) string {
	var encoded string
	for _, char := range data {
		encoded += codes[char]
	}
	return encoded
}

func HuffmanDecode(encoded string, root *HuffmanNode) string {
	var decoded string
	node := root
	for _, bit := range encoded {
		if bit == '0' {
			node = node.left
		} else {
			node = node.right
		}

		if node.left == nil && node.right == nil {
			decoded += string(node.char)
			node = root
		}
	}
	return decoded
}
