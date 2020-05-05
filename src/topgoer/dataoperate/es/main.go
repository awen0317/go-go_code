package main

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"golang.org/x/net/context"
	"log"
	"os"
	"reflect"
)

//33
var client *elastic.Client
var host ="http://127.0.0.1:9200/"

type Employee struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Age int `json:"age"`
	About string `json:"about"`
	Interests []string `json:"interests"`
}

func init()  {
	errorlog :=log.New(os.Stdout,"APP",log.LstdFlags)
	var err error
	client,err =elastic.NewClient(elastic.SetErrorLog(errorlog),elastic.SetURL(host))
	if err !=nil{
		panic(err)
	}
	info,code,err :=client.Ping(host).Do(context.Background())
	if err !=nil{
		panic(err)
	}
	fmt.Printf("Elasticsearch returned whith code %d and version %s\n",code,info.Version.Number)

	esversion,err :=client.ElasticsearchVersion(host)
	if err !=nil{
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s",esversion)
}

/*下面是简单的curd*/
func create(){
	e1 := Employee{
		FirstName: "Jane",
		LastName:  "Smith",
		Age:       32,
		About:     "I like to collect rock albums",
		Interests: []string{"music"},
	}
	put1 ,err := client.Index().
		Index("megocorp").
		Type("employee").
		Id("1").
		BodyJson(e1).
		Do(context.Background())
	if err !=nil{
		panic(err)
	}
	fmt.Printf("index tweet %s to index s %s ,type %s \n",put1.Id,put1.Index,put1.Type)

	//使用字符串
	e2 :=`{"first_name":"John","last_name":"Smith","age":25,"about":"I love to go rock climbing","interests":["sports","music"]}`
	put2,err := client.Index().
		Index("megacorp").
		Type("employee").
		Id("2").
		BodyJson(e2).
		Do(context.Background())
	if nil != nil{
		panic(err)
	}
	fmt.Printf("Index tweet %s to index %s ,type %s\n",put2.Id,put2.Index,put2.Type)

	e3 := `{"first_name":"Douglas","last_name":"Fir","age":35,"about":"I like to build cabinets","interests":["forestry"]}`
	put3, err := client.Index().
		Index("megacorp").
		Type("employee").
		Id("3").
		BodyJson(e3).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index s%s, type %s\n", put3.Id, put3.Index, put3.Type)
}

func delete()  {
	res,err :=client.Delete().Index("megacorp").
		Type("employee").
		Id("1").
		Do(context.Background())
	if err !=nil{
		println(err.Error())
		return
	}
	fmt.Printf("delete result %s\n",res.Result)

}
//修改
func update()  {
	res,err :=client.Update().
		Index("megacory").
		Id("2").
		Doc(map[string]interface{}{"age":88}).
		Do(context.Background())
	if err !=nil{
		println(err.Error())
	}
	fmt.Printf("update age %s \n",res.Result)
}
func gets(){
	//通过id查找
	get1,err := client.Get().Index("megacorp").Type("employee").Id("2").Do(context.Background())
	if err != nil {
		panic(err)
	}
	if get1.Found {
		fmt.Printf("Got document %s in version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)
	}
}

func query(){
	var res *elastic.SearchResult
	var err error
	res,err = client.Search("megacorp").Type("employee").Do(context.Background())
	if err !=nil{
		println(err.Error())
	}
	printEmployee(res,err)
	//字段相等
	q :=elastic.NewQueryStringQuery("last_name:Smith")
	res ,err =client.Search("megacorp").Type("employee").Query(q).Do(context.Background())
	if err:=nil{
		println(err.Error())
	}
	printEmployee(res,err)
	//条件查询,年纪大于30岁的
	boolQ :=elastic.NewBoolQuery()
	boolQ.Must(elastic.NewMatchQuery("last_name","smith"))
	boolQ.Filter(elastic.NewRngeQuery("age").Gt(30))
	res,err =client.Search("megacorp").Type("employee").Query(q).Do(context.Background())
	printEmployee(res,err)

	//短语搜索,搜索about中有rock clibming
	matchPhraseQuery :=elastic.NewMatchPhraseQuery("about","rock climbing")
	res ,err =client.Search("megacorp").Type("employee").Query(matchPhraseQuery).Do(context.Background())
	printEmployee(res,err)

	//分析iinterests
	aggs :=elastic.NewTermsAggregation().Field("interests")
	res,err =client.Search("megacorp").Type("employee").Aggregation("all_interests",aggs).Do(context.Background())
	printEmployee(res,err)
}

func list(size,page int){
	if size < 0 || page <1{
		fmt.Printf("params error")
	}
	res ,err :=client.Search("megacorp").
		Type("employee").
		Size(size).
		From((page-1)*size).
		Do(context.Background())
		printEmployee(res,err)
}

//打印查询到的employee
func printEmployee(res *elastic.SearchResult,err error){
	if err!=nil{
		println(err.Error())
		return
	}
	var typ Employee
	for _,item :=range res.Each(reflect.TypeOf(typ)){
		t := item.(Employee)
		fmt.Printf("%#v\n",t)
	}

}
func main() {
	//client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	//if err != nil {
	//	//	//Handle error
	//	panic(err)
	//}
	//fmt.Println("connet to es success")
	//p1 := Person{Name: "1mh", Age: 18, Married: false}
	//put1, err := client.Index().Index("user").BodyJson(p1).Do(context.Background())
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Indexed user %s to index %s,type %s\n",put1.Id,put1.Index,put1.Type)

}
