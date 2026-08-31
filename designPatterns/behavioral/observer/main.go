package main

import (
	"fmt"
)

/*
🧠 Observer Pattern
-------------------
One object (subject) maintains a list of others (observers) that need to be notified when it changes.
- Like YouTube: when a channel uploads a video, all subscribers (observers) get notified.
*/

// 👁️ Observer Interface
type Observer interface {
	Update(productName string)     // When product is available, call this
	GetID() string                 // Helps identify the observer
}

// 👤 Concrete Observer: Customer
type Customer struct {
	ID string
}

// 🧠 When product is available, this is called
func (c *Customer) Update(productName string) {
	fmt.Printf("📢 Notifying Customer [%s]: Product [%s] is now in stock!\n", c.ID, productName)
}

// 🧠 Get customer ID
func (c *Customer) GetID() string {
	return c.ID
}

// 📦 Subject: Product
type product struct {
	observerList []Observer
	name         string
	inStock      bool
}

// 🔧 Constructor for product
func NewProduct(_name string) *product {
	return &product{name: _name}
}

// 🧠 Notify all observers
func (p *product) NotifyALL() {
	for _, observer := range p.observerList {
		observer.Update(p.name)
	}
}

// 🧠 Change product availability and notify everyone
func (p *product) UpdateAvailability() {
	fmt.Printf("✅ Product [%s] is now in stock!\n\n", p.name)
	p.inStock = true
	p.NotifyALL()
}

// ➕ Add one observer
func (p *product) Register(o Observer) {
	p.observerList = append(p.observerList, o)
}

// ➕ Add a list of observers
func (p *product) RegisterList(olist []Observer) {
	for _, o := range olist {
		p.Register(o)
	}
}

// ❌ Remove an observer
func RemoveFromList(observerList []Observer, removeObserver Observer) []Observer {
	for index, observer := range observerList {
		if removeObserver.GetID() == observer.GetID() {
			return append(observerList[:index], observerList[index+1:]...)
		}
	}
	return observerList
}

// ❌ Unregister an observer
func (p *product) Unregister(o Observer) []Observer {
	p.observerList = RemoveFromList(p.observerList, o) // ✅ Fixed typo: 'oberverList' to 'observerList'
	return p.observerList
}

func main() {
	// 🛍️ Create a new product
	shoe := NewProduct("Nike Air Max")

	// 👥 Create some customers (observers)
	c1 := &Customer{ID: "cust_001"}
	c2 := &Customer{ID: "cust_002"}
	c3 := &Customer{ID: "cust_003"}

	// ➕ Register single customer
	shoe.Register(c1)

	// ➕ Register multiple customers
	shoe.RegisterList([]Observer{c2, c3})

	// ❌ Unregister one customer before notifying
	shoe.Unregister(c2)

	// ✅ Product is now available - notify all registered customers
	shoe.UpdateAvailability()
}
