package main
import (
	"fmt"
	"strconv"
)
func main(){
	var fn func(uint) uint
	fn=func(n uint) uint{
		s:=fmt.Sprint(n)
		result:= ""
		for i:=0;i<len(s);i++{
			chislo:=s[i]-'0'
			if chislo%2==0 && chislo!=0 {
				result+=string(s[i])
			}
		}
		if result=="" {
			return 100
		}
		num, _:=strconv.Atoi(result)
		return uint(num)
	}
	var num uint
	fmt.Scan(&num)
	fmt.Println(fn(num))
