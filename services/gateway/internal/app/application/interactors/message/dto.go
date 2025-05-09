package message

type Message struct {
	Id   string `json:"id"`
	Text string `json:"text"`
}

type GetMessagesByUSerIdResponse struct {
	Messages []Message `json:"messages"`
}

type SendMessageRequest struct {
	UserId string
	Text   string
}
type SendMessageResponse struct {
	Success bool `json:"success"`
}
