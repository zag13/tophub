package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/bitly/go-simplejson"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"tophub/app/topdata/dal/model"
)

var (
	group       sync.WaitGroup
	methodNames = []string{"ZhiHu", "WeiBo"}

	ZhiHuUrl = "https://www.zhihu.com/api/v3/feed/topstory/hot-lists/total?limit=50&desktop=true"
	WeiBoUrl = "https://s.weibo.com/top/summary"
	// WeiBoUrl2 = "https://m.weibo.cn/api/container/getIndex?containerid=106003type%3D25%26t%3D3%26disable_hot%3D1%26filter_type%3Drealtimehot"
	GitHubUrl = "https://github.com/trending"
)

type Spider struct {
	DB         *gorm.DB
	MethodName string
}

func (s Spider) Run() {
	group.Add(len(methodNames))
	for _, v := range methodNames {
		s.MethodName = v
		go ExecMethod(s)
	}
	group.Wait()
}

func ExecMethod(s Spider) {
	defer group.Done()
	rv := reflect.ValueOf(s)
	m := rv.MethodByName(s.MethodName)
	val := m.Call(nil)

	data, err := val[0].Interface().([]model.Topdatum), val[1].Interface()
	if err != nil || data == nil {
		fmt.Printf("%s ERROR: %s, DATA: %v\n", s.MethodName, err, data)
		return
	}

	if err := s.DB.Create(&data).Error; err != nil {
		fmt.Printf("DB CREATE ERROR：%s\n", err)
		return
	}
}

func (s Spider) V2EX() []map[string]interface{} {
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

func (s Spider) ZhiHu() (result []model.Topdatum, err error) {
	req, err := http.NewRequest("GET", ZhiHuUrl, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{Timeout: 3 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	j, err := simplejson.NewJson(body)
	if err != nil {
		return nil, err
	}

	data := j.Get("data")
	for i := range data.MustArray() {
		info := data.GetIndex(i)
		us := strings.Split(info.Get("target").Get("url").MustString(), "/")

		extra, err := json.Marshal(map[string]string{
			"description": info.Get("target").Get("excerpt").MustString(),
			"image":       info.Get("children").GetIndex(0).Get("thumbnail").MustString(),
		})
		if err != nil {
			return nil, err
		}

		result = append(result, model.Topdatum{
			Tab:      "zhihu",
			Position: cast.ToInt64(i + 1),
			Title:    info.Get("target").Get("title").MustString(),
			URL:      "https://www.zhihu.com/question/" + us[len(us)-1],
			Extra:    cast.ToString(extra),
		})
	}
	return
}

func (s Spider) WeiBo() (result []model.Topdatum, err error) {
	req, err := http.NewRequest("GET", WeiBoUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36")
	req.Header.Set("cookie", "SUB=_2AkMV0R1pdcPxrAVTn_EQyWvgZY5H-jymBHSfAn7uJhIyOhhu7nECqSVutBF-XDXcz_XKfP8ConVYk8o_32dN56GH; SUBP=0033WrSXqPxfM72wWs9jqgMF555t")

	client := &http.Client{Timeout: 3 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	doc.Find("#pl_top_realtimehot table tbody tr").Each(func(i int, selection *goquery.Selection) {
		pos := cast.ToUint8(selection.Find(".td-01").Text())
		if pos == 0 {
			return
		}

		extra, err := json.Marshal(map[string]string{})
		if err != nil {
			return
		}

		u, _ := selection.Find(".td-02 a").Attr("href")
		result = append(result, model.Topdatum{
			Tab:      "weibo",
			Position: int64(pos),
			Title:    selection.Find(".td-02 a").Text(),
			URL:      "https://s.weibo.com" + u,
			Extra:    cast.ToString(extra),
		})
	})
	return
}

func (s Spider) GitHub() (result []model.Topdatum, err error) {
	_, err = http.NewRequest("GET", GitHubUrl, nil)
	if err != nil {
		return nil, err
	}
	return
}

func (s Spider) Bilibili() (result []model.Topdatum, err error) {
	return
}
