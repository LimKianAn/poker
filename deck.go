package deck

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	Club Suite = iota
	Diamond
	Spade
	Heart
)

const (
	Ace Rank = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

type Suite int

type Rank int

type Deck []int

type Card struct {
	Suite
	Rank
}

func NewDeck() *Deck {
	nums := make([]int, 52)
	for i := range nums {
		nums[i] = i
	}
	d := Deck(nums)
	return &d
}

func (d *Deck) Draw() *Card {
	i := rand.Intn(len(*d))
	n := (*d)[i]
	suite := n / 13
	rank := n - 13*suite
	*d = append((*d)[:i], (*d)[i+1:]...)
	return &Card{Suite(suite), Rank(rank)}
}

func (d *Deck) Reset() {
	nums := make([]int, 52)
	for i := range nums {
		nums[i] = i
	}
	*d = nums
}

func (d *Deck) Shuffle() {
	d.shuffle0()
}

func (d *Deck) shuffle0() {
	rand.Shuffle(len(*d), func(i, j int) {
		(*d)[i], (*d)[j] = (*d)[j], (*d)[i]
	})
}

func (d *Deck) shuffle1() {
	for i := len(*d); i > 0; i-- {
		j := rand.Intn(i + 1)
		(*d)[i], (*d)[j] = (*d)[j], (*d)[i]
	}
}

func (d *Deck) shuffle2() {
	dLen := len(*d)
	ans := make([]int, 0, dLen)
	for i := 0; i < dLen; i++ {
		n := rand.Intn(dLen - i)
		ans = append(ans, (*d)[n])
		*d = append((*d)[:n], (*d)[n+1:]...)
	}
	*d = ans
}
