package entity

import "time"

type StatusType string

const (
	STATUS_PENDING        StatusType = "pending"
	STATUS_SETTLEMENT     StatusType = "settlement"
	STATUS_CENCEL         StatusType = "cencel"
	STATUS_DENY           StatusType = "deny"
	STATUS_REFUND         StatusType = "refund"
	STATUS_PARTIAL_REFUND StatusType = "partial_refund"
	STATUS_CHARGEBACK     StatusType = "chargeback"
	STATUS_EXPIRE         StatusType = "expire"
	STATUS_FAILURE        StatusType = "failure"
)

type Order struct {
	Id        string     `json:"id"`
	UserId    string     `json:"user_id"`
	ProductId int        `json:"product_id"`
	Price     int        `json:"price"`
	Quantity  int        `json:"quantity"`
	Status    StatusType `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
