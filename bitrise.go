package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/bitrise-io/go-utils/v2/log"
)

/*
*

	{
	  "data": {
	    "abort_reason": {
	      "string": "string",
	      "valid": true
	    },
	    "branch": {
	      "string": "string",
	      "valid": true
	    },
	    "build_number": 0,
	    "commit_hash": {
	      "string": "string",
	      "valid": true
	    },
	    "commit_message": {
	      "string": "string",
	      "valid": true
	    },
	    "commit_view_url": {
	      "string": "string",
	      "valid": true
	    },
	    "credit_cost": {
	      "int64": 0,
	      "valid": true
	    },
	    "environment_prepare_finished_at": "string",
	    "finished_at": "string",
	    "is_on_hold": true,
	    "is_processed": true,
	    "is_status_sent": true,
	    "log_format": "string",
	    "machine_type_id": {
	      "string": "string",
	      "valid": true
	    },
	    "original_build_params": [
	      0
	    ],
	    "pipeline_workflow_id": "string",
	    "pull_request_id": 0,
	    "pull_request_target_branch": {
	      "string": "string",
	      "valid": true
	    },
	    "pull_request_view_url": {
	      "string": "string",
	      "valid": true
	    },
	    "slug": "string",
	    "stack_identifier": {
	      "string": "string",
	      "valid": true
	    },
	    "started_on_worker_at": "string",
	    "status": 0,
	    "status_text": "string",
	    "tag": {
	      "string": "string",
	      "valid": true
	    },
	    "triggered_at": "string",
	    "triggered_by": {
	      "string": "string",
	      "valid": true
	    },
	    "triggered_workflow": "string"
	  }
	}
*/

// StringValid represents a nested object with "string" and "valid" fields.
type StringValid struct {
	String string `json:"string,omitempty"`
	Valid  bool   `json:"valid,omitempty"`
}

// Int64Valid represents a nested object with "int64" and "valid" fields.
type Int64Valid struct {
	Int64 int64 `json:"int64,omitempty"`
	Valid bool  `json:"valid,omitempty"`
}

// BuildInfo represents the main structure of your JSON response.
type BuildInfo struct {
	AbortReason                  *StringValid `json:"abort_reason,omitempty"`
	Branch                       *StringValid `json:"branch,omitempty"`
	BuildNumber                  int          `json:"build_number,omitempty"`
	CommitHash                   *StringValid `json:"commit_hash,omitempty"`
	CommitMessage                *StringValid `json:"commit_message,omitempty"`
	CommitViewURL                *StringValid `json:"commit_view_url,omitempty"`
	CreditCost                   *Int64Valid  `json:"credit_cost,omitempty"`
	EnvironmentPrepareFinishedAt string       `json:"environment_prepare_finished_at,omitempty"`
	FinishedAt                   string       `json:"finished_at,omitempty"`
	IsOnHold                     bool         `json:"is_on_hold,omitempty"`
	IsProcessed                  bool         `json:"is_processed,omitempty"`
	IsStatusSent                 bool         `json:"is_status_sent,omitempty"`
	LogFormat                    string       `json:"log_format,omitempty"`
	MachineTypeID                *StringValid `json:"machine_type_id,omitempty"`
	OriginalBuildParams          []int        `json:"original_build_params,omitempty"`
	PipelineWorkflowID           string       `json:"pipeline_workflow_id,omitempty"`
	PullRequestID                int          `json:"pull_request_id,omitempty"`
	PullRequestTargetBranch      *StringValid `json:"pull_request_target_branch,omitempty"`
	PullRequestViewURL           *StringValid `json:"pull_request_view_url,omitempty"`
	Slug                         string       `json:"slug,omitempty"`
	StackIdentifier              *StringValid `json:"stack_identifier,omitempty"`
	StartedOnWorkerAt            string       `json:"started_on_worker_at,omitempty"`
	Status                       int          `json:"status,omitempty"`
	StatusText                   string       `json:"status_text,omitempty"`
	Tag                          StringValid  `json:"tag,omitempty"`
	TriggeredAt                  string       `json:"triggered_at,omitempty"`
	TriggeredBy                  StringValid  `json:"triggered_by,omitempty"`
	TriggeredWorkflow            string       `json:"triggered_workflow,omitempty"`
}

/*
{"triggered_at": "2024-01-03T02:06:49Z",
	"started_on_worker_at": "2024-01-03T02:06:51Z",
	"environment_prepare_finished_at": "2024-01-03T02:06:51Z",
	"finished_at": "2024-01-03T02:13:08Z",
	"slug": "6852c4a0-fa17-4faa-8e31-065bd43dc7cb",
	"status": 1,
	"status_text": "success",
	"abort_reason": null,
	"is_on_hold": false,
	"is_processed": true,
	"is_status_sent": false,
	"branch": "dev/24.0111",
	"build_number": 5619,
	"commit_hash": "",
	"commit_message": "",
	"tag": null,
	"triggered_workflow": "publish_debug_with_JDK11",
	"triggered_by": "manual-guangyu.wang",
	"machine_type_id": "g2-m1.8core",
	"stack_identifier": "osx-xcode-14.1.x-ventura",
	"original_build_params": {
		"branch": "dev/24.0111",
		"branch_dest": "",
		"branch_dest_repo_owner": "",
		"branch_repo_owner": "",
		"commit_hash": "",
		"commit_message": "",
		"pull_request_author": "",
		"pull_request_merge_branch": "",
		"pull_request_unverified_merge_branch": "",
		"pull_request_head_branch": "",
		"pull_request_id": null,
		"pull_request_repository_url": "",
		"tag": null,
		"environments": null,
		"workflow_id": "publish_debug_with_JDK11"
	},
	"pipeline_workflow_id": null,
	"pull_request_id": 0,
	"pull_request_target_branch": null,
	"pull_request_view_url": null,
	"commit_view_url": "https://bitbucket.org/xxxx/xxxx/commits/",
	"credit_cost": 28,
	"log_format": "json"
}*/
// define a struct to represent the BuildInfo2

// OriginalBuildParams represents the nested "original_build_params" object.
type OriginalBuildParams struct {
	Branch                           string        `json:"branch,omitempty"`
	BranchDest                       string        `json:"branch_dest,omitempty"`
	BranchDestRepoOwner              string        `json:"branch_dest_repo_owner,omitempty"`
	BranchRepoOwner                  string        `json:"branch_repo_owner,omitempty"`
	CommitHash                       string        `json:"commit_hash,omitempty"`
	CommitMessage                    string        `json:"commit_message,omitempty"`
	PullRequestAuthor                string        `json:"pull_request_author,omitempty"`
	PullRequestMergeBranch           string        `json:"pull_request_merge_branch,omitempty"`
	PullRequestUnverifiedMergeBranch string        `json:"pull_request_unverified_merge_branch,omitempty"`
	PullRequestHeadBranch            string        `json:"pull_request_head_branch,omitempty"`
	PullRequestID                    *int          `json:"pull_request_id,omitempty"`
	PullRequestRepositoryURL         string        `json:"pull_request_repository_url,omitempty"`
	Tag                              *string       `json:"tag,omitempty"`
	Environments                     []Environment `json:"environments,omitempty"`
	WorkflowID                       string        `json:"workflow_id,omitempty"`
}

type Environment struct {
	IsExpand bool   `json:"is_expand"`
	Key      string `json:"key"`
	Value    string `json:"value"`
}

// BuildInfo represents the main structure of your JSON response.
type Build struct {
	TriggeredAt                  string               `json:"triggered_at,omitempty"`
	StartedOnWorkerAt            string               `json:"started_on_worker_at,omitempty"`
	EnvironmentPrepareFinishedAt string               `json:"environment_prepare_finished_at,omitempty"`
	FinishedAt                   string               `json:"finished_at,omitempty"`
	Slug                         string               `json:"slug,omitempty"`
	Status                       int                  `json:"status,omitempty"`
	StatusText                   string               `json:"status_text,omitempty"`
	AbortReason                  *string              `json:"abort_reason,omitempty"`
	IsOnHold                     bool                 `json:"is_on_hold,omitempty"`
	IsProcessed                  bool                 `json:"is_processed,omitempty"`
	IsStatusSent                 bool                 `json:"is_status_sent,omitempty"`
	Branch                       string               `json:"branch,omitempty"`
	BuildNumber                  int                  `json:"build_number,omitempty"`
	CommitHash                   string               `json:"commit_hash,omitempty"`
	CommitMessage                string               `json:"commit_message,omitempty"`
	Tag                          *string              `json:"tag,omitempty"`
	TriggeredWorkflow            string               `json:"triggered_workflow,omitempty"`
	TriggeredBy                  string               `json:"triggered_by,omitempty"`
	MachineTypeID                string               `json:"machine_type_id,omitempty"`
	StackIdentifier              string               `json:"stack_identifier,omitempty"`
	OriginalBuildParams          *OriginalBuildParams `json:"original_build_params,omitempty"`
	PipelineWorkflowID           *string              `json:"pipeline_workflow_id,omitempty"`
	PullRequestID                int                  `json:"pull_request_id,omitempty"`
	PullRequestTargetBranch      *string              `json:"pull_request_target_branch,omitempty"`
	PullRequestViewURL           *string              `json:"pull_request_view_url,omitempty"`
	CommitViewURL                string               `json:"commit_view_url,omitempty"`
	CreditCost                   int                  `json:"credit_cost,omitempty"`
	LogFormat                    string               `json:"log_format,omitempty"`
}

func (b *Build) ChangeTimeZone() *Build {
	b.TriggeredAt = formatTime(b.TriggeredAt)
	b.FinishedAt = formatTime(b.FinishedAt)
	return b
}

func formatTime(timeStr string) string {
	// utcTime, err := time.Parse(time.RFC3339, dateTimeStr)
	utcTime, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return timeStr
	}
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return timeStr
	}
	estTime := utcTime.In(location)

	return estTime.String()
}

const (
	BitriseApiUrl = "https://api.bitrise.io/v0.1/"
)

type Response[T any] struct {
	Data    *T     `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

type BitriseApi struct {
	PersonalToken string
	logger        log.Logger
}

func (b *BitriseApi) GetBuildInfo(appSlug string, buildSlug string) (*Build, error) {

	req, err := http.NewRequest("GET", BitriseApiUrl+"apps/"+appSlug+"/builds/"+buildSlug, nil)
	if err != nil {
		b.logger.Errorf("Error occurred while creating HTTP request: %s", err)
	}
	res, err := newRequest[Build](b, req)
	if err != nil {
		b.logger.Errorf("Error occurred: %s", err)
		return nil, err
	}

	return res.Data, nil
}

// Set the content type to application/json
// Send the request
func newRequest[T any](b *BitriseApi, req *http.Request) (*Response[T], error) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", b.PersonalToken)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		b.logger.Errorf("Error occurred while sending request to webhook: %s", err)
		return nil, err
	}
	defer resp.Body.Close()
	bodyStr, _ := io.ReadAll(resp.Body)
	if BitriseDebugable {
		b.logger.Printf(string(bodyStr))
	}

	var data Response[T]
	if err := json.NewDecoder(bytes.NewReader(bodyStr)).Decode(&data); err != nil {
		b.logger.Errorf("Error occurred while decoding response: %s", err)
		return nil, err
	}
	return &data, nil
}
