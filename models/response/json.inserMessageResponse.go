package response

type MessageResponse struct {
	Response
	Id uint64
}

func NewResponseMessage(result uint8, message string, id uint64) *MessageResponse {
	return &MessageResponse{
		Response: Response{
			Result:  result,
			Message: message,
		},
		Id: id,
	}
}
