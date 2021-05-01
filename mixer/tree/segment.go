package tree

func NewSegment(s string) (Segment, bool) {
	if len(s) > 0 {
		return s, true
	} else {
		return nil, false
	}
}
