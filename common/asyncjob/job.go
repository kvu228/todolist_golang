package asyncjob

import (
	"context"
	"log"
	"time"
)

type Job interface {
	Execute(ctx context.Context) error
	Retry(ctx context.Context) error
	State() JobState
	SetRetryDurations(times []time.Duration)
}

const (
	// defautlMaxTimeOut for 10 seconds
	defaultMaxTimeout = time.Second * 10
)

var (
	defaultRetryTime = []time.Duration{time.Second, time.Second * 2, time.Second * 4}
)

type JobHandler func(ctx context.Context) error

type JobState int

const (
	StateInit        JobState = iota
	StateRunning              = 1
	StateFailed               = 2
	StateTimeout              = 3
	StateCompleted            = 4
	StateRetryFailed          = 5
)

func (js JobState) String() string {
	return [6]string{"init", "Running", "Failed", "Timeout", "Complete", "RetryFailed"}[js]
}

type jobConfig struct {
	Name       string
	MaxTimeout time.Duration
	Retries    []time.Duration
}
type job struct {
	config     jobConfig
	handler    JobHandler
	state      JobState
	retryIndex int
	stopChan   chan bool
}

func NewJob(jobHandler JobHandler, options ...OptionHdl) *job {
	j := job{
		config: jobConfig{
			MaxTimeout: defaultMaxTimeout,
			Retries:    defaultRetryTime,
		},
		handler:    jobHandler,
		state:      StateInit,
		retryIndex: -1,
		stopChan:   make(chan bool),
	}
	for i := range options {
		options[i](&j.config)
	}
	return &j
}

func (j *job) Execute(ctx context.Context) error {
	log.Printf("execute %s\n", j.config.Name)
	j.state = StateRunning

	var err error
	err = j.handler(ctx)

	if err != nil {
		j.state = StateFailed
		return err
	}

	j.state = StateCompleted

	return nil
}

func (j *job) Retry(ctx context.Context) error {
	if j.retryIndex == len(j.config.Retries)-1 {
		return nil
	}

	j.retryIndex += 1
	time.Sleep(j.config.Retries[j.retryIndex])

	err := j.Execute(ctx)

	if err == nil {
		j.state = StateCompleted
		return nil
	}

	if j.retryIndex == len(j.config.Retries)-1 {
		j.state = StateRetryFailed
		return err
	}

	j.state = StateFailed
	return err
}

func (j *job) State() JobState { return j.state }
func (j *job) RetryIndex() int { return j.retryIndex }

func (j *job) SetRetryDurations(durations []time.Duration) {
	if len(durations) == 0 {
		return
	}

	j.config.Retries = durations
}

type OptionHdl func(*jobConfig)

func WithName(name string) OptionHdl {
	return func(cf *jobConfig) {
		cf.Name = name
	}
}

func WithRetriesDuration(items []time.Duration) OptionHdl {
	return func(cf *jobConfig) {
		cf.Retries = items
	}
}
