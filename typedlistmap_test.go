package snippets

import "testing"

func TestMapToTYPEB(t *testing.T) {
	h := NewTestHelper(t)
	a := TYPEAList{TYPEA{5}, TYPEA{4}, TYPEA{6}, TYPEA{3}}
	b := a.MapToTYPEB(func(val TYPEA) TYPEB {
		return TYPEB{val.age}
	})
	h.Assert(len(b) == 4, "Incorrect Map length")
	h.Assert(b[1].cost == 4, "Incorrect Map value")
}
