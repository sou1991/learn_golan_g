package main

import (
	"fmt"
	"regexp"	
)

func main(){
	//半角数値チェック空白Ok
	fmt.Println(DoCheckRegexp("123", `^[0-9]*$`))
	//半角数値チェック空白NG
	fmt.Println(DoCheckRegexp("", `^[0-9]+$`))

	fmt.Println(DoCheckRegexp("ase", `^[a-zA-Z]+$`))//true

	fmt.Println(DoCheckRegexp("as1e", `^[a-zA-Z]+$`))//false
}

func DoCheckRegexp (target, rexp string) bool{
	r := regexp.MustCompile(rexp)
	return r.MatchString(target)
}
