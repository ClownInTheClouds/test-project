package mapper

type CsvMapper[T any] interface {
	Mapper[T]

	fromCsv(row []string) (*T, error)
	ToCsv(object *T) []string
}
