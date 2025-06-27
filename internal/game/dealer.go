package game

import "fmt"

type Dealer struct {
	Hand []Card // 莊家的手牌
}

// ReceiveCard 莊家抽牌
func (d *Dealer) ReceiveCard(c Card) {
	d.Hand = append(d.Hand, c)
}

// ResetHand 重置莊家手牌（新回合前使用）
func (d *Dealer) ResetHand() {
	d.Hand = []Card{}
}

// Score 計算莊家目前手牌的點數
func (d *Dealer) Score() int {
	total := 0
	aces := 0

	for _, card := range d.Hand {
		switch card.Value {
		case "J", "Q", "K":
			total += 10
		case "A":
			total += 1
			aces++
		default:
			var val int
			fmt.Sscanf(card.Value, "%d", &val)
			total += val
		}
	}

	for aces > 0 && total+10 <= 21 {
		total += 10
		aces--
	}

	return total
}

// IsBust 判斷玩家是否爆牌（超過 21 點）
func (d *Dealer) IsBust() bool {
	return d.Score() > 21
}
