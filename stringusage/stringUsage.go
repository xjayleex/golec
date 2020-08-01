package stringusage

import (
	"fmt"
	"strings"
	"unsafe"
)

func StringUsage1() {
	s1 := []string{"apple","banana","cheese"}
	s2 := strings.Join(s1, " ")
	fmt.Println(s2)
}
func StringUsage2() {
	str := "hello, world!"
	vec := *(*[]byte)(unsafe.Pointer(&str))
	fmt.Println(vec)
	fmt.Println(string(vec))
	vec1 := []byte(str)
	vec1 = append(vec1, 'a')
	vec1[0] += 2
	fmt.Println(string(vec1))
}
func StrBuilder(){
	var sb strings.Builder
	sb.Grow(100)
	fmt.Println("Cap :",sb.Cap())
	sb.WriteString("ABCDE")
	sb.Write([]byte{'a','b','c','d'})
	fmt.Println(sb.String())
	sb.WriteByte('e')
	fmt.Println(sb.String())
}

func CorrectTextModf(){
	s := "We went to eat at multiple cafe"
	fmt.Println(&s)
	cafe := "cafe"
	if p := strings.Index(s,cafe); p != -1 {
		p += len(cafe)
		s = s[:p] + "s" + s[p:]
	}
	/////
	s1 := "We went to eat at multiple cafe"

}
