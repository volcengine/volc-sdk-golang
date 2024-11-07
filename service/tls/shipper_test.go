package tls

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"net/http"
	"os"
	"testing"
	"time"
)

type SDKShipperTestSuite struct {
	suite.Suite

	cli                  Client
	CreateShipperRequest CreateShipperRequest
	iamProjectName       string
	projectId            string
	topicId              string
	tosBucketName        string
	tosBucketModifyName  string
	tosObjectKey         string
	tosObjectModifyKey   string
	kafkaInstance        string
	kafkaTopicName       string
	KafkaTopicNameModify string
}

func (suite *SDKShipperTestSuite) SetupSuite() {
	suite.cli = NewClientWithEnv()
	suite.Equal("/CreateShipper", PathCreateShipper)
	suite.Equal("/DeleteShipper", PathDeleteShipper)
	suite.Equal("/ModifyShipper", PathModifyShipper)
	suite.Equal("/DescribeShipper", PathDescribeShipper)
	suite.Equal("/DescribeShippers", PathDescribeShippers)

	suite.iamProjectName = os.Getenv("IAM_PROJECT_NAME")
	suite.tosBucketName = os.Getenv("SHIPPER_TOS_BUCKET")
	suite.tosBucketModifyName = os.Getenv("SHIPPER_TOS_MODIFY_BUCKET")
	suite.tosObjectKey = os.Getenv("SHIPPER_TOS_OBJECT_KEY")
	suite.tosObjectModifyKey = os.Getenv("SHIPPER_TOS_OBJECT_MODIFY_KEY")
	suite.kafkaInstance = os.Getenv("SHIPPER_KAFKA_INSTANCE")
	suite.kafkaTopicName = os.Getenv("SHIPPER_KAFKA_TOPIC")
	suite.KafkaTopicNameModify = os.Getenv("SHIPPER_KAFKA_MODIFY_TOPIC")

	suite.createProject()
	suite.createTopic()
}

func (suite *SDKShipperTestSuite) TearDownSuite() {
	_, err := suite.cli.DeleteTopic(&DeleteTopicRequest{TopicID: suite.topicId})
	suite.NoError(err)
	_, err = suite.cli.DeleteProject(&DeleteProjectRequest{ProjectID: suite.projectId})
	suite.NoError(err)
}

func (suite *SDKShipperTestSuite) SetupTest() {
	suite.CreateShipperRequest = CreateShipperRequest{
		TopicId:          suite.topicId,
		ShipperName:      "test-shipper-name" + uuid.New().String(),
		ShipperType:      "",
		ShipperStartTime: int(time.Now().UnixMilli()),
		ShipperEndTime:   int(time.Now().UnixMilli()) + 3600*1000,
		TosShipperInfo:   nil,
		KafkaShipperInfo: nil,
		ContentInfo:      nil,
	}
}

func TestSDKShipperTestSuite(t *testing.T) {
	suite.Run(t, new(SDKShipperTestSuite))
}

func (suite *SDKShipperTestSuite) TestShipperWholeProcessOfTos() {
	shipperId, createParams := suite.testCreateShipperTosNormally()
	suite.testDescribeShipperNormally(shipperId, "tos", createParams)
	suite.testModifyShipperTosNormally(shipperId)
	suite.testDeleteShipperNormally([]string{shipperId})
}

func (suite *SDKShipperTestSuite) TestShipperWholeProcessOfKafka() {
	shipperId, createParams := suite.testCreateShipperKafkaNormally()
	suite.testDescribeShipperNormally(shipperId, "kafka", createParams)
	suite.testModifyShipperKafkaNormally(shipperId)
	suite.testDeleteShipperNormally([]string{shipperId})
}

func (suite *SDKShipperTestSuite) TestCreateShipperKafkaWithoutKafkaTimeNormally() {
	createParams := suite.CreateShipperRequest
	createParams.ShipperType = "kafka"
	createParams.KafkaShipperInfo = &KafkaShipperInfo{
		Instance:   suite.kafkaInstance,
		KafkaTopic: suite.kafkaTopicName,
		Compress:   "none",
	}
	createParams.ContentInfo = &ContentInfo{
		Format: "json",
		CsvInfo: &CsvInfo{
			Keys:            []string{"time", "level", "msg"},
			Delimiter:       " ",
			EscapeChar:      "",
			PrintHeader:     false,
			NonFieldContent: "****PADDING****",
		},
		JsonInfo: &JsonInfo{
			Keys:   []string{"__content__", "__source__", "__time__"},
			Enable: true,
			Escape: false,
		},
	}
	res, err := suite.cli.CreateShipper(&createParams)
	suite.NoError(err)
	suite.NotNil(res.ShipperId)

	suite.testDescribeShipperNormally(res.ShipperId, "kafka", createParams)
	suite.testDeleteShipperNormally([]string{res.ShipperId})
}

func (suite *SDKShipperTestSuite) TestModifyShipperTosWithOutChangeNormally() {
	shipperId, createParams := suite.testCreateShipperTosNormally()

	status := true
	modifyParams := &ModifyShipperRequest{
		ShipperId:        shipperId,
		ShipperName:      createParams.ShipperName,
		ShipperType:      createParams.ShipperType,
		Status:           &status,
		ContentInfo:      createParams.ContentInfo,
		TosShipperInfo:   createParams.TosShipperInfo,
		KafkaShipperInfo: createParams.KafkaShipperInfo,
	}

	_, err := suite.cli.ModifyShipper(modifyParams)
	suite.NoError(err)

	suite.testDescribeShipperNormally(shipperId, "tos", createParams)

	suite.testDeleteShipperNormally([]string{shipperId})
}

func (suite *SDKShipperTestSuite) TestModifyShipperKafkaWithOutChangeNormally() {
	shipperId, createParams := suite.testCreateShipperKafkaNormally()

	status := true
	modifyParams := &ModifyShipperRequest{
		ShipperId:        shipperId,
		ShipperName:      createParams.ShipperName,
		ShipperType:      createParams.ShipperType,
		Status:           &status,
		ContentInfo:      createParams.ContentInfo,
		TosShipperInfo:   createParams.TosShipperInfo,
		KafkaShipperInfo: createParams.KafkaShipperInfo,
	}

	_, err := suite.cli.ModifyShipper(modifyParams)
	suite.NoError(err)

	suite.testDescribeShipperNormally(shipperId, "kafka", createParams)

	suite.testDeleteShipperNormally([]string{shipperId})
}

func (suite *SDKShipperTestSuite) TestDescribeShippersNormally() {
	params := &DescribeShippersRequest{
		CommonRequest:  CommonRequest{},
		IamProjectName: "",
		ProjectId:      "",
		ProjectName:    "",
		TopicId:        suite.topicId,
		TopicName:      "",
		ShipperId:      "",
		ShipperName:    "",
		ShipperType:    "",
		PageNumber:     10,
		PageSize:       1,
	}

	res, err := suite.cli.DescribeShippers(params)
	suite.NoError(err)

	suite.NotNil(res.Shippers)
	suite.Equal(0, res.Total)

	tosShipperId, tosCreateParams := suite.testCreateShipperTosNormally()
	suite.CreateShipperRequest.ShipperName = "test-shipper-name" + uuid.New().String()
	kafkaShipperId, kafkaCreateParams := suite.testCreateShipperKafkaNormally()

	res, err = suite.cli.DescribeShippers(params)
	suite.NoError(err)

	suite.NotNil(res.Shippers)
	suite.Equal(2, res.Total)

	var (
		shipperId    string
		shipperType  string
		createParams CreateShipperRequest
	)

	for _, shipperInfo := range res.Shippers {

		switch shipperInfo.ShipperType {
		case "tos":
			shipperId = tosShipperId
			createParams = tosCreateParams
		case "kafka":
			shipperId = kafkaShipperId
			createParams = kafkaCreateParams
		default:
			suite.T().Errorf("unexpected shipper type: %s", shipperInfo.ShipperType)
		}

		suite.testDescribeShipperReturnParams(shipperInfo, shipperId, shipperType, createParams)
	}

	suite.testDeleteShipperNormally([]string{tosShipperId, kafkaShipperId})
}

func (suite *SDKShipperTestSuite) TestCreateShipperTosAbnormally1() {
	createParams := suite.CreateShipperRequest
	createParams.ShipperType = "tos"
	createParams.TosShipperInfo = &TosShipperInfo{
		Bucket:          "no-exist-bucket-name",
		Prefix:          suite.tosObjectKey,
		MaxSize:         10,
		Compress:        "none",
		Interval:        900,
		PartitionFormat: "%Y/%m/%d/%H",
	}
	createParams.ContentInfo = &ContentInfo{
		Format: "json",
		CsvInfo: &CsvInfo{
			Keys:            []string{"time", "level", "msg"},
			Delimiter:       " ",
			EscapeChar:      "",
			PrintHeader:     false,
			NonFieldContent: "****PADDING****",
		},
		JsonInfo: &JsonInfo{
			Keys:   []string{"time", "level", "msg"},
			Enable: true,
			Escape: false,
		},
	}

	res, err := suite.cli.CreateShipper(&createParams)
	expectedErr := &Error{
		HTTPCode: http.StatusBadRequest,
		Code:     "TosBucketNotExist",
		Message:  "The tos bucket " + createParams.TosShipperInfo.Bucket + " not exist, region " + os.Getenv("LOG_SERVICE_REGION"),
	}
	suite.validateError(err, expectedErr)

	if err == nil {
		suite.testDeleteShipperNormally([]string{res.ShipperId})
	}
}

func (suite *SDKShipperTestSuite) TestCreateShipperTosAbnormally2() {
	createParams := suite.CreateShipperRequest
	createParams.ShipperType = "tos1"
	createParams.TosShipperInfo = &TosShipperInfo{
		Bucket:          suite.tosBucketName,
		Prefix:          suite.tosObjectKey,
		MaxSize:         10,
		Compress:        "none",
		Interval:        900,
		PartitionFormat: "%Y/%m/%d/%H",
	}
	createParams.ContentInfo = &ContentInfo{
		Format: "json",
		CsvInfo: &CsvInfo{
			Keys:            []string{"time", "level", "msg"},
			Delimiter:       " ",
			EscapeChar:      "",
			PrintHeader:     false,
			NonFieldContent: "****PADDING****",
		},
		JsonInfo: &JsonInfo{
			Keys:   []string{"time", "level", "msg"},
			Enable: true,
			Escape: false,
		},
	}

	res, err := suite.cli.CreateShipper(&createParams)
	expectedErr := &Error{
		HTTPCode: http.StatusBadRequest,
		Code:     "InvalidArgument",
		Message:  "Invalid argument key ShipperType, value tos1, please check argument.",
	}
	suite.validateError(err, expectedErr)

	if err == nil {
		suite.testDeleteShipperNormally([]string{res.ShipperId})
	}
}

func (suite *SDKShipperTestSuite) TestCreateShipperTosAbnormally3() {
	shipperId, createParams := suite.testCreateShipperTosNormally()
	res, err := suite.cli.CreateShipper(&createParams)
	expectedErr := &Error{
		HTTPCode: http.StatusConflict,
		Code:     "ShipperAlreadyExist",
		Message:  "Shipper " + createParams.ShipperName + " already exist",
	}
	suite.validateError(err, expectedErr)

	deleteIds := []string{shipperId}

	if err == nil {
		deleteIds = append(deleteIds, res.ShipperId)
	}

	suite.testDeleteShipperNormally(deleteIds)
}

func (suite *SDKShipperTestSuite) TestCreateShipperKafkaAbnormally1() {
	createParams := suite.CreateShipperRequest
	createParams.ShipperType = "kafka"
	createParams.KafkaShipperInfo = &KafkaShipperInfo{
		Instance:   "no-exist-kafka-instance",
		KafkaTopic: suite.kafkaTopicName,
		Compress:   "none",
		StartTime:  int(time.Now().UnixMilli()),
		EndTime:    int(time.Now().UnixMilli()) + 3600*1000,
	}
	createParams.ContentInfo = &ContentInfo{
		Format: "json",
		CsvInfo: &CsvInfo{
			Keys:            []string{"time", "level", "msg"},
			Delimiter:       " ",
			EscapeChar:      "",
			PrintHeader:     false,
			NonFieldContent: "****PADDING****",
		},
		JsonInfo: &JsonInfo{
			Keys:   []string{"__content__", "__source__", "__time__"},
			Enable: true,
			Escape: false,
		},
	}

	res, err := suite.cli.CreateShipper(&createParams)
	expectedErr := &Error{
		HTTPCode: http.StatusBadRequest,
		Code:     "InvalidArgument",
		Message:  "Invalid argument key Instance, value no-exist-kafka-instance, please check argument.",
	}
	suite.validateError(err, expectedErr)

	if err == nil {
		suite.testDeleteShipperNormally([]string{res.ShipperId})
	}
}

func (suite *SDKShipperTestSuite) TestCreateShipperKafkaAbnormally2() {
	shipperId, createParams := suite.testCreateShipperKafkaNormally()
	res, err := suite.cli.CreateShipper(&createParams)
	expectedErr := &Error{
		HTTPCode: http.StatusConflict,
		Code:     "ShipperAlreadyExist",
		Message:  "Shipper " + createParams.ShipperName + " already exist",
	}
	suite.validateError(err, expectedErr)

	deleteIds := []string{shipperId}

	if err == nil {
		deleteIds = append(deleteIds, res.ShipperId)
	}

	suite.testDeleteShipperNormally(deleteIds)
}

func (suite *SDKShipperTestSuite) TestModifyShipperTosAbnormally1() {
	status := false
	modifyParams := &ModifyShipperRequest{
		ShipperId:   "no-exist-modify-shipperId",
		ShipperName: "test-shipper-name-modify" + uuid.New().String(),
		ShipperType: "tos",
		Status:      &status,
		ContentInfo: &ContentInfo{
			Format: "csv",
			CsvInfo: &CsvInfo{
				Keys:            []string{"time", "level", "msg"},
				Delimiter:       " ",
				EscapeChar:      "",
				PrintHeader:     false,
				NonFieldContent: "****PADDING****",
			},
			JsonInfo: nil,
		},
		TosShipperInfo: &TosShipperInfo{
			Bucket:          suite.tosBucketName,
			Prefix:          suite.tosObjectModifyKey,
			MaxSize:         20,
			Compress:        "snappy",
			Interval:        300,
			PartitionFormat: "%Y/%m/%d",
		},
		KafkaShipperInfo: nil,
	}

	_, err := suite.cli.ModifyShipper(modifyParams)
	expectedErr := &Error{
		HTTPCode: http.StatusBadRequest,
		Code:     "InvalidArgument",
		Message:  "Invalid argument key ShipperId, value " + modifyParams.ShipperId + ", please check argument.",
	}
	suite.validateError(err, expectedErr)
}

func (suite *SDKShipperTestSuite) TestModifyShipperTosAbnormally2() {

	shipperId, _ := suite.testCreateShipperTosNormally()

	modifyParams := &ModifyShipperRequest{
		ShipperId:   shipperId,
		ShipperName: "test-shipper-name-modify" + uuid.New().String(),
		ShipperType: "kafka",
		Status:      nil,
		ContentInfo: &ContentInfo{
			Format:   "original",
			CsvInfo:  nil,
			JsonInfo: nil,
		},
		KafkaShipperInfo: &KafkaShipperInfo{
			Instance:   suite.kafkaInstance,
			KafkaTopic: suite.kafkaTopicName,
			Compress:   "none",
			StartTime:  int(time.Now().UnixMilli()),
			EndTime:    int(time.Now().UnixMilli()) + 3600*1000,
		},
	}

	_, err := suite.cli.ModifyShipper(modifyParams)
	expectedErr := &Error{
		HTTPCode: http.StatusBadRequest,
		Code:     "InvalidArgument",
		Message:  "Invalid argument key ShipperType, value kafka, please check argument.",
	}
	suite.validateError(err, expectedErr)

	suite.testDeleteShipperNormally([]string{shipperId})
}

func (suite *SDKShipperTestSuite) TestModifyShipperTosAbnormally3() {
	shipperId, _ := suite.testCreateShipperTosNormally()

	modifyParams := &ModifyShipperRequest{
		ShipperId:   shipperId,
		ShipperName: "test-shipper-name-modify" + uuid.New().String(),
		ShipperType: "tos",
		Status:      nil,
		ContentInfo: &ContentInfo{
			Format:   "original",
			CsvInfo:  nil,
			JsonInfo: nil,
		},
		TosShipperInfo: &TosShipperInfo{
			Bucket:          suite.tosBucketName,
			Prefix:          suite.tosObjectModifyKey,
			MaxSize:         20,
			Compress:        "snappy",
			Interval:        300,
			PartitionFormat: "%Y/%m/%d",
		},
	}

	_, err := suite.cli.ModifyShipper(modifyParams)
	expectedErr := &Error{
		HTTPCode: http.StatusBadRequest,
		Code:     "InvalidArgument",
		Message:  "Invalid argument key Format, value original, please check argument.",
	}
	suite.validateError(err, expectedErr)

	suite.testDeleteShipperNormally([]string{shipperId})
}

func (suite *SDKShipperTestSuite) TestModifyShipperKafkaAbnormally1() {
	shipperId, _ := suite.testCreateShipperKafkaNormally()

	modifyParams := &ModifyShipperRequest{
		ShipperId:   shipperId,
		ShipperName: "test-shipper-name-modify" + uuid.New().String(),
		ShipperType: "tos",
		Status:      nil,
		ContentInfo: &ContentInfo{
			Format:  "json",
			CsvInfo: nil,
			JsonInfo: &JsonInfo{
				Keys:   []string{"__content__", "__source__", "__time__", "__namespace__", "__pod_ip__"},
				Enable: true,
				Escape: true,
			},
		},
		TosShipperInfo: &TosShipperInfo{
			Bucket:          suite.tosBucketModifyName,
			Prefix:          suite.tosObjectModifyKey,
			MaxSize:         20,
			Compress:        "snappy",
			Interval:        300,
			PartitionFormat: "%Y/%m/%d",
		},
		KafkaShipperInfo: nil,
	}

	_, err := suite.cli.ModifyShipper(modifyParams)
	expectedErr := &Error{
		HTTPCode: http.StatusBadRequest,
		Code:     "InvalidArgument",
		Message:  "Invalid argument key ShipperType, value tos, please check argument.",
	}
	suite.validateError(err, expectedErr)

	suite.testDeleteShipperNormally([]string{shipperId})
}

func (suite *SDKShipperTestSuite) TestModifyShipperKafkaAbnormally2() {
	shipperId, _ := suite.testCreateShipperKafkaNormally()

	modifyParams := &ModifyShipperRequest{
		ShipperId:   shipperId,
		ShipperName: "test-shipper-name-modify" + uuid.New().String(),
		ShipperType: "kafka",
		Status:      nil,
		ContentInfo: &ContentInfo{
			Format:   "original",
			CsvInfo:  nil,
			JsonInfo: nil,
		},
		TosShipperInfo: nil,
		KafkaShipperInfo: &KafkaShipperInfo{
			Instance:   "no-exist-kafka-instance",
			KafkaTopic: suite.KafkaTopicNameModify,
			Compress:   "lz4",
			StartTime:  0,
			EndTime:    0,
		},
	}

	_, err := suite.cli.ModifyShipper(modifyParams)
	expectedErr := &Error{
		HTTPCode: http.StatusBadRequest,
		Code:     "InvalidArgument",
		Message:  "Invalid argument key Instance, value no-exist-kafka-instance, please check argument.",
	}
	suite.validateError(err, expectedErr)

	suite.testDeleteShipperNormally([]string{shipperId})
}

func (suite *SDKShipperTestSuite) TestModifyShipperKafkaAbnormally3() {
	shipperId, _ := suite.testCreateShipperKafkaNormally()

	modifyParams := &ModifyShipperRequest{
		ShipperId:   shipperId,
		ShipperName: "test-shipper-name-modify" + uuid.New().String(),
		ShipperType: "kafka",
		Status:      nil,
		ContentInfo: &ContentInfo{
			Format:   "original",
			CsvInfo:  nil,
			JsonInfo: nil,
		},
		TosShipperInfo: nil,
		KafkaShipperInfo: &KafkaShipperInfo{
			Instance:   suite.kafkaInstance,
			KafkaTopic: "no-exist-kafka-topic",
			Compress:   "lz4",
			StartTime:  0,
			EndTime:    0,
		},
	}

	_, err := suite.cli.ModifyShipper(modifyParams)
	expectedErr := &Error{
		HTTPCode: http.StatusBadRequest,
		Code:     "InvalidArgument",
		Message:  "Invalid argument key KafkaTopic, value no-exist-kafka-topic, please check argument.",
	}
	suite.validateError(err, expectedErr)

	suite.testDeleteShipperNormally([]string{shipperId})
}

func (suite *SDKShipperTestSuite) TestModifyShipperKafkaAbnormally4() {
	shipperId, _ := suite.testCreateShipperKafkaNormally()

	modifyParams := &ModifyShipperRequest{
		ShipperId:   shipperId,
		ShipperName: "test-shipper-name-modify" + uuid.New().String(),
		ShipperType: "kafka",
		Status:      nil,
		ContentInfo: &ContentInfo{
			Format:   "csv",
			CsvInfo:  nil,
			JsonInfo: nil,
		},
		TosShipperInfo: nil,
		KafkaShipperInfo: &KafkaShipperInfo{
			Instance:   suite.kafkaInstance,
			KafkaTopic: suite.kafkaTopicName,
			Compress:   "lz4",
			StartTime:  0,
			EndTime:    0,
		},
	}

	_, err := suite.cli.ModifyShipper(modifyParams)
	expectedErr := &Error{
		HTTPCode: http.StatusBadRequest,
		Code:     "InvalidArgument",
		Message:  "Invalid argument key Format, value csv, please check argument.",
	}
	suite.validateError(err, expectedErr)

	suite.testDeleteShipperNormally([]string{shipperId})
}

func (suite *SDKShipperTestSuite) TestDescribeShipperAbnormally() {
	params := &DescribeShipperRequest{
		ShipperId: "no-exist-shipperId",
	}

	_, err := suite.cli.DescribeShipper(params)
	expectedErr := &Error{
		HTTPCode: http.StatusBadRequest,
		Code:     "InvalidArgument",
		Message:  "Invalid argument key ShipperId, value " + params.ShipperId + ", please check argument.",
	}
	suite.validateError(err, expectedErr)
}

func (suite *SDKShipperTestSuite) TestDescribeShippersAbnormally() {
	params := &DescribeShippersRequest{
		IamProjectName: suite.iamProjectName,
		ProjectId:      "",
		ProjectName:    "",
		TopicId:        "",
		TopicName:      "",
		ShipperId:      "no-exist-shipperId",
		ShipperName:    "",
		ShipperType:    "",
		PageNumber:     0,
		PageSize:       0,
	}

	_, err := suite.cli.DescribeShippers(params)
	expectedErr := &Error{
		HTTPCode: http.StatusBadRequest,
		Code:     "InvalidArgument",
		Message:  "Invalid argument key ShipperId, value " + params.ShipperId + ", please check argument.",
	}
	suite.validateError(err, expectedErr)
}

func (suite *SDKShipperTestSuite) TestDeleteShippersAbnormally() {
	params := &DeleteShipperRequest{
		ShipperId: "no-exist-shipperId",
	}
	_, err := suite.cli.DeleteShipper(params)
	expectedErr := &Error{
		HTTPCode: http.StatusBadRequest,
		Code:     "InvalidArgument",
		Message:  "Invalid argument key ShipperId, value " + params.ShipperId + ", please check argument.",
	}
	suite.validateError(err, expectedErr)
}

func (suite *SDKShipperTestSuite) testCreateShipperTosNormally() (string, CreateShipperRequest) {
	createParams := suite.CreateShipperRequest
	createParams.ShipperType = "tos"
	createParams.TosShipperInfo = &TosShipperInfo{
		Bucket:          suite.tosBucketName,
		Prefix:          suite.tosObjectKey,
		MaxSize:         10,
		Compress:        "none",
		Interval:        900,
		PartitionFormat: "%Y/%m/%d/%H",
	}
	createParams.ContentInfo = &ContentInfo{
		Format: "json",
		CsvInfo: &CsvInfo{
			Keys:            []string{"time", "level", "msg"},
			Delimiter:       " ",
			EscapeChar:      "",
			PrintHeader:     false,
			NonFieldContent: "****PADDING****",
		},
		JsonInfo: &JsonInfo{
			Keys:   []string{"time", "level", "msg"},
			Enable: true,
			Escape: false,
		},
	}
	res, err := suite.cli.CreateShipper(&createParams)
	suite.NoError(err)
	suite.NotNil(res.ShipperId)

	return res.ShipperId, createParams
}

func (suite *SDKShipperTestSuite) testCreateShipperKafkaNormally() (string, CreateShipperRequest) {
	createParams := suite.CreateShipperRequest
	createParams.ShipperStartTime = 0
	createParams.ShipperEndTime = 0
	createParams.ShipperType = "kafka"
	createParams.KafkaShipperInfo = &KafkaShipperInfo{
		Instance:   suite.kafkaInstance,
		KafkaTopic: suite.kafkaTopicName,
		Compress:   "none",
		StartTime:  int(time.Now().UnixMilli()),
		EndTime:    int(time.Now().UnixMilli()) + 3600*1000,
	}
	createParams.ContentInfo = &ContentInfo{
		Format: "json",
		CsvInfo: &CsvInfo{
			Keys:            []string{"time", "level", "msg"},
			Delimiter:       " ",
			EscapeChar:      "",
			PrintHeader:     false,
			NonFieldContent: "****PADDING****",
		},
		JsonInfo: &JsonInfo{
			Keys:   []string{"__content__", "__source__", "__time__"},
			Enable: true,
			Escape: false,
		},
	}
	res, err := suite.cli.CreateShipper(&createParams)
	suite.NoError(err)
	suite.NotNil(res.ShipperId)

	return res.ShipperId, createParams
}

func (suite *SDKShipperTestSuite) testDescribeShipperNormally(shipperId string, shipperType string, createParams CreateShipperRequest) {
	res, err := suite.cli.DescribeShipper(&DescribeShipperRequest{ShipperId: shipperId})
	suite.NoError(err)
	suite.testDescribeShipperReturnParams(res, shipperId, shipperType, createParams)
}

func (suite *SDKShipperTestSuite) testDescribeShipperReturnParams(response *DescribeShipperResponse, shipperId string, shipperType string, createParams CreateShipperRequest) {
	suite.Equal(shipperId, response.ShipperId)
	suite.Equal(shipperType, response.ShipperType)
	suite.Equal(suite.projectId, response.ProjectId)
	suite.Equal(suite.topicId, response.TopicId)
	suite.Equal(createParams.ContentInfo, response.ContentInfo)
	suite.Equal(createParams.ShipperName, response.ShipperName)
	suite.NotNil(response.CreateTime)
	suite.NotNil(response.DashboardId)
	suite.NotNil(response.ModifyTime)
	suite.NotNil(response.ProjectName)
	suite.NotNil(response.TopicName)
	suite.True(response.Status)

	switch shipperType {
	case "tos":
		suite.Equal(createParams.ShipperEndTime/1000, response.ShipperEndTime/1000)
		suite.Equal(createParams.ShipperStartTime/1000, response.ShipperStartTime/1000)
		suite.Equal(createParams.TosShipperInfo, response.TosShipperInfo)
	case "kafka":
		suite.Equal(createParams.KafkaShipperInfo.Instance, response.KafkaShipperInfo.Instance)
		suite.Equal(createParams.KafkaShipperInfo.KafkaTopic, response.KafkaShipperInfo.KafkaTopic)
		suite.Equal(createParams.KafkaShipperInfo.Compress, response.KafkaShipperInfo.Compress)

		if createParams.KafkaShipperInfo.StartTime == 0 {
			suite.Equal(createParams.ShipperStartTime/1000, response.ShipperStartTime/1000)
			suite.Equal(createParams.ShipperStartTime/1000, response.KafkaShipperInfo.StartTime/1000)
		}
		if createParams.KafkaShipperInfo.EndTime == 0 {
			suite.Equal(createParams.ShipperEndTime/1000, response.ShipperEndTime/1000)
			suite.Equal(createParams.ShipperEndTime/1000, response.KafkaShipperInfo.EndTime/1000)
		}

		suite.Equal(response.ShipperStartTime, response.KafkaShipperInfo.StartTime)
		suite.Equal(response.ShipperEndTime, response.KafkaShipperInfo.EndTime)

	default:
		suite.T().Errorf("Unexpected value for SourceType: %s", shipperType)
	}
}

func (suite *SDKShipperTestSuite) testModifyShipperTosNormally(shipperId string) {
	modifyParams := &ModifyShipperRequest{
		ShipperId:   shipperId,
		ShipperName: "test-shipper-name-modify" + uuid.New().String(),
		ShipperType: "tos",
		Status:      nil,
		ContentInfo: &ContentInfo{
			Format: "csv",
			CsvInfo: &CsvInfo{
				Keys:            []string{"time", "level", "msg"},
				Delimiter:       " ",
				EscapeChar:      "",
				PrintHeader:     false,
				NonFieldContent: "****PADDING****",
			},
			JsonInfo: nil,
		},
		TosShipperInfo: &TosShipperInfo{
			Bucket:          suite.tosBucketModifyName,
			Prefix:          suite.tosObjectModifyKey,
			MaxSize:         20,
			Compress:        "snappy",
			Interval:        300,
			PartitionFormat: "%Y/%m/%d",
		},
		KafkaShipperInfo: nil,
	}

	_, err := suite.cli.ModifyShipper(modifyParams)
	suite.NoError(err)

	res, err := suite.cli.DescribeShipper(&DescribeShipperRequest{ShipperId: shipperId})
	suite.NoError(err)
	suite.True(res.Status)
	suite.Equal(modifyParams.ShipperName, res.ShipperName)
	suite.Equal(modifyParams.ContentInfo, res.ContentInfo)
	suite.Equal(modifyParams.TosShipperInfo, res.TosShipperInfo)
}

func (suite *SDKShipperTestSuite) testModifyShipperKafkaNormally(shipperId string) {
	status := false
	modifyParams := &ModifyShipperRequest{
		ShipperId:   shipperId,
		ShipperName: "test-shipper-name-modify" + uuid.New().String(),
		ShipperType: "kafka",
		Status:      &status,
		ContentInfo: &ContentInfo{
			Format:  "original",
			CsvInfo: nil,
			JsonInfo: &JsonInfo{
				Keys:   []string{"__content__", "__source__", "__time__", "__namespace__", "__pod_ip__"},
				Enable: true,
				Escape: true,
			},
		},
		TosShipperInfo: nil,
		KafkaShipperInfo: &KafkaShipperInfo{
			Instance:   suite.kafkaInstance,
			KafkaTopic: suite.KafkaTopicNameModify,
			Compress:   "lz4",
			StartTime:  0,
			EndTime:    0,
		},
	}

	_, err := suite.cli.ModifyShipper(modifyParams)
	suite.NoError(err)

	res, err := suite.cli.DescribeShipper(&DescribeShipperRequest{ShipperId: shipperId})
	suite.NoError(err)
	suite.Equal(*modifyParams.Status, res.Status)
	suite.Equal(modifyParams.ShipperName, res.ShipperName)
	suite.Equal(modifyParams.ContentInfo, res.ContentInfo)
	suite.Equal(modifyParams.KafkaShipperInfo.KafkaTopic, res.KafkaShipperInfo.KafkaTopic)
	suite.Equal(modifyParams.KafkaShipperInfo.Compress, res.KafkaShipperInfo.Compress)
	suite.True(modifyParams.KafkaShipperInfo.StartTime < res.KafkaShipperInfo.StartTime)
	suite.True(modifyParams.KafkaShipperInfo.EndTime < res.KafkaShipperInfo.EndTime)
}

func (suite *SDKShipperTestSuite) testDeleteShipperNormally(shipperIds []string) {
	for _, shipperId := range shipperIds {
		if shipperId == "" {
			continue
		}
		_, err := suite.cli.DeleteShipper(&DeleteShipperRequest{ShipperId: shipperId})
		suite.NoError(err)
	}
}

func (suite *SDKShipperTestSuite) createProject() {
	createProjectReq := &CreateProjectRequest{
		ProjectName:    "shipper-test-project" + uuid.New().String(),
		Description:    "go sdk投递需求测试项目",
		Region:         os.Getenv("LOG_SERVICE_REGION"),
		IamProjectName: &suite.iamProjectName,
	}

	createProjectResp, err := suite.cli.CreateProject(createProjectReq)
	suite.NoError(err)
	suite.projectId = createProjectResp.ProjectID
}

func (suite *SDKShipperTestSuite) createTopic() {
	createTopicReq := &CreateTopicRequest{
		ProjectID:      suite.projectId,
		TopicName:      "import-test-topic" + uuid.New().String(),
		Ttl:            1,
		Description:    "go sdk导入需求测试topic",
		ShardCount:     1,
		EnableTracking: BoolPtr(true),
	}

	createTopicResp, err := suite.cli.CreateTopic(createTopicReq)
	suite.NoError(err)
	suite.topicId = createTopicResp.TopicID
}

func (suite *SDKShipperTestSuite) validateError(err error, expectErr *Error) {
	sdkErr, ok := err.(*Error)

	if sdkErr == nil {
		suite.Nil(sdkErr)
		return
	}

	suite.Equal(true, ok)
	suite.Equal(expectErr.HTTPCode, sdkErr.HTTPCode)
	suite.Equal(expectErr.Code, sdkErr.Code)
	suite.Equal(expectErr.Message, sdkErr.Message)
}
