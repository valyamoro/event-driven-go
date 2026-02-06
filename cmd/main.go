package main

import (
	"time"

	"event-driven-go/internal/dispatcher"
	"event-driven-go/internal/events"
	"event-driven-go/internal/generator"
	"event-driven-go/internal/handlers"
)

func main() {
	dispatcher := dispatcher.NewEventDispatcher(50)

	userManager := handlers.NewUserManager("UserManager")
	orderProcessor := handlers.NewOrderProcessor("OrderProcessor")
	paymentGateway := handlers.NewPaymentGateway("PaymentGateway")
	inventoryManager := handlers.NewInventoryManager("InventoryManager")
	analyticsService := handlers.NewAnalyticsService("Analytics")
	notificationService := handlers.NewNotificationService("NotificationService")

	dispatcher.RegisterHandler(events.UserCreated, userManager)
	dispatcher.RegisterHandler(events.UserUpdated, userManager)
	dispatcher.RegisterHandler(events.UserDeleted, userManager)

	dispatcher.RegisterHandler(events.OrderPlaced, orderProcessor)
	dispatcher.RegisterHandler(events.OrderCancelled, orderProcessor)
	dispatcher.RegisterHandler(events.OrderShipped, orderProcessor)

	dispatcher.RegisterHandler(events.PaymentSuccess, paymentGateway)
	dispatcher.RegisterHandler(events.PaymentFailed, paymentGateway)
	dispatcher.RegisterHandler(events.PaymentRefunded, paymentGateway)

	dispatcher.RegisterHandler(events.InventoryLow, inventoryManager)

	dispatcher.RegisterHandler(events.ReviewAdded, notificationService)
	dispatcher.RegisterHandler(events.PromoCodeUsed, notificationService)
	dispatcher.RegisterHandler(events.OrderShipped, notificationService)
	dispatcher.RegisterHandler(events.PaymentSuccess, notificationService)

	allEventTypes := []events.EventType{
		events.UserCreated, events.UserUpdated, events.UserDeleted,
		events.OrderPlaced, events.OrderCancelled, events.OrderShipped,
		events.PaymentSuccess, events.PaymentFailed, events.PaymentRefunded,
		events.InventoryLow, events.ReviewAdded, events.PromoCodeUsed,
		events.SystemAlert,
	}
	for _, eventType := range allEventTypes {
		dispatcher.RegisterHandler(eventType, analyticsService)
	}

	dispatcher.Start()

	time.Sleep(100 * time.Millisecond)

	generator.GenerateEvents(dispatcher)

	time.Sleep(5 * time.Second)

	dispatcher.Stop()
}
