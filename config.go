package main

import "github.com/bitrise-io/go-steputils/stepconf"

type Inputs struct {
	WebhookUrl stepconf.Secret `env:"google_chat_webhook_url,required"`

	AppTitle string `env:"bitrise_app_title,required"`

	MessageHeader            string `env:"message_header"`
	MessageText              string `env:"message_text"`
	MessageDecoratedTextList string `env:"message_decorated_text_list"`
	MessageButtonList        string `env:"message_button_list"`

	PipelineBuildStatus    string `env:"bitriseio_pipeline_build_status"`
	BuildStatus            string `env:"bitrise_build_status,required"`
	TriggeredWorkflowTitle string `env:"bitrise_triggered_workflow_title,required"`

	PersonalToken stepconf.Secret `env:"bitrise_personal_access_token"`
	AppSlug       string          `env:"bitrise_app_slug,required"`
	BuildSlug     string          `env:"bitrise_build_slug,required"`
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
