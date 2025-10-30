package main

import "fmt"

type paymenter interface{
	pay(amount float32)
	refund(amount float32, account string)
}



type payment struct{
	gateway paymenter
}


// Open close principal 
func(p payment) makePayment(amount float32){

	// razorpayPaymentGw := razorpay{}
	// razorpayPaymentGw.pay(amount)
	p.gateway.pay(amount)
}

type razorpay struct {}

func (r razorpay) pay(amount float32){

	//logic to make payment
	fmt.Println("making payment using razorpay", amount)
}


// type stripe struct{}

// func(s stripe) pay(amount float32){
// 	fmt.Println("making payemnt using stripe", amount)
// }

type fakePayment struct{}

func (f fakePayment)pay(amount float32){
	fmt.Println("making payment using fake gateway for testing purpose")
}


type paypal struct{}

func (p paypal) pay(amount float32){
	fmt.Println("making payment using paypal", amount)
}

func (p paypal) refund(amount float32, account string){

}

func main(){
	// stripePaymentGw := stripe{}
    // razorpayPaymentGw := razorpay{}
	// fakeGw := fakePayment{}
	paypalGw := paypal{}
	newPayment := payment{
		gateway: paypalGw,
	}
	newPayment.makePayment(100)
}

