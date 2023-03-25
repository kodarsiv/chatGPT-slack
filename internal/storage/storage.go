package storage

import "kodarsiv/chatGPT-slack/internal/types"

type Storage interface {
	PutMessage(message types.Message) (bool, error)
}
