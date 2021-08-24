package saver

import (
	"errors"
	"sync"
	"time"

	"github.com/ozonva/ova-obligation-api/internal/entity"
	"github.com/ozonva/ova-obligation-api/internal/flusher"
)

var ErrorClosedSaver = errors.New("saver already closed")

type Saver interface {
	Save(entity entity.Obligation) error
	Close() error
}

func NewSaver(
	capacity uint,
	flusher flusher.Flusher,
	intervalToFlush time.Ticker,
) Saver {
	instance := saver{
		mu:       sync.Mutex{},
		entities: make([]entity.Obligation, 0, capacity),
		tiker:    intervalToFlush,
		flusher:  flusher,
		closeCh:  make(chan bool),
	}

	instance.startTimer()

	return &instance
}

type saver struct {
	mu       sync.Mutex
	entities []entity.Obligation
	tiker    time.Ticker
	flusher  flusher.Flusher
	closeCh  chan bool
}

func (s *saver) Save(entity entity.Obligation) error {
	if s.isClosed() {
		return ErrorClosedSaver
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if cap(s.entities) == len(s.entities) {
		s.flush()
	}

	s.entities = append(s.entities, entity)

	return nil
}

func (s *saver) Close() error {
	if s.isClosed() {
		return ErrorClosedSaver
	}

	close(s.closeCh)

	return nil
}

func (s *saver) isClosed() bool {
	select {
	case <-s.closeCh:
		return true
	default:
	}

	return false
}

func (s *saver) startTimer() {
	go func() {
		for {
			select {
			case <-s.tiker.C:
				s.mu.Lock()
				s.flush()
				s.mu.Unlock()
			case <-s.closeCh:
				s.mu.Lock()
				s.flush()
				s.mu.Unlock()
				return
			}
		}
	}()
}

func (s *saver) flush() {
	if len(s.entities) == 0 {
		return
	}

	s.flusher.Flush(s.entities)

	s.entities = s.entities[:0]
}
