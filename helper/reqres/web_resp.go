package reqres

type WebResp struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
