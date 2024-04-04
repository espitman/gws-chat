package http

import "errors"

func (h *MessageHandler) createValidator(reqDto *messageCreateRequestDto) error {
	if err := h.validate.Struct(reqDto); err != nil {
		return err
	}
	return nil
}

func (h *MessageHandler) getValidator(ID int) error {
	err := h.validate.Var(ID, "required")
	if err != nil {
		return errors.New("ID is required")
	}
	return nil
}

func (h *MessageHandler) updateValidator(ID int, reqDto *messageUpdateRequestDto) error {
	err := h.validate.Var(ID, "required")
	if err != nil {
		return errors.New("ID is required")
	}
	if err := h.validate.Struct(reqDto); err != nil {
		return err
	}
	return nil
}

func (h *MessageHandler) deleteValidator(ID int) error {
	err := h.validate.Var(ID, "required")
	if err != nil {
		return errors.New("ID is required")
	}
	return nil
}


