// Command toolbox-web is the standalone binary form of the codefly
// web toolbox plugin. Loaded via the standard agent loader
// (core/agents/manager.Load); registers a Toolbox server through
// agents.Serve.
//
// Configuration:
//
//	CODEFLY_TOOLBOX_VERSION         — Identity version. Default "0.0.0-dev".
//	CODEFLY_TOOLBOX_ALLOWED_DOMAINS — comma-separated allowlist. The
//	                                  toolbox starts with no allowed
//	                                  domains (deny by default) if
//	                                  unset.
package main

import (
	"github.com/codefly-dev/core/agents"
	coretoolbox "github.com/codefly-dev/core/toolbox"
	web "github.com/codefly-dev/toolbox-web"
)

func main() {
	server := web.New(coretoolbox.Version()).WithAllowedDomains(
		coretoolbox.EnvironmentList("CODEFLY_TOOLBOX_ALLOWED_DOMAINS")...,
	)
	agents.ServeToolbox(server)
}
