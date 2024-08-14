package db

type DbService struct {
	Client *dbClient.Client
}

func NewDBService(client *dbClient.Client) *DbService {
	return &DbService{Client: client}
}
