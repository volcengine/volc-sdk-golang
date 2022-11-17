package cloudtrail

import "github.com/volcengine/volc-sdk-golang/base"

type NullResultResp struct {
	ResponseMetadata *base.ResponseMetadata
}

// events
type LookupEventsReq struct {
	NextToken        string            `json:",omitempty"`
	MaxResults       int               `json:",omitempty"`
	StartTime        int64             `json:",omitempty"`
	EndTime          int64             `json:",omitempty"`
	LookupConditions []LookupCondition `json:",omitempty"`
}

type LookupCondition struct {
	LookupConditionKey   string
	LookupConditionValue string
}

type LookupEventsResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *EventList `json:",omitempty"`
}

type EventList struct {
	NextToken string
	Trails    []*TrailEvent
}

type TrailEvent struct {
	EventID         string
	EventTime       string
	EventName       string
	RequestID       string
	SourceIPAddress string
	EventSource     string
	UserName        string
	AccessKeyID     string
	EventDetail     string
	Region          string
	ErrorCode       string
}
