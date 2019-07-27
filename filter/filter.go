package filter

import (
	"github.com/polpetta/datlabel"
)

// FIXME we should return the Result type here

func ByValueEqual(label *datlabel.Label, match string) *datlabel.Label {
	if label.Value() == match {
		return label
	}

	return nil
}

func ByValueNotEqual(label *datlabel.Label, match string) *datlabel.Label {
	if label.Value() != match {
		return label
	}

	return nil
}

func ByNameEqual(label *datlabel.Label, match string) *datlabel.Label {
	if label.Name() == match {
		return label
	}

	return nil
}

func ByNameNotEqual(label *datlabel.Label, match string) *datlabel.Label {
	if label.Name() != match {
		return label
	}

	return nil
}
