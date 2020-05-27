package main

import (
	"time"
)

// APIAlert represents an notifier.AlertingRule state
// for WEB view
type APIAlert struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	GroupID     string            `json:"group_id"`
	Expression  string            `json:"expression"`
	State       string            `json:"state"`
	Value       string            `json:"value"`
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
	ActiveAt    time.Time         `json:"activeAt"`
}

type APIGroup struct {
	Name           string
	ID             string
	File           string
	AlertingRules  []APIAlertingRule
	RecordingRules []APIRecordingRule
}

type APIAlertingRule struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	GroupID     string            `json:"group_id"`
	Expression  string            `json:"expression"`
	For         time.Duration     `json:"for"`
	LastError   string            `json:"last_error"`
	LastExec    time.Time         `json:"last_exec"`
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
}

type APIRecordingRule struct {
	ID         string            `json:"id"`
	Name       string            `json:"name"`
	GroupID    string            `json:"group_id"`
	Expression string            `json:"expression"`
	LastError  string            `json:"last_error"`
	LastExec   time.Time         `json:"last_exec"`
	Labels     map[string]string `json:"labels"`
}
