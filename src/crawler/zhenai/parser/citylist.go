package parser

import (
	"crawler/engine"
	"regexp"
)
const cityListRe =  `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`
func PaseCityList(contents []byte) engine.ParseResult {
	re :=regexp.MustCompile(cityListRe)
	//matches := re.FindAll(contents,-1)
	matches := re.FindAllSubmatch(contents,-1)
	result := engine.ParseResult{}

	for _,m := range matches{
		//for _,subMatch := range  m{
		//fmt.Printf("City:%s URL:%s",m[2],m[1])
		result.Items =append(result.Items,m[2])
		result.Requests = append(result.Requests,engine.Request{
			Url: string(m[1]),
		})
	}
	return result
}
