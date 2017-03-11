package snippets

func (x TYPEAList) MapToTYPEB(fn func(TYPEA) TYPEB) []TYPEB {
	y := make([]TYPEB, 0)
	for _, val := range x {
		y = append(y, fn(val))
	}
	return y
}
