package pathFinder

type tracker struct {
	fl     []Flight
	starts []string
	ends   []string
}

func New(fl []Flight) *tracker {
	t := &tracker{
		fl: fl,
	}
	starts := make([]string, len(t.fl))
	ends := make([]string, len(t.fl))
	t.starts = starts
	t.ends = ends
	return t
}

func (t *tracker) PathFinder() Flight {
	if len(t.fl) == 1 {
		return t.fl[0]
	}
	t.setSlices()
	return Flight{
		Source:      DiffSlice(t.starts, t.ends),
		Destination: DiffSlice(t.ends, t.starts),
	}
}

func (t *tracker) setSlices() {
	i := 0
	for _, v := range t.fl {
		t.starts[i] = v.Source
		t.ends[i] = v.Destination
		i++
	}
}

func DiffSlice(a, b []string) string {
	bMap := make(map[string]bool, len(b))
	for _, v := range b {
		bMap[v] = true
	}
	diff := ""
	for _, v := range a {
		if !bMap[v] {
			diff = v
			break
		}
	}
	return diff
}
