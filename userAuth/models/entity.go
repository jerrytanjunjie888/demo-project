package model

type Response struct {
	Code int32 `json:"code"`

	Message string `json:"message,omitempty"`

	Data string `json:"message,omitempty"`
}

type request struct {
	SystemId string `json:"system_id"`

	SequenceId string `json:"sequence_id"`

	CountryCode string `json:"country_code"`

	BizData map[string]interface{} `json:"biz_data"`
}

type BizDataPushMessage struct {
	EventId string `json:"event_id"`

	EventName string `json:"event_name"`

	Channel string `json:"channel"`

	LanguageCode string `json:"language_code"`

	ContaictId string `json:"contact_id"`

	MobileCountryCode string `json:"mobile_country_code,omitempty`

	MobileNo string `json:"mobile_no,omitempty"`

	Email string `json:"email,omitempty"`

	WeChatOpenId string `json:"we_chat_open_id,omitempty"`

	FirstName string `json:"first_name"`

	LastName string `json:"last_name"`
}

type BizDataRegisterEvent struct {
	EventId string `json:"event_id"`

	EventName string `json:"event_name"`

	NotityUrl string `json:"notify_url"`

	EventParam []string `json:"event_param"`
}

type BizDataEventParamRegisterEvent struct {
	Code string `json:"code"`

	Desc string `json:"desc"`
}
