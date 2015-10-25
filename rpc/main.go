package main

import(
	"log"
	"github.com/maxzerbini/packagemain/rpc/model"
	"github.com/maxzerbini/packagemain/rpc/client"
)
// Test the RPC client
func main(){
	endpoint := "localhost:9000"
	c := client.NewOrderClient(endpoint)
	
	product := &model.Product{Id:100,Name:"test",Quantity:1,Cost:3.8}
	if err := c.SendSingleProduct(product); err == nil {
		log.Println("Product sent")
	} else {
		log.Printf("Error: %v\r\n", err)
	}
	
	o := &model.Order{IdCustomer:101,Items:make([]*model.Product,10)}
	for i := 0; i<10; i++ {
		o.Items[i] =&model.Product{Id:i,Name:"product",Quantity:1,Cost:5.3}
	}  
	
	if ref, err := c.SendOrder(o); err == nil {
		log.Printf("Order reference: %v\r\n", ref)
	} else {
		log.Printf("Error: %v", err)
	}
}