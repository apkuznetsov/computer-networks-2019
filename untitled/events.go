package main

type EventType byte

const (
	_                  = iota
	EventPut EventType = iota
	EventDelete
)

type Event struct {
	Sequence  uint64
	EventType EventType
	Key       string
	Value     string
}
