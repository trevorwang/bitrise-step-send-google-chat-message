package main

import (
	"strings"

	exit "github.com/bitrise-io/go-utils/v2/exitcode"
)

func (r ChatRunner) Run(cfg *Config) exit.ExitCode {
	if cfg.WebhookUrl == "" {
		r.logger.Errorf("Google chat webhook url is empty")
		return exit.Failure
	}

	ops := []Option{}
	ops = r.createCardHeader(cfg, ops)
	ops = r.createCartText(cfg, ops)
	ops = r.createDecoratedTextList(cfg, ops)
	ops = r.createBuildInfo(cfg, ops)
	ops = r.createButtonList(cfg, ops)

	chatMsg := r.NewMessageWithOption(ops...)

	r.SendMessage(string(cfg.WebhookUrl), chatMsg)
	return exit.Success
}

func (ChatRunner) createCartText(cfg *Config, ops []Option) []Option {
	if cfg.MessageText != "" {
		ops = append(ops, WithCardText(cfg.MessageText))
	}
	return ops
}

func (ChatRunner) createCardHeader(cfg *Config, ops []Option) []Option {
	if cfg.MessageHeader != "" {
		ops = append(ops, WithCardHeader(cfg.MessageHeader))
	}
	return ops
}

func (r ChatRunner) createDecoratedTextList(cfg *Config, ops []Option) []Option {
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
	return ops
}

func (r ChatRunner) createBuildInfo(cfg *Config, ops []Option) []Option {
	if cfg.PersonalToken != "" {
		api := BitriseApi{string(cfg.PersonalToken), r.logger}

		build, err := api.GetBuildInfo(cfg.AppSlug, cfg.BuildSlug)
		if err != nil {
			r.logger.Errorf("Failed to get build info: %s", err)
		}
		r.logger.Debugf("build: %#v", build)
		build = build.ChangeTimeZone()

		clickIcon := "CLOCK"
		triggeredByIcon := "PERSON"
		ops = append(ops,
			WithCardDecoratedText(build.TriggeredBy, "Triggered By", &triggeredByIcon),
			WithCardDecoratedText(build.TriggeredAt, "Triggered At", &clickIcon),
		)
	}
	return ops
}

func (r ChatRunner) createButtonList(cfg *Config, ops []Option) []Option {
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
				link := strings.TrimSpace(decoratedText[1])
				if link == "" {
					r.logger.Errorf("invalid button link: %s", link)
				}
				ops = append(ops, WithCardButton(decoratedText[0], link))
			}
		}
	}
	return ops
}
