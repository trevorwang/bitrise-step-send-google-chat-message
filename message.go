package main

import "strings"

// Message is the message to be sent to Google Chat.
// It contains a slice of CardCell, which contains a Card.
type Message struct {
	Name string `json:"name,omitempty"`
	// Sender          *User        `json:"sender,omitempty"`
	// CreatedTime     string       `json:"createTime,omitempty"`
	// LastUpdatedTime string       `json:"lastUpdateTime,omitempty"`
	// DeleteTime      string       `json:"deleteTime,omitempty"`
	Text string `json:"text,omitempty"`
	// FormattedText string       `json:"formattedText,omitempty"`
	CardMessage []CardWithId `json:"cardsV2,omitempty"`
}

// User is the user who sends the message.
type User struct {
	/*Resource name for a Google Chat user.

	  Format: users/{user}. users/app can be used as an alias for the calling app bot user.

	  For human users, {user} is the same user identifier as:

	  the id for the Person in the People API. For example, users/123456789 in Chat API represents the same person as the 123456789 Person profile ID in People API.

	  the id for a user in the Admin SDK Directory API.

	  the user's email address can be used as an alias for {user} in API requests. For example, if the People API Person profile ID for user@example.com is 123456789, you can use users/user@example.com as an alias to reference users/123456789. Only the canonical resource name (for example users/123456789) will be returned from the API.
	*/
	Name        string   `json:"name,omitempty"`
	DisplayName string   `json:"displayName,omitempty"`
	DomainId    string   `json:"domainId,omitempty"`
	IsAnonymous bool     `json:"isAnonymous,omitempty"`
	Type        UserType `json:"type,omitempty"`
}

type UserType string

const (
	Human            UserType = "HUMAN"
	Bot              UserType = "BOT"
	TYPE_UNSPECIFIED UserType = "TYPE_UNSPECIFIED" // Default value for the enum. DO NOT USE.
)

// CardWithId is a cell in the message.
// It contains a Card.
type CardWithId struct {
	CardID string `json:"cardId,omitempty"`
	Card   Card   `json:"card"`
}

// Widget is the widget to be sent to Google Chat.
// It contains a TextParagraph, an Image, a DecoratedText, a ButtonList, a TextInput, a SelectionInput, a DateTimePicker, a Divider, a Grid, and a Columns.
type Widget struct {
	// Union field data can be only one of the following:
	TextParagraph *TextParagraph `json:"textParagraph,omitempty"`
	Image         *Image         `json:"image,omitempty"`
	DecoratedText *DecoratedText `json:"decoratedText,omitempty"`
	ButtonList    *ButtonList    `json:"buttonList,omitempty"`
	// TextInput      TextInput      `json:"textInput"`
	// SelectionInput SelectionInput `json:"selectionInput"`
	// DateTimePicker DateTimePicker `json:"dateTimePicker"`
	// Divider        Divider        `json:"divider"`
	// Grid           Grid           `json:"grid"`
	// Columns        Columns        `json:"columns"`
	// Union field data can be only one of the following:
}

// type SelectionInput struct {
// }
// type Columns struct{}
// type Grid struct{}

// type DateTimePicker struct{}
// type Divider struct {
// }

// type TextInput struct {
// 	Name     string `json:"name"`
// 	Label    string `json:"label"`
// 	HintText string `json:"hintText"`
// 	Value    string `json:"value"`
// }

type Image struct {
	ImageUrl string   `json:"imageUrl"`
	OnClick  *OnClick `json:"onClick,omitempty"`
	AltText  string   `json:"altText,omitempty"`
}

// Card is the card to be sent to Google Chat.
// It contains a CardHeader and a slice of CardSection.
// CardHeader is the header of the card.
// CardSection is the section of the card.
type Card struct {
	Header   *CardHeader `json:"header,omitempty"`
	Sections []Section   `json:"sections,omitempty"`
}

// Section is the section of the card.
// It contains a slice of interface{}.
type Section struct {
	Header                    string   `json:"header,omitempty"`
	Collapsible               bool     `json:"collapsible,omitempty"`
	UncollapsibleWidgetsCount int      `json:"uncollapsibleWidgetsCount,omitempty"`
	Widgets                   []Widget `json:"widgets,omitempty"`
}

// TextParagraph is the text paragraph to be sent to Google Chat.
// It contains a string. The string is the text to be sent.
type TextParagraph struct {
	Text string `json:"text"`
}

// DecoratedText is the decorated text to be sent to Google Chat.
// It contains a string. The string is the text to be sent.
// It also contains a TopLabel and a StartIcon.
// TopLabel is the label on the top of the text.
// StartIcon is the icon on the left of the text.
type DecoratedText struct {
	Icon      *Icon    `json:"icon,omitempty"`
	StartIcon *Icon    `json:"startIcon,omitempty"`
	TopLabel  *string  `json:"topLabel,omitempty"`
	Text      string   `json:"text,omitempty"`
	WrapText  bool     `json:"wrapText,omitempty"`
	OnClick   *OnClick `json:"onClick,omitempty"`

	// Union field control can be only one of the following:
	Button  *Button `json:"button,omitempty"`
	EndIcon *Icon   `json:"endIcon,omitempty"`
	// End of list of possible types for union field control.
}

// ButtonList is the button list to be sent to Google Chat.
// It contains a slice of Button.
type ButtonList struct {
	Buttons []Button `json:"buttons"`
}

// Button is the button to be sent to Google Chat.
// It contains a string. The string is the text of the button.
// It also contains an OnClick.
type Button struct {
	Text     string   `json:"text"`
	OnClick  *OnClick `json:"onClick"`
	Color    *Color   `json:"color"`
	Disabled bool     `json:"disabled"`
	Icon     *Icon    `json:"icon"`
	AltText  string   `json:"altText"`
}

type ImageType string

const (
	Square ImageType = "SQUARE"
	Circle ImageType = "CIRCLE"
)

// Icon is the icon to be sent to Google Chat.
// It contains an AltText and an ImaegType.
// AltText is the text to be shown when the icon is hovered.
// ImaegType is the type of the icon.
type Icon struct {
	AltText   string    `json:"altText,omitempty"`
	ImaegType ImageType `json:"imaegType,omitempty"`
	// Union field icons can be only one of the following:
	IconUrl   string `json:"iconUrl,omitempty"`
	KnownIcon string `json:"knownIcon,omitempty"`
	// Union field icons can be only one of the following:
}

type Color struct {
	Red   float64 `json:"red"`
	Green float64 `json:"green"`
	Blue  float64 `json:"blue"`
	Alpha float64 `json:"alpha"`
}

// OnClick is the action to be performed when the button is clicked.
// It contains an OpenLink.
type OnClick struct {
	// Union field data can be only one of the following:
	OpenLink *OpenLink `json:"openLink,omitempty"`
	Card     *Card     `json:"card,omitempty"`
	// End of list of possible types for union field data.
}

// OpenLink is the link to be opened when the button is clicked.
// It contains a string. The string is the url of the link.
type OpenLink struct {
	Url string `json:"url"`
}

// CardHeader is the header of the card.
// It contains a string. The string is the title of the header.
type CardHeader struct {
	Title     string     `json:"title"`
	Subtitle  *string    `json:"subtitle,omitempty"`
	ImageUrl  *string    `json:"imageUrl,omitempty"`
	ImageType *ImageType `json:"imageType,omitempty"`
}

type Option func(*Message)

func (c *Message) assertCardCell() {
	if c.CardMessage == nil {
		c.CardMessage = []CardWithId{
			{},
		}
	}
}

func (c *Message) assertCardSection() {
	c.assertCardCell()
	if c.CardMessage[0].Card.Sections == nil {
		c.CardMessage[0].Card.Sections = []Section{
			{},
		}
	}
}
func (c *Message) assertWidgets() {
	c.assertCardSection()
	if c.CardMessage[0].Card.Sections[0].Widgets == nil {
		c.CardMessage[0].Card.Sections[0].Widgets = []Widget{}
	}
}

func WithMessageText(text string) Option {
	return func(c *Message) {
		c.Text = text
	}
}

func WithCardHeader(title string, iconUrl string) Option {
	return func(c *Message) {
		c.assertCardCell()
		c.CardMessage[0].Card.Header = &CardHeader{
			Title:    title,
			ImageUrl: &iconUrl,
		}
	}
}

func WithCardText(text string) Option {
	return func(c *Message) {
		c.assertWidgets()
		c.CardMessage[0].Card.Sections[0].Widgets = append(c.CardMessage[0].Card.Sections[0].Widgets,
			Widget{TextParagraph: &TextParagraph{Text: text}},
		)
	}
}

func WithCardImage(url string) Option {
	return func(c *Message) {
		c.assertWidgets()
		c.CardMessage[0].Card.Sections[0].Widgets = append(c.CardMessage[0].Card.Sections[0].Widgets,
			Widget{Image: &Image{ImageUrl: url}},
		)
	}
}

func WithCardDecoratedText(text string, label string, icon *string) Option {
	return func(c *Message) {
		c.assertWidgets()
		var startIcon *Icon
		if icon != nil {
			startIcon = &Icon{
				KnownIcon: strings.TrimSpace(*icon),
			}
		}

		c.CardMessage[0].Card.Sections[0].Widgets = append(c.CardMessage[0].Card.Sections[0].Widgets, Widget{
			DecoratedText: &DecoratedText{
				Text:      text,
				TopLabel:  &label,
				StartIcon: startIcon,
			},
		})
	}
}

func WithCardButton(text string, link string) Option {
	return func(c *Message) {
		c.assertWidgets()
		c.CardMessage[0].Card.Sections[0].Widgets = append(c.CardMessage[0].Card.Sections[0].Widgets, Widget{
			ButtonList: &ButtonList{
				Buttons: []Button{
					{
						Text: text,
						OnClick: &OnClick{
							OpenLink: &OpenLink{
								Url: link,
							},
						},
					},
				},
			},
		})
	}
}

func (r ChatRunner) NewMessageWithOption(ops ...Option) *Message {
	msg := &Message{}
	for _, op := range ops {
		op(msg)
	}
	return msg
}

func (r ChatRunner) NewMessage(content Card) *Message {
	return &Message{
		CardMessage: []CardWithId{
			{
				Card: content,
			},
		},
	}
}

func (r ChatRunner) NewMessageWithText(text string, header string) *Message {
	return r.NewMessage(Card{
		Header: &CardHeader{Title: header},
		Sections: []Section{
			{
				Widgets: []Widget{
					{
						TextParagraph: &TextParagraph{Text: text},
					},
				},
			},
		},
	})
}
