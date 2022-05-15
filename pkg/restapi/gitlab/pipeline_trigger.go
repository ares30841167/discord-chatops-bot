package gitlab

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"guanyu.dev/chatopsbot/pkg/restapi/gitlab/model"
)

type PipelineTriggerService struct {
	client *Client
}

func init() {
	// Check whether gitlab base url exist or not
	if os.Getenv("GITLAB_BASE_URL") == "" {
		fmt.Fprintln(os.Stderr, "No gitlab base url provided.")
		os.Exit(1)
	}

	// Check whether trigger token exist or not
	if os.Getenv("CI_TRIGGER_TOKEN") == "" {
		fmt.Fprintln(os.Stderr, "No trigger token provided.")
		os.Exit(1)
	}

	// Check whether project id exist or not
	if os.Getenv("PROJECT_ID") == "" {
		fmt.Fprintln(os.Stderr, "No project id provided.")
		os.Exit(1)
	}
}

// Initial a new pipeline trigger service instance
func NewPipelineTriggerService() *PipelineTriggerService {
	client, err := NewClient(
		&url.URL{
			Path: fmt.Sprintf("https://%s/api/v4/projects/%s/trigger/pipeline",
				os.Getenv("GITLAB_BASE_URL"), os.Getenv("PROJECT_ID")),
		},
	)
	if err != nil {
		fmt.Println(err.Error())
	}
	return &PipelineTriggerService{
		client: client,
	}
}

// Trigger a pipeline within a certain branch with variables
func (pts *PipelineTriggerService) TriggerPipeline(branch string, job string, variables map[string]string) (*model.PipelineTriggerResponse, error) {
	form := map[string]string{
		"token":          os.Getenv("CI_TRIGGER_TOKEN"),
		"ref":            branch,
		"variables[JOB]": job,
	}
	// Fill the variables into form in variables[$VAR] format
	for k, v := range variables {
		varKey := fmt.Sprintf("variables[%s]", k)
		form[varKey] = v
	}
	// Send HTTPS request
	res, err := pts.client.PostForm(form)
	if err != nil {
		return nil, err
	}
	// Unmarshal response body
	response := &model.PipelineTriggerResponse{}
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bodyBytes, response)
	if err != nil {
		return nil, err
	}
	// Return the result
	return response, nil
}
