package tests

import "errors"

var (
	ErrLengthPhone = errors.New("длина номера телефона должна быть 11 символов")
	ErrPhoneStart  = errors.New("номер телефона должен начинаться с 7")
	ErrPhoneDigit  = errors.New("номер телефона должен состоять только из цифр")
)

func Validate(phone string) error {
	if len(phone) != 11 {
		return ErrLengthPhone
	}

	if phone[0] != '7' {
		return ErrPhoneStart
	}

	runes := []rune(phone)

	for _, v := range runes {
		if v < '0' || v > '9' {
			return ErrPhoneDigit
		}
	}

	return nil
}
