package main

import (
	"sort"
	"testing"
)

// NewCard конструктор новой карты
func TestNewCard(t *testing.T) {

	c := NewCard(1, Diamonds)
	if c.Rank != 1 || c.Suit != Diamonds {
		t.Errorf("NewCard: expected [A:D] got [%s:%s]", c.Suit.String(), c.RankString())
	}
}

func TestRankString(t *testing.T) {

	c := NewCard(1, Diamonds)
	str := c.RankString()
	if str != "A" {
		t.Errorf("RankString: expectd [A] got [%s]", str)
	}
}

// String as is
func TestCardString(t *testing.T) {

	c := NewCard(1, Diamonds)
	str := c.String()

	if str != "[D:A]" {
		t.Errorf("RankString: expectd [D:A] got %s", str)
	}
}

func TestNewDeckOfCards(t *testing.T) {

	cards := NewDeckOfCards()
	len := len(cards)
	if len != 52 {
		t.Errorf("NewDeckOfCards: expected len=52 got %d", len)
	}
	if cards[0].Rank != 1 || cards[0].Suit != 1 {
		t.Errorf("NewDeckOfCards: expected Rank=1 Suit=1 go Rank=%d Suit=%d",
			cards[0].Rank, cards[0].Suit)
	}
	if cards[51].Rank != 13 || cards[51].Suit != 4 {
		t.Errorf("NewDeckOfCards: expected Rank=13 Suit=4 go Rank=%d Suit=%d",
			cards[51].Rank, cards[51].Suit)
	}
}

func TestCardsSortedByRank(t *testing.T) {

	cards := NewDeckOfCards()
	cards.Shuffle()
	sort.Sort(CardsSortedByRank(cards))
	str := cards[0].String()
	if str != "[D:A]" {
		t.Errorf("CardsSortedByRank: expectd [D:A] got %s", str)
	}
}

func TestCardsSortedBySuitRank(t *testing.T) {

	cards := NewDeckOfCards()
	cards.Shuffle()
	sort.Sort(CardsSortedByRank(cards))
	str := cards[0].String()
	if str != "[H:A]" {
		t.Errorf("CardsSortedBySuitRank: expectd [H:A] got %s", str)
	}
}
