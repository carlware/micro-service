package version

import "fmt"

// Build information. Populated at build-time.
var (
	Version   string
	Revision  string
	BuildDate string
	GoVersion string
)

// Map provides the iterable version information.
var Map = map[string]string{
	"version":   Version,
	"revision":  Revision,
	"buildDate": BuildDate,
	"goVersion": GoVersion,
}

// Ver returns version as version
func Ver() string {
	return Version
}

// Full returns full composed version string
func Full() string {
	return fmt.Sprintf("%s [%s] (Go: %s, Date: %s)", Version, Revision, GoVersion, BuildDate)
}
