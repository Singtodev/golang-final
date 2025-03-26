package dtos

type CustomerAddItemDto struct {
	CartName  string `json:"cart_name"`
	ProductID uint   `json:"product_id"`
	Quantity  int    `json:"quantity"`
}
