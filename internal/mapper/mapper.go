package mapper

type Mapper struct {
	CategoryMapper CategoryMapping
	ProductMapper  ProductMapping
	OrderMapper    OrderMapping
	CartMapper     CartMapping
}

func NewMapper() *Mapper {
	return &Mapper{
		CategoryMapper: NewCategoryMapper(),
		ProductMapper:  NewProductMapper(),
		OrderMapper:    NewOrderMapper(),
		CartMapper:     NewCartMapper(),
	}
}
