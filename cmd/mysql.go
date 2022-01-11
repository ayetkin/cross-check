package cmd

import (
	"cross-check/cfg"
	"cross-check/check"
	"cross-check/model"
	"cross-check/router"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var mysqlCmd = &cobra.Command{
	Use:   "mysql",
	Short: "Switch master&slave check for MySQL nodes.",
	Long: `This tool helps to making master&slave check for MySQL servers.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(flagCONFIG) == 0 {
			_ = cmd.Help()
			os.Exit(0)
		}else {
			model.DATABASE = "mysql"
			model.CONFIG = flagCONFIG
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cfg.NewConfig(model.CONFIG); err != nil {
			log.Fatal("Failed to initialize config file. ", err)
		}

		go check.Retry(model.DATABASE)

		r := router.Router()
		log.Warning("Starting web server on "+ cfg.Values.Server.Host + ":" + cfg.Values.Server.Port + " ...")
		log.Fatal(http.ListenAndServe(cfg.Values.Server.Host+":"+cfg.Values.Server.Port, r))
		return  nil
	},
}

func init() {
	rootCmd.AddCommand(mysqlCmd)
	mysqlCmd.Flags().StringVarP(&flagCONFIG, "config", "c", "", "Give a config file. \t ex: --config /etc/cross-check/config.yaml")
}
