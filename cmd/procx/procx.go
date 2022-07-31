package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/robertlestak/procx/internal/flags"
	"github.com/robertlestak/procx/pkg/drivers"
	"github.com/robertlestak/procx/pkg/procx"
	log "github.com/sirupsen/logrus"
)

var (
	Version      = "dev"
	AppName      = "procx"
	EnvKeyPrefix = fmt.Sprintf("%s_", strings.ToUpper(AppName))
)

func init() {
	ll, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		ll = log.InfoLevel
	}
	log.SetLevel(ll)
}

func printVersion() {
	fmt.Printf(AppName+" version %s\n", Version)
}

func LoadEnv(prefix string) error {
	if os.Getenv(prefix+"DRIVER") != "" {
		d := os.Getenv(prefix + "DRIVER")
		flags.Driver = &d
	}
	if os.Getenv(prefix+"HOSTENV") != "" {
		h := os.Getenv(prefix + "HOSTENV")
		t := h == "true"
		flags.HostEnv = &t
	}
	if os.Getenv(prefix+"PASS_WORK_AS_ARG") != "" {
		r := os.Getenv(prefix + "PASS_WORK_AS_ARG")
		t := r == "true"
		flags.PassWorkAsArg = &t
	}
	if os.Getenv(prefix+"DAEMON") != "" {
		r := os.Getenv(prefix + "DAEMON")
		t := r == "true"
		flags.Daemon = &t
	}
	return nil
}

func run() {
	l := log.WithFields(log.Fields{
		"app": AppName,
	})
	l.Debug("start")
	j := &procx.ProcX{
		DriverName:    drivers.DriverName(*flags.Driver),
		HostEnv:       *flags.HostEnv,
		PassWorkAsArg: *flags.PassWorkAsArg,
	}
	if err := j.Init(EnvKeyPrefix); err != nil {
		l.WithError(err).Error("InitDriver")
		os.Exit(1)
	}
	if err := j.DoWork(); err != nil {
		l.Errorf("failed to do work: %s", err)
		os.Exit(1)
	}
}

func main() {
	l := log.WithFields(log.Fields{
		"app": AppName,
	})
	l.Debug("start")
	if len(os.Args) > 1 {
		if os.Args[1] == "--version" || os.Args[1] == "-v" {
			printVersion()
			os.Exit(0)
		}
	}
	flags.FlagSet.Parse(os.Args[1:])
	if err := LoadEnv(EnvKeyPrefix); err != nil {
		l.Error(err)
		os.Exit(1)
	}
	l.Debug("parsed flags")
	args := flags.FlagSet.Args()
	if len(args) == 0 {
		printVersion()
		flags.FlagSet.PrintDefaults()
		os.Exit(1)
	}
	if *flags.Daemon {
		l.Debug("running as daemon")
		for {
			run()
		}
	} else {
		run()
	}
	l.Debug("exited")
}