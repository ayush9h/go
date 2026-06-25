package main

import "fmt"

type OrderStatus string

const (
	//variable type = value
	Received  OrderStatus = "received"
	Confirmed OrderStatus = "confirmed"
	Prepared  OrderStatus = "prepared"
	Delieverd OrderStatus = "delievered"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("changing order status", status)
}

func main() {
	changeOrderStatus(Prepared)
}
