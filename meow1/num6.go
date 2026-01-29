package main
import "fmt"
type Battery struct{
	stroka string
}
func (b Battery) String() string {
	kolEd:=0
	for i:=0;i<len(b.stroka);i++{
		if b.stroka[i]=='1'{
			kolEd++
		}
	}
	otv:="["
	for i:=0;i<10-kolEd;i++{
		otv=otv+" "
	}
	for i:=0;i<kolEd;i++{
		otv=otv+"X"
	}
	otv=otv+"]"
	return otv
	}
func main(){
	var meow string
	fmt.Scan(&meow)
	battery:=Battery{stroka:meow}
	fmt.Println(battery)
