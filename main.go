package main

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

// Hand 代表一手牌（玩家或莊家）
type Hand []Card

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

// 從牌堆頂端抽一張牌，並回傳抽出牌與剩餘牌堆
func (d *Deck) Draw() Card {
	card := (*d)[0] // 取得牌堆頂端的牌(第一張)
	*d = (*d)[1:]   // 刪除牌堆頂端的牌，剩餘牌堆從第二張開始
	if len(*d) == 0 {
		fmt.Println("牌堆已空，無法抽牌！")
		return Card{} // 回傳空牌
	}

	// fmt.Printf("抽出牌: %s of %s\n", card.Value, card.Suit)

	// 回傳抽出的牌
	return card
}

// 計算手牌點數（A可算1或11點），回傳最高不超過21的點數
func (h Hand) CalculateScore() int {
	total := 0 // 初始點數
	aces := 0  // A的數量

	// 計算初步點數，J/Q/K算10點，A先算1點，記錄A數量
	for _, card := range h {
		switch card.Value {
		case "J", "Q", "K":
			total += 10 // J、Q、K都算10點
		case "A":
			total += 1 // A先算1點
			aces++     // A先算1點，記錄A的數量
		default:
			// 2~10轉成數字加總
			var val int
			fmt.Sscanf(card.Value, "%d", &val)
			total += val
		}
	}

	// 判斷A要不要加10點（11點），只要不爆牌就加
	for aces > 0 && total+10 <= 21 {
		total += 10
		aces--
	}

	return total
}

// 判斷是否爆牌（超過21點）
func (h Hand) IsBust() bool {
	return h.CalculateScore() > 21
}

// 輸出手牌的字串（方便印出）
func (h Hand) String() string {
	result := ""
	for _, c := range h {
		result += fmt.Sprintf("%s of %s, ", c.Value, c.Suit)
	}
	return result
}

// 模擬21點遊戲回合
func main() {
	deck := NewDeck() // 新牌堆
	deck.Shuffle()    // 洗牌

	playerHand := Hand{} // 玩家手牌
	dealerHand := Hand{} // 莊家手牌

	// 玩家、莊家各發兩張牌
	playerHand = append(playerHand, deck.Draw(), deck.Draw())
	dealerHand = append(dealerHand, deck.Draw(), deck.Draw())

	fmt.Printf("玩家手牌: %s (點數: %d)\n", playerHand.String(), playerHand.CalculateScore())
	fmt.Printf("莊家手牌: %s (點數: %d)\n", dealerHand.String(), dealerHand.CalculateScore())

	// 玩家行動範例：玩家持續要牌直到點數 >= 17（簡化版）
	for playerHand.CalculateScore() < 17 {
		card := deck.Draw() // 抽一張牌
		fmt.Printf("玩家抽牌: %s of %s\n", card.Value, card.Suit)
		playerHand = append(playerHand, card)
		if playerHand.IsBust() {
			fmt.Println("玩家爆牌，莊家勝利！")
			return
		}
	}

	// 莊家行動：莊家持續要牌直到點數 >= 17
	for dealerHand.CalculateScore() < 17 {
		card := deck.Draw()
		fmt.Printf("莊家抽牌: %s of %s\n", card.Value, card.Suit)
		dealerHand = append(dealerHand, card)
		if dealerHand.IsBust() {
			fmt.Println("莊家爆牌，玩家勝利！")
			return
		}
	}

	// 判斷勝負
	playerScore := playerHand.CalculateScore()
	dealerScore := dealerHand.CalculateScore()

	fmt.Printf("玩家最終點數: %d\n", playerScore)
	fmt.Printf("莊家最終點數: %d\n", dealerScore)

	switch {
	case playerScore > dealerScore:
		fmt.Println("玩家勝利！")
	case playerScore < dealerScore:
		fmt.Println("莊家勝利！")
	default:
		fmt.Println("平手！")
	}
}
