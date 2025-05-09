package send_message

type Request struct {
	UserId string `json:"user_id"`
	Text   string `json:"message"`
}

type Response struct {
	Success bool
}
