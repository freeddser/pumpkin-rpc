package repository

import (
	"flag"
	"fmt"
	"github.com/coreos/pkg/flagutil"
	"github.com/freeddser/rs-common/logging"
)

var (
	psqldatasource *DataSource
	psqldbname     string
	psqluser       string
	psqlpassword   string
	psqlhost       string
	psqlport       string
)

var log = logging.MustGetLogger()

func InitFactory() error {
	flags := flag.NewFlagSet("user-auth", flag.PanicOnError)
	flags.StringVar(&psqluser, "DB_USER", "", "PSQL username")
	flags.StringVar(&psqlpassword, "DB_PASSWORD", "", "PSQL username")
	flags.StringVar(&psqldbname, "DB_NAME", "", "PSQL db name")
	flags.StringVar(&psqlhost, "DB_HOST", "", "PSQL username")
	flags.StringVar(&psqlport, "DB_PORT", "", "PSQL db name")
	err := flagutil.SetFlagsFromEnv(flags, "DG_PSQL")

	if err != nil {
		return err
	}

	if psqlhost == "" {
		psqlhost = "localhost"
	}

	if psqlport == "" {
		psqlport = "5432"
	}

	if psqluser == "" || psqlpassword == "" || psqldbname == "" {
		return fmt.Errorf("Unable to get environtment variable, make sure you already set it!")
	}

	psqldatasource, err = NewDatabaseConnection(psqlhost, psqlport, psqluser, psqlpassword, psqldbname)
	if err != nil {
		return err
	}
	log.Info("Database Connection Started")
	fmt.Println(psqlhost, psqluser, psqldbname)

	return nil
}
