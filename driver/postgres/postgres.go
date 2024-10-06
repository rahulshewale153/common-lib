package postgres

import (
	"fmt"
	"time"

	"github.com/rahulshewale153/common-lib/log"

	"database/sql"

	"github.com/go-sql-driver/mysql"
)

const (
	defaultCollation        = "utf8mb4_general_ci"
	defaultMaxAllowedPacket = 4 << 20 // 4 MiB
	defaultConnectTimeout   = 5 * time.Second
	defaultReadTimeout      = 1 * time.Second
	defaultWriteTimeout     = 5 * time.Second
	defaultMaxIdleConns     = 1
	defaultMaxOpenConns     = 5
	defaultMaxLifetime      = 0 // no expiry

	SQLTimeFormatLayout = "2006-01-02 15:04:05"
)

type PGConnector struct {
	client *sql.DB
}

type PGConfig struct {
	Host             string
	Port             int
	UserName         string
	Password         string
	DBName           string
	Collation        string // e.g. utf8_general_ci
	MaxAllowedPacket int    // in byts
	Location         *time.Location
	ConnectTimeout   time.Duration
	ReadTimeout      time.Duration
	WriteTimeout     time.Duration
	MaxIdleConns     int
	MaxOpenConns     int
	ConnMaxLifetime  time.Duration
	ParseTime        bool
}

// return PGConnector with setting client to given connection
// its expected to provide mocked conn while calling this api
func NewMockedPGConnector(conn *sql.DB) *PGConnector {
	return &PGConnector{
		client: conn,
	}
}

// return the new MySQL connector
func NewPGConnector(cfg PGConfig) (*PGConnector, error) {
	collation := cfg.Collation
	if "" == collation {
		collation = defaultCollation
	}
	location := cfg.Location
	if nil == location {
		location = time.UTC
	}

	connectTimeout := cfg.ConnectTimeout
	if 0 == connectTimeout {
		connectTimeout = defaultConnectTimeout
	}

	readTimeout := cfg.ReadTimeout
	if 0 == readTimeout {
		readTimeout = defaultReadTimeout
	}
	writeTimeout := cfg.WriteTimeout
	if 0 == readTimeout {
		writeTimeout = defaultWriteTimeout
	}

	maxAllowedPacket := cfg.MaxAllowedPacket
	if maxAllowedPacket <= 0 {
		maxAllowedPacket = defaultMaxAllowedPacket
	}

	mCfg := mysql.Config{
		User:                 cfg.UserName,
		Passwd:               cfg.Password,
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		DBName:               cfg.DBName,
		Collation:            collation,
		Loc:                  location,
		Timeout:              connectTimeout,
		ReadTimeout:          readTimeout,
		WriteTimeout:         writeTimeout,
		MaxAllowedPacket:     maxAllowedPacket,
		AllowNativePasswords: true,
		CheckConnLiveness:    true,
		ParseTime:            cfg.ParseTime,
	}

	connector, err := mysql.NewConnector(&mCfg)
	if err != nil {
		log.Error("Failed to create MySQL connector.", err.Error())
		return nil, err
	}

	conn := sql.OpenDB(connector)
	if err := conn.Ping(); err != nil {
		log.Error("Failed to connect MySQL Server.", err.Error())
		return nil, err
	}

	maxIdleConns := cfg.MaxIdleConns
	if maxIdleConns <= 0 {
		maxIdleConns = defaultMaxIdleConns
	}
	conn.SetMaxIdleConns(maxIdleConns)

	maxOpenConns := cfg.MaxOpenConns
	if maxOpenConns <= 0 {
		maxOpenConns = defaultMaxOpenConns
	}
	conn.SetMaxOpenConns(maxOpenConns)

	maxLifetime := cfg.ConnMaxLifetime
	if maxLifetime < 0 {
		maxLifetime = defaultMaxLifetime
	}
	conn.SetConnMaxLifetime(maxLifetime)

	return &PGConnector{
		client: conn,
	}, nil
}

// execute the select queries
func (conn *PGConnector) ExecuteSelect(query string, args ...any) (*sql.Rows, error) {
	return conn.client.Query(query, args...)
}

func (conn *PGConnector) Execute(query string, args ...any) (sql.Result, error) {
	return conn.client.Exec(query, args...)
}

// close the Aerospike client connection
func (conn *PGConnector) Close() {
	if err := conn.client.Close(); err != nil {
		log.Error("Error while closing MySQL connection!!! ", err.Error())
	} else {
		log.Error("Closed MySQL connection!!!")
	}
}
