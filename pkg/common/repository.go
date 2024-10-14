package common

type Repository[T any] interface {
	GetAll() []T
	GetAllBy(params Params, modifier int) (entities []T)
	Create(entity T) (createdEntity T)
}
