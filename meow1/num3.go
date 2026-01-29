package main
import (
	"strconv"
)
func adding (a,b string) int64{
	tolkoCifri:= func(s string) string{
		result:=""
		for i:=0;i<len(s);i++{
			if s[i]>='0' && s[i]<='9'{
				result=result+string(s[i])
			}
		}
		return result
	}
	clean1:=tolkoCifri(a)
	clean2:=tolkoCifri(b)
	chislo1,_:=strconv.ParseInt(clean1,10,64)
	chislo2,_:=strconv.ParseInt(clean2,10,64)
	return chislo1+chislo2
