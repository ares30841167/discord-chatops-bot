package model

type PipelineTriggerResponse struct {
	ID             int64          `json:"id"`
	Iid            int64          `json:"iid"`
	ProjectID      int64          `json:"project_id"`
	SHA            string         `json:"sha"`
	Ref            string         `json:"ref"`
	Status         string         `json:"status"`
	Source         string         `json:"source"`
	CreatedAt      string         `json:"created_at"`
	UpdatedAt      string         `json:"updated_at"`
	WebURL         string         `json:"web_url"`
	BeforeSHA      string         `json:"before_sha"`
	Tag            bool           `json:"tag"`
	YAMLErrors     interface{}    `json:"yaml_errors"`
	User           User           `json:"user"`
	StartedAt      interface{}    `json:"started_at"`
	FinishedAt     interface{}    `json:"finished_at"`
	CommittedAt    interface{}    `json:"committed_at"`
	Duration       interface{}    `json:"duration"`
	QueuedDuration interface{}    `json:"queued_duration"`
	Coverage       interface{}    `json:"coverage"`
	DetailedStatus DetailedStatus `json:"detailed_status"`
}

type DetailedStatus struct {
	Icon         string      `json:"icon"`
	Text         string      `json:"text"`
	Label        string      `json:"label"`
	Group        string      `json:"group"`
	Tooltip      string      `json:"tooltip"`
	HasDetails   bool        `json:"has_details"`
	DetailsPath  string      `json:"details_path"`
	Illustration interface{} `json:"illustration"`
	Favicon      string      `json:"favicon"`
}

type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Name      string `json:"name"`
	State     string `json:"state"`
	AvatarURL string `json:"avatar_url"`
	WebURL    string `json:"web_url"`
}
