package models

type Orders struct {
	OrderId  int    `json:"orderId:int"`
	Type     string `json:"orderType:string"`
	Name     string `json:"name:string"`
	Quantity int    `json:"quantity:int"`
}
