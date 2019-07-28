// +build unit

package datlabel

import (
	"testing"
)

// This test checks that the same number of labels are returned after being
// inserted into a resultImpl struct
func TestItShouldReturnRightNumberOfLabels(t *testing.T) {
	label1 := &Label{
		name:  "something1",
		value: "abc",
	}
	label2 := &Label{
		name:  "something2",
		value: "efg",
	}
	labels := []Label{
		*label1,
		*label2,
	}

	resultImpl := &resultImpl{
		labels: labels,
	}

	if len(resultImpl.Labels()) != len(labels) {
		t.Fatalf("Size of labels is %d, but got %d from Result",
			len(labels),
			len(resultImpl.Labels()))
	}
}

// This test checks that, given two labels,
// they are retrieved correctly after being inserted into a resultImpl struct
func TestItShouldReturnRightLabels(t *testing.T) {
	label1 := &Label{
		name:  "something1",
		value: "abc",
	}
	label2 := &Label{
		name:  " something2",
		value: "efg",
	}
	labels := []Label{
		*label1,
		*label2,
	}

	resultImpl := &resultImpl{
		labels: labels,
	}

	if resultImpl.Labels()[0] != *label1 {
		t.Fatalf("Label from result value is %s, %s, but label1 was %s %s",
			resultImpl.Labels()[0].Name(),
			resultImpl.Labels()[0].Value(),
			label1.Name(),
			label1.Value())
	}

	t.Logf("resultImpl[0]: %s %s, label1: %s %s",
		resultImpl.Labels()[0].Name(),
		resultImpl.Labels()[0].Value(),
		label1.Name(),
		label1.Value())

	if resultImpl.Labels()[1] != *label2 {
		t.Fatalf("Label from result value is %s, %s, but label2 was %s %s",
			resultImpl.Labels()[0].Name(),
			resultImpl.Labels()[0].Value(),
			label2.Name(),
			label2.Value())
	}

	t.Logf("resultImpl[1]: %s %s, label1: %s %s",
		resultImpl.Labels()[1].Name(),
		resultImpl.Labels()[1].Value(),
		label2.Name(),
		label2.Value())
}

// This test checks the NewResult function,
// assuring that the conversion between Labels struct and map[string]string
// is correctly performed
func TestNewResult(t *testing.T) {
	stringsOfLabels := make(map[string]string)

	stringsOfLabels["something"] = "abc"
	stringsOfLabels["something2"] = "efg"
	stringsOfLabels["something3"] = "lmn"

	result := NewResult(stringsOfLabels)

	convertedLabel1 := result.Labels()[0]
	convertedLabel2 := result.Labels()[1]
	convertedLabel3 := result.Labels()[2]

	if convertedLabel1.Name() == "something" &&
		convertedLabel1.Value() != stringsOfLabels["something"] {
		t.Fatalf("convertedLabel1 value is %s, "+
			"stringOfLabels['something'] is %s",
			convertedLabel1.Value(),
			stringsOfLabels["something"])
	}

	if convertedLabel2.Name() == "something2" &&
		convertedLabel2.Value() != stringsOfLabels["something2"] {
		t.Fatalf("convertedLabel2 value is %s, "+
			"stringOfLabels['something2'] is %s",
			convertedLabel2.Value(),
			stringsOfLabels["something2"])
	}

	if convertedLabel3.Name() == "something3" &&
		convertedLabel3.Value() != stringsOfLabels["something3"] {
		t.Fatalf("convertedLabel3 value is %s, "+
			"stringOfLabels['something3'] is %s",
			convertedLabel3.Value(),
			stringsOfLabels["something3"])
	}
}
