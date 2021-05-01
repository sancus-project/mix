package tree

type Path struct{}

func Compile(pattern string) (*Path, error) {
	return nil, InvalidPattern(pattern)
}
