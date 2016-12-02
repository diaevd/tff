package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
)

// Все названия взяты отсюда http://ok-english.ru/kartochnyie-masti-na-angliyskom-suits/

// suit масть
type Suit uint8

// rank ранг
type Rank uint8

const (
	// Hearts черви
	Hearts Suit = 1 + iota
	// Diamonds бубны
	Diamonds
	// Clubs трефы
	Clubs
	// Spades пики
	Spades
)

// suitName имена мастей
var suitName = map[Suit]string{
	Hearts:   "H",
	Diamonds: "D",
	Clubs:    "C",
	Spades:   "S",
}

// String as is
func (s Suit) String() string {

	return suitName[s]
}

// Card ранг и масть карты
type Card struct {
	Rank Rank
	Suit Suit
}

// rankName имена некоторых рангов
var rankName = map[Rank]string{
	// Да, туз именно 1, я играю в покер ;)
	1: "A", 11: "J", 12: "Q", 13: "K",
}

// NewCard конструктор новой карты
func NewCard(r Rank, s Suit) *Card {

	if r < 1 || r > 13 {
		panic("Rank is out of range")
	}
	if s < 1 || s > 4 {
		panic("Suit is out of range")
	}

	return &Card{Rank: r, Suit: s}
}

// String as is
func (c *Card) String() string {

	return fmt.Sprintf("[%s:%s]", c.Suit.String(), c.RankString())
}

// RankString as is
func (c *Card) RankString() string {
	str := strconv.Itoa(int(c.Rank))
	if v, ok := rankName[c.Rank]; ok {
		str = v
	}

	return str
}

// Cards колода
type Cards []*Card

func (c Cards) String() string {
	var strs []string

	for _, v := range c {
		strs = append(strs, v.String())
	}

	return strings.Join(strs, " ")
}

// NewDeckOfCards конструктор новой колоды
func NewDeckOfCards() Cards {
	cards := Cards{}

	// Набьем колоду
	for _, s := range []Suit{Hearts, Diamonds, Clubs, Spades} {
		for i := 1; i <= 13; i++ {
			cards = append(cards, NewCard(Rank(i), s))
		}
	}

	return cards
}

// Shuffle перемешивает карты
func (c Cards) Shuffle() {
	for i := range c {
		j := rand.Intn(i + 1)
		c[i], c[j] = c[j], c[i]
	}
}

// PrintDeck просто печать по рангам
// красиво, если отсортировано :)
func (c Cards) PrintDeck() {
	i := 0

	for _, card := range c {
		fmt.Print(card.String(), " ")
		i++
		if i >= 13 {
			fmt.Println("")
			i = 0
		}
	}
}

// CardsSortedByRank хелпер для сортировки по рангу
type CardsSortedByRank Cards

// Len ...
func (c CardsSortedByRank) Len() int {

	return len(c)
}

// Swap ...
// названия l => left и r => right - это привычка из C/C++
func (c CardsSortedByRank) Swap(l, r int) {

	c[l], c[r] = c[r], c[l]
}

// Less ...
func (c CardsSortedByRank) Less(l, r int) bool {

	return c[l].Rank < c[r].Rank
}

// CardsSortedBySuitRank хелпер для сортировки по масти/рангу
type CardsSortedBySuitRank Cards

// Len ...
func (c CardsSortedBySuitRank) Len() int {

	return len(c)
}

// Swap ...
func (c CardsSortedBySuitRank) Swap(l, r int) {

	c[l], c[r] = c[r], c[l]
}

// Less ...
func (c CardsSortedBySuitRank) Less(l, r int) bool {

	if c[l].Suit < c[r].Suit {
		return true
	}
	if c[l].Suit > c[r].Suit {
		return false
	}

	return c[l].Rank < c[r].Rank
}

// main ...
func main() {

	cards := NewDeckOfCards()
	// log.Println(cards)
	fmt.Println("New deck of cards:")
	cards.PrintDeck()
	fmt.Println("")

	// перемешаем
	cards.Shuffle()
	// log.Println(cards)
	fmt.Println("Shuffled deck of cards:")
	cards.PrintDeck()
	fmt.Println("")

	// Сортировка по рангу
	sort.Sort(CardsSortedByRank(cards))
	// тут не особо красиво
	fmt.Println("Sorted by rank:")
	cards.PrintDeck()
	fmt.Println("")
	// Сортировка по масти и рангу
	sort.Sort(CardsSortedBySuitRank(cards))
	// тут выглдит лучше
	fmt.Println("Sorted by Suit and Rank:")
	cards.PrintDeck()
	fmt.Println("")
}
