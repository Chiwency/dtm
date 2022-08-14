/*
 * Copyright (c) 2021 yedf. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package test

import (
	"testing"
	"time"

	"github.com/dtm-labs/dtm/client/dtmcli"
	"github.com/dtm-labs/dtm/client/dtmcli/dtmimp"
	"github.com/dtm-labs/dtm/dtmutil"
	"github.com/dtm-labs/dtm/test/busi"
	"github.com/stretchr/testify/assert"
)

func TestMsgNormal(t *testing.T) {
	msg := genMsg(dtmimp.GetFuncName())
	msg.Submit()
	assert.Equal(t, StatusSubmitted, getTransStatus(msg.Gid))
	waitTransProcessed(msg.Gid)
	assert.Equal(t, []string{StatusSucceed, StatusSucceed, StatusSucceed, StatusSucceed}, getBranchesStatus(msg.Gid))
	assert.Equal(t, StatusSucceed, getTransStatus(msg.Gid))
}

func TestMsgTimeoutSuccess(t *testing.T) {
	gid := dtmimp.GetFuncName()
	msg := genMsg(gid)
	msg.Prepare("")
	assert.Equal(t, StatusPrepared, getTransStatus(msg.Gid))
	busi.MainSwitch.QueryPreparedResult.SetOnce(dtmcli.ResultOngoing)
	cronTransOnceForwardNow(t, gid, 180)
	assert.Equal(t, StatusPrepared, getTransStatus(msg.Gid))
	busi.MainSwitch.TransInResult.SetOnce(dtmcli.ResultOngoing)
	cronTransOnceForwardNow(t, gid, 180)
	assert.Equal(t, StatusSubmitted, getTransStatus(msg.Gid))
	cronTransOnce(t, gid)
	assert.Equal(t, []string{StatusSucceed, StatusSucceed, StatusSucceed, StatusSucceed}, getBranchesStatus(msg.Gid))
	assert.Equal(t, StatusSucceed, getTransStatus(msg.Gid))
}

func TestMsgTimeoutFailed(t *testing.T) {
	gid := dtmimp.GetFuncName()
	msg := genMsg(gid)
	msg.Prepare("")
	assert.Equal(t, StatusPrepared, getTransStatus(msg.Gid))
	busi.MainSwitch.QueryPreparedResult.SetOnce(dtmcli.ResultOngoing)
	cronTransOnceForwardNow(t, gid, 360)
	assert.Equal(t, StatusPrepared, getTransStatus(msg.Gid))
	busi.MainSwitch.QueryPreparedResult.SetOnce(dtmcli.ResultFailure)
	cronTransOnceForwardNow(t, gid, 180)
	assert.Equal(t, []string{StatusPrepared, StatusPrepared, StatusPrepared, StatusPrepared}, getBranchesStatus(msg.Gid))
	assert.Equal(t, StatusFailed, getTransStatus(msg.Gid))
}

func TestMsgAbnormal(t *testing.T) {
	msg := genMsg(dtmimp.GetFuncName())
	msg.Prepare("")
	err := msg.Prepare("")
	assert.Nil(t, err)
	err = msg.Submit()
	assert.Nil(t, err)
	waitTransProcessed(msg.Gid)
	err = msg.Prepare("")
	assert.Error(t, err)
}

func genMsg(gid string) *dtmcli.Msg {
	subscribeTopic()
	req := busi.GenReqHTTP(30, false, false)
	msg := dtmcli.NewMsg(dtmutil.DefaultHTTPServer, gid).
		AddTopic("trans", &req).
		Add(busi.Busi+"/TransOut", &req).
		Add(busi.Busi+"/TransIn", &req)
	msg.QueryPrepared = busi.Busi + "/QueryPrepared"
	return msg
}

func subscribeTopic() {
	dtmcli.GetRestyClient().R().SetQueryParams(map[string]string{
		"topic":  "trans",
		"url":    busi.Busi + "/TransOut",
		"remark": "trans test",
	}).Get(dtmutil.DefaultHTTPServer + "/subscribe")
	dtmcli.GetRestyClient().R().SetQueryParams(map[string]string{
		"topic":  "trans",
		"url":    busi.Busi + "/TransIn",
		"remark": "trans test",
	}).Get(dtmutil.DefaultHTTPServer + "/subscribe")
	// wait for the topic configuration to take effect
	time.Sleep(time.Second * time.Duration(conf.ConfigUpdateInterval+1))
}
