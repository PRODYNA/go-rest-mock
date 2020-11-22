package handler

import "strings"

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

func Match(path, template string) bool {
	return MatchTemplate(SplitPath(path), SplitPath(template))
}

func SplitPath(s string) []string {
	return strings.Split(strings.Trim(s, "/"), "/")
}
