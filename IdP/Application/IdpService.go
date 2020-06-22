package Application

import (
	"BablarAPI/Config"
	Domain "BablarAPI/IdP/Domain"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"log"
	"time"
)

type IdPService struct {
	domainIdPService *Domain.IdPService
	logger           *log.Logger
}

func NewIdPService(domainIdPService *Domain.IdPService, logger *log.Logger) *IdPService {
	return &IdPService{domainIdPService: domainIdPService, logger: logger}
}

type UnAuthorizedError struct{}

func (u UnAuthorizedError) Error() string {
	return "Un authorized login attempt"
}

type SigningFailed struct {
	Inner error
}

func (f SigningFailed) Error() string {
	return "Failed to sign token: " + f.Inner.Error()
}

func (i IdPService) GetName(id uuid.UUID) (error, string) {
	return i.domainIdPService.GetName(id)
}

func (i IdPService) SetName(id uuid.UUID, name string) error {
	return i.domainIdPService.SetName(id, name)
}
func (i IdPService) New(username string, password string) (error, uuid.UUID) {
	return i.domainIdPService.New(username, password)
}
func (i IdPService) Login(username string, passward string) (error, string) {

	ok, identety := i.domainIdPService.Login(username, passward)
	if !ok {
		i.logger.Print(UnAuthorizedError{})
		return UnAuthorizedError{}, ""
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = identety.Id
	claims["exp"] = time.Now().Add(Config.TOKEN_TIMEOUT).Unix()
	t, err := token.SignedString([]byte(Config.SECRET))
	if err != nil {
		i.logger.Print(SigningFailed{err})
		return SigningFailed{err}, ""
	}
	return nil, t
}
