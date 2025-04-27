package siem

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

type SyslogClient struct {
	conn                 net.Conn
	SslSettings          *tls.Config
	Host                 string
	Port                 int
	Protocol             string
	ConnTimeout          time.Duration
	SendTimeout          time.Duration
	MaxMessageLength     *int
	RetrySendLogWaitTime time.Duration
}

func NewSyslogClient() *SyslogClient {
	defaultMaxMessageLength := 1024

	return &SyslogClient{
		Protocol:             "udp",
		ConnTimeout:          30 * time.Second,
		SendTimeout:          5 * time.Second,
		MaxMessageLength:     &defaultMaxMessageLength,
		RetrySendLogWaitTime: 200 * time.Millisecond,
	}
}

func (c *SyslogClient) validBaseClientSettings() error {
	if c.Host == "" || c.Port == 0 {
		return errors.New("host or port is not set")
	}

	protocol := strings.ToLower(c.Protocol)
	if protocol != "tcp" && protocol != "udp" {
		return fmt.Errorf("unsupported protocol: %s", c.Protocol)
	}
	c.Protocol = protocol

	if *c.MaxMessageLength <= 0 {
		return fmt.Errorf("invalid max message length: %d", *c.MaxMessageLength)
	}

	if c.ConnTimeout <= 0 {
		return fmt.Errorf("invalid conn timeout: %d", c.ConnTimeout)
	}

	if c.SendTimeout <= 0 {
		return fmt.Errorf("invalid send timeout: %d", c.SendTimeout)
	}

	return nil
}

func (c *SyslogClient) Connect() error {
	if c.conn != nil {
		return nil
	}

	err := c.validBaseClientSettings()
	if err != nil {
		return fmt.Errorf("failed to valid client setting, err: %v", err)
	}

	address := net.JoinHostPort(c.Host, strconv.Itoa(c.Port))

	dialer := &net.Dialer{
		Timeout: c.ConnTimeout,
	}

	var conn net.Conn

	if c.SslSettings != nil && c.Protocol == "tcp" {
		conn, err = tls.DialWithDialer(dialer, "tcp", address, c.SslSettings)
	} else {
		conn, err = dialer.Dial(c.Protocol, address)
	}

	if err != nil {
		return fmt.Errorf("failed to connect to %s over %s: %w", address, c.Protocol, err)
	}

	c.conn = conn

	return nil
}

func (c *SyslogClient) Close() error {
	if c.conn == nil {
		return nil
	}

	err := c.conn.Close()

	c.conn = nil

	return err
}

func (c *SyslogClient) Send(message string) error {
	return c.sendWithRetry(message)
}

func (c *SyslogClient) sendWithRetry(message string) error {
	if c.conn == nil {
		return errors.New("syslog client not initialized")
	}

	byteMessage := []byte(message)

	if c.MaxMessageLength != nil && len(byteMessage) > *c.MaxMessageLength {
		byteMessage = byteMessage[:*c.MaxMessageLength]
	}

	if err := c.conn.SetWriteDeadline(time.Now().Add(c.SendTimeout)); err != nil {
		return fmt.Errorf("failed to set write deadline: %v", err)
	}

	_, err := c.conn.Write(byteMessage)
	if err != nil {

		time.Sleep(c.RetrySendLogWaitTime)

		_, err = c.conn.Write(byteMessage)
		if err != nil {
			return fmt.Errorf("failed to write message: %v", err)
		}
	}

	return nil
}
