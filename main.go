package main

import (
	"os"

	"github.com/bitrise-io/go-steputils/v2/stepconf"
	"github.com/bitrise-io/go-utils/pathutil"
	"github.com/bitrise-io/go-utils/v2/command"
	"github.com/bitrise-io/go-utils/v2/env"
	exit "github.com/bitrise-io/go-utils/v2/exitcode"
	"github.com/bitrise-io/go-utils/v2/log"
)

func main() {
	exitCode := run()
	os.Exit(int(exitCode))
}

func run() exit.ExitCode {
	logger := log.NewLogger()
	logger.EnableDebugLog(Debugable)
	runner := createStep(logger)
	cfg, err := runner.ProcessConfig()
	if err != nil {
		logger.Errorf("Found error: %s", err)
		return exit.Failure
	}
	logger.Debugf("cfg: %#v", cfg)
	runner.Run(&cfg)
	return exit.Success
}

func createStep(logger log.Logger) *ChatRunner {
	envRepository := env.NewRepository()
	inputParser := stepconf.NewInputParser(envRepository)
	cmdFactory := command.NewFactory(envRepository)
	cmdLocator := env.NewCommandLocator()
	pathModifier := pathutil.NewPathModifier()

	return &ChatRunner{
		logger:       logger,
		inputParser:  inputParser,
		cmdFactory:   cmdFactory,
		cmdLocator:   cmdLocator,
		pathModifier: pathModifier,
	}

}

type ChatRunner struct {
	logger       log.Logger
	cmdFactory   command.Factory
	inputParser  stepconf.InputParser
	cmdLocator   env.CommandLocator
	pathModifier pathutil.PathModifier
}
