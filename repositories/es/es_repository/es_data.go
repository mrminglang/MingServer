package es_repository

import (
	"context"
	"server/utils/esdb"
	"server/utils/log"
)

// 创建ES数据
func SetESData(indexName string, typ string, id string, body interface{}) (ret int32, err error) {
	log.Es.Infof("{SetESData start indexName::%s, id::%s, body::%s}", indexName, id, body)
	rsp, err := esdb.Client.Index().
		Index(indexName).
		Type(typ).
		Id(id).
		BodyJson(body).
		Do(context.Background())
	if err != nil {
		log.Es.Errorf("{SetESData error::%s}", err.Error())
		return
	}

	log.Es.Infof("{SetESData rsp::%s}", &rsp)
	return
}

// 查找ES数据 BY id
func GetESDataById(indexName string, typ string, id string) (source string, err error) {
	log.Es.Infof("{GetESDataById start indexName::%s, id::%s}", indexName, id)
	rsp, err := esdb.Client.Get().
		Index(indexName).
		Type(typ).
		Id(id).
		Do(context.Background())
	if err != nil {
		log.Es.Errorf("{GetESDataById error::%s}", err.Error())
		return
	}
	log.Es.Infof("{GetESDataById rsp::%s}", rsp)
	if rsp.Found {
		by, _ := rsp.Source.MarshalJSON()
		//if err != nil {
		//	log.Es.Errorf("{GetESDataById rsp.Source json MarshalJSON error::%s}", err.Error())
		//	return
		//}
		source = string(by)
	}
	log.Es.Infof("GetESDataById end source::%s", source)
	return
}
