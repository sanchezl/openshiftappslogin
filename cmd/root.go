/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/sanchezl/openshiftappslogin/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "openshiftappslogin",
	Short: "Retrieves a bearer token from an OpenShift CI cluster.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(pkg.RetrieveBearerToken(
			viper.GetString("url"),
			viper.GetString("username"),
			viper.GetString("password"),
			log.Print,
		))
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $XDG_CONFIG_HOME/openshiftappslogin/config.yaml)")
	rootCmd.Flags().StringP("url", "o", "", "Token request URL")
	rootCmd.Flags().StringP("username", "u", "", "User name")
	rootCmd.Flags().StringP("password", "p", "", "User password")
	rootCmd.MarkFlagRequired("url")
	rootCmd.MarkFlagRequired("username")
	rootCmd.MarkFlagRequired("password")
	viper.BindPFlags(rootCmd.Flags())
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Use config file in $XDG_CONFIG_HOME (or $HOME/.config if not set)
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		configPath := path.Join(home, ".config/openshiftappslogin")
		if configHome := os.Getenv("XDG_CONFIG_HOME"); len(configHome) != 0 {
			configPath = path.Join(configHome, "openshiftappslogin")
		}
		viper.AddConfigPath(configPath)
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	if !viper.IsSet("password") {
		if viper.IsSet("secret") {
			p, err := pkg.RedHatInternalPassword(viper.GetString("secret"), viper.GetString("prefix"))
			if err != nil {
				panic(err)
			}
			viper.Set("password", p)
		}
	}
	rootCmd.Flags().VisitAll(func(f *pflag.Flag) {
		if viper.IsSet(f.Name) {
			rootCmd.Flags().SetAnnotation(f.Name, cobra.BashCompOneRequiredFlag, []string{"false"})
		}
	})
}
