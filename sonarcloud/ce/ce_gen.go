package ce

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// ActivityRequest Search for tasks.<br> Either componentId or component can be provided, but not both.<br> Requires the project administration permission if componentId or component is set.
type ActivityRequest struct {
	Component      string `form:"component,omitempty"`      // Key of the component (project) to filter on
	ComponentId    string `form:"componentId,omitempty"`    // Id of the component (project) to filter on
	MaxExecutedAt  string `form:"maxExecutedAt,omitempty"`  // Maximum date of end of task processing (inclusive)
	MinSubmittedAt string `form:"minSubmittedAt,omitempty"` // Minimum date of task submission (inclusive)
	OnlyCurrents   string `form:"onlyCurrents,omitempty"`   // Filter on the last tasks (only the most recent finished task by project)
	Q              string `form:"q,omitempty"`              // Limit search to: <ul><li>component names that contain the supplied string</li><li>component keys that are exactly the same as the supplied string</li><li>task ids that are exactly the same as the supplied string</li></ul>Must not be set together with componentId
	Status         string `form:"status,omitempty"`         // Comma separated list of task statuses
	Type           string `form:"type,omitempty"`           // Task type
}

// ActivityResponse is the response for ActivityRequest
type ActivityResponse struct {
	Tasks []struct {
		AnalysisId         string  `json:"analysis_id,omitempty"`
		ComponentId        string  `json:"component_id,omitempty"`
		ComponentKey       string  `json:"component_key,omitempty"`
		ComponentName      string  `json:"component_name,omitempty"`
		ComponentQualifier string  `json:"component_qualifier,omitempty"`
		ExecutedAt         string  `json:"executed_at,omitempty"`
		ExecutionTimeMs    float64 `json:"execution_time_ms,omitempty"`
		HasErrorStacktrace bool    `json:"has_error_stacktrace,omitempty"`
		HasScannerContext  bool    `json:"has_scanner_context,omitempty"`
		Id                 string  `json:"id,omitempty"`
		Logs               bool    `json:"logs,omitempty"`
		Organization       string  `json:"organization,omitempty"`
		StartedAt          string  `json:"started_at,omitempty"`
		Status             string  `json:"status,omitempty"`
		SubmittedAt        string  `json:"submitted_at,omitempty"`
		SubmitterLogin     string  `json:"submitter_login,omitempty"`
		Type               string  `json:"type,omitempty"`
		ErrorMessage       string  `json:"error_message,omitempty"`
		TaskType           string  `json:"task_type,omitempty"`
	} `json:"tasks,omitempty"`
}

// ActivityStatusRequest Returns CE activity related metrics.<br>Requires 'Administer' permission on the specified project.
type ActivityStatusRequest struct {
	ComponentId  string `form:"componentId,omitempty"`  // Id of the component (project) to filter on
	ComponentKey string `form:"componentKey,omitempty"` // Key of the component (project) to filter on
}

// ActivityStatusResponse is the response for ActivityStatusRequest
type ActivityStatusResponse struct {
	Failing     float64 `json:"failing,omitempty"`
	InProgress  float64 `json:"in_progress,omitempty"`
	Pending     float64 `json:"pending,omitempty"`
	PendingTime float64 `json:"pending_time,omitempty"`
}

// ComponentRequest Get the pending tasks, in-progress tasks and the last executed task of a given component (usually a project).<br>Requires the following permission: 'Browse' on the specified component.<br>Either 'componentId' or 'component' must be provided.
type ComponentRequest struct {
	Component   string `form:"component,omitempty"`   //
	ComponentId string `form:"componentId,omitempty"` //
}

// ComponentResponse is the response for ComponentRequest
type ComponentResponse struct {
	Current struct {
		AnalysisId         string  `json:"analysis_id,omitempty"`
		ComponentId        string  `json:"component_id,omitempty"`
		ComponentKey       string  `json:"component_key,omitempty"`
		ComponentName      string  `json:"component_name,omitempty"`
		ComponentQualifier string  `json:"component_qualifier,omitempty"`
		ErrorMessage       string  `json:"error_message,omitempty"`
		ErrorType          string  `json:"error_type,omitempty"`
		ExecutionTimeMs    float64 `json:"execution_time_ms,omitempty"`
		FinishedAt         string  `json:"finished_at,omitempty"`
		HasErrorStacktrace bool    `json:"has_error_stacktrace,omitempty"`
		HasScannerContext  bool    `json:"has_scanner_context,omitempty"`
		Id                 string  `json:"id,omitempty"`
		Logs               bool    `json:"logs,omitempty"`
		Organization       string  `json:"organization,omitempty"`
		StartedAt          string  `json:"started_at,omitempty"`
		Status             string  `json:"status,omitempty"`
		SubmittedAt        string  `json:"submitted_at,omitempty"`
		Type               string  `json:"type,omitempty"`
	} `json:"current,omitempty"`
	Queue []struct {
		ComponentId        string `json:"component_id,omitempty"`
		ComponentKey       string `json:"component_key,omitempty"`
		ComponentName      string `json:"component_name,omitempty"`
		ComponentQualifier string `json:"component_qualifier,omitempty"`
		Id                 string `json:"id,omitempty"`
		Logs               bool   `json:"logs,omitempty"`
		Organization       string `json:"organization,omitempty"`
		Status             string `json:"status,omitempty"`
		SubmittedAt        string `json:"submitted_at,omitempty"`
		Type               string `json:"type,omitempty"`
	} `json:"queue,omitempty"`
}

// TaskRequest Give Compute Engine task details such as type, status, duration and associated component.<br />Requires 'Execute Analysis' permission.
type TaskRequest struct {
	AdditionalFields string `form:"additionalFields,omitempty"` // Comma-separated list of the optional fields to be returned in response.
	Id               string `form:"id,omitempty"`               // Id of task
}

// TaskResponse is the response for TaskRequest
type TaskResponse struct {
	Task struct {
		AnalysisId         string  `json:"analysis_id,omitempty"`
		ComponentId        string  `json:"component_id,omitempty"`
		ComponentKey       string  `json:"component_key,omitempty"`
		ComponentName      string  `json:"component_name,omitempty"`
		ComponentQualifier string  `json:"component_qualifier,omitempty"`
		ErrorMessage       string  `json:"error_message,omitempty"`
		ErrorStacktrace    string  `json:"error_stacktrace,omitempty"`
		ExecutedAt         string  `json:"executed_at,omitempty"`
		ExecutionTimeMs    float64 `json:"execution_time_ms,omitempty"`
		HasErrorStacktrace bool    `json:"has_error_stacktrace,omitempty"`
		HasScannerContext  bool    `json:"has_scanner_context,omitempty"`
		Id                 string  `json:"id,omitempty"`
		Logs               bool    `json:"logs,omitempty"`
		Organization       string  `json:"organization,omitempty"`
		ScannerContext     string  `json:"scanner_context,omitempty"`
		StartedAt          string  `json:"started_at,omitempty"`
		Status             string  `json:"status,omitempty"`
		SubmittedAt        string  `json:"submitted_at,omitempty"`
		Type               string  `json:"type,omitempty"`
	} `json:"task,omitempty"`
}
