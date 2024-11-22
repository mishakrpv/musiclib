package cmd

import "github.com/mishakrpv/musiclib/pkg/config"

type CmdConfiguration struct {
	config.Configuration
}

func NewCmdConfiguration() *CmdConfiguration {
	return &CmdConfiguration{
		Configuration: config.Configuration{},
	}
}
