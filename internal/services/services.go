package service

type Services struct {
	Repositories Repositories
}

type Repositories interface {
}

func NewService(repositories Repositories) *Services {
	return &Services{
		Repositories: repositories,
	}
}
