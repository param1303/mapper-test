package mapper

import (
//	"fmt"
	"unicode"
	"regexp"
)

// replaceStrFlg  variable to check position and update the case accordingly 
var replaceStrFlg int

// isAlphaAndNum  check alpha and num pattern
var isAlphaAndNum = regexp.MustCompile("^[a-zA-Z0-9_]*$").MatchString

type Interface interface {
   TransformRune(pos int)
   GetValueAsRuneSlice() []rune
}

func MapString(i Interface) {
   for pos, _ := range i.GetValueAsRuneSlice() {
      i.TransformRune(pos)
   }
}

// And change your code to look like this:

// SkipString struct 
type SkipString struct {
Pos int
Str string
}


func (s SkipString) GetValueAsRuneSlice() []rune{
	return  []rune(s.Str)
}


func  ReplaceIndexChar(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)

}

func (s *SkipString) TransformRune(pos int) {
        char :=[]rune(s.Str) [pos]
	if isAlphaAndNum(string(char)){
        replaceStrFlg = replaceStrFlg+1
//	fmt.Println("--->>>",replaceStrFlg)
		
	}
if(replaceStrFlg==s.Pos){
	rs :=unicode.ToUpper(char)
	s.Str = ReplaceIndexChar(s.Str,rs,pos)
	replaceStrFlg=0
}else{
       rs :=unicode.ToLower(char)
       s.Str = ReplaceIndexChar(s.Str,rs,pos)
}
}


func NewSkipString(p int,s string) SkipString {
return SkipString{Pos:p,Str:s}
}
