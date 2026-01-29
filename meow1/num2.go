package main
import "fmt"
func convert(x int64) uint16 {
    if x<0 {
		panic("отрицательное число")
    }
    if x>65535 {  
		panic("число слишком большое")
    }
    return uint16(x)
}
func main(){
	var chislo int64
	fmt.Scan(&chislo)
	fmt.Print(convert(chislo))
