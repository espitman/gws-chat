package http

import (
	"github.com/espitman/gws-chat/chat-service/internal/core/domain"
)

func messageCreateRequestDtoToMessageDomainMapper(dto messageCreateRequestDto) domain.Message {
	return domain.Message{
		Name:   dto.Name,
	}
}

func messageDomainToMessageGetResponseDtoMapper(d *domain.Message) messageGetResponseDto {
	var resp messageGetResponseDto
	resp.Payload.Message = messageDto{
		ID:   d.ID,
		Name:	d.Name,
	}
	return resp
}


func messageUpdateRequestDtoToMessageDomainMapper(dto messageUpdateRequestDto) domain.Message {
	return domain.Message{
		Name: dto.Name,
	}
}

func messageDomainToMessageUpdateResponseDtoMapper(d *domain.Message) messageUpdateResponseDto {
	var resp messageUpdateResponseDto
	resp.Payload.Message = messageDto{
		ID:   d.ID,
		Name: d.Name,
	}
	return resp
}

func messageDomainToMessageDeleteResponseDtoMapper(d *domain.Message) messageDeleteResponseDto {
	var resp messageDeleteResponseDto
	resp.Payload.Message = messageDto{
		ID:   d.ID,
		Name: d.Name,
	}
	return resp
}

