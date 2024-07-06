package player

import "github.com/HexArq/internal/domain"

func (s Service) Retrieve() (players []domain.Player, err error) {
	players, err = s.Repo.FindAll()
	if err != nil {
		return nil, err
	}
	return players, nil
}
