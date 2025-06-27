package game

import (
	"fmt"
)

type Game struct {
	Deck   Deck    // 一副牌
	Player *Player // 玩家
	Dealer *Dealer // 莊家
}

func NewGame(playerName string, initialBalance int) *Game {
	player := &Player{Name: playerName, Balance: initialBalance} // 初始化玩家，設定名稱和初始餘額
	dealer := &Dealer{}                                          // 初始化莊家，手牌為空
	deck := NewDeck()                                            // 建立一副新牌
	deck.Shuffle()                                               // 洗牌

	return &Game{
		Deck:   deck,
		Player: player,
		Dealer: dealer,
	}
}

// StartRound 開始一個新的回合
// 玩家可以下注，然後發牌給玩家和莊家
func (g *Game) StartRound(bet int) {
	g.Player.ResetHand() // 重置玩家手牌
	g.Dealer.ResetHand() // 重置莊家手牌

	if !g.Player.PlaceBet(bet) { // 嘗試下注，如果餘額不足則下注失敗
		fmt.Println("下注失敗：餘額不足")
		return
	}

	// 初始發牌
	g.Player.ReceiveCard(g.Deck.Draw()) // 玩家抽一張牌
	g.Dealer.ReceiveCard(g.Deck.Draw()) // 莊家抽一張牌
	g.Player.ReceiveCard(g.Deck.Draw()) // 玩家再抽一張牌
	g.Dealer.ReceiveCard(g.Deck.Draw()) // 莊家再抽一張牌

	fmt.Printf("玩家手牌: %+v\n", g.Player.Hand)
	fmt.Printf("莊家手牌: %+v\n", g.Dealer.Hand[:1]) // 一張明牌
}

// PlayerTurn 玩家點數判斷
func (g *Game) PlayerTurn() {
	for g.Player.Score() < 17 {
		g.Player.ReceiveCard(g.Deck.Draw()) // 玩家持續抽牌直到點數 >= 17
		if g.Player.IsBust() {
			fmt.Println("玩家爆牌，莊家勝！")
			return
		}
	}
}

// DealerTurn 莊家點數判斷
func (g *Game) DealerTurn() {
	for g.Dealer.Score() < 17 { // 莊家持續抽牌直到點數 >= 17
		g.Dealer.ReceiveCard(g.Deck.Draw())
		if g.Dealer.IsBust() {
			fmt.Println("莊家爆牌，玩家勝！")
			g.Player.Balance += g.Player.Bet * 2 // 玩家贏得下注金額的兩倍
			// fmt.Println("g.Player.Bet * 2:", g.Player.Bet*2)
			// fmt.Println("玩家餘額:", g.Player.Balance)
			return
		}
	}
}

func (g *Game) Judge() {
	ps := g.Player.Score()
	ds := g.Dealer.Score()

	switch {
	case ps > ds:
		fmt.Println("玩家勝！")
		g.Player.Balance += g.Player.Bet * 2 // 玩家贏得下注金額的兩倍
		// fmt.Println("g.Player.Bet * 2:", g.Player.Bet*2)
		// fmt.Println("玩家餘額:", g.Player.Balance)
	case ps < ds:
		fmt.Println("莊家勝！")
	// 前面已經扣過下注金額了，所以這裡不需要再扣
	default:
		fmt.Println("平手，退回下注！")
		g.Player.Balance += g.Player.Bet // 平手，退回下注金額
		// fmt.Println("g.Player.Bet :", g.Player.Bet)
		// fmt.Println("玩家餘額:", g.Player.Balance)
	}
}
