package events

import (
	"time"
)

type EventType string

const (
	UserCreated 	EventType = "USER_CREATED"
	UserUpdated 	EventType = "USER_UPDATED"
	UserDeleted 	EventType = "USER_DELETED"
	OrderPlaced 	EventType = "ORDER_PLACED"
	OrderCancelled 	EventType = "ORDER_CANCELLED"
	OrderShipped 	EventType = "ORDER_SHIPPED"
	PaymentSuccess 	EventType = "PAYMENT_SUCCESS"
	PaymentFailed 	EventType = "PAYMENT_FAILED"
	PaymentRefunded EventType = "PAYMENT_REFUNDED"
	SystemAlert 	EventType = "SYSTEM_ALERT"
	InventoryLow 	EventType = "INVENTORY_LOW"
	ReviewAdded 	EventType = "REVIEW_ADDED"
	PromoCodeUsed 	EventType = "PROMO_CODE_USERD"
)

type Event struct {
	ID 		  string
	Type 	  EventType
	Timestamp time.Time
	Source	  string
	Payload   interface{}
}

type EventHandler interface {
	Handle(event Event) error
	Name() string
}
