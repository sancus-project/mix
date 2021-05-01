package tree

type Segment interface{}

func NewSegment(s string) (Segment, bool) {
	if len(s) > 0 {
		return s, true
	} else {
		return nil, false
	}
}
