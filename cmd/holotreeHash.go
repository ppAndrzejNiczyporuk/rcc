package cmd

import (
	"fmt"

	"github.com/robocorp/rcc/common"
	"github.com/robocorp/rcc/htfs"
	"github.com/robocorp/rcc/operations"
	"github.com/robocorp/rcc/pretty"

	"github.com/spf13/cobra"
)

var (
	holotreeShowBlueprint bool
)

var holotreeHashCmd = &cobra.Command{
	Use:   "hash <conda.yaml+>",
	Short: "Calculates a blueprint hash for managed holotree virtual environment from conda.yaml files.",
	Long:  "Calculates a blueprint hash for managed holotree virtual environment from conda.yaml files.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if common.DebugFlag() {
			defer common.Stopwatch("Conda YAML hash calculation lasted").Report()
		}
		_, holotreeBlueprint, err := htfs.ComposeFinalBlueprint(args, "", common.DevDependencies)
		pretty.Guard(err == nil, 1, "Blueprint calculation failed: %v", err)
		hash := common.BlueprintHash(holotreeBlueprint)

		common.Log("Blueprint hash for %v is %v.", args, hash)
		if holotreeShowBlueprint {
			common.Log("Blueprint:\n%s", holotreeBlueprint)
		}

		if holotreeJson {
			result := make(map[string]interface{})
			result["hash"] = hash
			if holotreeShowBlueprint {
				result["blueprint"] = string(holotreeBlueprint)
			}

			out, err := operations.NiceJsonOutput(result)
			pretty.Guard(err == nil, 6, "%s", err)
			fmt.Println(out)
		} else {
			if common.Silent() {
				common.Stdout("%s\n", hash)
				if holotreeShowBlueprint {
					common.Stdout("Blueprint:\n%s", holotreeBlueprint)
				}
			}
		}
	},
}

func init() {
	holotreeCmd.AddCommand(holotreeHashCmd)
	holotreeHashCmd.Flags().BoolVarP(&common.DevDependencies, "devdeps", "", false, "Include dev-dependencies from the `package.yaml` when calculating the hash (only valid when dealing with a `package.yaml` file).")
	holotreeHashCmd.Flags().BoolVarP(&holotreeJson, "json", "j", false, "Show environment as JSON.")
	holotreeHashCmd.Flags().BoolVarP(&holotreeShowBlueprint, "show-blueprint", "", false, "Show full blueprint, not just hash.")
}
