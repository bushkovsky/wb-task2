package main

import "fmt"

/*
Реализовать паттерн «адаптер» на любом примере.
*/

/*
Решение:
Главная суть в том, что Client ничего не знает о структуре Srvice, информацию об этом хранит
лишь адаптер, интерфейс которого хранится в структуре Клиента, тем самым, мы можем писать
разные адаптеры, реализующие интерфейс ClientInterface для работы с разными сторонними сервисами
*/

// Клиент — это структура, которая содержит существующую бизнес-логику программы.
type Client struct {
	val     int
	adapter ClientInterface
}

func (c *Client) Connect() {
	c.adapter.Send(c.val)
}

// Клиентский интерфейс описывает протокол, через который клиент может работать с другими классами.
type ClientInterface interface {
	Send(val int)
	Convert(val int) CustomInt
}

// Адаптер — это структура, которая может одновременно работать и с клиентом, и с сервисом.
// Она реализует клиентский интерфейс и содержит ссылку на объект сервиса.
// Адаптер получает вызовы от клиента через методы клиентского интерфейса,
// а затем переводит их в вызовы методов обёрнутого объекта в правильном формате.
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

// Сервис — это какая-то полезная структура, обычно сторонняя.
// Клиент не может использовать эту структуру напрямую, так как сервис имеет непонятный ему интерфейс.
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
