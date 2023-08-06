package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"net/url"
)

type DbParams struct {
	dbName   string
	host     string
	user     string
	password string
}

type TransactionLoggerPostgres struct {
	events chan<- Event
	errors <-chan error
	db     *sql.DB
}

func (log *TransactionLoggerPostgres) WritePut(key, value string) {
	log.events <- Event{EventType: EventPut, Key: key, Value: url.QueryEscape(value)}
}

func (log *TransactionLoggerPostgres) WriteDelete(key string) {
	log.events <- Event{EventType: EventDelete, Key: key}
}

func (log *TransactionLoggerPostgres) Err() <-chan error {
	return log.errors
}

func NewTransactionLoggerPostgres(params DbParams) (*TransactionLoggerPostgres, error) {
	connStr := fmt.Sprintf("host=%s dbname=%s user=%s password=%s",
		params.host, params.dbName, params.user, params.password)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to open db connection: %w", err)
	}

	logger := &TransactionLoggerPostgres{db: db}
	exists, err := logger.validateTableExists()
	if err != nil {
		return nil, fmt.Errorf("failed to verify table exists: %w", err)
	}
	if !exists {
		if err = logger.createTable(); err != nil {
			return nil, fmt.Errorf("failed to create table: %w", err)
		}
	}

	return logger, nil
}

func (log *TransactionLoggerPostgres) Run() {
	events := make(chan Event, 16)
	log.events = events

	errors := make(chan error, 1)
	log.errors = errors

	go func() {
		query := `INSERT INTO transactions
 (event_type, key, value)
 VALUES ($1, $2, $3)`

		for e := range events {
			_, err := log.db.Exec(
				query,
				e.EventType, e.Key, e.Value)
			if err != nil {
				errors <- err
			}
		}
	}()
}

func (log *TransactionLoggerPostgres) Wait() {
}

func (log *TransactionLoggerPostgres) Close() error {
	log.Wait()

	if log.events != nil {
		close(log.events)
	}

	return nil
}

func (log *TransactionLoggerPostgres) ReadEvents() (<-chan Event, <-chan error) {
	outEvent := make(chan Event)
	outError := make(chan error, 1)

	go func() {
		defer close(outEvent)
		defer close(outError)

		query := `SELECT sequence, event_type, key, value FROM transactions
 ORDER BY sequence`
		rows, err := log.db.Query(query) // Выполнить запрос; получить набор результатов
		if err != nil {
			outError <- fmt.Errorf("sql query error: %w", err)
			return
		}
		defer rows.Close()

		e := Event{}
		for rows.Next() {
			err = rows.Scan(
				&e.Sequence, &e.EventType,
				&e.Key, &e.Value)
			if err != nil {
				outError <- fmt.Errorf("error reading row: %w", err)
				return
			}
			outEvent <- e
		}
		err = rows.Err()
		if err != nil {
			outError <- fmt.Errorf("transaction log read failure: %w", err)
		}
	}()

	return outEvent, outError
}

func (log *TransactionLoggerPostgres) validateTableExists() (bool, error) {
	return false, nil
}

func (log *TransactionLoggerPostgres) createTable() error {
	return nil
}
