package dns_test

import (
	"context"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/dns"
)

func TestTopZone(t *testing.T) {
	ctx := context.Background()
	c := dns.NewClient(dns.NewVolcCaller())

	// ListZones
	key := "nextDNSSDKTestGo.com"

	respListZone, err := c.ListZones(ctx, &dns.ListZonesRequest{
		Key: &key,
	})
	assert.Nil(t, err)
	assert.NotNil(t, respListZone)

	var zID *int64
	zID = nil
	if len(respListZone.Zones) > 0 {
		for _, i := range respListZone.Zones {
			if *i.ZoneName == key {
				zID = i.ZID
				break
			}
		}
	}

	if zID != nil {
		err = c.DeleteZone(ctx, &dns.DeleteZoneRequest{
			ZID: zID,
		})
		assert.Nil(t, err)
	}

	// CreateZone
	remark := "next dns sdk test"
	zoneName := "nextDNSSDKTestGo.com"

	respCreateZone, err := c.CreateZone(ctx, &dns.CreateZoneRequest{
		Remark:   &remark,
		ZoneName: zoneName,
	})
	assert.Nil(t, err)
	assert.NotNil(t, respCreateZone)

	// save ZID
	zID = respCreateZone.ZID

	// UpdateZone
	remark = "next dns sdk update test"

	respUpdateZone, err := c.UpdateZone(ctx, &dns.UpdateZoneRequest{
		Remark: &remark,
		ZID:    zID,
	})
	assert.Nil(t, err)
	assert.NotNil(t, respUpdateZone)
	assert.True(t, *respUpdateZone.Remark == remark)

	// DeleteZone
	err = c.DeleteZone(ctx, &dns.DeleteZoneRequest{
		ZID: zID,
	})
	assert.Nil(t, err)
}

func TestTopRecord(t *testing.T) {
	ctx := context.Background()
	c := dns.NewClient(dns.NewVolcCaller())

	// ListZones
	key := "nextDNSSDKTestGo.com"

	respListZone, err := c.ListZones(ctx, &dns.ListZonesRequest{
		Key: &key,
	})
	assert.Nil(t, err)
	assert.NotNil(t, respListZone)

	var zID *int64
	zID = nil
	if len(respListZone.Zones) > 0 {
		for _, i := range respListZone.Zones {
			if *i.ZoneName == key {
				zID = i.ZID
				break
			}
		}
	}

	if zID != nil {
		err = c.DeleteZone(ctx, &dns.DeleteZoneRequest{
			ZID: zID,
		})
		assert.Nil(t, err)
	}

	// CreateZone
	remark := "next dns sdk test"
	zoneName := "nextDNSSDKTestGo.com"

	respCreateZone, err := c.CreateZone(ctx, &dns.CreateZoneRequest{
		Remark:   &remark,
		ZoneName: zoneName,
	})
	assert.Nil(t, err)
	assert.NotNil(t, respCreateZone)

	// save ZID
	zID = respCreateZone.ZID
	stringzID := strconv.FormatInt(*zID, 10)

	host := "dont.change.client.test.volc.com"
	// List records
	respListRecords, err := c.ListRecords(ctx, &dns.ListRecordsRequest{
		ZID:  &stringzID,
		Host: &host,
	})
	assert.Nil(t, err)
	assert.NotNil(t, respListRecords)

	var recordID string
	if len(respListRecords.Records) > 0 {
		for _, i := range respListRecords.Records {
			if *i.Host == key {
				recordID = *i.RecordID
				break
			}
		}
	}

	if len(recordID) != 0 {
		err = c.DeleteRecord(ctx, &dns.DeleteRecordRequest{
			RecordID: &recordID,
		})
		assert.Nil(t, err)
	}

	// CreateRecord
	Type := "A"
	value := "1.1.1.1"

	respCreateRecord, err := c.CreateRecord(ctx, &dns.CreateRecordRequest{
		ZID:   zID,
		Host:  &host,
		Type:  &Type,
		Value: &value,
	})
	assert.Nil(t, err)
	assert.NotNil(t, respCreateRecord)

	// save recordID
	recordID = *respCreateRecord.RecordID
	// RecordSetID := *respListRecords.Records[0].RecordSetID

	// QueryRecord
	respQueryRecord, err := c.QueryRecord(ctx, &dns.QueryRecordRequest{
		RecordID: &recordID,
	})
	assert.Nil(t, err)
	assert.NotNil(t, respQueryRecord)

	// save RecordSetID
	recordSetID := *respQueryRecord.RecordSetID

	// UpdateRecord
	value = "1.1.2.3"
	line := *respQueryRecord.Line
	respUpdateRecord, err := c.UpdateRecord(ctx, &dns.UpdateRecordRequest{
		Host:     host,
		Line:     line,
		RecordID: recordID,
		Value:    &value,
	})
	assert.Nil(t, err)
	assert.NotNil(t, respUpdateRecord)

	// UpdateRecordStatus
	enable := true
	respUpdateRecordStatus, err := c.UpdateRecordStatus(ctx, &dns.UpdateRecordStatusRequest{
		RecordID: &recordID,
		Enable:   &enable,
	})
	assert.Nil(t, err)
	assert.NotNil(t, respUpdateRecordStatus)

	// UpdateRecordSet
	weightEnable := true
	respUpdateRecordSet, err := c.UpdateRecordSet(ctx, &dns.UpdateRecordSetRequest{
		// RecordSetID: &recordSetID,
		ID:            recordSetID,
		WeightEnabled: weightEnable,
	})
	assert.Nil(t, err)
	assert.NotNil(t, respUpdateRecordSet)

	// ListRecordSets
	respListRecordSets, err := c.ListRecordSets(ctx, &dns.ListRecordSetsRequest{
		// RecordSetID: &recordSetID,
		ZID: &stringzID,
	})
	assert.Nil(t, err)
	assert.NotNil(t, respListRecordSets)

	// DeleteRecord
	err = c.DeleteRecord(ctx, &dns.DeleteRecordRequest{
		RecordID: &recordID,
	})
	assert.Nil(t, err)

	// DeleteZone
	err = c.DeleteZone(ctx, &dns.DeleteZoneRequest{
		ZID: zID,
	})
	assert.Nil(t, err)
}

func TestTopStats(t *testing.T) {
	ctx := context.Background()
	c := dns.NewClient(dns.NewVolcCaller())

	// ListZones
	key := "nextDNSSDKTestGo.com"

	respListZone, err := c.ListZones(ctx, &dns.ListZonesRequest{
		Key: &key,
	})
	assert.Nil(t, err)
	assert.NotNil(t, respListZone)

	var zID *int64
	zID = nil
	if len(respListZone.Zones) > 0 {
		for _, i := range respListZone.Zones {
			if *i.ZoneName == key {
				zID = i.ZID
				break
			}
		}
	}

	if zID != nil {
		err = c.DeleteZone(ctx, &dns.DeleteZoneRequest{
			ZID: zID,
		})
		assert.Nil(t, err)
	}

	// CreateZone
	remark := "next dns sdk test"
	zoneName := "nextDNSSDKTestGo.com"

	respCreateZone, err := c.CreateZone(ctx, &dns.CreateZoneRequest{
		Remark:   &remark,
		ZoneName: zoneName,
	})
	assert.Nil(t, err)
	assert.NotNil(t, respCreateZone)

	// save ZID
	zID = respCreateZone.ZID
	stringzID := strconv.FormatInt(*zID, 10)

	// ListZoneStatistics
	name := zoneName
	pageNum := "1"
	pageSize := "3"

	respListZoneStatistics, err := c.ListZoneStatistics(ctx, &dns.ListZoneStatisticsRequest{
		Name:       &name,
		PageNumber: &pageNum,
		PageSize:   &pageSize,
	})
	assert.Nil(t, err)
	assert.NotNil(t, respListZoneStatistics)

	// QueryZoneStatistics
	_, err = c.QueryZoneStatistics(ctx, &dns.QueryZoneStatisticsRequest{
		ZID: &stringzID,
	})
	assert.Nil(t, err)

	// ListDomainStatistics
	respListDomainStatistics, err := c.ListDomainStatistics(ctx, &dns.ListDomainStatisticsRequest{
		ZID: &stringzID,
	})
	assert.Nil(t, err)
	assert.NotNil(t, respListDomainStatistics)

	// QueryDomainStatistics
	respQueryDomainStatistics, err := c.QueryDomainStatistics(ctx, &dns.QueryDomainStatisticsRequest{
		ZID: &stringzID,
	})
	assert.Nil(t, err)
	assert.NotNil(t, respQueryDomainStatistics)

	// DeleteZone
	err = c.DeleteZone(ctx, &dns.DeleteZoneRequest{
		ZID: zID,
	})
	assert.Nil(t, err)
}

func TestTopOthers(t *testing.T) {
	ctx := context.Background()
	c := dns.NewClient(dns.NewVolcCaller())

	// CreateNewPreorder
	pageNumber := "1"
	pageSize := "10"

	respUpdateRecordStatus, err := c.ListLines(ctx, &dns.ListLinesRequest{
		PageNumber: &pageNumber,
		PageSize:   &pageSize,
	})
	assert.Nil(t, err)
	assert.NotNil(t, respUpdateRecordStatus)

}
