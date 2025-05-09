package authenticate

type RequestDto struct {
	Login    string
	Password string
}

type ResponseDto struct {
	Token string
}
