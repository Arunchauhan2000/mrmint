package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the mrmint blockchain",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🔧 Initializing chain...")
		runCommand("mrmintchain", "init", "mymint", "--chain-id", "os_9000-1", "--home", os.Getenv("HOME")+"/.mrmint")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func runCommand(name string, args ...string) {
	c := exec.Command(name, args...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	err := c.Run()
	if err != nil {
		fmt.Println("❌ Error:", err)
	}
}
