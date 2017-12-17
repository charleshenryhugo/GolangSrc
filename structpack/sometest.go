package structpack

type Day struct {
	Int int
}

var (
	WEEK = []string{"MON", "TUE", "WED", "THU", "FRI", "SAT", "SUN"}
)

const (
	MON = iota
	TUE
	WED
	THU
	FRI
	SAT
	SUN
)

func (day *Day) String() string {
	if day.Int > 7 {
		return "failure"
	}
	return WEEK[day.Int-1]
}
