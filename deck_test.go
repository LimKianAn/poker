package deck

import (
	"testing"
)

func TestDeck_Draw(t *testing.T) {
	m := map[int]bool{}
	for i := 0; i < 52; i++ {
		m[i] = true
	}

	d := NewDeck()
	for i := 0; i < 52; i++ {
		card := d.Draw()
		n := int(card.Suite*13) + int(card.Rank)
		delete(m, n)
	}

	if len(m) != 0 || len(*d) != 0 {
		t.Errorf("after having drawn 52 cards, got %v, want an empty deck", *d)

	}
}

func TestDeck_shuffle2(t *testing.T) {
	m := map[int]bool{}
	for i := 0; i < 52; i++ {
		m[i] = true
	}

	d := NewDeck()
	d.shuffle2()
	if len(*d) != 52 {
		t.Errorf("after a shuffle, %v cards reamin", len(*d))

	}

	for i := 0; i < 52; i++ {
		n := (*d)[i]
		delete(m, n)
	}

	if len(m) != 0 {
		keys := []int{}
		for k := range m {
			keys = append(keys, k)
		}
		t.Errorf("%v are missing in the shuffled deck", keys)
	}
}
