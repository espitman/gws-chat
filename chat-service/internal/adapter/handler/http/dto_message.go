package http

type messageDto struct {
	ID   int `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type messageCreateRequestDto struct {
    Name string `json:"name" validate:"required"`
}

type messageGetResponseDto struct {
	ResponseDto
	Payload struct {
		Message messageDto `json:"message"`
	} `json:"payload"`
}


type messageUpdateRequestDto struct {
	Name string `json:"name" validate:"required"`
}

type messageUpdateResponseDto struct {
	ResponseDto
	Payload struct {
		Message messageDto `json:"message"`
	} `json:"payload"`
}

type messageDeleteResponseDto struct {
	ResponseDto
	Payload struct {
		Message messageDto `json:"message"`
	} `json:"payload"`
}