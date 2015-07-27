package slack

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/RobotClubKut/SlackBot/bot"
)

type Field struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

type Attachment struct {
	Fallback   string  `json:"fallback"`
	Color      string  `json:"color"`
	Pretext    string  `json:"pretext"`
	AuthorName string  `json:"author_name"`
	AuthorLink string  `json:"author_link"`
	AuthorIcon string  `json:"author_icon"`
	Title      string  `json:"title"`
	TitleLink  string  `json:"title_link"`
	Text       string  `json:"text"`
	Fields     []Field `json:"fields"`
	ImageUrl   string  `json:"image_url"`
	ThumbUrl   string  `json:"thumb_url"`
}

type Attachments struct {
	Attachments []Attachment `json:"attachments"`
}

type Conf struct {
	ApiURL string
}

/*
func NewAttachments() *Attachments {
	field := Field{
		Title: "",
		Value: "",
		Short: false,
	}
	fields := make([]Field, 1)
	fields[0] = field
	attachment := Attachment{
		Fallback:   "",
		Color:      "#36a64f",
		Pretext:    "",
		AuthorName: "anime-ifomation",
		AuthorLink: "",
		AuthorIcon: "",
		Title:      "",
		TitleLink:  "",
		Text:       "",
		Fields:     fields,
		ImageUrl:   "",
		ThumbUrl:   "",
	}
	var attachments Attachments
	attachments.Attachments = make([]Attachment, 1)
	attachments.Attachments[0] = attachment
	return &attachments
}*/

func NewAttachments(n int) *Attachments {
	field := Field{
		Title: "",
		Value: "",
		Short: false,
	}
	fields := make([]Field, 1)
	fields[0] = field
	attachment := Attachment{
		Fallback:   "",
		Color:      "#36a64f",
		Pretext:    "",
		AuthorName: "anime-ifomation",
		AuthorLink: "",
		AuthorIcon: "",
		Title:      "",
		TitleLink:  "",
		Text:       "",
		Fields:     fields,
		ImageUrl:   "",
		ThumbUrl:   "",
	}
	var attachments Attachments
	attachments.Attachments = make([]Attachment, n)
	//attachments.Attachments[0] = attachment
	for i := 0; i < n; i++ {
		attachments.Attachments[i] = attachment
	}
	return &attachments
}

func createPostString(d bot.Data) string {
	var ret string
	ret = ret + "[取得日時]: " + d.Time.String() + "\n"
	ret = ret + "[Link]: " + d.URL + "\n"
	ret = ret + "[最新情報]: " + "http://www.nosub.tv/channel/anime/on_air"
	return ret
}

func PostSlack(d []bot.Data) {
	confJs, err := ioutil.ReadFile("../conf/incoming_webhooks_configure.json")
	if err != nil {
		log.Fatalln(err)
	}
	var conf Conf
	json.Unmarshal(confJs, &conf)
	index := bot.Diff(d)
	attachments := NewAttachments(index + 1)
	for i := 0; i <= index; i++ {
		attachments.Attachments[i].Text = createPostString(d[index-i])
		attachments.Attachments[i].ImageUrl = d[index-i].ImageURL
		attachments.Attachments[i].AuthorName = d[index-i].Title
	}
	js, err := json.Marshal(attachments)
	if err != nil {
		log.Fatalln(err)
	}

	client := &http.Client{}
	data := url.Values{"payload": {string(js)}}
	resp, err := client.Post(
		conf.ApiURL,
		"application/x-www-form-urlencoded",
		strings.NewReader(data.Encode()),
	)

	if err != nil {
		log.Fatalln(err)
	}
	ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
}

func CreateConf() {
	conf := Conf{ApiURL: ""}
	js, _ := json.Marshal(conf)
	ioutil.WriteFile("../conf/incoming_webhooks_configure-example.json", js, 0644)
}
