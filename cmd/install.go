/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/technikhil314/mono-cli/utils"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs one or more package(s)",
	Long: `
		Examples:

		mono install lodash
		mono install -g lodash
	`,
	Run: func(cmd *cobra.Command, args []string) {
		runtime := utils.DetectRuntime()
		packageManager := utils.DetectPackageManager(runtime)
		global, _ := cmd.Flags().GetBool("global")
		dev, _ := cmd.Flags().GetBool("dev")
		save, _ := cmd.Flags().GetBool("save")
		peer, _ := cmd.Flags().GetBool("peer")
		optional, _ := cmd.Flags().GetBool("optional")
		if global {
			executeGlobalInstall(args, packageManager)
		} else if dev {
			executeDevInstall(args, packageManager)
		} else if peer {
			executePeerInstall(args, packageManager)
		} else if optional {
			executeOptionalInstall(args, packageManager)
		} else if save {
			executeSaveInstall(args, packageManager)
		}
	},
}

func executeOptionalInstall(args []string, packageManager string) {
	switch packageManager {
	case "NPM":
		utils.ExecuteCommand("npm", "install", "-O", strings.Join(args, " "))
		break
	case "YARN":
		utils.ExecuteCommand("yarn", "add", "-O", strings.Join(args, " "))
		break
	case "PNPM":
		utils.ExecuteCommand("pnpm", "add", "-O", strings.Join(args, " "))
		break
	}
}

func executePeerInstall(args []string, packageManager string) {
	switch packageManager {
	case "NPM":
		fmt.Printf("%s %s", infoStyle, "npm does not support peer install but we will support soon")
		os.Exit(1)
		break
	case "YARN":
		utils.ExecuteCommand("yarn", "add", "-P", strings.Join(args, " "))
		break
	case "PNPM":
		utils.ExecuteCommand("pnpm", "add", "--save-peer", strings.Join(args, " "))
		break
	}
}

func executeSaveInstall(args []string, packageManager string) {
	switch packageManager {
	case "NPM":
		utils.ExecuteCommand("npm", "install", strings.Join(args, " "))
		break
	case "YARN":
		utils.ExecuteCommand("yarn", "add", strings.Join(args, " "))
		break
	case "PNPM":
		utils.ExecuteCommand("pnpm", "add", "-P", strings.Join(args, " "))
		break
	}
}

func executeDevInstall(args []string, packageManager string) {
	switch packageManager {
	case "NPM":
		utils.ExecuteCommand("npm", "install", "-D", strings.Join(args, " "))
		break
	case "YARN":
		utils.ExecuteCommand("yarn", "add", "-D", strings.Join(args, " "))
		break
	case "PNPM":
		utils.ExecuteCommand("pnpm", "add", "-D", strings.Join(args, " "))
		break
	}
}

func executeGlobalInstall(args []string, packageManager string) {
	switch packageManager {
	case "NPM":
		utils.ExecuteCommand("npm", "install", "-g", strings.Join(args, " "))
		break
	case "YARN":
		utils.ExecuteCommand("yarn", "global", "add", strings.Join(args, " "))
		break
	case "PNPM":
		utils.ExecuteCommand("pnpm", "add", "-g", strings.Join(args, " "))
		break
	}
}

func init() {
	rootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	installCmd.Flags().BoolP("global", "g", false, "Install the specified packages globally")
	installCmd.Flags().BoolP("dev", "d", false, "Install the specified packages locally and add them to devDependancies")
	installCmd.Flags().BoolP("save", "s", true, "Install the specified packages locally and add them to dependancies")
	installCmd.Flags().BoolP("peer", "p", false, "Install the specified packages locally and add them to peerDependancies")
	installCmd.Flags().BoolP("optional", "o", false, "Install the specified packages locally and add them to optionalDependancies")
}
