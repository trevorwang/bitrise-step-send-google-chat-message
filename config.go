package main

import "github.com/bitrise-io/go-steputils/stepconf"

type Inputs struct {
	WebhookUrl        stepconf.Secret `env:"webhook_url,required"`
	WebhookUrlOnError stepconf.Secret `env:"webhook_url_on_error"`

	AppTitle string `env:"app_title,required"`

	Text        string `env:"text"`
	TextOnError string `env:"text_on_error"`

	IconUrl               string `env:"icon_url"`
	IconUrlOnErorr        string `env:"icon_url_on_error"`
	CardHeader            string `env:"card_header"`
	CardText              string `env:"card_text"`
	ImageUrl              string `env:"image_url"`
	ImageUrlOnError       string `env:"image_url_on_error"`
	CardDecoratedTextList string `env:"card_decorated_text_list"`
	CardButtonList        string `env:"card_buttons"`

	PipelineBuildStatus    string `env:"pipeline_build_status"`
	BuildStatus            string `env:"build_status,required"`
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
