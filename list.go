package goalgo

type (
	List interface {
		Build()
		Find()
		Insert()
		Delete()
		Min()
		Max()
		Prev()
		Next()
	}

	list struct {
		s []Any
	}
)
