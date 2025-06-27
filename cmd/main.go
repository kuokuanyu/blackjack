package main

import (
	"fmt"

	"github.com/kuokuanyu/blackjack/internal/game"
)

func main() {
	g := game.NewGame("Kuo", 1000) // 初始化遊戲，玩家名稱為 "KuanYu"，初始餘額為 1000
	g.StartRound(200)              // 開始一個新的回合，玩家下注 200

	g.PlayerTurn() // 玩家回合，玩家可以持續抽牌直到點數 >= 17 或爆牌
	if g.Player.IsBust() {
		// fmt.Println("玩家爆牌，結束回合")
		fmt.Printf("玩家餘額: %d\n", g.Player.Balance)
		return
	}

	g.DealerTurn() // 莊家回合，莊家持續抽牌直到點數 >= 17 或爆牌
	if g.Dealer.IsBust() {
		// fmt.Println("莊家爆牌，結束回合")
		fmt.Printf("玩家餘額: %d\n", g.Player.Balance)
		return
	}

	g.Judge() // 判斷勝負

	fmt.Printf("玩家餘額: %d\n", g.Player.Balance)
}
