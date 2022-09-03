package test

import (
	"context"
	"testing"

	"github.com/dtm-labs/dtm/client/dtmcli"
	"github.com/dtm-labs/dtm/client/dtmcli/dtmimp"
	"github.com/dtm-labs/dtm/client/dtmgrpc"
	"github.com/dtm-labs/dtm/client/dtmgrpc/dtmgimp"
	"github.com/dtm-labs/dtm/client/dtmgrpc/dtmgpb"
	"github.com/dtm-labs/dtm/dtmsvr/storage/sql"
	"github.com/dtm-labs/dtm/dtmutil"
	"github.com/go-errors/errors"
	"github.com/stretchr/testify/assert"
)

func TestGeneralDB(t *testing.T) {
	if conf.Store.IsDB() {
		testSql(t)
		testDbAlone(t)
	}
}

func testSql(t *testing.T) {
	conf := conf.Store.GetDBConf()
	conf.Host = "127.0.0.1" // use a new host to trigger SetDBConn called
	db := dtmutil.DbGet(conf, sql.SetDBConn)
	err := func() (rerr error) {
		defer dtmimp.P2E(&rerr)
		db.Must().Exec("select a")
		return nil
	}()
	assert.NotEqual(t, nil, err)
}

func testDbAlone(t *testing.T) {
	db, err := dtmimp.StandaloneDB(conf.Store.GetDBConf())
	assert.Nil(t, err)
	_, err = dtmimp.DBExec(conf.Store.Driver, db, "select 1")
	assert.Equal(t, nil, err)
	_, err = dtmimp.DBExec(conf.Store.Driver, db, "")
	assert.Equal(t, nil, err)
	db.Close()
	_, err = dtmimp.DBExec(conf.Store.Driver, db, "select 1")
	assert.NotEqual(t, nil, err)
}

func TestMustGenGid(t *testing.T) {
	dtmgrpc.MustGenGid(dtmutil.DefaultGrpcServer)
	dtmcli.MustGenGid(dtmutil.DefaultHTTPServer)
}

func TestTopic(t *testing.T) {
	testSubscribe(t, httpSubscribe)
	testUnSubscribe(t, httpUnSubscribe)
	testDeleteTopic(t, httpDeleteTopic)

	testSubscribe(t, grpcSubscribe)
	testUnSubscribe(t, grpcUnSubscribe)
	testDeleteTopic(t, grpcDeleteTopic)
}

func testSubscribe(t *testing.T, subscribe func(topic, url string) error) {
	assert.Nil(t, subscribe("test_topic", "http://dtm/test1"))
	assert.Error(t, subscribe("test_topic", "http://dtm/test1"))
	assert.Nil(t, subscribe("test_topic", "http://dtm/test2"))
}

func testUnSubscribe(t *testing.T, unSubscribe func(topic, url string) error) {
	assert.Nil(t, unSubscribe("test_topic", "http://dtm/test1"))
	assert.Error(t, unSubscribe("test_topic", "http://dtm/test1"))
	assert.Nil(t, unSubscribe("test_topic", "http://dtm/test2"))
	assert.Error(t, unSubscribe("test_topic", "http://dtm/test2"))
}

func testDeleteTopic(t *testing.T, deleteTopic func(topic string) error) {
	assert.Error(t, deleteTopic("123"))
	assert.Nil(t, deleteTopic("test_topic"))
}

func httpSubscribe(topic, url string) error {
	resp, err := dtmcli.GetRestyClient().R().SetQueryParams(map[string]string{
		"topic":  topic,
		"url":    url,
		"remark": "for test",
	}).Get(dtmutil.DefaultHTTPServer + "/subscribe")
	e2p(err)
	if resp.StatusCode() != 200 {
		err = errors.Errorf("Http Request Error. Resp:%v", resp.String())
	}
	return err
}

func httpUnSubscribe(topic, url string) error {
	resp, err := dtmcli.GetRestyClient().R().SetQueryParams(map[string]string{
		"topic": topic,
		"url":   url,
	}).Get(dtmutil.DefaultHTTPServer + "/unsubscribe")
	e2p(err)
	if resp.StatusCode() != 200 {
		err = errors.Errorf("Http Request Error. Resp:%+v", resp.String())
	}
	return err
}

func httpDeleteTopic(topic string) error {
	resp, err := dtmcli.GetRestyClient().R().Delete(dtmutil.DefaultHTTPServer + "/topic/" + topic)
	e2p(err)
	if resp.StatusCode() != 200 {
		err = errors.Errorf("Http Request Error. Resp:%+v", resp.String())
	}
	return err
}

func grpcSubscribe(topic, url string) error {
	_, err := dtmgimp.MustGetDtmClient(dtmutil.DefaultGrpcServer).Subscribe(context.Background(),
		&dtmgpb.DtmTopicRequest{
			Topic:  topic,
			URL:    url,
			Remark: "for test"})
	return err
}

func grpcUnSubscribe(topic, url string) error {
	_, err := dtmgimp.MustGetDtmClient(dtmutil.DefaultGrpcServer).UnSubscribe(context.Background(),
		&dtmgpb.DtmTopicRequest{
			Topic: topic,
			URL:   url})
	return err
}

func grpcDeleteTopic(topic string) error {
	_, err := dtmgimp.MustGetDtmClient(dtmutil.DefaultGrpcServer).DeleteTopic(context.Background(),
		&dtmgpb.DtmTopicRequest{
			Topic: topic})
	return err
}
