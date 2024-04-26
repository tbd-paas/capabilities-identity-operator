/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package commands

import (
	"github.com/spf13/cobra"

	// common imports for subcommands
	cmdgenerate "github.com/tbd-paas/capabilities-certificates-operator/cmd/identityctl/commands/generate"
	cmdinit "github.com/tbd-paas/capabilities-certificates-operator/cmd/identityctl/commands/init"
	cmdversion "github.com/tbd-paas/capabilities-certificates-operator/cmd/identityctl/commands/version"

	// specific imports for workloads
	generatecapabilities "github.com/tbd-paas/capabilities-certificates-operator/cmd/identityctl/commands/generate/capabilities"
	initcapabilities "github.com/tbd-paas/capabilities-certificates-operator/cmd/identityctl/commands/init/capabilities"
	versioncapabilities "github.com/tbd-paas/capabilities-certificates-operator/cmd/identityctl/commands/version/capabilities"
	//+kubebuilder:scaffold:operator-builder:subcommands:imports
)

// IdentityctlCommand represents the base command when called without any subcommands.
type IdentityctlCommand struct {
	*cobra.Command
}

// NewIdentityctlCommand returns an instance of the IdentityctlCommand.
func NewIdentityctlCommand() *IdentityctlCommand {
	c := &IdentityctlCommand{
		Command: &cobra.Command{
			Use:   "identityctl",
			Short: "Manage the identity capability",
			Long:  "Manage the identity capability",
		},
	}

	c.addSubCommands()

	return c
}

// Run represents the main entry point into the command
// This is called by main.main() to execute the root command.
func (c *IdentityctlCommand) Run() {
	cobra.CheckErr(c.Execute())
}

func (c *IdentityctlCommand) newInitSubCommand() {
	parentCommand := cmdinit.GetParent(c.Command)
	_ = parentCommand

	// add the init subcommands
	initcapabilities.NewIdentityCapabilitySubCommand(parentCommand)
	//+kubebuilder:scaffold:operator-builder:subcommands:init
}

func (c *IdentityctlCommand) newGenerateSubCommand() {
	parentCommand := cmdgenerate.GetParent(c.Command)
	_ = parentCommand

	// add the generate subcommands
	generatecapabilities.NewIdentityCapabilitySubCommand(parentCommand)
	//+kubebuilder:scaffold:operator-builder:subcommands:generate
}

func (c *IdentityctlCommand) newVersionSubCommand() {
	parentCommand := cmdversion.GetParent(c.Command)
	_ = parentCommand

	// add the version subcommands
	versioncapabilities.NewIdentityCapabilitySubCommand(parentCommand)
	//+kubebuilder:scaffold:operator-builder:subcommands:version
}

// addSubCommands adds any additional subCommands to the root command.
func (c *IdentityctlCommand) addSubCommands() {
	c.newInitSubCommand()
	c.newGenerateSubCommand()
	c.newVersionSubCommand()
}
