package main

import (
	"fmt"
	"time"
)

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecond precision
}

func newOrder(id string, amount float32, status string) *order {

	//initial setup goes here.....
	myOrder := order{
		id:        id,
		amount:    amount,
		status:    status,
	}
	return &myOrder
}


//receiver  type
func (o *order) changeStatus(status string){

	o.status = status

}

func (o *order) getAmount() float32{
	return o.amount
}
func main() {


     myOrder := newOrder("1", 30.50, "received")
	 fmt.Println(myOrder)
	 fmt.Println(myOrder.getAmount())

   // if you dont set any field, default value in zero value
   // int => 0,, float => 0, string "", bool=> false
	// myOrder := order{
	// 	id:        "12345",
	// 	amount:    250.75,
	// 	status:    "received",
	// }
    // myOrder.changeStatus("confirmed")
	// fmt.Println(myOrder)
	// myOrder.createdAt = time.Now()
	// fmt.Println(myOrder.status)

	// myOrder2 := order{
	// 	id:        "2",
	// 	amount:    100,
	// 	status:    "delivered",
	// 	createdAt: time.Now(),
	// }
	// myOrder.status = "paid"
	// fmt.Println("Order struct", myOrder2)
	// fmt.Println("Order struct", myOrder)

}
