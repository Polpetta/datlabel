package utils

func FilterLabelsByString(
	listOfLabels map[string]string,
	filter string) []string {
	var result []string
	for key, value := range listOfLabels {
		if value != filter {
			result = append(result, key)
		}
	}

	return result
}
