package filter

import (
	"github.com/Polpetta/datlabel/result"
)

// FIXME we should return the LabelResult type here

func ByValueEqual(label *result.Label, match string) *result.Label {
	if label.Value() == match {
		return label
	}

	return nil
}

func ByValueNotEqual(label *result.Label, match string) *result.Label {
	if label.Value() != match {
		return label
	}

	return nil
}

func ByNameEqual(label *result.Label, match string) *result.Label {
	if label.Name() == match {
		return label
	}

	return nil
}

func ByNameNotEqual(label *result.Label, match string) *result.Label {
	if label.Name() != match {
		return label
	}

	return nil
}
