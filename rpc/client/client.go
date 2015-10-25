package client

import(
	"log"
	"errors"
	"net/rpc"
	"github.com/maxzerbini/packagemain/rpc/model"
)

type OrderClient struct {
	endpoint string
	client *rpc.Client
}

func (nc *OrderClient) ConnectClient(){
	defer func() {
		// Println executes normally even if there is a panic
		if err := recover(); err != nil {
			log.Println("run time panic: %v", err)
		}
	}()
	client, err := rpc.DialHTTP("tcp", nc.endpoint)
	if err != nil {
		log.Printf("Error dialing server: %v \r\n", err)
	} else {
		log.Printf("Dialing server on %s done.\r\n",nc.endpoint)
		nc.client = client
	}
}

// Send an order
func (nc *OrderClient) SendOrder(order *model.Order) ( reference *model.OrderReference, err error) {
	defer func() {
		// executes normally even if there is a panic
		if e := recover(); e != nil {
			log.Printf("Error %v\r\n",e)
			reference = nil
			err = errors.New("Client runtime error.")
		}
	}()
	
	reference = new (model.OrderReference)
	err = nc.client.Call("OrderServer.RegisterOrder", order, reference)
	if err != nil {
		reference = nil
		log.Println("OrderServer.RegisterOrder error: ", err)
	}
	return reference, err
}

// Send new product
func (nc *OrderClient) SendSingleProduct(product *model.Product) ( err error) {
	defer func() {
		// executes normally even if there is a panic
		if e := recover(); e != nil {
			log.Printf("Error %v\r\n",e)
			err = errors.New("Client runtime error.")
		}
	}()
	
	var ref = new (float64)
	err = nc.client.Call("OrderServer.InsertProduct", product, ref)
	if err != nil {
		log.Println("OrderServer.InsertProduct error: ", err)
	}
	log.Printf("Cost of the product %f", *ref)
	return err
}

func NewOrderClient(endpoint string) *OrderClient{
	c := &OrderClient{endpoint:endpoint}
	c.ConnectClient()
	return c
}

