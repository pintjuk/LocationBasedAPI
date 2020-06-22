package Domain

import (
	"github.com/google/uuid"
)

/*
Dont implement your own logon, as OWASP says. This is just a toy solution met to be swapped for a real IdP product
*/

type IdPService struct {
	idPRepsitory IdPRepsitory
}

func NewIdPService(idPRepsitory IdPRepsitory) *IdPService {
	return &IdPService{idPRepsitory: idPRepsitory}
}

func (idPService IdPService) Login(username string, password string) (ok bool, identity Principal) {
	error, identity := idPService.idPRepsitory.GetByUsername(username)
	if error != nil {
		return false, Principal{}
	}
	if identity.Password == password {
		return true, identity
	}
	return false, Principal{}
}

func (idPService IdPService) GetName(id uuid.UUID) (error, string) {
	err, identity := idPService.idPRepsitory.Get(id)
	return err, identity.Username
}

func (idPService IdPService) SetName(id uuid.UUID, name string) error {
	return idPService.idPRepsitory.UpdateUsername(id, name)
}

func (idPService IdPService) New(username string, password string) (error, uuid.UUID) {
	newId := uuid.New()
	return idPService.idPRepsitory.Add(Principal{
		Id:       newId,
		Username: username,
		Password: password,
	}), newId
}
