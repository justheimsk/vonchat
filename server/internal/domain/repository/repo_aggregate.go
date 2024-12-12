package domain_repo

type RepositoryAggregate struct {
	Health HealthRepository
	User   UserRepository
	Auth   AuthRepository
}
