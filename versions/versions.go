package versions

var commit string
var version string

// Commit returns the commit hash for this build (injected at build).
func Commit() string {
	return commit
}

// Version returns the version for this build (injected at build).
func Version() string {
	return version
}
