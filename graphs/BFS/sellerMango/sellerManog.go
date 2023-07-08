package sellerMango

import "algorithms/queue"

type Seller struct {
	Sell      bool
	Neighbors []*Seller
	Level     int
}

func FindSellerMango(seller *Seller) int {
	visited := make(map[*Seller]bool)
	queueSellers := queue.CreateQueue()
	queueSellers.Enqueue(seller)
	for queueSellers.NotEmpty() {
		seller := queueSellers.Dequeue().(*Seller)
		if seller.Sell {
			return seller.Level
		}
		for _, neighbor := range seller.Neighbors {
			if !visited[neighbor] {
				visited[neighbor] = true
				queueSellers.Enqueue(neighbor)
			}
		}
	}
	return -1
}
