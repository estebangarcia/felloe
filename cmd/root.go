package cmd

import (
	"felloe/logger"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var logLevel string
var KubeNamespace string
var KubeConfigContext string

var rootCmd = &cobra.Command{
	Use:   "felloe",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		var level logrus.Level

		if err := level.UnmarshalText([]byte(logLevel)); err != nil {
			return fmt.Errorf("loglevel has to be one of trace, debug, info, error, fatal")
		}

		logger.GetLogger().SetLevel(level)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&KubeNamespace, "namespace", "n", "", "namespace scope")
	rootCmd.PersistentFlags().StringVar(&KubeConfigContext, "kube-context","", "name of the kubeconfig context to use")
	rootCmd.PersistentFlags().StringVarP(&logLevel, "loglevel", "l", "error", "set logger level [trace, debug, info, error, fatal]")

	viper.BindPFlags(rootCmd.PersistentFlags())
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}