package repository

type Repository interface {
	UserRepository() UserRepository
}
