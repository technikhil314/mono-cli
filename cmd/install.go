package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

var infoStyle = chalk.Yellow.NewStyle().WithBackground(chalk.Black)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs all your dependancies based on package manager you/your team is using on this project",
	Long: `This command detects which package manager you or your team members are using.
based on presence of package manager lock files. and will use the same package manager 
if it is not able to detect then it will fallback to npm as pacakge manager	
`,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		executeMonorepoManagers()
		executePackageManagers()
		executeDefault()
	},
}

func executeMonorepoManagers() {
	cwd, _ := os.Getwd()
	monorepoManagerFile := fmt.Sprint(cwd, "/lerna.json")
	_, statErr := os.Stat(monorepoManagerFile)
	if statErr == nil {
		fmt.Printf("%s%s", infoStyle, "running lerna bootstrap\n")
		exec.Command("npx lerna bootstrap")
		os.Exit(0)
	}
}

func executePackageManagers() {
	cwd, _ := os.Getwd()
	lockfilePath := fmt.Sprint(cwd, "/package-lock.json")
	_, statErr := os.Stat(lockfilePath)
	if statErr == nil {
		fmt.Printf("%s%s", infoStyle, "running npm install\n")
		exec.Command("npm install")
		os.Exit(0)
	}
	lockfilePath = fmt.Sprint(cwd, "/yarn.lock")
	_, statErr = os.Stat(lockfilePath)
	if statErr == nil {
		fmt.Printf("%s%s", infoStyle, "running yarn install\n")
		exec.Command("yarn install")
		os.Exit(0)
	}
	lockfilePath = fmt.Sprint(cwd, "/pnpm-lock.yaml")
	_, statErr = os.Stat(lockfilePath)
	if statErr == nil {
		fmt.Printf("%s%s", infoStyle, "running pnpm install\n")
		exec.Command("pnpm install")
		os.Exit(0)
	}
}

func executeDefault() {
	cwd, _ := os.Getwd()
	packageManagerFile := fmt.Sprint(cwd, "/package.json")
	_, statErr := os.Stat(packageManagerFile)
	if statErr == nil {
		fmt.Println("%s%s", infoStyle, "No lock file found. Running npm install")
		exec.Command("npm install")
		os.Exit(0)
	}
	fmt.Printf("%sNo relevant files found in %s\n", chalk.Red, cwd)
	os.Exit(1)
}

func init() {
	rootCmd.AddCommand(installCmd)
}
