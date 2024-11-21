package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/iambrkkdn/go-clean-project/internal/config"
	_ "github.com/lib/pq"
)

const (
	defaultMaxOpenConns    = 30
	defaultMaxIdleConns    = 10
	defaultLifeTimeConn    = time.Minute
	defaultMaxIdleTimeConn = 5 * time.Minute
)

const (
	LocaleRU = "ru"
	LocaleEN = "en"
	LocaleAZ = "az"
)

type application struct {
	db     *sql.DB
	server *http.Server
}

func newApplication(c config.Config) (*application, error) {
	d, err := provideDB(c.DB)
	if err != nil {
		return nil, fmt.Errorf("failed to establish connection to PostgreSQL server: %w", err)
	}

	api := gin.New()

	server := &http.Server{
		Addr:           c.API.Address,
		Handler:        api,
		MaxHeaderBytes: 1 << 20,
	}

	return &application{
		db:     d,
		server: server,
	}, nil
}

func (a *application) Run() {
	if err := a.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		//log.Fatal().Err(err).Msg("server listening failed")
	}
}

func provideDB(c config.DB) (*sql.DB, error) {
	var d *sql.DB
	var err error
	driverName := "postgres"
	dataSourceName := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s&search_path=%s",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Name,
		c.SSLMode,
		c.Schema,
	)
	d, err = sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	if c.MaxIdleConnections == nil {
		d.SetMaxIdleConns(defaultMaxIdleConns)
	} else {
		d.SetMaxIdleConns(*c.MaxIdleConnections)
	}

	if c.MaxOpenConnections == nil {
		d.SetMaxOpenConns(defaultMaxOpenConns)
	} else {
		d.SetMaxOpenConns(*c.MaxOpenConnections)
	}

	d.SetConnMaxLifetime(defaultLifeTimeConn)
	d.SetConnMaxIdleTime(defaultMaxIdleTimeConn)

	if err := d.Ping(); err != nil {
		return nil, err
	}
	return d, nil
}
