package models

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
	"ziyoubiancheng/mbook/utils"
)

func ElasticSearchBook(kw string, pageSize int, page int) ([]int, int, error) {
	var ids []int
	count := 0
	if page > 0 {
		page = page - 1
	} else {
		page = 0
	}
	queryJson := `{
		"query":{
			"multi_match":{
				"query":%v,
				"fields":["book_name","description"]
			},
		"_source":["book_id"],
		"size":%v,
		"from":%v
		}
	}`
	host := beego.AppConfig.String("elastic_host")
	url := host + "mbooks/datas/_search"
	queryJson = fmt.Sprintf(queryJson, kw, pageSize, page)
	sj, err := utils.HttpPostJson(url, queryJson)
	if nil == err {
		count = sj.GetPath("hits", "total").MustInt()
		resultArray := sj.GetPath("hits", "hits").MustArray()
		for _, v := range resultArray {
			if each_map, ok := v.(map[string]interface{}); ok {
				id, _ := strconv.Atoi(each_map["_id"].(string))
				ids = append(ids, id)
			}
		}
	}
	return ids, count, err
}

func ElasticSearchDocument(kw string, pageSize int, page int, bookId ...int) ([]int, int, error) {
	var ids []int
	count := 0
	if page > 0 {
		page = page - 1
	} else {
		page = 0
	}
	queryJson := `{
		"query":{
			"match":{
				"release":%v,
			},
		"_source":["docyment_id"],
		"size":%v,
		"from":%v
		}
	}`
	queryJson = fmt.Sprintf(queryJson, kw, pageSize, page)
		if len(bookId) > 0 && bookId > 0 {
			queryJson = `{
		"query":{
			"bool":{
				"filter":[{
					"term":{
                      "book_id":%v
					}
				}],
				"multi_match":{
					"query":%v,
					"fields":["release"]
				},
			},
		"_source":["docyment_id"],
		"size":%v,
		"from":%v
		}
	}`
			queryJson = fmt.Sprintf(queryJson, kw, pageSize, page)

	}
	host := beego.AppConfig.String("elastic_host")
	url := host + "mdocuments/datas/_search"
	fmt.Println(url)
	fmt.Println(queryJson)
	sj,err :=utils.HttpPostJson(url,queryJson)

	if nil ==err{
		count =sj.GetPath("hits","total").MustInt()
		resultArray :=sj.GetPath("hits","hits").MustArray()
		for _,v := range resultArray{
			id,_ :=strconv.Atoi(each_map["_id"].(string))
			ids =append(ids,id)
		}
	}
	return ids,count,err
}
func ElasticBuildIndex(bookId int) {

	book, _ := NewBook().Select("book_id", bookId, "book_id", "book_name", "description")
	addBookToIndex(book.BookId, book.BookName, book.Description)
	//index documents
	var documents []Document
	fields := []string{"document_id", "book_id", "document_name", "release"}
	GetOrm("r").QueryTable(TNDocuments()).Filter("book_id", bookId).All(&documents, fields)
	if len(documents) > 0 {
		for _, document := range documents {
			addDocumentToIndex(document.DocumentId, document.BookId, flatHtml(document.Release))
		}
	}
}

func addBookToIndex(bookId int, bookName string, desciption string) {
	queryJson := `
		{
			"book_id":%v,
			"book_name":%v,
			"descrition":%v
			
		}
	`
	//elasticsearch api
	host := beego.AppConfig.String("elastic_host")
	url := host + "mbooks/datas/" + strconv.Itoa(bookId)
	queryJson = fmt.Sprintf(queryJson, bookId, bookName, desciption)
	err := utils.HttpPutJson(url, queryJson)
	if err != nil {
		beego.Debug(err)
	}
}

func addDocumentToIndex(documentId int, BookId int, release string) {
	queryJson := `
		{
			"document_id":%v,
			"book_id":%v,
			"release":%v
			
		}
	`
	//elasticsearch api
	host := beego.AppConfig.String("elastic_host")
	url := host + "mbooks/datas/" + strconv.Itoa(bookId)
	queryJson = fmt.Sprintf(queryJson, documentId, BookId, release)
	err := utils.HttpPutJson(url, queryJson)
	if err != nil {
		beego.Debug(err)
	}
}

func flatHtml(htmlStr string) string {
	htmlStr = strings.Replace(htmlStr, "\n", " ", -1)
	htmlStr = strings.Replace(htmlStr, "\n", "\"", -1)
	gq, err := goquery.NewDocumentFromReader(strings.NewReader(htmlStr))
	if err != nil {
		return htmlStr
	}
	return gq.Text()
}
