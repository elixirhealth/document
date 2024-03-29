package cmd

import (
	"log"

	"github.com/drausin/libri/libri/common/errors"
	"github.com/drausin/libri/libri/common/logging"
	"github.com/elixirhealth/document/pkg/server"
	"github.com/elixirhealth/document/version"
	"github.com/elixirhealth/service-base/pkg/cmd"
	bserver "github.com/elixirhealth/service-base/pkg/server"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	serviceNameLower = "document"
	serviceNameCamel = "Document"
	envVarPrefix     = "DOCUMENT"
	logLevelFlag     = "logLevel"

	// TODO uncomment or delete
	//storageMemoryFlag    = "storageMemory"
	//storageDataStoreFlag = "storageDataStore"
	//storagePostgresFlag  = "storagePostgres"
	//dbURLFlag            = "dbURL"
)

var (
	rootCmd = &cobra.Command{
		Short: "TODO", // TODO
	}
)

func init() {
	rootCmd.PersistentFlags().String(logLevelFlag, bserver.DefaultLogLevel.String(),
		"log level")

	cmd.Start(serviceNameLower, serviceNameCamel, rootCmd, version.Current, start,
		func(flags *pflag.FlagSet) {
			// TODO define other flags here if needed, e.g.,
			//flags.Bool(storageMemoryFlag, true, "use in-memory storage")
			//flags.Bool(storageDataStoreFlag, false, "use GCP DataStore storage")
			//flags.Bool(storagePostgresFlag, false, "use Postgres DB storage")
			//flags.String(dbURLFlag, "", "Postgres DB URL")
		})

	testCmd := cmd.Test(serviceNameLower, rootCmd)
	cmd.TestHealth(serviceNameLower, testCmd)
	cmd.TestIO(serviceNameLower, testCmd, testIO, func(flags *pflag.FlagSet) {
		// TODO define other flags here if needed
	})

	cmd.Version(serviceNameLower, rootCmd, version.Current)

	// bind viper flags
	viper.SetEnvPrefix(envVarPrefix) // look for env vars with prefix
	viper.AutomaticEnv()             // read in environment variables that match
	errors.MaybePanic(viper.BindPFlags(rootCmd.Flags()))
}

// Execute runs the root document command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func start() error {
	config, err := getDocumentConfig()
	if err != nil {
		return err
	}
	return server.Start(config, make(chan *server.Document, 1))
}

func getDocumentConfig() (*server.Config, error) {
	c := server.NewDefaultConfig()
	c.WithServerPort(uint(viper.GetInt(cmd.ServerPortFlag))).
		WithMetricsPort(uint(viper.GetInt(cmd.MetricsPortFlag))).
		WithProfilerPort(uint(viper.GetInt(cmd.ProfilerPortFlag))).
		WithLogLevel(logging.GetLogLevel(viper.GetString(logLevelFlag))).
		WithProfile(viper.GetBool(cmd.ProfileFlag))
	// TODO set other config elements here

	return c, nil
}
