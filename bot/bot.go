package bot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/moovweb/gokogiri"
	"github.com/moovweb/gokogiri/xpath"
	"github.com/nlopes/slack"
)

type Conf struct {
	Token       string
	Room        string
	UserName    string
	UnfurlMedia bool
}

type Data struct {
	Title    string
	Url      string
	Time     time.Time
	ImageUrl string
}

type List struct {
	Data []Data
}

func writeErr(err error) {
	ioutil.WriteFile("../log/error.log", []byte(err.Error()), 0644)
	log.Fatalln(err)
}

func CreateConfigure() {
	conf := Conf{Token: "", Room: "", UserName: "", UnfurlMedia: true}

	bin, err := json.Marshal(conf)
	if err != nil {
		log.Println("can not make configure.")
		writeErr(err)
	}
	err = ioutil.WriteFile("../conf/configure-example.json", bin, 0644)
	if err != nil {
		log.Println("can not write configure.")
		writeErr(err)
	}
}

func readConfigure() *Conf {
	var conf Conf
	bin, err := ioutil.ReadFile("../conf/configure.json")
	if err != nil {
		writeErr(err)
	}
	err = json.Unmarshal(bin, &conf)
	if err != nil {
		writeErr(err)
	}
	return &conf
}

func writeDatabase(d []Data) {
	l := List{Data: d}
	bin, err := json.Marshal(l)
	if err != nil {
		log.Println("can not write database.")
		writeErr(err)
	}
	ioutil.WriteFile("../database/anime.db", bin, 0644)
}

func readDatabase() []Data {
	var l List
	bin, err := ioutil.ReadFile("../database/anime.db")
	if err != nil {
		log.Println("can not read database.")
		writeErr(err)
	}
	json.Unmarshal(bin, &l)
	ret := l.Data
	return ret
}

func InitialDatabase() {
	writeDatabase(getNoSubData())
}

func diff(data []Data) int {
	oldData := readDatabase()
	index := 0
	for i, d := range data {
		index = i
		if d.Title == oldData[0].Title {
			break
		}
	}
	if len(data) > 0 {
		writeDatabase(data)
	}
	return index - 1
}

func getNoSubData() []Data {
	resp, err := http.Get("http://www.nosub.tv/channel/anime/on_air")
	if err != nil {
		writeErr(err)
	}
	page, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		writeErr(err)
	}
	doc, err := gokogiri.ParseHtml(page)
	if err != nil {
		writeErr(err)
	}
	defer doc.Free()
	xps := xpath.Compile("//body/div/div[@class='content']/div[@class='margin_bottom']/div[@class='content_710 cat cat_sub']/ul/li")
	newDatas, _ := doc.Root().Search(xps)
	a := xpath.Compile("./a/@href")
	title := xpath.Compile("./a/@title")
	imgsrc := xpath.Compile("./a/img/@src")
	ret := make([]Data, len(newDatas))

	for i, newData := range newDatas {
		urls, _ := newData.Search(a)
		imgs, _ := newData.Search(imgsrc)
		texts, _ := newData.Search(title)
		for _, url := range urls {
			ret[i].Url = url.String()
		}
		for _, img := range imgs {
			ret[i].ImageUrl = img.String()
		}
		for _, text := range texts {
			ret[i].Title = text.String()
		}
		ret[i].Time = time.Now()
	}
	return ret
}

func createPostData(d Data) string {
	var ret string

	ret += d.ImageUrl
	ret += "\n"
	ret += d.Title
	ret += "\n"
	ret = ret + "[動画URL]: " + d.Url
	ret += "\n"
	ret += "取得日時: "
	ret += d.Time.String()
	ret += "\n"
	ret += "==\n"
	ret = ret + "[最新情報]: http://www.nosub.tv/channel/anime/on_air\n"
	return ret
}

func PostNoSubNews(room string) {
	if room == "" {
		room = readConfigure().Room
	}

	data := getNoSubData()

	api := slack.New(readConfigure().Token)
	var param slack.PostMessageParameters
	param.Username = readConfigure().UserName
	param.UnfurlMedia = readConfigure().UnfurlMedia

	channels, err := api.GetChannels(true)
	if err != nil {
		writeErr(err)
	}
	for _, channel := range channels {
		//fmt.Println(channel.Name)
		if channel.Name == room {
			fmt.Println(channel.Id)
			for index := diff(data); index >= 0; index-- {
				info := createPostData(data[index])
				api.PostMessage(channel.Id, info, param)
			}
		}
	}
}
