package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile  string
	scrtFile string
)

var rootCmd = &cobra.Command{
	Use:   "taurus",
	Short: "Project Taurus is a secret project",
	Long: `We will define the long description later.
                It is own by https://wigataintech.com`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running Taurus...")
	},
}

// Execute executes the Go function.
//
// It has no parameters and returns an error.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// init initializes the application configuration.
//
// No parameters.
// No return type.
func init() {
	cobra.OnInitialize(initAppConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config.yaml)")
	rootCmd.PersistentFlags().StringVar(&scrtFile, "secret", "", "secret file (default is $HOME/.secret.yaml)")
	// rootCmd.PersistentFlags().StringVar(&migrationFile, "migration", "", "migration file (default is $HOME/migration/*.sql)")
	// rootCmd.PersistentFlags().StringVar(&cronJob, "cronJob", "", "cronJob (example like calculate-all-merchant-eod-balance)")
	// rootCmd.PersistentFlags().StringVar(&consoleName, "consoleName", "", "consoleName (example like migrate-transaction-fee)")

}

// initAppConfig initializes the application configuration.
func initAppConfig() {
	rootCmd.PersistentFlags().StringP("author", "a", "Dhira Wigata", "author name for copyright attribution")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")

	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "Dhira Wigata <github.com/mrdhira>")
}
