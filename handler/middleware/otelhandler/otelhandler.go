package otelhandler

import "runtime/debug"

var (
	Module = "github.com/disgoorg/disgo.handler/middleware/otelhandler"

	Version = getVersion()
)

func getVersion() string {
	bi, ok := debug.ReadBuildInfo()
	if ok {
		for _, dep := range bi.Deps {
			if dep.Path == Module {
				return dep.Version
			}
		}
	}
	return "unknown"
}

func SemVersion() string {
	return "semver:" + Version
}
