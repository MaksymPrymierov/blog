package session

type Session struct {
	username string
}

type Session struct {
	data map[string]*sessionData
}

func NewSession() *Session {
	s := new(Session)

	s.data = make(map[string]*sessionData)

	return s
}
