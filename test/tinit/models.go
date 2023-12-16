package tinit

func IgnoreTime(a, b CreatedAtOwner) {
	a.SetCreatedAt(0)
	b.SetCreatedAt(0)
}

func Contains(s []*string, e *string) bool {
	for _, a := range s {
		if *a == *e {
			return true
		}
	}
	return false

}

func ContainsAll(s []*string, e ...*string) bool {
	for _, a := range e {
		if !Contains(s, a) {
			return false
		}
	}
	return true
}

type CreatedAtOwner interface {
	GetCreatedAt() int64
	SetCreatedAt(int64)
}
