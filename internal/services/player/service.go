package player

import "github.com/HexArq/internal/ports"

type Service struct {
	Repo ports.PlayerRepository
}

var _  ports.PlayerService = &Service{}
