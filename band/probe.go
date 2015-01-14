package band

type Probe struct{}

func (p Probe) Run() {
	currentStatus = []Status{Status{0, 0}, Status{1, 0}, Status{2, 0}, Status{3, 0}, Status{4, 1}}
}

func (p Probe) Interval() string {
	return "300m"
}
