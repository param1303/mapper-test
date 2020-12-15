package main

import  (
	"fmt"
	"test/mapper"
	"regexp"
	"unicode"
)

func CapitalizeEveryThirdAlphanumericChar(s string,p int) string{

var replaceStrFlg  int

var isAlphaAndNum = regexp.MustCompile("^[a-zA-Z0-9_]*$").MatchString

stringSlice :=[]rune(s)
var newString string

for pos,_ := range stringSlice{
        char :=[]rune(s) [pos]
        if isAlphaAndNum(string(char)) {
	replaceStrFlg = replaceStrFlg+1

}

if replaceStrFlg == p {
       replaceStrFlg = 0
       newString = newString + string(unicode.ToUpper(char))
     }else{
	newString = newString + string(unicode.ToLower(char))
     }
}

return newString
}
func main(){

nString :=CapitalizeEveryThirdAlphanumericChar ("Aspiration.com",3)

fmt.Println(nString)

s := mapper.NewSkipString(3, "Aspiration.com")

mapper.MapString(&s)
   fmt.Println(s.Str)
}
