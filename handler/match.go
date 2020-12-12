package handler

import "strings"

// Matches templates
func MatchTemplate(pathElements, templateElements []string) bool {

	pathSize := len(pathElements)
	templateSize := len(templateElements)

	if pathSize == 0 && templateSize == 0 {
		return true
	}

	var max int

	if pathSize >= templateSize {
		max = templateSize
	} else {
		max = pathSize
	}

	if pathSize > max || templateSize > max {
		return false
	}

	for i := 0; i < max; i++ {
		t := templateElements[i]
		p := pathElements[i]

		if !strings.Contains(t, "{") {
			if t != p {
				return false
			}
		}
	}

	return true
}

// Match templates
func Match(path, template string) bool {
	return MatchTemplate(splitPath(path), splitPath(template))
}

func splitPath(s string) []string {
	return strings.Split(strings.Trim(s, "/"), "/")
}
