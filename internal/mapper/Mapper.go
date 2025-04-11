package mapper

type Mapper[T any] interface {
	mapRow(row []string) (*T, error)
}
