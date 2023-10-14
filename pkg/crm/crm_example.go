package crm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dedicio/sisgares-integrations-service/internal/dto"
)

type Config struct {
	Username string
	Token    string
}

type CRMExample struct {
	Config Config
}

type OrderItemsPayload struct {
	Product  string  `json:"product"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

type OrderPayload struct {
	ID            string               `json:"id"`
	CompanyID     string               `json:"company_id"`
	Items         []*OrderItemsPayload `json:"items"`
	Discount      float64              `json:"discount"`
	PaymentMethod string               `json:"payment_method"`
}

func NewCRMExample(config Config) *CRMExample {
	return &CRMExample{
		Config: config,
	}
}

func (crm *CRMExample) IntegrateOrderCreated(order dto.OrderMessagingDTO) error {
	var items []*OrderItemsPayload
	for _, item := range order.Items {
		items = append(items, &OrderItemsPayload{
			Product:  item.ProductId,
			Quantity: item.Quantity,
			Price:    item.Price,
		})
	}

	payload := &OrderPayload{
		ID:            order.ID,
		CompanyID:     order.CompanyID,
		Items:         items,
		Discount:      order.Discount,
		PaymentMethod: order.PaymentMethod,
	}

	var body bytes.Buffer
	err := json.NewEncoder(&body).Encode(payload)
	if err != nil {
		return err
	}
	log.Println("Integrating order created to CRM", body.String())

	client := &http.Client{}
	req, err := http.NewRequest(
		"POST",
		"https://jsonplaceholder.typicode.com/posts",
		&body,
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", crm.Config.Token))

	_, err = client.Do(req)
	if err != nil {
		return err
	}
	log.Println("Order created integrated to CRM")

	return nil
}
