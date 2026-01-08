package tls

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type SDKAlarmWebhookIntegrationTestSuite struct {
	suite.Suite
	cli Client
}

func TestSDKAlarmWebhookIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(SDKAlarmWebhookIntegrationTestSuite))
}

func (suite *SDKAlarmWebhookIntegrationTestSuite) SetupTest() {
	suite.cli = NewClientWithEnv()
}

func (suite *SDKAlarmWebhookIntegrationTestSuite) TestDescribeAlarmWebhookIntegrations() {
	resp, err := suite.cli.DescribeAlarmWebhookIntegrations(&DescribeAlarmWebhookIntegrationsRequest{
		PageNumber: 1,
		PageSize:   20,
	})
	suite.NoError(err)
	suite.GreaterOrEqual(resp.Total, 0)
}
