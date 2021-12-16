package monitor

func stringSetUnion(first, second []string) []string {
	items := make(map[string]bool)
	out := []string{}

	for i := range first {
		items[first[i]] = true
	}

	for i := range second {
		items[second[i]] = true
	}

	for key := range items {
		out = append(out, key)
	}

	return out
}

func stringSetSubtract(first, second []string) []string {
	items := make(map[string]bool)
	out := []string{}

	for i := range second {
		items[second[i]] = true
	}

	for i := range first {
		if !(items[first[i]]) {
			out = append(out, first[i])
		}
	}

	return out
}
