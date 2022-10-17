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
	"tophub/pkg/timez"
)

var (
	group       sync.WaitGroup
	methodNames = []string{"WeiBo"}
	ZhiHuUrl    = "https://www.zhihu.com/api/v3/feed/topstory/hot-lists/total?limit=50&desktop=true"
	WeiBoUrl    = "https://s.weibo.com/top/summary"
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
	defer group.Done()
	rv := reflect.ValueOf(s)
	m := rv.MethodByName(s.MethodName)
	val := m.Call(nil)

	data, err := val[0].Interface().([]datapkg.Data), val[1].Interface()
	if err != nil || data == nil {
		fmt.Printf("%s ERROR: %s, DATA: %v\n", s.MethodName, err, data)
		return
	}

	if err := s.DB.Create(&data).Error; err != nil {
		fmt.Printf("DB CREATE ERROR：%s\n", err)
		return
	}
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

func (s *Spider) ZhiHu() (result []datapkg.Data, err error) {
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

	u, _ := url.Parse(ZhiHuUrl)
	host := u.Host
	spiderTime := timez.UnixSecondNow()

	data := j.Get("data")
	for i := range data.MustArray() {
		info := data.GetIndex(i)
		us := strings.Split(info.Get("target").Get("url").MustString(), "/")
		result = append(result, datapkg.Data{
			Host:        host,
			Position:    cast.ToUint8(i + 1),
			Title:       info.Get("target").Get("title").MustString(),
			Description: info.Get("target").Get("excerpt").MustString(),
			Url:         "https://www.zhihu.com/question/" + us[len(us)-1],
			Image:       info.Get("children").GetIndex(0).Get("thumbnail").MustString(),
			Extra:       "{}",
			SpiderTime:  spiderTime,
		})
	}
	return
}

func (s *Spider) WeiBo() (result []datapkg.Data, err error) {
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

	u, _ := url.Parse(WeiBoUrl)
	host := u.Host
	spiderTime := timez.UnixSecondNow()

	doc.Find("#pl_top_realtimehot table tbody tr").Each(func(i int, selection *goquery.Selection) {
		pos := cast.ToUint8(selection.Find(".td-01").Text())
		if pos == 0 {
			return
		}

		u, _ := selection.Find(".td-02 a").Attr("href")
		result = append(result, datapkg.Data{
			Host:        host,
			Position:    pos,
			Title:       selection.Find(".td-02 a").Text(),
			Description: "",
			Url:         "https://s.weibo.com" + u,
			Image:       "",
			Extra:       "{}",
			SpiderTime:  spiderTime,
		})
	})
	return
}
