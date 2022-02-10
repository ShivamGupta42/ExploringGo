package structs

type StructTest struct {
	X int
	Y int
}

func (s *StructTest) SUM() int {
	return (*s).X + (*s).Y
}
