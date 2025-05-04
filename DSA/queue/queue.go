package queue

import "fmt"

type Queue []*Order

type Order struct {
	priority   int
	quantity   int
	product    string
	customName string
}

// contstructor

func (order *Order) New(priority, quantity int, product, customName string) {
	order.priority = priority
	order.quantity = quantity
	order.product = product
	order.customName = customName
}

// add

func (queue *Queue) Add(order *Order) {
	if len(*queue) == 0 {
		*queue = append(*queue, order)
	} else {
		var appended bool
		appended = false
		var i int
		var addedOrder *Order
		for i, addedOrder = range *queue {
			if order.priority > addedOrder.priority {
				*queue = append((*queue)[:i], append(Queue{order}, (*queue)[i:]...)...)
				appended = true
				break
			}
		}
		if !appended {
			*queue = append(*queue, order)
		}
	}
}

func main() {
	var queue Queue
	queue = make(Queue, 0)
	var order1 *Order = &Order{}
	var priority1 int = 2
	var quantity1 int = 20
	var product1 string = "Computer"
	var customName1 string = "Greg White"
	order1.New(priority1, quantity1, product1, customName1)

	var order2 *Order = &Order{}
	var priority2 int = 1
	var quantity2 int = 10
	var product2 string = "Monitor"
	var customName2 string = "John Smith"
	order2.New(priority2, quantity2, product2, customName2)
	queue.Add(order1)
	queue.Add(order2)
	var i int
	for i = 0; i < len(queue); i++ {
		fmt.Println(queue[i])
	}
}
