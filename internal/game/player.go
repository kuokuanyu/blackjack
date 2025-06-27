package game

import "fmt"

// Player 玩家的資料結構
type Player struct {
	Name    string // 玩家名稱（或 ID）
	Hand    []Card // 玩家手上的牌
	Bet     int    // 玩家本回合下注的金額
	Balance int    // 玩家帳戶餘額
}

// PlaceBet 玩家下注
// 輸入下注金額，回傳是否成功（餘額足夠才可以下注）
func (p *Player) PlaceBet(amount int) bool {
	if amount > p.Balance { // 檢查下注金額是否超過餘額
		// 餘額不足，下注失敗
		return false
	}

	p.Bet = amount      // 記錄下注金額
	p.Balance -= amount // 從餘額中扣除下注金額
	return true         // 下注成功
}

// ReceiveCard 玩家抽牌
func (p *Player) ReceiveCard(c Card) {
	p.Hand = append(p.Hand, c) // 將新抽到的牌加入玩家手牌中
}

// ResetHand 重置玩家手牌（新回合前使用）
func (p *Player) ResetHand() {
	p.Hand = []Card{} // 清空手牌
	p.Bet = 0         // 重置下注
}

// Score 計算玩家目前手牌的點數
// A 可當作 1 或 11，只要不爆就盡量當作 11
func (p *Player) Score() int {
	total := 0 // 總點數
	aces := 0  // A 的數量，用來後續判斷是否可以+10

	for _, card := range p.Hand {
		switch card.Value {
		case "J", "Q", "K":
			total += 10 // J/Q/K 各算 10 點
		case "A":
			total += 1 // 先當作 1 點處理
			aces++     // 記錄有幾張 A
		default:
			var val int
			fmt.Sscanf(card.Value, "%d", &val) // 將 "2"~"10" 轉成整數
			total += val
		}
	}

	// 若有 A，且總點數 +10 不會爆牌，則將 A 當作 11 點使用
	for aces > 0 && total+10 <= 21 {
		total += 10
		aces--
	}

	return total
}

// IsBust 判斷玩家是否爆牌（超過 21 點）
func (p *Player) IsBust() bool {
	return p.Score() > 21
}
