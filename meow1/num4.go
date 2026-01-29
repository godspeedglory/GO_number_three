package main
import (
	"fmt"
	"strconv"
	"strings"
)
func main(){
	var stroka string
	fmt.Scan(&stroka)
	chasti:=strings.Split(stroka,";")
	meow:= func(s string) (float64, error){
		s=strings.ReplaceAll(s," ","")
		s=strings.ReplaceAll(s,",",".")
		return strconv.ParseFloat(s,64)
	}
	chislo1,err1:=meow(chasti[0])
	chislo2,err2:=meow(chasti[1])
	if err1!=nil || err2!=nil {
		return
	}
	result:=chislo1/chislo2
	fmt.Printf("%.4f\n",result)
