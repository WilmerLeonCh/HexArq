package player

import (
	"fmt"
	"log"
	"time"

	"github.com/HexArq/internal/domain"
)

/*
	Set creation time
	Save to repo
	Return with created id
*/

func (s Service) Create(player domain.Player) (id interface{}, err error) {
	player.CreationTime = time.Now().UTC()

	insertedId, errSave := s.Repo.Insert(player)
	if errSave != nil {
		log.Println(errSave.Error())
		return nil, fmt.Errorf("error creating player: %w", errSave)
	}

	return insertedId, nil
}
