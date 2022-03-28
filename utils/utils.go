package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

var cwd, cwdErr = os.Getwd()

func ExecuteCommand(cmdName string, args ...string) {
	out, err := exec.Command(cmdName, args...).Output()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Println(string(out))
	os.Exit(0)
}

func DetectRuntime() string {
	var runtime = ""
	packageManagerFile := fmt.Sprint(cwd, "/package.json")
	_, statErr := os.Stat(packageManagerFile)
	if statErr == nil {
		runtime = "NODE"
	}
	packageManagerFile = fmt.Sprint(cwd, "/pom.xml")
	_, statErr = os.Stat(packageManagerFile)
	if statErr == nil {
		runtime = "JAVA"
	}
	packageManagerFile = fmt.Sprint(cwd, "/build.gradle")
	_, statErr = os.Stat(packageManagerFile)
	if statErr == nil {
		runtime = "JAVA"
	}
	return runtime
}

func DetectPackageManager(runtime string) string {
	var packageManager string = ""
	switch runtime {
	case "JAVA":
		packageManagerFile := fmt.Sprint(cwd, "/pom.xml")
		println(packageManagerFile)
		_, statErr := os.Stat(packageManagerFile)
		if statErr == nil {
			packageManager = "MAVEN"
		}
		packageManagerFile = fmt.Sprint(cwd, "/build.gradle")
		_, statErr = os.Stat(packageManagerFile)
		if statErr == nil {
			packageManager = "GRADLE"
		}
		break
	case "NODE":
		lockfilePath := fmt.Sprint(cwd, "/package-lock.json")
		_, statErr := os.Stat(lockfilePath)
		if statErr == nil {
			packageManager = "NPM"
		}
		lockfilePath = fmt.Sprint(cwd, "/yarn.lock")
		_, statErr = os.Stat(lockfilePath)
		if statErr == nil {
			packageManager = "YARN"
		}
		lockfilePath = fmt.Sprint(cwd, "/pnpm-lock.yaml")
		_, statErr = os.Stat(lockfilePath)
		if statErr == nil {
			packageManager = "PNPM"
		}
		monorepoManagerFile := fmt.Sprint(cwd, "/lerna.json")
		_, statErr = os.Stat(monorepoManagerFile)
		if statErr == nil {
			packageManager = "LERNA"
		}
		break
	case "GO":
		packageManagerFile := fmt.Sprintf(cwd, "/go.mod")
		_, statErr := os.Stat(packageManagerFile)
		if statErr == nil {
			packageManager = "GOMOD"
		}
		break
	}
	if packageManager == "" {
		packageManager = runtime
	}
	return packageManager
}
