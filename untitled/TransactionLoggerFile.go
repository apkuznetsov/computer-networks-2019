package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"sync"
)

type TransactionLoggerFile struct {
	events       chan<- Event
	errors       <-chan error
	lastSequence uint64
	file         *os.File
	wg           *sync.WaitGroup
}

func (log *TransactionLoggerFile) WritePut(key, value string) {
	log.wg.Add(1)
	log.events <- Event{EventType: EventPut, Key: key, Value: url.QueryEscape(value)}
}

func (log *TransactionLoggerFile) WriteDelete(key string) {
	log.wg.Add(1)
	log.events <- Event{EventType: EventDelete, Key: key}
}

func (log *TransactionLoggerFile) Err() <-chan error {
	return log.errors
}

func NewTransactionLoggerFile(filename string) (*TransactionLoggerFile, error) {
	var log = TransactionLoggerFile{wg: &sync.WaitGroup{}}
	var err error

	log.file, err = os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		return nil, fmt.Errorf("cannot open transaction log file: %w", err)
	}

	return &log, nil
}

func (log *TransactionLoggerFile) Run() {
	events := make(chan Event, 16)
	log.events = events

	errors := make(chan error, 1)
	log.errors = errors

	go func() {
		for e := range events {
			log.lastSequence++

			_, err := fmt.Fprintf(
				log.file,
				"%d\t%d\t%s\t%s\n",
				log.lastSequence, e.EventType, e.Key, e.Value)

			if err != nil {
				errors <- fmt.Errorf("cannot write to log file: %w", err)
			}

			log.wg.Done()
		}
	}()
}

func (log *TransactionLoggerFile) Wait() {
	log.wg.Wait()
}

func (log *TransactionLoggerFile) Close() error {
	log.Wait()

	if log.events != nil {
		close(log.events)
	}

	return log.file.Close()
}

func (log *TransactionLoggerFile) ReadEvents() (<-chan Event, <-chan error) {
	scanner := bufio.NewScanner(log.file)
	outEvent := make(chan Event)
	outError := make(chan error, 1)

	go func() {
		var e Event

		defer close(outEvent)
		defer close(outError)

		for scanner.Scan() {
			line := scanner.Text()

			fmt.Sscanf(
				line, "%d\t%d\t%s\t%s",
				&e.Sequence, &e.EventType, &e.Key, &e.Value)

			if log.lastSequence >= e.Sequence {
				outError <- fmt.Errorf("transaction numbers out of sequence")
				return
			}

			uv, err := url.QueryUnescape(e.Value)
			if err != nil {
				outError <- fmt.Errorf("value decoding failure: %w", err)
				return
			}

			e.Value = uv
			log.lastSequence = e.Sequence

			outEvent <- e
		}

		if err := scanner.Err(); err != nil {
			outError <- fmt.Errorf("transaction log read failure: %w", err)
		}
	}()

	return outEvent, outError
}
