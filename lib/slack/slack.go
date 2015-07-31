package slack

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/RobotClubKut/SlackBot/lib/conf"
	"github.com/RobotClubKut/SlackBot/lib/log"
	"github.com/RobotClubKut/SlackBot/lib/nosub"
)

//slackの投稿とかの制御関係

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

//CreatePostString is template
func CreatePostString(d nosub.Data) string {
	var ret string
	ret = ret + "[取得日時]: " + d.Time + "\n"
	ret = ret + "[Link]: " + d.URL + "\n"
	ret = ret + "[最新情報]: " + "http://www.nosub.tv/channel/anime/on_air"
	return ret
}

//Post Slack
func Post(postData string) {
	configure := conf.ReadConfigure()
	client := &http.Client{}
	data := url.Values{"payload": {postData}}
	resp, err := client.Post(
		configure.IncomingSlackConf.Token,
		"application/x-www-form-urlencoded",
		strings.NewReader(data.Encode()),
	)
	log.WriteErrorLog(err)
	ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
}
