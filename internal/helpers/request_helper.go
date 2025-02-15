package helpers

func IsPathInIgnoreList(path string, ignorePaths []string) bool {
	for _, ignorePath := range ignorePaths {
		if path == ignorePath {
			return true
		}
	}
	return false
}