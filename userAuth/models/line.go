package model

type RequestParaSystem struct {
	Timestamp string `json:"timestamp"`
	Sign      string `json:"sign"`
	Version   string `json:"version"`
}

type ReuqestParamsRegisterEvent struct {
	SystemId   string                            `json:"system_id"`
	EventId    string                            `json:"event_id"`
	EventName  string                            `json:"event_name"`
	EventParam []RequestEventParamsRegisterEvent `json:"event_param"`
	NotifyURL  string                            `json:"notify_url"`
}

type RequestEventParamsRegisterEvent struct {
	Code string `json:"code"`
	Desc string `json:"desc"`
}

type ResponseRegisteEvent struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type RequestParamAddEventParams struct {
	SystemId   string `json:"systemId"`
	EventId    string `json:"eventId"`
	EventParam string `json:"eventParam"`
}

type RequestParamsDeleteEventParams struct {
	SystemId    string   `json:"systemId"`
	EventId     string   `json:"eventId"`
	DeleteCodes []string `json:"deleteCodes"`
}

type ResponseDeleteEventParamsDelParamFailedData struct {
	SceneName   string                            `json:"sceneName"`
	SchemeDatas []ResponseDeleteEventParamsScheme `json:"schemeDatas"`
}

type RequestParamsPushMessage struct {
}
