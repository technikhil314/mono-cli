package cmd

import (
	"github.com/spf13/cobra"
	"github.com/technikhil314/mono-cli/utils"
	"github.com/ttacon/chalk"
)

var infoStyle = chalk.Yellow.NewStyle().WithBackground(chalk.Black)

var bootstrapCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "Installs all your dependancies based on listed in your package manager file",
	Run: func(cmd *cobra.Command, args []string) {
		runtime := utils.DetectRuntime()
		packageManager := utils.DetectPackageManager(runtime)
		executeBootstrapCommand(packageManager)
	},
}

func executeBootstrapCommand(packageManager string) {
	switch packageManager {
	case "LERNA":
		utils.ExecuteCommand("npx", "lerna", "bootstrap")
		break
	case "NPM":
		utils.ExecuteCommand("npm", "install")
		break
	case "YARN":
		utils.ExecuteCommand("yarn", "install")
		break
	case "PNPM":
		utils.ExecuteCommand("pnpm", "install")
		break
	case "MAVEN":
		utils.ExecuteCommand("mvn", "install", "-Dmaven.test.skip.exec", "-Dmaven.test.skip=true")
		break
	case "GRADLE":
		utils.ExecuteCommand("gradle", "assemble", "--refresh-dependencies")
		break
	case "GOMOD":
		utils.ExecuteCommand("go", "install")
		break
	}
}
func init() {
	rootCmd.AddCommand(bootstrapCmd)
}
