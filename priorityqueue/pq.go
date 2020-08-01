package priorityqueue

import "container/heap"

type Item struct {
	Value    interface{}
	Priority int
	Index    int
}

type itemHeap []*Item

func (ih *itemHeap) Len() int{
	return len(*ih)
}

func (ih *itemHeap) Less(i, j int) bool {
	return (*ih)[i].Priority < (*ih)[j].Priority
}

func (ih *itemHeap) Front() (interface{}){
	if len(*ih) == 0{
		return nil
	}
	return (*ih)[0]
}

func (ih *itemHeap) Swap(i, j int) {
	(*ih)[i], (*ih)[j] = (*ih)[j], (*ih)[i]
	(*ih)[i].Index = i
	(*ih)[j].Index = j
}

func (ih *itemHeap) Push(x interface{}){
	item := x.(*Item)
	item.Index = len(*ih)
	*ih = append(*ih,item)
}

func (ih *itemHeap) Pop()(x interface{}){
	old := *ih
	item := old[len(old)-1]
	*ih = old[0 : len(old)-1]
	return item
}

/* it's minheap by default. If u want use as a maxheap, set the priorty value as minus value.*/
type PriorityQueue struct{
	itemHeap *itemHeap

}

func New() PriorityQueue{
	return PriorityQueue{
		itemHeap: &itemHeap{},
	}
}

func (p *PriorityQueue) Len() int {
	return p.itemHeap.Len()
}
/* it's minheap by default. If u want use as a maxheap, set the priorty value as minus value.*/
func (p *PriorityQueue) Push(v interface{}, priority int){

	newItem := &Item{
		Value:    v,
		Priority: priority,
	}
	heap.Push(p.itemHeap, newItem)
}
/* it's minheap by default. If u want use as a maxheap, set the priorty value as minus value.*/
func (p *PriorityQueue) Pop() (interface{}){
	if len(*p.itemHeap) == 0 {
		return nil
	}
	item := heap.Pop(p.itemHeap).(*Item)
	return item.Value
}


func (p *PriorityQueue) Front() (interface{}){
	if len(*p.itemHeap) == 0 {
		return nil
	}
	return p.itemHeap.Front().(*Item)
}

func (p *PriorityQueue) IsEmpty() bool{
	if len(*p.itemHeap) == 0 {
		return true
	} else {
		return false
	}
}