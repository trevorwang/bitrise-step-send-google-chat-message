package main

import (
	"fmt"
	"strings"

	exit "github.com/bitrise-io/go-utils/v2/exitcode"
)

func (r ChatRunner) Run(cfg *Config) exit.ExitCode {
	if cfg.WebhookUrl == "" {
		r.logger.Errorf("Google chat webhook url is empty")
		return exit.Failure
	}

	ops := []Option{}

	if cfg.MessageHeader != "" {
		ops = append(ops, WithCardHeader(cfg.MessageHeader))
	}

	if cfg.MessageText != "" {
		ops = append(ops, WithCardText(cfg.MessageText))
	}

	if cfg.MessageDecoratedTextList != "" {
		lines := strings.Split(cfg.MessageDecoratedTextList, "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" {
				break
			}
			decoratedText := strings.Split(line, "|")
			if len(decoratedText) < 2 {
				r.logger.Errorf("Invalid decorated text format: %s", line)
				break
			}
			if len(decoratedText) == 2 {
				ops = append(ops, WithCardDecoratedText(decoratedText[0], decoratedText[1], nil))
			} else if len(decoratedText) == 3 {
				ops = append(ops, WithCardDecoratedText(decoratedText[0], decoratedText[1], &decoratedText[2]))
			}

		}
	}

	if cfg.MessageButtonList != "" {
		lines := strings.Split(cfg.MessageButtonList, "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			r.logger.Debugf("button item : %s", line)
			if line == "" {
				break
			}
			decoratedText := strings.Split(line, "|")
			if len(decoratedText) < 2 {
				r.logger.Errorf("Invalid button format: %s", line)
				break
			}
			if len(decoratedText) == 2 {

				if strings.TrimSpace(decoratedText[1]) == "" {
					r.logger.Errorf("invalid button link: %s", decoratedText[1])
				}

				ops = append(ops, WithCardButton(decoratedText[0], decoratedText[1]))
				// } else if len(decoratedText) == 3 {
				// ops = append(ops, WithCardDecoratedText(decoratedText[0], &decoratedText[1], &decoratedText[2]))
			}

		}
	}

	if cfg.PersonalToken != "" {
		api := BitriseApi{string(cfg.PersonalToken), r.logger}

		// Get build info
		build, err := api.GetBuildInfo(cfg.AppSlug, cfg.BuildSlug)
		if err != nil {
			r.logger.Errorf("Failed to get build info: %s", err)
		}
		r.logger.Debugf("build: %#v", build)
		build = build.ChangeTimeZone()

		ops = append(ops,
			WithCardDecoratedText(build.TriggeredBy, "Triggered By", nil),
			WithCardDecoratedText(build.TriggeredAt, "Triggered At", nil),
		)
	}

	r.logger.Debugf("webhook url: %#v", cfg.WebhookUrl)

	text := fmt.Sprintf("%s is built successfully", cfg.AppTitle)

	// if cfg.AppCenterId != "" {
	// 	ops = append(ops, WithCardText(fmt.Sprintf("%s has been released to AppCenter. And the download link is here: %s", cfg.AppTitle, cfg.AppCenterUrl)))
	// 	ops = append(ops, WithCardIconText("AppCenter App", cfg.AppCenterId))
	// }
	ops = append(ops,
		// WithMessageText(text),
		WithCardHeader(text),
		// WithCardDecoratedText(cfg.GitBranch, "Branch", nil),
		// WithCardDecoratedText(cfg.GitMessage, "Last Commit Message", nil),
	)

	chatMsg := r.NewMessageWithOption(ops...)

	r.SendMessage(string(cfg.WebhookUrl), chatMsg)
	return exit.Success
}
