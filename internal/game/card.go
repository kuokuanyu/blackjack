package game

import (
	"fmt"
	"math/rand"
)

// Card 代表一張牌
type Card struct {
	Suit  string // 花色（黑桃、紅心、梅花、方塊）
	Value string // 點數（2~10、J、Q、K、A）
}

// Deck 代表一副牌（多張Card）
type Deck []Card

// 建立一副標準撲克牌
func NewDeck() Deck {
	suits := []string{"Spades", "Hearts", "Clubs", "Diamonds"}
	values := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
	deck := Deck{}

	// 產生52張牌
	for _, suit := range suits {
		for _, val := range values {
			deck = append(deck, Card{Suit: suit, Value: val})
		}
	}
	return deck
}

// 洗牌，隨機打亂牌組順序
func (d Deck) Shuffle() {
	// rand.Seed(time.Now().UnixNano()) // 設定亂數種子

	// 使用rand.Shuffle打亂牌組
	// len(d)是牌組長度，func(i, j int)是交換兩張牌的函式
	// 索引 i 和 j 的兩張牌互換
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}

// // 從牌堆頂端抽一張牌，並回傳抽出牌與剩餘牌堆
func (d *Deck) Draw() Card {
	if len(*d) == 0 {
		fmt.Println("牌堆已空，無法抽牌！")
		return Card{} // 回傳空牌
	}

	card := (*d)[0] // 取得牌堆頂端的牌(第一張)
	*d = (*d)[1:]   // 刪除牌堆頂端的牌，剩餘牌堆從第二張開始

	// fmt.Printf("抽出牌: %s of %s\n", card.Value, card.Suit)

	// 回傳抽出的牌
	return card
}
