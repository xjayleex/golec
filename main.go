package main
import(
	pqc "./priorityqueue"
	"fmt"
)
func main(){
	pq := pqc.New()
	pq.Push("jay",1)
	pq.Push("foo",3)
	pq.Push("bar",2)
	pq.Push("kim",4)
	pq.Push("park",5)
	for !pq.IsEmpty() {
		fmt.Println(pq.Pop().(string))
	}

}

