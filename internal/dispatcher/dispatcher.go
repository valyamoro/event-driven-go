package dispatcher

import (
	"fmt"
	"sync"

	"event-driven-go/internal/events"
)

type EventDispatcher struct {
	handlers map[events.EventType][]events.EventHandler
	eventCh  chan events.Event
	stats    struct {
		eventsDispatched 	int
		eventsProcessed 	int
		mu 					sync.RWMutex
	}
}

func NewEventDispatcher(bufferSize int) *EventDispatcher {
	return &EventDispatcher {
		handlers: make(map[events.EventType][]events.EventHandler),
		eventCh: make(chan events.Event, bufferSize),
	}
}

func (d *EventDispatcher) RegisterHandler(eventType events.EventType, handler events.EventHandler) {
	d.handlers[eventType] = append(d.handlers[eventType], handler)
}

func (d *EventDispatcher) Dispatch(event events.Event) {
	d.stats.mu.Lock()
	d.stats.eventsDispatched++
	d.stats.mu.Unlock()

	fmt.Printf("\nДиспетчр: ПОЛУЧЕНО событие [%s] от %s\n", event.Type, event.Source)

	go func() {
		d.eventCh <- event
	}()
}

func (d *EventDispatcher) Start() {
	fmt.Println("Диспетчер: ЗАПУСК системы обработки событий...")

	go func() {
		for event := range d.eventCh {
			d.processEvent(event)
		}
		fmt.Println("Обработка событий завершена")
	}()
}

func (d *EventDispatcher) processEvent(event events.Event) {
	fmt.Printf("Диспетчер: ищу обработчики для %s...\n", event.Type)

	if handlers, exists := d.handlers[event.Type]; exists {
		fmt.Printf("Найдено обработчиков: %d\n", len(handlers))

		var wg sync.WaitGroup
		for _, handler := range handlers {
			wg.Add(1)

			go func(h events.EventHandler, e events.Event) {
				defer wg.Done()

				if err := h.Handle(e); err != nil {
					fmt.Printf("Ошибка в обработчике %s: %v\n", h.Name(), err)
				}

				d.stats.mu.Lock()
				d.stats.eventsProcessed++
				d.stats.mu.Unlock()
			}(handler, event)
		}

		wg.Wait()
	} else {
		fmt.Printf("Внимание: нет обработчиков для соыбтия %s\n", event.Type)
	}

	fmt.Printf("Диспетчер: событие %s полностью обработано\n", event.ID)
}

func (d *EventDispatcher) Stop() {
	close(d.eventCh)
	fmt.Println("\nДиспетчер: остановка...")

	d.stats.mu.RLock()
	fmt.Printf("Статистика диспетчера:\n")
	fmt.Printf("Отправлено событий: %d\n", d.stats.eventsDispatched)
	fmt.Printf("Обработано событий: %d\n", d.stats.eventsProcessed)
	d.stats.mu.RUnlock()
}
