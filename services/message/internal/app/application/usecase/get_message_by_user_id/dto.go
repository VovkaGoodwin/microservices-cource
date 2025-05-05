package get_message_by_user_id

type Message struct {
	Id   string `json:"id"`
	Text string `json:"text"`
}

type GetMessageByUserIdRequest struct {
	UserId string `json:"user_id"`
}

type GetMessageByUserIdResponse struct {
	Messages []Message `json:"messages"`
}
