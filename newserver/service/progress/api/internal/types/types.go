// Code generated by goctl. DO NOT EDIT.
package types

type Result struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type Notification struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Library struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Url   string `json:"url"`
	Owner string `json:"owner"`
}

type TopicReport struct {
	ID          int64  `json:"id"`
	TopicID     int64  `json:"topicID"`
	Description string `json:"description"`
	ReportUrl   string `json:"resultUrl"`
	StageID     int64  `json:"stageID"`
}

type Event struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	SchoolYear string `json:"schoolYear"`
	IsCurrent  int64  `json:"isCurrent"`
}

type Stage struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	FacultyID   int64  `json:"facultyID"`
	Url         string `json:"url"`
	EventID     int64  `json:"eventID"`
	TimeStart   int64  `json:"timeStart"`
	TimeEnd     int64  `json:"timeEnd"`
}

type CreateTopicReportReq struct {
	TopicID     int64  `json:"topicID"`
	Description string `json:"description"`
	ReportUrl   string `json:"reportUrl"`
	StageID     int64  `json:"stageID"`
}

type CreateTopicReportRes struct {
	Result Result `json:"result"`
}

type GetTopicReportsReq struct {
	StageID int64 `json:"stageID"`
	EventID int64 `json:"eventID"`
	TopicID int64 `json:"topicID"`
	Limit   int64 `json:"limit"`
	Offset  int64 `json:"offset"`
}

type GetTopicReportsRes struct {
	Result       Result        `json:"result"`
	Total        int64         `json:"total"`
	TopicReports []TopicReport `json:"topicReports"`
}

type CreateEventReq struct {
	Name       string `json:"name"`
	SchoolYear string `json:"schoolYear"`
}

type CreateEventRes struct {
	Result       Result `json:"result"`
	CurrentEvent Event  `json:"currentEvent"`
}

type GetEventsReq struct {
}

type GetEventsRes struct {
	Result Result  `json:"result"`
	Total  int64   `json:"total"`
	Events []Event `json:"events"`
}

type UpdateCurrentEventReq struct {
	EventID int64 `path:"eventID"`
}

type UpdateCurrentEventRes struct {
	Result Result `json:"result"`
}

type CreateStageReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Url         string `json:"url"`
	EventID     int64  `json:"eventID"`
	TimeStart   int64  `json:"timeStart"`
	TimeEnd     int64  `json:"timeEnd"`
	FacultyID   int64  `json:"facultyID"`
}

type CreateStageRes struct {
	Result Result `json:"result"`
}

type GetStagesReq struct {
	EventID   int64 `json:"eventID"`
	FacultyID int64 `json:"facultyID"`
}

type GetStagesRes struct {
	Result Result  `json:"result"`
	Total  int64   `json:"total"`
	Stages []Stage `json:"stages"`
}
