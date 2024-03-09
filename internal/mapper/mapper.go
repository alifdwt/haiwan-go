package mapper

type Mapper struct {
	CategoryMapper CategoryMapping
	ProductMapper  ProductMapping
}

func NewMapper() *Mapper {
	return &Mapper{
		CategoryMapper: NewCategoryMapper(),
		ProductMapper:  NewProductMapper(),
	}
}
