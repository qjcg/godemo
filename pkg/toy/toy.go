// Package toy contains functions accounting for toy sales.
package toy

// Toy represents a children's toy for sale.
type Toy struct {
	Name   string
	Weight int
	onHand int
	sold   int
}

// New is a factory function for producing new Toy values.
func New(name string, weight int) Toy {
	return Toy{
		Name:   name,
		Weight: weight,
	}
}

// OnHand returns the number of units of this toy on hand.
func (t *Toy) OnHand() int {
	return t.onHand
}

// UpdateOnHand returns the number of units of this toy on hand after adding "count" units.
func (t *Toy) UpdateOnHand(count int) int {
	t.onHand += count
	return t.onHand
}

// Sold returns the number of units sold.
func (t *Toy) Sold() int {
	return t.sold
}

// UpdateSold increases the number of units sold, and returns the final value.
func (t *Toy) UpdateSold(n int) int {
	t.sold += n
	return t.sold
}
