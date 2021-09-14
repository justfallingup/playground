package main

import (
	"fmt"
	"strconv"
	"strings"
)

type PQueue struct {
	bheap []int
}

func New() *PQueue {
	pq := PQueue{make([]int, 0)}
	return &pq
}

func (q *PQueue) Peek() int {
	return q.bheap[0]
}

func (q PQueue) Poll() {
	lastIdx := len(q.bheap)-1
	q.bheap[0] = q.bheap[lastIdx]
	q.bheap = q.bheap[:lastIdx]
	q.sink(0)
}

func (q *PQueue) sink(idx int)  {
	for i := idx; 2*i+2 < len(q.bheap)-1; {
		leftCh := q.bheap[2*i+1]
		rightCh := q.bheap[2*i+2]
		if leftCh < q.bheap[i] || rightCh < q.bheap[i]{
			if leftCh <= rightCh {
				q.bheap[2*i+1] = q.bheap[i]
				q.bheap[i] = leftCh
				i = 2*i+1
			} else {
				q.bheap[2*i+2] = q.bheap[i]
				q.bheap[i] = rightCh
				i = 2*i+2
			}
		}
	}
}

func (q *PQueue) Contains(n int) (int, bool) {
	for idx, i := range q.bheap {
		if i == n {
			return idx, true
		}
	}
	return 0, false
}

func (q *PQueue) Remove(idx int)  {
	lastIdx := len(q.bheap)-1
	q.bheap[idx] = q.bheap[lastIdx]
	q.bheap = q.bheap[:lastIdx]
	pIdx := (idx - 1)/2
	if pIdx >= 0 {
		if q.bheap[idx] < q.bheap[pIdx] {
			q.swim(idx)
		}

	}
	q.sink(idx)
}

func (q *PQueue) Add(n int)  {
	q.bheap = append(q.bheap, n)
	q.swim(len(q.bheap)-1)
}

func (q *PQueue) swim(idx int) {
	pIdx := (idx - 1)/2
	for idx > 0 && q.bheap[pIdx] > q.bheap[idx] {
		el := q.bheap[idx]
		q.bheap[idx] = q.bheap[pIdx]
		q.bheap[pIdx] = el
		idx = pIdx
		pIdx = (pIdx - 1)/2
	}
}

func printTree(t []int)  {
	//height := int(math.Ceil(math.Log2(float64(len(t) + 1))))
	start, fin := 0, 1
	lvls := make([][]int, 0)
	dels := [][]int{{start, fin}}
	for fin < len(t) {
		lvls = append(lvls, t[start:fin])
		start = fin
		fin = start*2 + 1
		dels = append(dels, []int{start, fin})
	}
	lvls = append(lvls, t[start:])

	for i := 0; i < len(dels)/2; i++ {
		dels[i], dels[len(dels) - 1 - i] = dels[len(dels) - 1 - i], dels[i]
	}
	for i, lvl := range lvls {
		del1, del2 := strings.Repeat(" ", 2*dels[i][0]), strings.Repeat(" ", 2*dels[i][1])
		row := del1
		for _, j := range lvl {
			row += strconv.Itoa(j) + del2
		}
		fmt.Println(row)
	}
}

func (q PQueue) MinHeap() bool {
 return false
}

func less(i, j int) bool {
 return false
}


func main()  {
	q := []int{6, 5, 12, 8, 2, 14, 19, 13, 12, 7, 1, 13, 4, 0, 10}
	fmt.Println("input: ", q)
	pq := New()

	for _, i := range q{
		pq.Add(i)
	}
	fmt.Printf("output: %v \n", pq.bheap)
	printTree(pq.bheap)
	fmt.Println("peek", pq.Peek())
	pq.Poll()
	fmt.Println("peek", pq.Peek())
	printTree(pq.bheap)

	idx, ok := pq.Contains(7)
	if ok {
		fmt.Println("7 is at index ", idx)
	}
	pq.Remove(idx)
	printTree(pq.bheap)
}
