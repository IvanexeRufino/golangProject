package domain

//DirectMessage estructura
type DirectMessage struct {
	From    string
	Message string
	readed  bool
}

//NewDirectMessage crea un tweet
func NewDirectMessage(from, message string) *DirectMessage {

	dm := DirectMessage{
		from,
		message,
		false,
	}

	return &dm
}
