package main

import (
	"log"
	"os/exec"
)

func upgradeGo() {
	if err := exec.Command("make", "upgrade-go").Run(); err != nil {
		panic(err)
	}
}

func upgradeNode() {
	if err := exec.Command("make", "upgrade-node").Run(); err != nil {
		panic(err)
	}
}

func runLint() {
	if err := exec.Command("make", "lint-check").Run(); err != nil {
		panic(err)
	}
}

func runTests() {
	if err := exec.Command("make", "compose-test").Run(); err != nil {
		panic(err)
	}
}

func commit() {
	if err := exec.Command("git", "add", ".", "--update").Run(); err != nil {
		panic(err)
	}

	if err := exec.Command("git", "commit", "-m", "fix(deps): upgrade dependencies").Run(); err != nil {
		panic(err)
	}
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Upgrade failed!")
			log.Println("Please, check 'Error' or undo changes with: make upgrade-reset")
			log.Println("Error:", r)
		}
	}()

	log.Println("Upgrading dependencies")

	log.Println("Running Go Tidy")
	runLint()

	log.Println("Upgrading Go dependencies")
	upgradeGo()

	log.Println("Upgrading Node dependencies")
	upgradeNode()

	log.Println("Running Lint")
	runLint()

	log.Println("Running Tests")
	runTests()

	log.Println("Commit changes")
	commit()

	log.Println("Upgrade completed!")
}
