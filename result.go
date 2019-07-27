package datlabel

// A label represent a pair of two strings: one identifying the label itself,
// the other assigning a value to the label
type Label struct {
	name  string
	value string
}

// Getter method to return the value of a label
func (l *Label) Value() string {
	return l.value
}

// Getter method to return the name of the label
func (l *Label) Name() string {
	return l.name
}

// A result is the competition of a listing or filtering operation
type Result interface {
	Labels() []Label
	Filter(filter func(label *Label) *Label) (Result, error)
}

// This is the a real implementation of the Result interface
type resultImpl struct {
	Result
	labels []Label
}

// Getter method to return a list of labels
func (r *resultImpl) Labels() []Label {
	return r.labels
}

// Allows to perform filtering operation,
// returning a new Result containing only the list of labels that have be
// returned by the filter function.
func (r *resultImpl) Filter(filter func(label *Label) *Label) (Result, error) {
	var result []Label
	for _, value := range r.labels {
		filterResult := filter(&value)
		if filterResult != nil {
			result = append(result, *filterResult)
		}
	}
	return &resultImpl{
		labels: result,
	}, nil
}

// Converts the data returned by the Docker library into a Result struct,
// that can be used to filter the labels or to perform more complex operations
func NewResult(toStructure map[string]string) Result {
	var labelList []Label
	for key, value := range toStructure {
		labelList = append(labelList, Label{
			name:  key,
			value: value,
		})
	}

	return &resultImpl{
		labels: labelList,
	}
}
