package slack

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
}
