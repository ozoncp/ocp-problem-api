package saver

import (
	"github.com/ozoncp/ocp-problem-api/internal/flusher"
	"github.com/ozoncp/ocp-problem-api/internal/utils"
	"time"
)

type Saver interface {
	Save(problem *utils.Problem)
	GetErrorReader() <-chan utils.Problem
	Close()
}

type saver struct {
	flusher flusher.Flusher
	ticker *time.Ticker
	chUpdate  chan utils.Problem
	isRunning bool
	chClose   chan struct{}
	chErrorReader chan utils.Problem
}

func (s *saver) Save(problem *utils.Problem) {
	if !s.isRunning {
		return
	}

	select {
	case s.chUpdate <-*problem:
	default:
		s.flush()
		s.chUpdate <-*problem
	}
}

func (s *saver) Close() {
	if s.isRunning {
		s.chClose <- struct{}{}
	}
}

func (s *saver) GetErrorReader() <-chan utils.Problem {
	if s.chErrorReader == nil {
		s.chErrorReader = make(chan utils.Problem, cap(s.chUpdate))
	}

	return s.chErrorReader
}


func (s *saver) flush() {
	size := len(s.chUpdate)
	problemList := make([]utils.Problem, 0, size)
	for i := 0; i < size; i++ {
		select {
		case problem := <-s.chUpdate:
			problemList = append(problemList, problem)
		default:
		}
	}

	for _, errProblem := range s.flusher.Flush(problemList) {
		if s.chErrorReader != nil {
			s.chErrorReader <- errProblem
		}
	}
}

func (s *saver) cleanup() {
	s.flush()
	s.isRunning = false
	s.ticker.Stop()
	close(s.chClose)
	close(s.chUpdate)
}

func (s *saver) init() {
	if s.isRunning {
		return
	}

	s.isRunning = true
	go func() {
		defer s.cleanup()

		for {
			select {
			case <-s.ticker.C:
				s.flush()
			case <-s.chClose:
				return
			}
		}
	}()
}

func NewSaver(capacity uint, f flusher.Flusher, timeout time.Duration) Saver {
	s := &saver{
		flusher: f,
		ticker: time.NewTicker(timeout),
		chUpdate: make(chan utils.Problem, int(capacity)),
		chClose: make(chan struct{}),
	}
	s.init()

	return s
}
