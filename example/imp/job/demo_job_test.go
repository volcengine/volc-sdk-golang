package job

import (
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/imp"
	"github.com/volcengine/volc-sdk-golang/service/imp/models/business"
	"github.com/volcengine/volc-sdk-golang/service/imp/models/request"
)

func TestImp_SubmitJob(t *testing.T) {
	impService := imp.NewInstance()

	// call below method if you don't set ak and sk in $HOME/.vcloud/config
	impService.SetAccessKey("your AK")
	impService.SetSecretKey("your SK")

	// SubmitJob
	req := &request.ImpSubmitJobRequest{
		InputPath: &business.InputPath{
			Type:         "VOD",
			VodSpaceName: "your vod space",
			FileId:       "your file id",
		},
		TemplateId:   "your template id",
		CallbackArgs: "your callback args",
	}

	resp, status, err := impService.SubmitJob(req)
	assert.NoError(t, err)
	assert.Equal(t, status, http.StatusOK)
	assert.NotNil(t, resp)
	assert.NotZero(t, len(resp.Result))
	log.Println(resp.Result)
}

func TestImp_KillJob(t *testing.T) {
	impService := imp.NewInstance()

	// call below method if you don't set ak and sk in $HOME/.vcloud/config
	impService.SetAccessKey("your AK")
	impService.SetSecretKey("your SK")

	// KillJob
	req := &request.ImpKillJobRequest{
		JobId: "your job id",
	}

	resp, status, err := impService.KillJob(req)
	assert.NoError(t, err)
	assert.Equal(t, status, http.StatusOK)
	assert.NotNil(t, resp)
}

func TestImp_RetrieveJob(t *testing.T) {
	impService := imp.NewInstance()

	// call below method if you don't set ak and sk in $HOME/.vcloud/config

	impService.SetAccessKey("your AK")
	impService.SetSecretKey("your SK")

	// RetrieveJob
	req := &request.ImpRetrieveJobRequest{
		JobIds: []string{
			"your job 1",
			"your job 2",
		},
	}

	resp, status, err := impService.RetrieveJob(req)
	assert.NoError(t, err)
	assert.Equal(t, status, http.StatusOK)
	assert.NotNil(t, resp)
	assert.Equal(t, 2, len(resp.Result))

	log.Println(resp.Result)
}
