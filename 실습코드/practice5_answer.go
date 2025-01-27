package main

import "fmt"

func main(){
	/*
	실습5) 
	GoCafe에 메뉴가 5개로 늘었습니다.
	어떤 VIP 손님이 메뉴에 있는 음료를 처음부터 끝까지 주문했습니다. 
	for문을 사용하여 메뉴와 가격을 알려주고, 총 금액을 계산해봅시다.
	*/
	drinkNames  := []string{"Americano","Latte","CafeMocha","GreenTeaLatte","HotChoco"}
	drinkPrices := []int{1500,1700,1900,2300,2000}

	totalPrice := 0

	for i, val := range drinkNames {
		fmt.Println(val, " : ", drinkPrices[i], "원")
	}
	
	for _, val := range drinkPrices {
		totalPrice += val
	}
	fmt.Println("총 금액은",totalPrice,"입니다.")
}
