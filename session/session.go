package session

import (
	"github.com/connor41/blog/utils"
)

/* Data stored session */
type sessionData struct {
	Username string
}

type Session struct {
	data map[string]*sessionData
}

/* Init session */
func NewSession() *Session {
	s := new(Session)

	s.data = make(map[string]*sessionData)

	return s
}

/* Init session on username */
func (s *Session) Init(username string) string {
	sessionId := utils.GenerateId()

	data := &sessionData{Username: username}
	s.data[sessionId] = data

	return sessionId
}

/* Get username on id session */
func (s *Session) Get(sessionId string) string {
	data := s.data[sessionId]

	if data == nil {
		return ""
	}

	return data.Username
}

/* Delete session from memory */
func (s *Session) Delete(sessionId string) {
	delete(s.data, sessionId)
}
