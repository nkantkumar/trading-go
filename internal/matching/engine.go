package matching

import (
	"fmt"
	"sync"
	"trading-go/internal/order"
)

type MatchingEngine struct {
	mu         sync.Mutex
	buyOrders  []order.Order
	sellOrders []order.Order
}

func (e *MatchingEngine) AddOrder(o order.Order) {
	e.mu.Lock()
	defer e.mu.Unlock()

	if o.Type == order.Buy {
		e.buyOrders = append(e.buyOrders, o)
	} else {
		e.sellOrders = append(e.sellOrders, o)
	}

	e.matchOrders()
}

func (e *MatchingEngine) matchOrders() {
	// Very basic matching logic (price-time priority)
	for len(e.buyOrders) > 0 && len(e.sellOrders) > 0 {
		buy := &e.buyOrders[0]
		sell := &e.sellOrders[0]

		if buy.Price >= sell.Price {
			fmt.Printf("Matched Order: Buy %s - Sell %s at %.2f\n", buy.ID, sell.ID, sell.Price)

			e.buyOrders = e.buyOrders[1:]
			e.sellOrders = e.sellOrders[1:]
		} else {
			break
		}
	}
}
