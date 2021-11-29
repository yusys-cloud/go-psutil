// Author: yangzq80@gmail.com
// Date: 2021-11-29
//
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:  "psutil [net]",
	Long: `retrieving information on running processes and system utilization (CPU, memory, disks, network, sensors)`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rootCmd")
	},
	Example: stressTestExample(),
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func stressTestExample() string {
	return `psutil net`
}
