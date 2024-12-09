package cmd

import (
	"fmt"
	"os"

	"github.com/HaroldObasi/multi-term/window"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mt",
	Short: "multi term is a cool little text editor",
	Long: `Generate and deploy static site portfolio
	
	Example: mt -f testing.txt`,
	
	Run:     RunCommand,
	PreRunE: PreRunChecks,
}

func init() {
	rootCmd.Flags().StringP("fileName", "f", "", "Filename")
}

func RunCommand(cmd *cobra.Command, args []string) {
	fileName, _ := cmd.Flags().GetString("fileName")
	fmt.Println("Editing file: ", fileName)

	win, err := window.NewWindow()

	if err != nil {
		fmt.Println(err)
	}

	err = win.Start()

	if err != nil {
		panic(err)
	}

}

func PreRunChecks(cmd *cobra.Command, args []string) error {
	fileName, _ := cmd.Flags().GetString("fileName")

	if fileName == "" {
		return fmt.Errorf("filename is required")
	}

	return nil
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}