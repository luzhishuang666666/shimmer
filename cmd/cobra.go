package cmd

import (
	"errors"
	"fmt"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	"github.com/spf13/cobra"
	"os"
	"shimmer/cmd/app"

	"shimmer/cmd/api"
	"shimmer/cmd/config"
	"shimmer/cmd/migrate"
	"shimmer/cmd/version"
)

var rootCmd = &cobra.Command{
	Use:          "shimmer",
	Short:        "shimmer",
	SilenceUsage: true,
	Long:         `shimmer`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New(pkg.Red("启动失敗"))
		}
		return nil
	},
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
}

func tip() {
	usageStr := `欢迎使用 微光`
	fmt.Printf("%s\n", usageStr)
}

func init() {
	rootCmd.AddCommand(api.StartCmd)
	rootCmd.AddCommand(migrate.StartCmd)
	rootCmd.AddCommand(version.StartCmd)
	rootCmd.AddCommand(config.StartCmd)
	rootCmd.AddCommand(app.StartCmd)
}

//Execute : apply commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
