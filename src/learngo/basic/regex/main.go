package main

import (
	"fmt"
	"regexp"
)

const text  = `
My email is ccmouse@gmail.com@abc.com
email1 is abc@def.org
email2 is      kkk@qq.com
`

func main()  {
	//re := regexp.MustCompile("ccmouse@gmail.com")
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)\.([a-zA-Z0-9]+)`)
	//match :=re.FindAllString(text,-1)
	match :=re.FindAllStringSubmatch(text,-1)
	for _,m := range match{
		fmt.Println(m)
	}
}
