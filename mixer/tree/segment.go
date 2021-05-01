package tree

type Segment interface {
	Match(string) (Match, bool)
}

type Literal string

func (v Literal) Match(s string) (Match, bool) {
	if string(v) == s {
		return s, true
	}
	return nil, false
}

type CatchAll struct{}

func (_ CatchAll) Match(s string) (Match, bool) {
	return s, true
}

func NewSegment(s string) (Segment, bool) {
	if s == "*" {
		return CatchAll{}, true
	} else if len(s) > 0 {
		return Literal(s), true
	} else {
		return nil, false
	}
}
