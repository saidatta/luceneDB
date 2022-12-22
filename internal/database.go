package main

import (
	"github.com/cockroachdb/pebble"
)

type database struct {
	data  *pebble.DB // Primary data
	index *pebble.DB
	port  string
}

func newDatabase(hostname string) (*pebble.DB, error) {
	return pebble.Open(hostname, &pebble.Options{})
}

func reindex() {

}
