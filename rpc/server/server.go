package main

import (
	"fmt"
)

import(
	"log"
	"errors"
	"net"
	"net/rpc"
	"net/http"
	"github.com/maxzerbini/packagemain/rpc/model"
)


type OrderServer struct {
	endpoint string
	referenceId int
}

// Register a new order.
func (srv *OrderServer) RegisterOrder(order *model.Order, reply *model.OrderReference) (err error) {
	defer func() {
		// Executes normally even if there is a panic
		if e:= recover(); e != nil {
			log.Println("Run time panic: %v", e)
			err = errors.New("Runtime error.")
		}
	}()
	// simulate order management creating a order reference id
	// (in a true case the order will be saved in a database)
	log.Printf("Order received %v", order)
	reply.Id = srv.referenceId
	for _,item := range order.Items {
		reply.Total += item.Cost * float64(item.Quantity)
	}
	reply.IdCustomer = order.IdCustomer
	return err
}

// Register a new order.
func (srv *OrderServer) InsertProduct(product *model.Product, reply *float64) (err error) {
	defer func() {
		// Executes normally even if there is a panic
		if e:= recover(); e != nil {
			log.Println("Run time panic: %v", e)
			err = errors.New("Runtime error.")
		}
	}()
	log.Printf("Product received %v", product)
	*reply = product.Cost * float64(product.Quantity)
	return nil
}
// Start listening commands.
func (srv *OrderServer)Do(){
	rpc.Register(srv)
	rpc.HandleHTTP()
	listener, e := net.Listen("tcp", srv.endpoint)
	if e != nil {
		log.Fatal("Starting RPC-server -listen error:", e)
	}
	log.Println("Server started ...")
	http.Serve(listener, nil)
}

func NewOrderServer(endpoint string) *OrderServer{
	s := &OrderServer{endpoint:endpoint, referenceId:1}
	return s
}

func main(){
	endpoint := "localhost:9000"
	s := NewOrderServer(endpoint)
	go s.Do()
	_,_ = fmt.Scanln()
	
}