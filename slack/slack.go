package slack

type Field struct {
	Title string `title`
	Value string `value`
	Short bool   `short`
}

type Attachment struct {
	Fallback   string  `fallback`
	Color      string  `color`
	Pretext    string  `pretext`
	AuthorName string  `author_name`
	AuthorLink string  `author_link`
	AuthorIcon string  `author_icon`
	Title      string  `title`
	TitleLink  string  `title_link`
	Text       string  `text`
	Fields     []Field `fields`
	ImageUrl   string  `image_url`
	ThumbUrl   string  `thumb_url`
}

type Attachments struct {
	Attachments []Attachment
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
		AuthorName: "",
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
