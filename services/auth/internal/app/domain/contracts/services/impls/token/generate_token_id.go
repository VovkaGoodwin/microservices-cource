package token

import (
	"crypto/rand"
	"fmt"
)

func (s *Service) generateTokenId(length int32) (id string, err error) {
	buf := make([]byte, length)
	_, err = rand.Read(buf)
	return fmt.Sprintf("%x", buf), err
}
