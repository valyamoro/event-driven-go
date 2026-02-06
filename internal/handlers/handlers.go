package handlers

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"event-driven-go/internal/events"
)

type UserManager struct {
	name string
	mu sync.Mutex
	userCount int
}

func NewUserManager(name string) *UserManager {
	return &UserManager{
		name: name,
		userCount: 0,
	}
}

func (h *UserManager) Name() string {
	return h.name
}

func (h *UserManager) Handle(event events.Event) error {
	fmt.Printf("[%s] Получено событие: %s\n", h.Name(), event.Type)

	switch event.Type {
	case events.UserCreated:
		if data, ok := event.Payload.(events.UserData); ok {
			h.mu.Lock()
			h.userCount++
			h.mu.Unlock()
			
			fmt.Printf("Создан пользователь: %s (%s)\n", data.Username, data.Email)
			fmt.Printf("Всего пользователей в системе: %d\n", h.userCount)

			time.Sleep(time.Duration(rand.Intn(100)+50) * time.Millisecond)
		}
	case events.UserUpdated:
		if data, ok := event.Payload.(events.UserData); ok {
			fmt.Printf("Обновлён пользователь: %s\n", data.UserID)
			
			time.Sleep(time.Duration(rand.Intn(80)+30) * time.Millisecond)
		}
	case events.UserDeleted:
		if data, ok := event.Payload.(events.UserData); ok {
			h.mu.Lock()
			h.userCount--
			h.mu.Unlock()

			fmt.Printf("Удален пользователь: %s\n", data.UserID)
			fmt.Printf("Осталось пользователей: %d\n", h.userCount)
			
			time.Sleep(time.Duration(rand.Intn(120)+60) * time.Millisecond)
		}
	}

	return nil
}

type OrderProcessor struct {
	name string
}

func NewOrderProcessor(name string) *OrderProcessor {
	return &OrderProcessor{name: name}
}

func (h *OrderProcessor) Name() string {
	return h.name
}

func (h *OrderProcessor) Handle(event events.Event) error {
	fmt.Printf("[%s] Получено событие: %s\n", h.Name(), event.Type)

	if data, ok := event.Payload.(events.OrderData); ok {
		switch event.Type {
		case events.OrderPlaced:
			fmt.Printf("Новый заказ #%s на сумму $%.2f\n", data.OrderID, data.TotalAmount)
			fmt.Printf("Доставка: %s\n", data.ShippingAddress)
			fmt.Printf("Товаров %d\n", len(data.Items))
		case events.OrderCancelled:
			fmt.Printf("Заказ #%s отменён\n", data.OrderID)
			if data.Reason != "" {
				fmt.Printf("Причина: %s\n", data.Reason)
			}
		case events.OrderShipped:
			fmt.Printf("Заказ #%s отправлен!\n", data.OrderID)
		}

		time.Sleep(time.Duration(rand.Intn(200)+100) * time.Millisecond)
	}

	return nil
}

type PaymentGateway struct {
	name string
}

func NewPaymentGateway(name string) *PaymentGateway {
	return &PaymentGateway{name: name}
}

func (h *PaymentGateway) Name() string {
	return h.name
}

func (h *PaymentGateway) Handle(event events.Event) error {
	fmt.Printf("[%s] Получено событие: %s\n", h.Name(), event.Type)

	if data, ok := event.Payload.(events.PaymentData); ok {
		switch data.Status {
		case "SUCCESS":
			fmt.Printf("Платеж #%s успешен: $%.2f %s\n", data.PaymentID, data.Amount, data.Currency)
			fmt.Printf("Метод оплаты: %s\n", data.Method)
		case "FAILED":
			fmt.Printf("Платеж #%s не прошел\n", data.PaymentID)
			if data.FailedReason != "" {
				fmt.Printf("Причина: %s\n", data.FailedReason)
			}
		case "REFUNDED":
			fmt.Printf("Возврат по платежу #%s: $%.2f\n", data.PaymentID, data.RefundAmount)
		}

		time.Sleep(time.Duration(rand.Intn(150)+50) * time.Millisecond)
	}

	return nil
}

type InventoryManager struct {
	name string
	lowStockAlerts int
}

func NewInventoryManager(name string) *InventoryManager {
	return &InventoryManager{
		name: name,
		lowStockAlerts: 0,
	}
}

func (h *InventoryManager) Name() string {
	return h.name
}

func (h *InventoryManager) Handle(event events.Event) error {
	fmt.Printf("[%s] Получено событие: %s\n", h.Name(), event.Type)

	if data, ok := event.Payload.(events.InventoryData); ok {
		h.lowStockAlerts++

		urgencyIcon := "⚠️"
		switch data.Urgency {
		case "high":
			urgencyIcon = "🚨"
		case "critical":
			urgencyIcon = "🔥"
		case "medium":
			urgencyIcon = "⚠️"
		case "low":
			urgencyIcon = "ℹ️"
		}

		fmt.Printf("%s Низкий запас товара: %s\n", urgencyIcon, data.ProductName)
		fmt.Printf("Осталось: %d из минимальных %d\n", data.CurrentQty, data.MinQty)
		fmt.Printf("Склад: %s\n", data.Warehouse)
		fmt.Printf("Всего алертов за сессию: %d\n", h.lowStockAlerts)

		if data.Urgency == "critical" {
			fmt.Printf("Автоматически создаем заказ поставщику!\n")
		}

		time.Sleep(time.Duration(rand.Intn(100)+30) * time.Millisecond)
	}

	return nil
}

type AnalyticsService struct {
	name 			string
	eventsProcessed map[events.EventType]int
	mu 				sync.RWMutex
}

func NewAnalyticsService(name string) *AnalyticsService {
	return &AnalyticsService {
		name: 			 name,
		eventsProcessed: make(map[events.EventType]int),
	}
}

func (h *AnalyticsService) Name() string {
	return h.name
}

func (h *AnalyticsService) Handle(event events.Event) error {
	h.mu.Lock()
	h.eventsProcessed[event.Type]++
	h.mu.Unlock()

	h.mu.RLock()
	total := 0
	for _, count := range h.eventsProcessed {
		total += count
	}
	h.mu.RUnlock()

	fmt.Printf("[%s] Аналитика: обработано событий: %d\n", h.Name(), total)

	time.Sleep(time.Duration(rand.Intn(50)+10) * time.Millisecond)
	return nil
}

type NotificationService struct {
	name string
}

func NewNotificationService(name string) *NotificationService {
	return &NotificationService{name: name}
}

func (h *NotificationService) Name() string {
	return h.name
}

func (h *NotificationService) Handle(event events.Event) error {
	fmt.Printf("[%s] Отправка уведомления для: %s\n", h.Name(), event.Type)

	switch event.Type {
	case events.OrderShipped:
		fmt.Printf("SMS: Ваш заказ отправлен!\n")
		fmt.Printf("Email: Отправлено письмо с трек-номером\n")

	case events.PaymentSuccess:
		fmt.Printf("Email: Чек отправлен на email\n")

	case events.ReviewAdded:
		fmt.Printf("Push: Спасибо за ваш отзыв!\n")

	case events.PromoCodeUsed:
		fmt.Printf("Email: Вы успешно использовали промокод!\n")
	}

	time.Sleep(time.Duration(rand.Intn(80)+20) * time.Millisecond)
	return nil
}
