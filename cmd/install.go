package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/technikhil314/mono-cli/utils"
	"github.com/ttacon/chalk"
)

var infoStyle = chalk.Yellow.NewStyle().WithBackground(chalk.Black)
var cwd, cwdErr = os.Getwd()

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs all your dependancies based on package manager you/your team is using on this project",
	Long: `This command detects which package manager you or your team members are using.
based on presence of package manager lock files. and will use the same package manager 
if it is not able to detect then it will fallback to npm as pacakge manager	
`,
	Run: func(cmd *cobra.Command, args []string) {
		executeMonorepoManagers()
		executePackageManagers()
		executeDefault()
	},
}

func executeMonorepoManagers() {
	monorepoManagerFile := fmt.Sprint(cwd, "/lerna.json")
	_, statErr := os.Stat(monorepoManagerFile)
	if statErr == nil {
		fmt.Printf("%s%s", infoStyle, "running lerna bootstrap\n")
		utils.ExecuteCommand("npx", "lerna", "bootstrap")
	}
}

func executePackageManagers() {
	lockfilePath := fmt.Sprint(cwd, "/package-lock.json")
	_, statErr := os.Stat(lockfilePath)
	if statErr == nil {
		fmt.Printf("%s%s", infoStyle, "running npm install\n")
		utils.ExecuteCommand("npm", "install")
	}
	lockfilePath = fmt.Sprint(cwd, "/yarn.lock")
	_, statErr = os.Stat(lockfilePath)
	if statErr == nil {
		fmt.Printf("%s%s", infoStyle, "running yarn install\n")
		utils.ExecuteCommand("yarn", "install")
	}
	lockfilePath = fmt.Sprint(cwd, "/pnpm-lock.yaml")
	_, statErr = os.Stat(lockfilePath)
	if statErr == nil {
		fmt.Printf("%s%s", infoStyle, "running pnpm install\n")
		utils.ExecuteCommand("pnpm", "install")
	}
}

func executeDefault() {
	packageManagerFile := fmt.Sprint(cwd, "/package.json")
	_, statErr := os.Stat(packageManagerFile)
	if statErr == nil {
		fmt.Println("%s%s", infoStyle, "No lock file found. Running npm install")
		utils.ExecuteCommand("npm", "install")
	}
	fmt.Printf("%sNo relevant files found in %s\n", chalk.Red, cwd)
	os.Exit(1)
}

func init() {
	rootCmd.AddCommand(installCmd)
}
