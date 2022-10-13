package service

import (
	"fmt"
	"gorm.io/gorm"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/bitly/go-simplejson"
	"github.com/spf13/cast"
	datapkg "tophub/app/task/internal/data"
	"tophub/pkg/util"
)

var (
	group       sync.WaitGroup
	methodNames = []string{"ZhiHu"}
	ZhiHuUrl    = "https://www.zhihu.com/api/v3/feed/topstory/hot-lists/total?limit=50&desktop=true"
)

type Spider struct {
	DB         *gorm.DB
	MethodName string
}

func (s *Spider) Run() {
	group.Add(len(methodNames))
	for _, v := range methodNames {
		s.MethodName = v
		go ExecMethod(s)
	}
	group.Wait()
}

func ExecMethod(s *Spider) {
	rv := reflect.ValueOf(s)
	m := rv.MethodByName(s.MethodName)
	val := m.Call(nil)

	data := val[0].Interface().([]datapkg.Data)

	if err := s.DB.Create(&data).Error; err != nil {
		fmt.Printf("%s ERROR：%s", s.MethodName, err)
	}

	group.Done()
}

func (s *Spider) GetV2EX() []map[string]interface{} {
	url := "https://www.v2ex.com/?tab=hot"
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取" + s.MethodName + "失败")
		return []map[string]interface{}{}
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Mobile Safari/537.36`)
	res, err := client.Do(request)
	if err != nil {
		fmt.Println("抓取" + s.MethodName + "失败")
		return []map[string]interface{}{}
	}
	defer res.Body.Close()
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("抓取" + s.MethodName + "失败")
		return []map[string]interface{}{}
	}
	var allData []map[string]interface{}
	document.Find(".item_title").Each(func(i int, selection *goquery.Selection) {
		url, boolUrl := selection.Find("a").Attr("href")
		text := selection.Find("a").Text()
		if boolUrl {
			allData = append(allData, map[string]interface{}{"title": text, "url": "https://www.v2ex.com" + url})
		}
	})
	return allData
}

func (s *Spider) ZhiHu() []datapkg.Data {
	req, err := http.NewRequest("GET", ZhiHuUrl, nil)
	if err != nil {
		return nil
	}

	client := &http.Client{Timeout: 3 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	j, err := simplejson.NewJson(body)
	if err != nil {
		return nil
	}

	var result []datapkg.Data

	u, _ := url.Parse(ZhiHuUrl)
	host := u.Host
	spiderTime := util.UnixSecondNow()

	data := j.Get("data")
	for i := range data.MustArray() {
		info := data.GetIndex(i)
		us := strings.Split(info.Get("target").Get("url").MustString(), "/")
		result = append(result, datapkg.Data{
			Host:        host,
			Rank:        cast.ToUint8(i + 1),
			Title:       info.Get("target").Get("title").MustString(),
			Description: info.Get("target").Get("excerpt").MustString(),
			Image:       info.Get("children").GetIndex(0).Get("thumbnail").MustString(),
			Url:         "https://www.zhihu.com/question/" + us[len(us)-1],
			Extra:       "{}",
			SpiderTime:  spiderTime,
		})
	}

	return result
}
