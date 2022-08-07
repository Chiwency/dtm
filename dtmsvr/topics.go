package dtmsvr

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/dtm-labs/dtm/client/dtmcli/dtmimp"
	"github.com/dtm-labs/dtm/client/dtmcli/logger"
)

const (
	topicsCat = "topics"
)

var (
	topicsMap = map[string]*Topic{}
	firstInit = true
)

type Topic struct {
	Name    string    `json:"k"`
	Urls    []UrlInfo `json:"v"`
	Version uint64    `json:"version"`
}

type UrlInfo struct {
	Url    string `json:"url"`
	Remark string `json:"remark"`
}

// Subscribe subscribe topic, create topic if not exist
func Subscribe(topic, url, remark string) error {
	urlToAdd := UrlInfo{
		Url:    url,
		Remark: remark,
	}
	kvs, err := GetStore().FindKeyValues(topicsCat, topic)
	if err != nil {
		return err
	}
	if len(kvs) == 0 {
		return GetStore().CreateKeyValue(topicsCat, topic, dtmimp.MustMarshalString([]UrlInfo{urlToAdd}))
	}
	urlInfos := []UrlInfo{}
	dtmimp.MustUnmarshalString(kvs[0].V, &urlInfos)
	for _, urlInfo := range urlInfos {
		if urlInfo.Url == url {
			return errors.New("this url has already exists")
		}
	}
	urlInfos = append(urlInfos, urlToAdd)
	return GetStore().UpdateKeyValue(&kvs[0], "", dtmimp.MustMarshalString(urlInfos))
}

func UnSubscribe(topic, url string) error {
	kvs, err := GetStore().FindKeyValues(topicsCat, topic)
	if err != nil {
		return err
	}
	if len(kvs) == 0 {
		return errors.New("no such a topic")
	}
	urlInfos := []UrlInfo{}
	dtmimp.MustUnmarshalString(kvs[0].V, &urlInfos)
	if len(urlInfos) == 0 {
		return errors.New("this topic has is empty")
	}
	n := len(urlInfos)
	for k, urlInfo := range urlInfos {
		if urlInfo.Url == url {
			urlInfos = append(urlInfos[:k], urlInfos[k+1:]...)
			break
		}
	}
	if len(urlInfos) == n {
		return errors.New("no such an url ")
	}
	return GetStore().UpdateKeyValue(&kvs[0], "", dtmimp.MustMarshalString(urlInfos))
}

func UpdateTopicsMap() {
	kvs, err := GetStore().FindKeyValues(topicsCat, "")
	if err != nil {
		logger.FatalfIf(firstInit, "init topicsMap failed. error: %v", err)
		logger.Errorf("update topicsMap error: %v", err)
		return
	}

	for _, kv := range kvs {
		topic, exist := topicsMap[kv.K]
		if exist && topic.Version >= kv.Version {
			continue
		}
		newTopic := &Topic{}
		newTopic.Name = kv.K
		newTopic.Version = kv.Version
		err = json.Unmarshal([]byte(kv.V), &newTopic.Urls)
		if err != nil {
			logger.FatalfIf(firstInit, "init topicsMap failed. error: %v, topicInfo: %v", err, kv)
			logger.Errorf("update topicsMap error: %v, topicInfo: %v", err, kv)
			continue
		}
		logger.Infof("topic updated. old topic:%v new topic:%v", topicsMap[kv.K], newTopic)
		topicsMap[kv.K] = newTopic
	}

	firstInit = false
	return
}

func CronUpdateTopicsMap() {
	for {
		time.Sleep(time.Duration(conf.ConfigUpdateInterval) * time.Second)
		UpdateTopicsMap()
	}
}
