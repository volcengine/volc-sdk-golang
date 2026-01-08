package tls

import (
	"net"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type SDKClientTestSuite struct {
	suite.Suite

	cli Client
}

func TestSDKClientTestSuite(t *testing.T) {
	suite.Run(t, new(SDKClientTestSuite))
}

func (suite *SDKClientTestSuite) SetupTest() {
	suite.cli = NewClientWithEnv()
}

func (suite *SDKClientTestSuite) TearDownTest() {

}

func (suite *SDKClientTestSuite) TestNewClient() {
	suite.Equal(defaultHttpClient, suite.cli.GetHttpClient())
}

func (suite *SDKClientTestSuite) TestSetHttpClient() {

	timeout := time.Second * 1
	transport := &http.Transport{
		MaxIdleConns:        10,
		MaxIdleConnsPerHost: 5,
		IdleConnTimeout:     1 * time.Second,
		DisableCompression:  false,
		DialContext: (&net.Dialer{
			Timeout:   1 * time.Second,
			KeepAlive: 1 * time.Second,
		}).DialContext,
	}

	suite.cli.SetHttpClient(&http.Client{
		Timeout:   timeout,
		Transport: transport,
	})

	suite.NotEqual(defaultHttpClient, suite.cli.GetHttpClient())
	suite.Equal(timeout, suite.cli.GetHttpClient().Timeout)
	suite.Equal(transport, suite.cli.GetHttpClient().Transport)

	suite.cli.SetHttpClient(defaultHttpClient)

}

func (suite *SDKClientTestSuite) TestTlsHttpRequestNormally() {

	err := suite.cli.SetHttpClient(&http.Client{
		Timeout: time.Second * 30,
		Transport: &http.Transport{
			MaxIdleConns:        1000,
			MaxIdleConnsPerHost: 50,
			IdleConnTimeout:     10 * time.Second,
			DisableCompression:  true,
			DialContext: (&net.Dialer{
				Timeout:   10 * time.Second,
				KeepAlive: 10 * time.Second,
			}).DialContext,
		},
	})
	suite.NoError(err)

	createProjectResponse, err := suite.cli.CreateProject(&CreateProjectRequest{
		ProjectName: "go-sdk-test" + time.Now().Format("20060102150405"),
		Region:      os.Getenv("LOG_SERVICE_REGION"),
	})
	suite.NoError(err)

	_, err = suite.cli.DeleteProject(&DeleteProjectRequest{
		ProjectID: createProjectResponse.ProjectID,
	})
}

func (suite *SDKClientTestSuite) TestTlsHttpRequestAbnormally() {
	err := suite.cli.SetHttpClient(nil)
	suite.Error(err)
	suite.Equal("client can not be nil", err.Error())
}
