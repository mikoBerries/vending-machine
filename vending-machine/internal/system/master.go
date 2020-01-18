package system

import (
	"errors"
	"fmt"

	"../../model"
)

//insert coin
func CalculationInsertCoin(balance int, allCoin []*model.Coin, inputedCoins []int, newCoin int) (int, []int) {
	//validasi Coin nominal
	err := validateNewCoin(newCoin)
	if err != nil {
		fmt.Println(err)
		return balance, inputedCoins
	}
	//  nominal == newCoin >> Count++
	if newCoin == 10 { //10 JPY
		allCoin[0].Count++
	} else if newCoin == 50 { //50 JPY
		allCoin[1].Count++
	} else if newCoin == 100 { //100 JPY
		allCoin[2].Count++
	} else if newCoin == 500 { //500 JPY
		allCoin[3].Count++
	}
	//add balance
	balance += newCoin
	//add coin
	inputedCoins = append(inputedCoins, newCoin)
	return balance, inputedCoins
}

func validateNewCoin(newCoin int) error {
	result := false
	validCoins := [...]int{10, 50, 100, 500}
loop:
	for _, validCoin := range validCoins {
		if newCoin == validCoin {
			result = true
			break loop
		}
	}
	if !result {
		return errors.New("Invalid Coin (Please insert 10, 50, 100, 500 JPY)")
	} else {
		return nil
	}

}

//Choose Item
func ChooseItem(selectedItems []model.Item, goods []*model.Good, balance int, allCoin []*model.Coin, goodsIndex int) ([]model.Item, int) {
	goodsIndex--
	if goodsIndex < 0 { //no item index below 0
		return selectedItems, balance
	}

	if len(goods) < goodsIndex { //out of index
		return selectedItems, balance
	}
	if balance >= goods[goodsIndex].Price && goods[goodsIndex].Quantity > 0 { //can buy selected goodsIndex
		if checkChange(balance-goods[goodsIndex].Price, allCoin) { //check can do change ?
			//update inputCoins
			balance -= goods[goodsIndex].Price
			//update goods Quantity
			goods[goodsIndex].Quantity--
			//input selectedItems
			selectedItems = append(selectedItems, model.Item{
				Name:  goods[goodsIndex].Name,
				Price: goods[goodsIndex].Price,
			})
		}
	}
	return selectedItems, balance
}

//change
func CalculateReturnCoin(returnCoins []int, balance int, allCoin []*model.Coin, buy bool, inputedCoins []int) ([]int, []int, int) {
	if !buy { //didn't buy anythings
		if len(returnCoins) == 0 {
			return inputedCoins, make([]int, 0), 0
		} else { //some coin in returnCoins
			for _, inputedCoin := range inputedCoins {
				returnCoins = append(returnCoins, inputedCoin)
			}
			return returnCoins, make([]int, 0), 0
		}
	}

	if balance == 0 { //nothings to do
		return returnCoins, inputedCoins, 0
	}

loop1:
	for i := len(allCoin) - 1; i >= 0; i-- {
	loop2:
		for allCoin[i].Count > 0 {
			if balance >= allCoin[i].Nominal { // reduce total  nominal coin
				balance -= allCoin[i].Nominal
				returnCoins = append(returnCoins, allCoin[i].Nominal)
				allCoin[i].Count--
			} else { //skip to smaller nominal
				break loop2
			}
			if balance == 0 { //done
				break loop1
			}
		}
	}
	return returnCoins, make([]int, 0), balance
}

func checkChange(balance int, allCoin []*model.Coin) bool {
	if !validateChange(allCoin) { //check all coin before
		return false
	}
	allCoinCopy := make([]*model.Coin, len(allCoin))
	copy(allCoinCopy, allCoin)
loop1:
	for i := len(allCoinCopy); i < 0; i-- {
	loop2:
		for allCoinCopy[i].Count > 0 {
			if balance >= allCoinCopy[i].Nominal { // reduce total  nominal coin
				balance -= allCoinCopy[i].Nominal
				allCoinCopy[i].Count--
			} else { //skip to smaller nominal
				break loop2
			}
			if balance == 0 { //done
				break loop1
			}
		}
	}

	if !validateChange(allCoinCopy) { //check all coin after
		return false
	}
	return true
}

func validateChange(allCoin []*model.Coin) bool {
	if allCoin[0].Count < 9 { //10 JPY < 9
		fmt.Println("10 JPY coin: Less than 9 coins")
		return false
	} else if allCoin[2].Count < 4 { //100 JPY < 4
		fmt.Println("100 JPY coin: Less than 4 coins")
		return false
	}
	return true
}
