package main

import(
	"fmt"
)

func main () {
	var price int
	var given int

	fmt.Println("Стоимость товара: ")
	fmt.Scan(&price)
	fmt.Println("Пользователь дал: ")
	fmt.Scan(&given)

	fmt.Println("Сдача: ", given - price)
	
	fmt.Println("Можно выдать: ")
	fmt.Println("Десятирублёвых монет: ", (given-price)/(10))
	fmt.Println("Пятирублёвых монет: ", ((given-price) - 10*((given-price)/(10)))/5)
	fmt.Println("Рублёвых монет: ",  ((given-price) - 10*((given-price)/(10))) - 5*(((given-price) - 10*((given-price)/(10)))/5))
}
