package Domain

import (
	"github.com/google/uuid"
)

type PlayerService struct {
	IdPService       IdPService
	playerRepository PlayerRepository
}

func NewPlayerService(idPService IdPService, playerRepository PlayerRepository) *PlayerService {
	return &PlayerService{IdPService: idPService, playerRepository: playerRepository}
}

func (playerService PlayerService) GetPlayer(id uuid.UUID) (error, Player) {
	err, player := playerService.playerRepository.Get(id)
	if err == nil {
		return nil, player
	}
	// if principal exists in IdP create a new player
	err2, _ := playerService.makePlayerFromPrincipal(id)
	if err2 != nil {
		return err2, Player{}
	}
	return playerService.playerRepository.Get(id)
}

func (playerService PlayerService) makePlayerFromPrincipal(id uuid.UUID) (error, Player) {
	err, name := playerService.IdPService.GetName(id)
	if err != nil {
		return err, Player{}
	}
	err = playerService.playerRepository.Add(Player{
		Id:           id,
		CashedName:   name,
		LastLocation: Location{},
		Score:        0,
	})
	return err, Player{}
}

func (playerService PlayerService) UpdateName(id uuid.UUID, name string) error {
	err := playerService.IdPService.SetName(id, name)
	if err != nil {
		return err
	}

	return playerService.playerRepository.Update(UpdatePlayerCommand{
		Id:         id,
		CashedName: &name,
	})
}

func (playerService PlayerService) SendLocation(id uuid.UUID, location Location) error {
	//just store for now
	return playerService.playerRepository.Update(UpdatePlayerCommand{
		Id:       id,
		Location: &location,
	})
}

func (playerService PlayerService) New(name string, password string) error {
	errNewPrincipal, id := playerService.IdPService.New(name, password)
	if errNewPrincipal != nil {
		return errNewPrincipal
	}
	errMakeplayer, _ := playerService.makePlayerFromPrincipal(id)
	return errMakeplayer
}
