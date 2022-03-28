package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/technikhil314/mono-cli/utils"
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
		runtime := utils.DetectRuntime()
		println(runtime)
		packageManager := utils.DetectPackageManager(runtime)
		executeCommand(packageManager)
	},
}

func executeCommand(packageManager string) {
	println(packageManager)
	switch packageManager {
	case "LERNA":
		fmt.Printf("%s%s", infoStyle, "running lerna bootstrap\n")
		utils.ExecuteCommand("npx", "lerna", "bootstrap")
		break
	case "NPM":
		fmt.Printf("%s%s", infoStyle, "running npm install\n")
		utils.ExecuteCommand("npm", "install")
		break
	case "YARN":
		fmt.Printf("%s%s", infoStyle, "running yarn install\n")
		utils.ExecuteCommand("yarn", "install")
		break
	case "PNPM":
		fmt.Printf("%s%s", infoStyle, "running pnpm install\n")
		utils.ExecuteCommand("pnpm", "install")
		break
	case "MAVEN":
		fmt.Printf("%s%s", infoStyle, "running maven install\n")
		utils.ExecuteCommand("mvn", "install", "-Dmaven.test.skip.exec", "-Dmaven.test.skip=true")
		break
	case "GRADLE":
		fmt.Printf("%s%s", infoStyle, "running gradle install\n")
		utils.ExecuteCommand("gradle", "build")
		break
	case "GOMOD":
		fmt.Printf("%s%s", infoStyle, "running go mod tidy\n")
		utils.ExecuteCommand("go", "mod", "tidy")
		break
	}
}
func init() {
	rootCmd.AddCommand(installCmd)
}
