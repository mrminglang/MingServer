package es_repository

import (
	"context"
	"errors"
	"github.com/olivere/elastic/v7"
	"server/utils/esdb"
	"server/utils/log"
)

// 检测健康状态
func ClusterHealth() (*elastic.ClusterHealthResponse, error) {
	log.Es.Infof("{ClusterHealth start......}")
	health, err := esdb.Client.ClusterHealth().Index().Do(context.TODO())
	if err != nil {
		log.Es.Errorf("{ClusterHealth error::}", err.Error())
		return nil, err
	}
	log.Es.Infof("{ClusterHealth ClusterName^Status|%s|%s}", health.ClusterName, health.Status)
	log.Es.Infof("{ClusterHealth success......}")
	return health, nil
}

// 查看索引 支持模糊匹配查询
func CatIndices(index string) (indices elastic.CatIndicesResponse, err error) {
	log.Es.Infof("{CatIndices start......index::}", index)
	if index != "" {
		indices, err = esdb.Client.CatIndices().Index("*" + index + "*").Do(context.TODO())
	} else {
		indices, err = esdb.Client.CatIndices().Do(context.TODO())
	}
	if err != nil {
		log.Es.Errorf("{CatIndices error::}", err.Error())
		return
	}
	log.Es.Infof("{CatIndices count:%d}", len(indices))
	log.Es.Infof("{CatIndices success......}")
	return
}

// 查看分片
func CatAllocation() (nodes elastic.CatAllocationResponse, err error) {
	log.Es.Infof("{CatAllocation start......}")
	nodes, err = esdb.Client.CatAllocation().Do(context.TODO())
	if err != nil {
		log.Es.Errorf("{CatAllocation error::}", err.Error())
		return
	}
	log.Es.Infof("{CatAllocation count:%d}", len(nodes))
	log.Es.Infof("{CatAllocation success......}")
	return
}

// 查看主分片
func CatMaster() (node elastic.CatMasterResponse, err error) {
	log.Es.Infof("{CatMaster start......}")
	node, err = esdb.Client.CatMaster().Do(context.TODO())
	if err != nil {
		log.Es.Errorf("{CatMaster error::}", err.Error())
		return
	}
	log.Es.Infof("{CatMaster count:%d}", len(node))
	log.Es.Infof("{CatMaster success......}")
	return
}

// 判断索引是否存在
func IndexExists(indices []string) (ok bool, err error) {
	log.Es.Infof("{IndexExists start......}")
	ok, err = esdb.Client.IndexExists(indices...).Do(context.TODO())
	if err != nil {
		log.Es.Errorf("{IndexExists error::}", err.Error())
		return
	}
	log.Es.Infof("{IndexExists ok::}", ok)
	log.Es.Infof("{IndexExists success......}")
	return
}

// 创建索引
func CreateIndex(name string) (rsp *elastic.IndicesCreateResult, err error) {
	log.Es.Infof("{CreateIndex start......}")
	if ok, _ := IndexExists([]string{name}); ok {
		err = errors.New(name + "索引已存在")
		log.Es.Errorf("{CreateIndex is exist error::}", err.Error())
		return
	}

	rsp, err = esdb.Client.CreateIndex(name).Do(context.TODO())
	if err != nil {
		log.Es.Errorf("{CreateIndex error::}", err.Error())
		return
	}
	log.Es.Infof("{CreateIndex Index|%s}", rsp.Index)
	log.Es.Infof("{CreateIndex success......}")
	return
}

// 删除索引，支撑批量删除
func DeleteIndex(indices []string) (rsp *elastic.IndicesDeleteResponse, err error) {
	log.Es.Infof("{DeleteIndex start......}")
	if ok, _ := IndexExists(indices); !ok {
		err = errors.New("索引不存在")
		log.Es.Errorf("{DeleteIndex is not exist error::}", err.Error())
		return nil, err
	}

	rsp, err = esdb.Client.DeleteIndex(indices...).Do(context.TODO())
	if err != nil {
		log.Es.Errorf("{DeleteIndex error::}", err.Error())
		return nil, err
	}
	log.Es.Infof("{DeleteIndex Acknowledged|%s}", rsp.Acknowledged)
	log.Es.Infof("{DeleteIndex success......}")
	return
}

// 创建ES数据
func SetESData(indexName string, typ string, id string, body interface{}) (ret int32, err error) {
	log.Es.Infof("{SetESData start indexName::%s, id::%s, body::%s}", indexName, id, body)
	esRsp, err := esdb.Client.Index().
		Index(indexName).
		Type(typ).
		Id(id).
		BodyJson(body).
		Do(context.Background())
	if err != nil {
		log.Es.Errorf("{SetESData error::%s}", err.Error())
		return
	}

	log.Es.Infof("{SetESData esRsp::%s}", &esRsp)
	return
}

// 查找ES数据 BY id
func GetESDataById(indexName string, typ string, id string) (source string, err error) {
	log.Es.Infof("{GetESDataById start indexName::%s, id::%s}", indexName, id)
	esRsp, err := esdb.Client.Get().
		Index(indexName).
		Type(typ).
		Id(id).
		Do(context.Background())
	if err != nil {
		log.Es.Errorf("{GetESDataById error::%s}", err.Error())
		return
	}
	log.Es.Infof("{GetESDataById esRsp::%s}", esRsp)
	if esRsp.Found {
		by, _ := esRsp.Source.MarshalJSON()
		source = string(by)
	}
	log.Es.Infof("GetESDataById end source::%s", source)
	return
}
