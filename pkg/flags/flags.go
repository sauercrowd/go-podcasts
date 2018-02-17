package flags

import (
	"flag"
	"os"
	"strconv"
)

// Vars for environment
type Vars struct {
	PostgresHost     string
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string
	PostgresPort     int
}

//Parse everything needed
func Parse() *Vars {
	var f Vars
	flag.StringVar(&f.PostgresHost, "phost", "127.0.0.1", "Postgres Host")
	flag.StringVar(&f.PostgresUser, "puser", "postgres", "Postgres User")
	flag.StringVar(&f.PostgresPassword, "ppass", "postgres", "Postgres Password")
	flag.StringVar(&f.PostgresDatabase, "pname", "podcasts", "Database name")
	flag.IntVar(&f.PostgresPort, "pport", 5432, "Postgres Port")
	env := flag.Bool("env", false, "Parse parameters from environment (Uppercase the option, e.g. -pport --> PPORT")
	flag.Parse()
	if *env {
		f = parseFromEnv()
		return &f
	}
	return &f
}

func parseFromEnv() Vars {
	var f Vars
	f.PostgresHost = os.Getenv("PHOST")
	f.PostgresPassword = os.Getenv("PPASS")
	f.PostgresUser = os.Getenv("PUSER")
	f.PostgresDatabase = os.Getenv("PNAME")
	port, err := strconv.Atoi(os.Getenv("PPORT"))
	if err != nil {
		panic(err)
	}
	f.PostgresPort = port
	return f
}
