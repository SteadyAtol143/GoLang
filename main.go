package main

import(
	"fmt"
	"math/rand"
	"time"
)

func main () {
	var command string
	var randomCoinVal int
	eagVal := 0
	nonEagVal := 0
	srnd1 := rand.NewSource(time.Now().UnixNano())
	rnd1 := rand.New(srnd1)

	fmt.Println("Программа: Давайте побросаем монетку?")
	fmt.Println("Вы: ")
	fmt.Scan(&command)
	if command == "No" {
		fmt.Println("Программа: Ну ладно ;(")
	} else if command == "Yes" {
		for  {
			fmt.Println("Программа: ведите th чтобы бросить монетку.")
			fmt.Println("Программа: ведите stop чтобы прекратить работу эмулятора.")
			fmt.Println("Вы: ")
			fmt.Scan(&command)
			if command == "stop" {
				fmt.Println("Всего орлов: ", eagVal)
				fmt.Println("Всего решек: ", nonEagVal)
				break
			} else {
				randomCoinVal = rnd1.Intn(2)
			    if randomCoinVal == 1 {
			    	fmt.Println("Выпал орёл!")
			    	eagVal += 1
			    } else {
			    	fmt.Println("Выпала решка!")
			    	nonEagVal += 1
			    }
			}
		}
	} else {
		fmt.Println("Программа: Я вас не поняла, но да ладно...")
	}
}

