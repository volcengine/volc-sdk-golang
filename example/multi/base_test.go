package multi

import (
	"testing"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/imp"
	"github.com/volcengine/volc-sdk-golang/service/vod"
)

func TestMulti_InitMulti(t *testing.T) {
	instanceVod := vod.NewInstance()

	instanceVod.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	instanceImp := imp.NewInstance()

	instanceImp.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

}
