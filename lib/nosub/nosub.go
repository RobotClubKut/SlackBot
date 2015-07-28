package nosub

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/RobotClubKut/SlackBot/lib/log"
	"github.com/moovweb/gokogiri"
	"github.com/moovweb/gokogiri/xpath"
)

// Data struct is used of NoSub Data seved
type Data struct {
	Title    string
	URL      string
	Time     time.Time
	ImageURL string
}

// GetNosubUpdate is get NoSub update
func GetNosubUpdate() []Data {
	resp, err := http.Get("http://www.nosub.tv/channel/anime/on_air")
	log.TerminateAndWriteMessage(err, "can not access NoSub.")

	page, err := ioutil.ReadAll(resp.Body)
	log.TerminateAndWriteMessage(err, "can not read page.")

	doc, err := gokogiri.ParseHtml(page)
	log.TerminateAndWriteMessage(err, "can not parse html.")

	defer doc.Free()
	xps := xpath.Compile("//body/div/div[@class='content']/div[@class='margin_bottom']/div[@class='content_710 cat cat_sub']/ul/li")
	newDatas, err := doc.Root().Search(xps)
	log.Terminate(err)
	a := xpath.Compile("./a/@href")
	title := xpath.Compile("./a/@title")
	imgsrc := xpath.Compile("./a/img/@src")
	ret := make([]Data, len(newDatas))

	for i, newData := range newDatas {
		urls, err := newData.Search(a)
		log.Terminate(err)
		imgs, err := newData.Search(imgsrc)
		log.Terminate(err)
		texts, err := newData.Search(title)
		log.Terminate(err)
		for _, url := range urls {
			ret[i].URL = url.String()
		}
		for _, img := range imgs {
			ret[i].ImageURL = img.String()
		}
		for _, text := range texts {
			ret[i].Title = text.String()
		}
		ret[i].Time = time.Now()
	}
	return ret
}
