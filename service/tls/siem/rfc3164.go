package siem

import (
	"errors"
	"fmt"
	"log/syslog"
	"strconv"
	"time"
)

func NewSyslogFormatRFC3164() *SyslogFormatRFC3164 {
	facility := syslog.LOG_USER
	severity := syslog.LOG_INFO

	return &SyslogFormatRFC3164{
		Facility:  &facility,
		Severity:  &severity,
		Timestamp: time.Now(),
		Hostname:  "-",
		Pid:       "-",
		Process:   "-",
	}
}

type SyslogFormatRFC3164 struct {
	Priority  string
	Facility  *syslog.Priority
	Severity  *syslog.Priority
	Timestamp time.Time
	Hostname  string
	Pid       string
	Process   string
	Message   string
}

func (f *SyslogFormatRFC3164) GetDefaultPriority() string {
	return strconv.Itoa(int(*f.Facility | *f.Severity))
}

func (f *SyslogFormatRFC3164) validate() error {
	if f.Priority == "" {
		if f.Facility != nil && f.Severity != nil {
			f.Priority = f.GetDefaultPriority()
		}
		return errors.New("syslog priority is required")
	}

	if f.Timestamp.IsZero() {
		return errors.New("syslog timestamp is required")
	}

	if f.Hostname == "" {
		return errors.New("syslog hostname is required")
	}

	if f.Pid == "" {
		return errors.New("syslog pid is required")
	}

	if f.Process == "" {
		return errors.New("syslog process is required")
	}

	if f.Message == "" {
		return errors.New("syslog message is required")
	}

	return nil
}

func (f *SyslogFormatRFC3164) FormatSyslogMessageWithRFC3164() (string, error) {
	if err := f.validate(); err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"<%s>%s %s %s[%s]: %s\n",
		f.Priority,
		f.Timestamp.Format(time.Stamp),
		f.Hostname,
		f.Process,
		f.Pid,
		f.Message,
	), nil
}
