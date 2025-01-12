package main

import "fmt"

/*
Реализовать паттерн «адаптер» на любом примере.
*/

type Client struct {
	val     int
	adapter ClientInterface
}

func (c *Client) Connect() {
	c.adapter.Send(c.val)
}

type ClientInterface interface {
	Send(val int)
	Convert(val int) CustomInt
}

type Adapter struct {
	service Service
}

func (a *Adapter) Send(val int) {
	a.service.getMessage(a.Convert(val))
}

func (a *Adapter) Convert(val int) CustomInt {
	return CustomInt{value: val}
}

type CustomInt struct {
	value int
}

type Service struct {
	data CustomInt
}

func (s *Service) getMessage(d CustomInt) {
	s.data = d
	fmt.Println("Got it", d.value)
}

func main() {
	cl := Client{val: 10, adapter: &Adapter{service: Service{}}}
	cl.Connect()
}
