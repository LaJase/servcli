/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/LaJase/servcli/internal"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "servcli",
	Short: "A brief description of your application",
	Long: ` This application provides a user-friendly command line interface that allows users to view and browse a list
  of available servers. Users will be able to specify criteria such as geographic location, availability, processing
  capacity, current load, or any other attribute relevant to their specific environment.

  Through a combination of intuitive commands and advanced filters, users will be able to quickly narrow down the list
  of available servers to their specific needs. Once the server is selected, actions such as SSH connection, remote
  command execution or resource management can be performed easily. `,

	// with to argument this is the default launched application
	Run: func(cmd *cobra.Command, args []string) {
		m := internal.Model{}
		m.InitLists()

		p := tea.NewProgram(m)

		if _, err := p.Run(); err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}
	},
}

// Return the cobra.command
// This is needed for documentation generation
func GetServCliCmd() *cobra.Command {
	return rootCmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $PWD/config/servcli-config.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		pwd, err := os.Getwd()
		cobra.CheckErr(err)

		// Search config in $PWD/resource repository, should be yaml file named servcli-config
		viper.AddConfigPath(fmt.Sprintf("%s/config", pwd))
		viper.SetConfigType("yaml")
		viper.SetConfigName("servcli-config")
	}

	if err := viper.ReadInConfig(); err == nil {
		fmt.Printf("\nUsing config file: %s\n\n", viper.ConfigFileUsed())

		if err := viper.Unmarshal(&internal.CfgGlobal); err != nil {
			fmt.Printf("ERROR - Unmarshalling config file %s\n%s", viper.ConfigFileUsed(), err)
			os.Exit(1)
		}
	} else {
		fmt.Printf("ERROR - Reading config file %s\n%s", viper.ConfigFileUsed(), err)
		os.Exit(1)
	}
}
