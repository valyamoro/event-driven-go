package generator

import (
	"fmt"
	"time"

	"event-driven-go/internal/dispatcher"
	"event-driven-go/internal/events"
)

func GenerateEvents(dispatcher *dispatcher.EventDispatcher) {
	eventList := []events.Event{
		{
			ID:        "evt_user_001",
			Type:      events.UserCreated,
			Timestamp: time.Now(),
			Source:    "auth-service",
			Payload: events.UserData{
				UserID:    "user_1001",
				Username:  "alexey_ivanov",
				Email:     "alexey@example.com",
				Action:    "create",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		},

		{
			ID:        "evt_user_002",
			Type:      events.UserUpdated,
			Timestamp: time.Now().Add(5 * time.Second),
			Source:    "user-service",
			Payload: events.UserData{
				UserID:    "user_1001",
				Username:  "alexey_ivanov_pro",
				Email:     "alexey.new@example.com",
				Action:    "update",
				Timestamp: time.Now().Add(5 * time.Second).Format(time.RFC3339),
			},
		},

		{
			ID:        "evt_order_001",
			Type:      events.OrderPlaced,
			Timestamp: time.Now().Add(10 * time.Second),
			Source:    "order-service",
			Payload: events.OrderData{
				OrderID:     "order_5001",
				UserID:      "user_1001",
				TotalAmount: 249.99,
				Status:      "placed",
				ShippingAddress: "Москва, ул. Тверская, д. 10",
				Items: []events.OrderItem{
					{ProductID: "prod_101", Name: "Книга 'Go на практике'", Quantity: 1, Price: 199.99},
					{ProductID: "prod_102", Name: "Футболка с логотипом Go", Quantity: 2, Price: 25.00},
				},
			},
		},

		{
			ID:        "evt_payment_001",
			Type:      events.PaymentSuccess,
			Timestamp: time.Now().Add(12 * time.Second),
			Source:    "payment-service",
			Payload: events.PaymentData{
				PaymentID: "pay_9001",
				OrderID:   "order_5001",
				Amount:    249.99,
				Currency:  "USD",
				Status:    "SUCCESS",
				Method:    "credit_card",
			},
		},

		{
			ID:        "evt_inventory_001",
			Type:      events.InventoryLow,
			Timestamp: time.Now().Add(15 * time.Second),
			Source:    "inventory-service",
			Payload: events.InventoryData{
				ProductID:   "prod_101",
				ProductName: "Книга 'Go на практике'",
				CurrentQty:  3,
				MinQty:      10,
				Warehouse:   "Московский склад",
				Urgency:     "high",
			},
		},

		{
			ID:        "evt_order_002",
			Type:      events.OrderShipped,
			Timestamp: time.Now().Add(20 * time.Second),
			Source:    "shipping-service",
			Payload: events.OrderData{
				OrderID:     "order_5001",
				UserID:      "user_1001",
				TotalAmount: 249.99,
				Status:      "shipped",
				ShippingAddress: "Москва, ул. Тверская, д. 10",
				Items: []events.OrderItem{
					{ProductID: "prod_101", Name: "Книга 'Go на практике'", Quantity: 1, Price: 199.99},
					{ProductID: "prod_102", Name: "Футболка с логотипом Go", Quantity: 2, Price: 25.00},
				},
			},
		},

		{
			ID:        "evt_review_001",
			Type:      events.ReviewAdded,
			Timestamp: time.Now().Add(25 * time.Second),
			Source:    "review-service",
			Payload: events.ReviewData{
				ReviewID:    "rev_001",
				ProductID:   "prod_101",
				UserID:      "user_1001",
				Rating:      5,
				Title:       "Отличная книга!",
				Comment:     "Помогла освоить Go за 2 недели",
				VerifiedPurchase: true,
			},
		},

		{
			ID:        "evt_promo_001",
			Type:      events.PromoCodeUsed,
			Timestamp: time.Now().Add(28 * time.Second),
			Source:    "order-service",
			Payload: events.PromoCodeData{
				Code:        "WELCOME10",
				UserID:      "user_1001",
				OrderID:     "order_5001",
				Discount:    24.99,
				DiscountPct: 10,
				MinAmount:   100,
			},
		},

		{
			ID:        "evt_alert_001",
			Type:      events.SystemAlert,
			Timestamp: time.Now().Add(30 * time.Second),
			Source:    "monitoring",
			Payload: events.AlertData{
				Severity: "WARNING",
				Service:  "payment-service",
				Message:  "Высокое время ответа API платежей (>500ms)",
				Code:     "PERF_SLOW",
				Action:   "Проверить логи, увеличить ресурсы",
			},
		},

		{
			ID:        "evt_payment_002",
			Type:      events.PaymentRefunded,
			Timestamp: time.Now().Add(35 * time.Second),
			Source:    "payment-service",
			Payload: events.PaymentData{
				PaymentID:   "pay_9002",
				OrderID:     "order_5000",
				Amount:      99.99,
				Currency:    "USD",
				Status:      "REFUNDED",
				Method:      "credit_card",
				RefundAmount: 99.99,
			},
		},
	}

	for i, event := range eventList {
		fmt.Printf("\n[Шаг %d/%d] ", i+1, len(eventList))
		dispatcher.Dispatch(event)
		time.Sleep(800 * time.Millisecond)
	}
}
