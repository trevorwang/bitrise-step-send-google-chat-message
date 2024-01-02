package main

import (
	"fmt"

	exit "github.com/bitrise-io/go-utils/v2/exitcode"
)

func (r ChatRunner) Run(cfg *Config) exit.ExitCode {
	if cfg.WebhookUrl == "" {
		r.logger.Errorf("Google chat webhook url is empty")
		return exit.Failure
	}
	if cfg.PersonalToken == "" {
		r.logger.Errorf("Personal access token is empty")
		return exit.Failure
	}

	api := BitriseApi{string(cfg.PersonalToken), r.logger}

	// Get build info
	build, err := api.GetBuildInfo(cfg.AppSlug, cfg.BuildSlug)
	if err != nil {
		r.logger.Errorf("Failed to get build info: %s", err)
	}
	r.logger.Debugf("build: %#v", build)
	build = build.ChangeTimeZone()

	r.logger.Debugf("webhook url: %#v", cfg.WebhookUrl)

	text := fmt.Sprintf("%s is built successfully", cfg.AppTitle)

	ops := []Option{}

	if cfg.AppCenterId != "" {
		ops = append(ops, WithCardText(fmt.Sprintf("%s has been released to AppCenter. And the download link is here: %s", cfg.AppTitle, cfg.AppCenterUrl)))
		ops = append(ops, WithCardIconText("AppCenter App", cfg.AppCenterId))
	}
	ops = append(ops,
		WithMessageText(text),
		WithCardHeader(text),
		WithCardIconText("Branch", cfg.GitBranch),
		WithCardIconText("Triggered By", build.TriggeredBy),
		WithCardIconText("Triggered At", build.TriggeredAt),
		WithCardIconText("Last Commit Message", cfg.GitMessage),
	)
	if cfg.AppCenterUrl != "" {
		ops = append(ops, WithCardButton("View build", cfg.AppCenterUrl))
	}

	chatMsg := r.NewMessageWithOption(ops...)

	r.SendMessage(cfg.WebhookUrl, chatMsg)
	return exit.Success
}
