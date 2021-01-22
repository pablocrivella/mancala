package integration

import (
	"database/sql"
	"fmt"
	"net"
	"net/url"
	"os"
	"runtime"
	"testing"
	"time"

	"github.com/ory/dockertest"
)

func TestMain(m *testing.M) {
	os.Exit(testMain(m))
}

// testMain returns an integer denoting an exit code to be returned and used in
// TestMain. The exit code 0 denotes success, all other codes denote failure (1
// and 2).
func testMain(m *testing.M) int {
	// pgURL, err, cleanUp := startDatabase()
	// defer cleanUp()
	// if err != nil {
	// 	log.Fatalf(err.Error())
	// 	return 1
	// }
	// db, err := dburl.Open(pgURL.String())
	// if err != nil {
	// 	return 1
	// }
	// if err := db.Ping(); err != nil {
	// 	return 1
	// }
	// ready := make(chan bool)
	// server := restapi.Server{
	// 	DB:    db,
	// 	Port:  "8080",
	// 	Ready: ready,
	// }
	// go server.Start()
	// <-ready

	return m.Run()
}

func startDatabase() (*url.URL, error, func()) {
	pgURL := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword("mancala_dev", "pass"),
		Path:   "mancala_test",
	}
	q := pgURL.Query()
	q.Add("sslmode", "disable")
	pgURL.RawQuery = q.Encode()

	pool, err := dockertest.NewPool("")
	if err != nil {
		return nil, fmt.Errorf("Could not connect to docker: %w", err), nil
	}

	pw, _ := pgURL.User.Password()
	env := []string{
		"POSTGRES_USER=" + pgURL.User.Username(),
		"POSTGRES_PASSWORD=" + pw,
		"POSTGRES_DB=" + pgURL.Path,
	}

	resource, err := pool.Run("postgres", "13-alpine", env)
	if err != nil {
		return nil, fmt.Errorf("Could not start postgres container: %w", err), nil
	}

	cleanUp := func() {
		if err = pool.Purge(resource); err != nil {
			fmt.Errorf("Could not purge resource: %w", err)
		}
	}

	pgURL.Host = resource.Container.NetworkSettings.IPAddress
	if runtime.GOOS == "darwin" {
		pgURL.Host = net.JoinHostPort(resource.GetBoundIP("5432/tcp"), resource.GetPort("5432/tcp"))
	}

	pool.MaxWait = 10 * time.Second
	err = pool.Retry(func() error {
		db, err := sql.Open("postgres", pgURL.String())
		if err != nil {
			return err
		}
		defer func() {
			cerr := db.Close()
			if err == nil {
				err = cerr
			}
		}()
		return db.Ping()
	})
	if err != nil {
		return nil, fmt.Errorf("Could not connect to docker: %w", err), cleanUp
	}

	return pgURL, nil, cleanUp
}
