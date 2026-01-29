package main

import (
	"fmt"
)
func work (x int) int{
	return x*x+2 //  рандомное написала
}//медленная функция
func main() {
	cash:=make(map[int]int) //кеш
	for i:=0;i<10;i++{
		var chislo int
		fmt.Scan(&chislo)
	//читаем наши входные 10 чисел
		if result,ok:=cash[chislo];ok{ //проверяем есть ли в нашем кеше результат
			fmt.Print(result)
		}else{ //рез-т не найден. находим значение через медленную функцию и записываем в кеш
			result=work(chislo)
			cash[chislo]=result
			fmt.Print(result)
		}
		if i<9{   //для вывода красивого)
			fmt.Print(" ")
		}
	}
	fmt.Println(" time limit ok")
}
