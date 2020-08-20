package algorithms

// Less ...
type Less interface {
	Less(other interface{}) bool
}

// Equal ...
type Equal interface {
	Equal(other interface{}) bool
}
