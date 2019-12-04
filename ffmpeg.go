package fftool

import (
	"context"
	"errors"
	"github.com/goextension/log"
	"os"
	"strings"
	"sync"
)

// FFMpeg ...
type FFMpeg struct {
	err    error
	config *Config
	cmd    *Command
	Name   string
}

func (ff *FFMpeg) init() error {
	if ff.cmd == nil {
		ff.cmd = New(ff.Name)
	}
	if ff.err != nil {
		return ff.err
	}
	return nil
}

// Version ...
func (ff *FFMpeg) Version() (string, error) {
	if err := ff.init(); err != nil {
		return "", err
	}

	return ff.cmd.Run("-version")
}

// OptimizeWithFormat ...
func (ff *FFMpeg) OptimizeWithFormat(sfmt *StreamFormat) (newFF *FFMpeg) {
	cfg := ff.config.Clone()
	newFF = NewFFMpeg(&cfg)
	newFF.Name = ff.Name
	e := OptimizeWithFormat(&cfg, sfmt)
	if e != nil {
		newFF.err = e
	}
	return
}

// Run ...
func (ff FFMpeg) Run(ctx context.Context, input string) (e error) {
	if err := ff.init(); err != nil {
		return errWrap(err, "init")
	}
	args := outputArgs(ff.config, input)

	stat, e := os.Stat(ff.config.Output())
	if e != nil {
		if os.IsNotExist(e) {
			_ = os.MkdirAll(ff.config.Output(), 0755)
		} else {
			return errWrap(e, "stat")
		}
	}
	if e == nil && !stat.IsDir() {
		return errors.New("target is not dir")
	}

	outlog := make(chan string)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		e = ff.cmd.RunContext(ctx, args, outlog)
	}()
	for i2 := range outlog {
		log.Infow("runmsg", "log", strings.TrimSpace(i2))
	}
	wg.Wait()
	return e
}

// Error ...
func (ff *FFMpeg) Error() error {
	return ff.err
}

// NewFFMpeg ...
func NewFFMpeg(config *Config) *FFMpeg {
	ff := &FFMpeg{
		config: config,
		Name:   "ffmpeg",
	}

	return ff
}
