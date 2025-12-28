package pkg

type Payload struct {
	Error   bool `json:"error"`
	Message any  `json:"message"`
}
