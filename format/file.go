package format

import (
	"sort"
	"strconv"
	"strings"
)

type filePart int

const (
	preImport filePart = iota
	importSection
	postImport
)

func File(file string, local string) string {
	contents, ok := extractImports(file)
	if !ok || len(contents[importSection]) == 0 {
		return file
	}
	contents[importSection] = formatImports(contents[importSection], local)
	contents[importSection] = append([]string{"import ("}, contents[importSection]...)
	contents[importSection] = append(contents[importSection], ")")

	var allLines []string
	for _, c := range contents {
		allLines = append(allLines, c...)
	}

	return strings.Join(allLines, "\n")
}

// extractImports splits file into three groups: pre-imports, imports (first group) and post-imports.
func extractImports(s string) ([3][]string, bool) {
	var result [3][]string
	lines := strings.Split(s, "\n")
	phase := preImport

	nextPart := func(line string, previousPhase filePart) filePart {
		if previousPhase == preImport && line == "import (" {
			return importSection
		}
		if previousPhase == importSection && line == ")" {
			return postImport
		}
		return previousPhase
	}

	for _, line := range lines {
		newPhase := nextPart(line, phase)
		if newPhase == phase {
			result[phase] = append(result[phase], line)
		}
		phase = newPhase
	}

	return result, phase == postImport
}

func formatImports(imports []string, local string) []string {
	group := func(s string) int {
		path := importPath(s)

		if !strings.Contains(path, ".") {
			return 0
		}
		if local != "" && strings.HasPrefix(path, local) {
			return 2
		}
		return 1
	}

	groups := [3][]string{}

	for _, imp := range imports {
		if strings.TrimSpace(imp) == "" {
			continue
		}
		groups[group(imp)] = append(groups[group(imp)], imp)
	}

	var result []string
	needEmptyLine := false

	for _, g := range groups {
		sort.Slice(g, func(i, j int) bool {
			a, b := importPath(g[i]), importPath(g[j])
			return a < b
		})

		if len(g) > 0 {
			if needEmptyLine {
				result = append(result, "")
			}
			result = append(result, g...)
			needEmptyLine = true
		}
	}

	return result
}

func importPath(s string) string {
	s = strings.TrimSpace(s)
	groups := strings.Split(s, " ")
	path := groups[len(groups)-1]
	path, _ = strconv.Unquote(path)
	return path
}
