package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)


const (
	Program = "ict Cmd"
	Version = "0.1.0"
	ConfigType = "yaml"
)

var ConfigFile string

const defaultConfigFile = "ict.yaml"

var rootCmd  = &cobra.Command{
	Use: Program,
	Short: "a simple tool to merge request, review and commit",
	Run: func (cmd *cobra.Command, args []string)  {
		fmt.Printf("rootCmd starts!")
	},
}

func Excute(){
	err := rootCmd.Execute()
	if err != nil{
		os.Exit(1)
	}
}

func init(){
	rootCmd.PersistentFlags().StringP("Log-level", "", "", "specify the Log level")
}

