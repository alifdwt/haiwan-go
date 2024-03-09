package mapper

type Mapper struct {
	CategoryMapper CategoryMapping
}

func NewMapper() *Mapper {
	return &Mapper{
		CategoryMapper: NewCategoryMapper(),
	}
}
