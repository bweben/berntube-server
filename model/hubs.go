package model

import (
	"errors"
	"github.com/bweben/berntube-server/web/socket"
)

type HubsErr string

func (e HubsErr) Error() string {
	return string(e)
}

const (
	ErrNotFound = HubsErr("could not find hub")
	ErrFound    = HubsErr("cannot add hub because it already exists")
)

type Hubs map[int64]*socket.Hub

func (h Hubs) Has(id int64) bool {
	_, ok := h[id]

	if !ok {
		return false
	}

	return true
}

func (h Hubs) Get(id int64) (*socket.Hub, error) {
	definition, ok := h[id]
	if !ok {
		return nil, errors.New("not found")
	}
	return definition, nil
}

func (h Hubs) Set(id int64, hub *socket.Hub) error {
	if !h.Has(id) {
		h[id] = hub
		return nil
	} else {
		return errors.New("found")
	}
}
