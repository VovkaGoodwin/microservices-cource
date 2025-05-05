package models

import (
	"github.com/golang-jwt/jwt/v5"
)

type TokenId string

type AccessTokenClaims struct {
	jwt.RegisteredClaims
	UserID string `json:"user_id"`
}

type AccessToken struct {
	id            TokenId
	unsignedToken *jwt.Token
	UserId        string
}

func NewAccessToken(id TokenId, token *jwt.Token) *AccessToken {
	return &AccessToken{
		id:            id,
		unsignedToken: token,
	}
}

func (a *AccessToken) Id() TokenId {
	return a.id
}

func (a *AccessToken) UnsignedToken() *jwt.Token {
	return a.unsignedToken
}

func (a *AccessToken) WithUserId(userId string) *AccessToken {
	a.UserId = userId
	return a
}
