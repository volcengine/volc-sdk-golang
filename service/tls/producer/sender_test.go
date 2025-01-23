package producer

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"github.com/volcengine/volc-sdk-golang/service/tls"
	"github.com/volcengine/volc-sdk-golang/service/tls/common"
	"strconv"
	"testing"
	"time"
)

type SDKSenderTestSuite struct {
	suite.Suite

	sender *Sender
	batch  *Batch
}

func (suite *SDKSenderTestSuite) initSender() {
	producerConfig := GetDefaultProducerConfig()

	errorStatusMap := func() map[int]struct{} {
		errorCodeMap := make(map[int]struct{})
		for _, v := range producerConfig.NoRetryStatusCodeList {
			errorCodeMap[v] = struct{}{}
		}

		return errorCodeMap
	}()

	suite.sender = initSender(
		newClientWithEnv(),
		newRetryQueue(),
		20,
		common.LogConfig(producerConfig.LoggerConfig),
		errorStatusMap,
		nil,
	)
}

func (suite *SDKSenderTestSuite) initBatch() {
	suite.batch = &Batch{}
}

func (suite *SDKSenderTestSuite) SetupTest() {
	suite.initSender()
	suite.initBatch()
}

func (suite *SDKSenderTestSuite) TearDownTest() {}

func TestSDKSenderTestSuite(t *testing.T) {
	suite.Run(t, new(SDKSenderTestSuite))
}

func (suite *SDKSenderTestSuite) TestHandleSuccessRetryTimesNormally() {

	suite.sender.producer = &producer{
		producerLogGroupSize: 1000,
	}

	suite.batch = &Batch{
		attemptCount:               9,
		maxRetryTimes:              10,
		maxReservedAttempts:        11,
		retryBackoffMs:             6700,
		baseRetryBackoffMs:         1000,
		baseIncreaseRetryBackoffMs: 1000,
		maxRetryIntervalInMs:       10 * 1000,
		result: &Result{
			Attempts:    []*Attempt{},
			SuccessFlag: false,
		},
		totalDataSize: 100,
	}

	commonRes := &tls.CommonResponse{
		RequestID: "Test-TLS-RequestID",
	}

	suite.sender.handleSuccess(suite.batch, commonRes)

	suite.Equal(int64(0), suite.batch.retryBackoffMs)
}

func (suite *SDKSenderTestSuite) TestHandleFailureRetryTimesNormally() {
	suite.batch = &Batch{
		attemptCount:               0,
		maxRetryTimes:              10,
		maxReservedAttempts:        11,
		retryBackoffMs:             0,
		baseRetryBackoffMs:         1000,
		baseIncreaseRetryBackoffMs: 1000,
		maxRetryIntervalInMs:       10 * 1000,
		result: &Result{
			Attempts:    []*Attempt{},
			SuccessFlag: false,
		},
	}

	err := &tls.Error{
		HTTPCode:  429,
		Code:      "TlsCode",
		Message:   "TlsMessage",
		RequestID: "TlsRequestID",
	}

	for i := suite.batch.attemptCount; i < suite.batch.maxRetryTimes; i++ {

		fmt.Println("#############" + strconv.Itoa(i) + "#############")

		suite.sender.handleFailure(suite.batch, err)
		now := time.Now().UnixNano()

		suite.False(suite.batch.result.SuccessFlag)
		suite.Equal(i+1, suite.batch.attemptCount)

		suite.Less(int64(1), suite.batch.retryBackoffMs)
		suite.Greater(now+suite.batch.retryBackoffMs*1000*1000, suite.batch.nextRetryMs*1000*1000)
		suite.LessOrEqual(suite.batch.retryBackoffMs, suite.batch.maxRetryIntervalInMs)
	}

	suite.Equal(int64(1000), suite.batch.baseRetryBackoffMs)
	suite.Equal(int64(1000), suite.batch.baseIncreaseRetryBackoffMs)
	suite.Equal(int64(10*1000), suite.batch.maxRetryIntervalInMs)
}
