package times

type Dates interface {
	Output() string
}

func (ds DateSet) Output() string {
	return ds.DateSetOut()
}

func (ds DateSpan) Output() string {
	return ds.DateSpanOut()
}

func (ds DateSpanSet) Output() string {
	return ds.DateSpanSetOut()
}
