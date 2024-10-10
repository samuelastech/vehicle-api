package common

type Repository[T any] interface {
	GetAll() []T
	Create(entity T) (createdEntity T)
}
