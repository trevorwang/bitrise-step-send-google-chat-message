package main

import "github.com/bitrise-io/go-steputils/stepconf"

type Inputs struct {
	WebhookUrl    string          `env:"google_chat_webhook_url,required"`
	PersonalToken stepconf.Secret `env:"bitrise_personal_access_token,required"`
	AppSlug       string          `env:"bitrise_app_slug,required"`
	BuildSlug     string          `env:"bitrise_build_slug,required"`
	GitMessage    string          `env:"bitrise_git_message"`
	AppTitle      string          `env:"bitrise_app_title,required"`
	GitBranch     string          `env:"bitrise_git_branch,required"`
	AppCenterId   string          `env:"appcenter_deploy_release_id"`
	AppCenterUrl  string          `env:"appcenter_release_page_url"`
}

type Config struct {
	Inputs
}

func (r ChatRunner) ProcessConfig() (Config, error) {
	var inputs Inputs
	if err := r.inputParser.Parse(&inputs); err != nil {
		return Config{}, err
	}
	stepconf.Print(inputs)
	config := Config{Inputs: inputs}
	return config, nil
}
