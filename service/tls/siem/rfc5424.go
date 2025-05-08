package siem

import (
	"errors"
	"fmt"
	"log/syslog"
	"strconv"
	"time"
)

func NewSyslogFormatRFC5424() *SyslogFormatRFC5424 {
	facility := syslog.LOG_USER
	severity := syslog.LOG_INFO

	return &SyslogFormatRFC5424{
		Facility:  &facility,
		Severity:  &severity,
		Version:   "1",
		Timestamp: time.Now(),
		Hostname:  "-",
		AppName:   "-",
		Pid:       "-",
		MsgId:     "-",
	}
}

type SyslogFormatRFC5424 struct {
	Priority  string
	Facility  *syslog.Priority
	Severity  *syslog.Priority
	Version   string
	Timestamp time.Time
	Hostname  string
	AppName   string
	Pid       string
	MsgId     string
	Message   string
}

func (f *SyslogFormatRFC5424) GetDefaultPriority() string {
	return strconv.Itoa(int(*f.Facility | *f.Severity))
}

func (f *SyslogFormatRFC5424) validate() error {
	if f.Priority == "" {
		if f.Facility != nil && f.Severity != nil {
			f.Priority = f.GetDefaultPriority()
		}
		return errors.New("syslog priority is required")
	}

	if f.Version == "" {
		return errors.New("syslog version is required")
	}

	if f.Timestamp.IsZero() {
		return errors.New("syslog timestamp is required")
	}

	if f.Hostname == "" {
		return errors.New("syslog hostname is required")
	}

	if f.AppName == "" {
		return errors.New("syslog app name is required")
	}

	if f.Pid == "" {
		return errors.New("syslog pid is required")
	}

	if f.MsgId == "" {
		return errors.New("syslog message id is required")
	}

	if f.Message == "" {
		return errors.New("syslog message is required")
	}

	return nil
}

func (f *SyslogFormatRFC5424) FormatSyslogMessageWithRFC5424() (string, error) {
	if err := f.validate(); err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"<%s>%s %s %s %s %s %s - %s\n",
		f.Priority,
		f.Version,
		f.Timestamp.Format(time.RFC3339),
		f.Hostname,
		f.AppName,
		f.Pid,
		f.MsgId,
		f.Message,
	), nil
}
