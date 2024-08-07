package ports

import "github.com/HexArq/internal/domain"

type PlayerService interface {
	Create(player domain.Player) (id interface{}, err error)
	Retrieve() (players []domain.Player, err error)
}

type PlayerRepository interface {
	Insert(player domain.Player) (id interface{}, err error)
	FindAll() (players []domain.Player, err error)
}
