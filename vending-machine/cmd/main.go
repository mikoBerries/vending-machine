package main

import (
	"fmt"

	"../internal/other"
	"../internal/system"
	"../internal/vending"
	"../model"
)

func main() {
	defer other.Catch()
	//init good
	goods := make([]*model.Good, 0)
	goods = append(goods, &model.Good{ //hardcode goods
		Name:     "Canned coffee",
		Price:    120,
		Quantity: 2,
	})
	goods = append(goods, &model.Good{
		Name:     "Water PET bottle",
		Price:    100,
		Quantity: 0,
	})
	goods = append(goods, &model.Good{
		Name:     "Sport drinks",
		Price:    150,
		Quantity: 1,
	})

	//init coins on vending machine
	allCoin := make([]*model.Coin, 0)
	allCoin = append(allCoin, &model.Coin{ //hardcode coins
		Nominal: 10,
		Count:   20,
	})
	allCoin = append(allCoin, &model.Coin{
		Nominal: 50,
		Count:   0,
	})
	allCoin = append(allCoin, &model.Coin{
		Nominal: 100,
		Count:   10,
	})
	allCoin = append(allCoin, &model.Coin{
		Nominal: 500,
		Count:   0,
	})

	//init items
	items := make([]model.Item, 0)

	//init returnCoins
	returnCoins := make([]int, 0)

	//my balance
	balance := 0

	//bool buy
	buy := false

	//inputed Coin
	inputedCoins := make([]int, 0)
	for {
		fmt.Println("----------------------------------")
		fmt.Println("[Input amount]\t", balance, " JPY")
		vending.DisplayChanges(allCoin)
		vending.DisplayReturnCoin(returnCoins)
		vending.DisplayGoods(goods, balance)
		vending.DisplayItems(items)
		fmt.Println("----------------------------------")

		command := make([]int, 2)
		fmt.Print("Command: ")
		fmt.Scanln(&command[0], &command[1])

		if command[0] == 1 { //insert coin to vending machine -> allCoin, balance
			balance, inputedCoins = system.CalculationInsertCoin(balance, allCoin, inputedCoins, command[1])
		} else if command[0] == 2 { //choose goods to buy -> Items
			items, balance = system.ChooseItem(items, goods, balance, allCoin, command[1])
		} else if command[0] == 3 { //get Items (buyed) (flush object)
			items = make([]model.Item, 0)
		} else if command[0] == 4 { //fill returnCoins
			returnCoins, inputedCoins, balance = system.CalculateReturnCoin(returnCoins, balance, allCoin, buy, inputedCoins)
		} else if command[0] == 5 { //get returnCoins (flush Object)
			returnCoins = make([]int, 0)
		}
	}
}
