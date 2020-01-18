package vending

import (
	"fmt"

	"../../model"
)

func DisplayGoods(goods []*model.Good, totalInputcoins int) {
	fmt.Print("[Items for sale]")
	if len(goods) == 0 {
		fmt.Println("\t Empty")
	} else {
		for i, good := range goods {
			goodLabel := fmt.Sprint("\t ", i+1, ". ", good.Name, "\t", good.Price, " JPY")
			if good.Quantity == 0 {
				goodLabel += " Sold out"
			} else if good.Price <= totalInputcoins {
				goodLabel += " Available for purchase"
			}
			fmt.Println(goodLabel)
		}
	}

}

func DisplayItems(items []model.Item) {
	fmt.Print("[Outlet]")
	if len(items) == 0 {
		fmt.Println("\t Empty")
	} else {
		for _, item := range items {
			fmt.Println("\t", item.Name)
		}
	}
}

func DisplayReturnCoin(returnCoins []int) {
	fmt.Print("[Return gate]")
	if len(returnCoins) == 0 {
		fmt.Println("\t Empty")
	} else {
		for _, returnCoin := range returnCoins {
			fmt.Println("\t", returnCoin, " JPY")
		}
	}
}

func DisplayChanges(coins []*model.Coin) {
	fmt.Print("[Change]\t")

	JpyString10 := "\t 10 JPY "
	JpyString100 := "\t 100 JPY "

	for _, coin := range coins {
		if coin.Nominal == 10 {
			if coin.Count > 0 {
				JpyString10 += "Change"
			} else {
				JpyString10 += "No Change"
			}
		} else if coin.Nominal == 100 {
			if coin.Count > 0 {
				JpyString100 += "Change"
			} else {
				JpyString100 += "No Change"
			}
		}
	}
	fmt.Println(JpyString10)
	fmt.Println(JpyString100)

}
