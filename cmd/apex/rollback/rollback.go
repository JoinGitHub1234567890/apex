// Package rollback implements function version rollback.
package rollback

import (
	"errors"

	"github.com/spf13/cobra"

	"github.com/apex/apex/cmd/apex/root"
)

// alias.
var alias string

// name of function.
var name string

// version target.
var version string

// example output.
const example = `  Rollback a function to the previous version
  $ apex rollback foo

  Rollback canary alias
  $ apex rollback foo --alias canary

  Rollback a function to the specified version
  $ apex rollback bar -v 3`

// Command config.
var Command = &cobra.Command{
	Use:     "rollback <name>",
	Short:   "Rollback a function",
	Example: example,
	PreRunE: preRun,
	RunE:    run,
}

// Initialize.
func init() {
	root.Register(Command)

	f := Command.Flags()
	f.StringVarP(&alias, "alias", "a", "current", "Function alias")
	f.StringVarP(&version, "version", "v", "", "version to which rollback is done")
}

// PreRun errors if the name is missing.
func preRun(c *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("Missing name argument")
	}

	name = args[0]
	return nil
}

// Run command.
func run(c *cobra.Command, args []string) error {
	root.Project.Alias = alias

	if err := root.Project.LoadFunctions(name); err != nil {
		return err
	}

	fn := root.Project.Functions[0]

	if version == "" {
		return fn.Rollback()
	}

	return fn.RollbackVersion(version)
}
