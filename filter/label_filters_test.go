// +build unit

package filter

import (
	"github.com/Polpetta/datlabel/result"
	"testing"
)

func createDummyLabelList() result.LabelResult {
	m := make(map[string]string)
	m["com.label.test1"] = "dummyValue1"
	m["com.label.test2"] = "dummyValue2"

	return result.NewLabelResult(m)
}

func TestByValueEqual(t *testing.T) {
	labelList := createDummyLabelList()

	nonNullValue := ByValueEqual(&labelList.Labels()[0], "dummyValue1")
	if nonNullValue == nil {
		t.Fatalf("%s value is erroneously filtered", "dummyValue1")
	}
	if nonNullValue.Value() != "dummyValue1" {
		t.Fatalf("Found label expected to have value %s, but found %s",
			"dummyValue1",
			nonNullValue.Name())
	}

	nullValue := ByValueEqual(&labelList.Labels()[1], "dummyValue1")
	if nullValue != nil {
		t.Fatalf("Expected to get back nil value, but got %s label"+
			"instead",
			nullValue.Name())
	}
}

func TestByValueNotEqual(t *testing.T) {
	labelList := createDummyLabelList()

	nullValue := ByValueNotEqual(&labelList.Labels()[0], "dummyValue1")
	if nullValue != nil {
		t.Fatalf("Expected to get back nil value, but got %s label"+
			"instead",
			nullValue.Name())
	}

	nonNullValue := ByValueNotEqual(&labelList.Labels()[1], "dummyValue1")
	if nonNullValue == nil {
		t.Fatalf("%s value is erroneously filtered", "dummyValue2")
	}
	if nonNullValue.Value() != "dummyValue2" {
		t.Fatalf("Found label expected to have value %s, but found %s",
			"dummyValue1",
			nonNullValue.Name())
	}
}

func TestByNameEqual(t *testing.T) {
	labelList := createDummyLabelList()

	nonNullValue := ByNameEqual(&labelList.Labels()[0], "com.label.test1")
	if nonNullValue == nil {
		t.Fatalf("%s label is erroneously filtered", "com.label.test1")
	}
	if nonNullValue.Name() != "com.label.test1" {
		t.Fatalf("Found label expected to have value %s, but found %s",
			"com.label.test1",
			nonNullValue.Name())
	}

	nullValue := ByNameEqual(&labelList.Labels()[1], "com.label.test1")
	if nullValue != nil {
		t.Fatalf("Expected to get back null value, but got %s label "+
			"instead",
			nullValue.Name())
	}
}

func TestByNameNotEqual(t *testing.T) {
	labelList := createDummyLabelList()

	nullValue := ByNameNotEqual(&labelList.Labels()[0], "com.label.test1")
	if nullValue != nil {
		t.Fatalf("Expected to get back nil value, but got %s label "+
			"instead",
			nullValue.Name())
	}

	nonNullValue := ByNameNotEqual(&labelList.Labels()[1], "com.label.test1")
	if nonNullValue == nil {
		t.Fatalf("%s label is erroneously filtered", "com.label.test2")
	}
	if nonNullValue.Name() != "com.label.test2" {
		t.Fatalf("Found label expected to have value %s, but found %s",
			"dummyValue1",
			nonNullValue.Name())
	}
}
