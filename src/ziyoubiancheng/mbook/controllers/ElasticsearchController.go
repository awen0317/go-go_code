package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"time"
	"ziyoubiancheng/mbook/models"
	"ziyoubiancheng/mbook/utils"
)

type ElasticsearchController struct {
	BaseController
}

func (c *ElasticsearchController) Search() {
	c.TplName = "search/search.html"
}

func (c *ElasticsearchController) Result() {
	wd := c.GetString("wd")
	if "" == wd {
		c.Redirect(beego.URLFor("ElasticsearchController.Search"), 302)
	}
	c.Data["Wd"] = wd
	tab := c.GetString("tab", "doc")
	c.Data["Tab"] = tab
	page, _ := c.GetInt("page", 1)
	if page < 1 {
		page = 1
	}
	size := 10
	now := time.Now()
	if "doc" == tab {
		ids, count, err := models.ElasticSearchDocument(wd, size, page)
		c.Data["totalRows"] = count
		if nil == err && len(ids) > 0 {
			c.Data["Docs"], _ = models.NewDocumentSearch().GetDocsById(ids)
		} else {
			ids, count, err := models.ElasticSearchBook(wd, size, page)
			c.Data["totalRows"] = count
			if nil == err && len(ids) > 0 {
				c.Data["Docs"], _ = models.NewBook().GetBooksByIds(ids)
			}
		}
		if count > size {
			urlSuffix :=fmt.Sprintf("&tab=%v&wd=%v",tab,wd)
			html :=utils.NewPaginations(4,count,size,page,beego.URLFor("ElasticsearchController.Result"),urlSuffix)
			c.Data["PageHtml"] = html
		}else {
			c.Data["PageHtml"] =""
		}
	}
	c.Data["SpeadTime"] = fmt.Sprintf("%.3f", time.Since(now).Seconds())
	c.TplName = "search/result.html"
}
