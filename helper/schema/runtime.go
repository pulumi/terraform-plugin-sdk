// Copied from https://github.com/pulumi/pulumi-terraform-bridge/blob/7c411fc807e709a85170cab0add5cb3a856c0c25/pkg/tfbridge/runtime.go#L1
package schema

import (
	"runtime/debug"
	"strings"
)

// Used internally for the code to distinguish if it is running at build-time or at runtime.
type runtimeStage int

const (
	unknownStage runtimeStage = iota
	buildingProviderStage
	runningProviderStage
)

var currentRuntimeStage = guessRuntimeStage()

func guessRuntimeStage() runtimeStage {
	buildInfo, _ := debug.ReadBuildInfo()
	stage := unknownStage
	if buildInfo != nil {
		if strings.Contains(buildInfo.Path, "pulumi-tfgen") {
			stage = buildingProviderStage
		} else if strings.Contains(buildInfo.Path, "pulumi-resource") {
			stage = runningProviderStage
		}
	}
	return stage
}
