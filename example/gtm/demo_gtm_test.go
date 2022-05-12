package gtm_test

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/gtm"
)

func TestTopGTM(t *testing.T) {
	comm := gtm.InitCommonParameter()
	ctx := context.Background()
	c := gtm.NewClient(gtm.NewVolcCaller())

	perf1 := uuid.NewString()[:8]
	perf2 := uuid.NewString()[:8]

	domain1 := perf1 + "-volc-test.com"
	cname1 := perf1 + ".gtm.volcdns.com"
	ip := "ip"
	accmode := "cname"

	// Create GTM
	respCreate, err := c.CreateGTM(ctx, &gtm.CreateGTMRequest{
		Version: &comm.Version,

		Domain:   &domain1,
		AddrType: &ip,
		AccMode:  &accmode,
		Cname:    &cname1,
	})
	assert.Nil(t, err)
	assert.True(t, len(*respCreate.ID) > 0)

	// Save ID
	gID := respCreate.ID

	// Modify GTM
	cname1 = perf1 + "-modify.gtm.volcdns.com"
	err = c.ModifyGTM(ctx, &gtm.ModifyGTMRequest{
		Version:  &comm.Version,
		ID:       gID,
		Cname:    &cname1,
		AddrType: &ip,
	})
	assert.Nil(t, err)

	// Read GTM
	respReadGTM, err := c.ReadGTM(ctx, &gtm.ReadGTMRequest{
		Version: &comm.Version,
		ID:      gID,
	})
	assert.Nil(t, err)
	assert.True(t, *respReadGTM.Cname == cname1)

	// Create view
	name := "地址池配置1"
	line := "default"
	weight := int64(1)

	addrTTL := int64(60)
	addrType := int64(0)
	addrValue := "8.8.8.8"
	addrWeight := int64(1)

	poolActThresh := int64(1)
	poolAddrs := []gtm.Address{{TTL: &addrTTL, Type: &addrType, Value: &addrValue, Weight: &addrWeight}}
	pool := gtm.Pool{ActThresh: &poolActThresh, Addrs: poolAddrs}

	respCreateView, err := c.CreateView(ctx, &gtm.CreateViewRequest{
		Version: &comm.Version,
		ID:      gID,

		Name:    &name,
		Line:    &line,
		Weight:  &weight,
		PriPool: &pool,
		SecPool: nil,
	})
	assert.Nil(t, err)
	if err != nil {
		fmt.Println("Fail to create view!")
		fmt.Println(err.Error())
		return
	}
	assert.True(t, len(*respCreateView.ViewID) > 0)

	// Save VID
	vid := *respCreateView.ViewID

	// Start GTM
	err = c.StartGTM(ctx, &gtm.StartGTMRequest{
		Version: &comm.Version,
		ID:      gID,
	})
	assert.Nil(t, err)

	// Read GTM
	respReadGTM, err = c.ReadGTM(ctx, &gtm.ReadGTMRequest{
		Version: &comm.Version,
		ID:      gID,
	})
	assert.Nil(t, err)
	assert.True(t, *respReadGTM.State == 1)

	// Stop GTM
	err = c.StopGTM(ctx, &gtm.StopGTMRequest{
		Version: &comm.Version,
		ID:      gID,
	})
	assert.Nil(t, err)

	// Read GTM
	respReadGTM, err = c.ReadGTM(ctx, &gtm.ReadGTMRequest{
		Version: &comm.Version,
		ID:      gID,
	})
	assert.Nil(t, err)
	assert.True(t, *respReadGTM.State == 0)

	// Create GTM
	domain2 := perf2 + "-volc-test.com"
	cname2 := perf2 + ".gtm.volcdns.com"
	respCreate, err = c.CreateGTM(ctx, &gtm.CreateGTMRequest{
		Version: &comm.Version,

		Domain:   &domain2,
		AddrType: &ip,
		AccMode:  &accmode,
		Cname:    &cname2,
	})
	assert.Nil(t, err)

	// Save ID2
	gID2 := respCreate.ID

	// Find
	respFindGTMs, err := c.FindGTMs(ctx, &gtm.FindGTMsRequest{
		Version: &comm.Version,
		Domain:  &domain2,
	})
	assert.Nil(t, err)
	assert.Equal(t, *respFindGTMs.Total, int64(1))

	// Delete GTM
	err = c.DeleteGTM(ctx, &gtm.DeleteGTMRequest{
		Version: &comm.Version,
		ID:      gID2,
	})
	assert.Nil(t, err)

	// Find
	respFindGTMs, err = c.FindGTMs(ctx, &gtm.FindGTMsRequest{
		Version: &comm.Version,
		Domain:  &domain2,
	})
	assert.Nil(t, err)
	assert.Equal(t, *respFindGTMs.Total, int64(0))

	// Delete View
	err = c.DeleteView(ctx, &gtm.DeleteViewRequest{
		Version: &comm.Version,
		ID:      gID,
		Vid:     &vid,
	})
	assert.Nil(t, err)

	// Delete GTM
	err = c.DeleteGTM(ctx, &gtm.DeleteGTMRequest{
		Version: &comm.Version,
		ID:      gID,
	})
	assert.Nil(t, err)
}

func TestTopGTMProbe(t *testing.T) {
	comm := gtm.InitCommonParameter()
	ctx := context.Background()
	c := gtm.NewClient(gtm.NewVolcCaller())

	perf1 := uuid.NewString()[:8]

	domain1 := perf1 + "-volc-test.com"
	cname1 := perf1 + ".gtm.volcdns.com"
	ip := "ip"
	accmode := "cname"

	// Create GTM
	respCreate, err := c.CreateGTM(ctx, &gtm.CreateGTMRequest{
		Version: &comm.Version,

		Domain:   &domain1,
		AddrType: &ip,
		AccMode:  &accmode,
		Cname:    &cname1,
	})
	assert.Nil(t, err)

	// Save ID
	gID := respCreate.ID

	// Modify Probe
	enable := true
	proto := "ping"
	interval := int64(60)
	count := int64(10)
	timeout := "5s"
	fail_count := int64(2)

	err = c.ModifyProbe(ctx, &gtm.ModifyProbeRequest{
		Version:   &comm.Version,
		ID:        gID,
		Enable:    &enable,
		Proto:     &proto,
		Interval:  &interval,
		Count:     &count,
		Timeout:   &timeout,
		FailCount: &fail_count,
	})
	assert.Nil(t, err)

	// Read Probe
	resp, err := c.ReadProbe(ctx, &gtm.ReadProbeRequest{
		Version: &comm.Version,
		ID:      gID,
	})
	assert.Nil(t, err)
	assert.Equal(t, *resp.Enable, enable)
	assert.Equal(t, *resp.Proto, proto)
	assert.Equal(t, *resp.Interval, interval)
	assert.Equal(t, *resp.Count, count)
	assert.Equal(t, *resp.Timeout, timeout)
	assert.Equal(t, *resp.FailCount, fail_count)

	// Delete GTM
	err = c.DeleteGTM(ctx, &gtm.DeleteGTMRequest{
		Version: &comm.Version,
		ID:      gID,
	})
	assert.Nil(t, err)
}

func TestTopGTMView(t *testing.T) {
	comm := gtm.InitCommonParameter()
	ctx := context.Background()
	c := gtm.NewClient(gtm.NewVolcCaller())

	perf := uuid.NewString()[:8]

	domain := perf + "-volc-test.com"
	cname := perf + ".gtm.volcdns.com"
	ip := "ip"
	accmode := "cname"

	// Create GTM
	respCreate, err := c.CreateGTM(ctx, &gtm.CreateGTMRequest{
		Version: &comm.Version,

		Domain:   &domain,
		AddrType: &ip,
		AccMode:  &accmode,
		Cname:    &cname,
	})
	assert.Nil(t, err)

	// Save ID
	gID := respCreate.ID

	// Create view
	name := "地址池配置1"
	line := "default"
	weight := int64(1)

	addrTTL := int64(60)
	addrType := int64(0)
	addrValue := "8.8.8.8"
	addrWeight := int64(1)

	poolActThresh := int64(1)
	poolAddrs := []gtm.Address{{TTL: &addrTTL, Type: &addrType, Value: &addrValue, Weight: &addrWeight}}

	pool := gtm.Pool{ActThresh: &poolActThresh, Addrs: poolAddrs}

	respCreateView, err := c.CreateView(ctx, &gtm.CreateViewRequest{
		Version: &comm.Version,
		ID:      gID,

		Name:    &name,
		Line:    &line,
		Weight:  &weight,
		PriPool: &pool,
		SecPool: nil,
	})
	assert.Nil(t, err)
	assert.True(t, len(*respCreateView.ViewID) > 0)

	// Save VID
	vid := *respCreateView.ViewID

	// List view
	resp, err := c.ListViews(ctx, &gtm.ListViewsRequest{
		Version: &comm.Version,
		ID:      gID,
	})
	assert.Nil(t, err)
	assert.Equal(t, len(resp), 1)

	// Update view
	name = "地址池配置2"
	line = "default"
	weight = int64(1)

	addrTTL = int64(60)
	addrType = int64(0)
	addrValue = "8.8.8.9"
	addrWeight = int64(1)

	poolActThresh = int64(1)
	poolAddrs = []gtm.Address{{TTL: &addrTTL, Type: &addrType, Value: &addrValue, Weight: &addrWeight}}

	pool = gtm.Pool{ActThresh: &poolActThresh, Addrs: poolAddrs}

	err = c.ModifyView(ctx, &gtm.ModifyViewRequest{
		Version: &comm.Version,
		ID:      gID,
		Vid:     &vid,

		Name:    &name,
		Line:    &line,
		Weight:  &weight,
		PriPool: &pool,
		SecPool: &pool,
	})
	assert.Nil(t, err)

	// Read view
	respReadView, err := c.ReadView(ctx, &gtm.ReadViewRequest{
		Version: &comm.Version,
		ID:      gID,
		Vid:     &vid,
	})
	assert.Nil(t, err)
	assert.Equal(t, *respReadView.Name, name)

	// Delete view
	err = c.DeleteView(ctx, &gtm.DeleteViewRequest{
		Version: &comm.Version,
		ID:      gID,
		Vid:     &vid,
	})
	assert.Nil(t, err)

	// Delete GTM
	err = c.DeleteGTM(ctx, &gtm.DeleteGTMRequest{
		Version: &comm.Version,
		ID:      gID,
	})
	assert.Nil(t, err)
}

func TestTopGTMStatus(t *testing.T) {
	comm := gtm.InitCommonParameter()
	ctx := context.Background()
	c := gtm.NewClient(gtm.NewVolcCaller())

	res, err := c.Stats(ctx, &gtm.StatsRequest{
		Version: &comm.Version,
	})
	assert.Nil(t, err)
	if err != nil {
		log.Println(err.Error())
		return
	}
	b, _ := json.Marshal(res)
	fmt.Println(string(b))
	assert.True(t, len(b) > 0)
}
