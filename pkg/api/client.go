package api

type ClientService interface{}

type ClientRepository interface{}

type clientService struct {
	storage ClientRepository
}

func NewUserService(clientRepo ClientRepository) ClientService {
	return &clientService{
		storage: clientRepo,
	}
}
