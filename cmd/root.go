/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	// "runtime"
	"github.com/spf13/cobra"
	// "seatbelt2/internals"
	fs "github.com/karrick/godirwalk"
	ps "github.com/mitchellh/go-ps"

	// "github.com/becheran/wildmatch-go"
	"github.com/iamacarpet/go-win64api"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "seatbelt2",
	Short: "A brief description of your application",
	Long: `Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("It's time to strap your seatbelt on! Use -h to get started.")
	},
}

var lightScan = &cobra.Command{
	Use:   "light [flags]",
	Short: "Run a Light detection check",
	Long:  `Enumerate Processes & File System.`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Starting Light scan...")

		proc, err := ps.Processes()
		if err != nil {
			fmt.Println("An error occured while enumerating processes: ", err)
		}

		// map to hold processes, used an empty struct to track the unique process names easier.

		uniqueProcesses := make(map[string]struct{})

		{
			for _, p := range proc {
				exec := p.Executable()
				if _, exists := uniqueProcesses[exec]; !exists {
					uniqueProcesses[exec] = struct{}{}
					fmt.Println("Process Found: ", exec)
				}
			}
		}

		time.Sleep(time.Second * 2)

		dirname := "/home/vscode"

		{
			err := fs.Walk(dirname, &fs.Options{
				Callback: func(osPathname string, de *fs.Dirent) error {

					// matcher := wildmatch.NewWildMatch(string(".*")) <- for wildcard shit, this should come in handy later.

					if strings.Contains(osPathname, ".git") || strings.Contains(osPathname, ".cache") || strings.Contains(osPathname, ".dotnet") || strings.Contains(osPathname, ".vscode-remote") || strings.Contains(osPathname, "go") {
						return fs.SkipThis
					}
					fmt.Printf("%s\n", osPathname)
					return nil
				},
				Unsorted: false, // (optional) set true for faster yet non-deterministic enumeration (see godoc)
			})

			if err != nil {
				fmt.Println("Error enumerating filesystem: ", err)
			}
		}
	},
}

var normalScan = &cobra.Command{
	Use:   "normal [flags]",
	Short: "Run a Normal detection check",
	Long:  `Enumerate Processes, File System, Services, & DNS Cache.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting Normal scan...")
		fmt.Println("Running Process Function")
		fmt.Println("Checking File System")

		// services
		{
			svc, err := winapi.GetServices()
			if err != nil {
				fmt.Println("Error getting services.")
			}
			for _, v := range svc {
				fmt.Printf("%-50s - %-75s - Status: %-20s - Accept Stop: %-5t, Running Pid: %d\r\n", v.SCName, v.DisplayName, v.StatusText, v.AcceptStop, v.RunningPid)
			}
		}
		// DNS Cache
		{
			
		}
	},
}

var OperatingSystem string

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.AddCommand(lightScan)
	rootCmd.AddCommand(normalScan)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.seatbelt2.yaml)")
	rootCmd.PersistentFlags().StringVar(&OperatingSystem, "os", "", "Options: windows, darwin (mac), linux")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}
