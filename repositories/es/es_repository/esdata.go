package es_repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/olivere/elastic/v7"
	"reflect"
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
func CreateIndex(name string, body string) (rsp *elastic.IndicesCreateResult, err error) {
	log.Es.Infof("{CreateIndex start......}")
	if ok, _ := IndexExists([]string{name}); ok {
		err = errors.New(name + "索引已存在")
		log.Es.Errorf("{CreateIndex is exist error::}", err.Error())
		return
	}
	if body == "" {
		rsp, err = esdb.Client.CreateIndex(name).Do(context.TODO())
	} else {
		rsp, err = esdb.Client.CreateIndex(name).BodyString(body).Do(context.TODO())
	}

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
func SetESData(index string, typ string, id string, body interface{}) (ret int32, err error) {
	log.Es.Infof("{SetESData start index::%s, id::%s, body::%s}", index, id, body)
	if typ == "" {
		typ = index
	}

	esRsp, err := esdb.Client.Index().
		Index(index).
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
func GetESDataById(index string, typ string, id string) (source string, err error) {
	log.Es.Infof("{GetESDataById start indexName::%s, id::%s}", index, id)
	esRsp, err := esdb.Client.Get().
		Index(index).
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

// 查询ES数据，支持对象，数组类型
func QueryESData(
	index string,
	query elastic.Query,
	from int,
	size int,
	result interface{},
	sorter ...elastic.Sorter,
) (totalCount int, rsp []interface{}, err error) {
	if size == 0 {
		size = 10 // 默认取10条数据
	}
	querySql, _ := QueryESSourceSql(query)
	log.Es.Infof("{QueryESData start index::%s, querySql::%s, from::%d, size::%d, querySql::%s}", index, querySql, from, size, query)
	esRsp, err := esdb.Client.Search().
		Index(index).
		Query(query).
		SortBy(sorter...). // 排序:支持多字段排序
		From(from).        // 分页:偏移量，从0开始
		Size(size).        // 分页:偏移数，每页数量
		Pretty(true).      // JSON格式
		Do(context.Background())
	if err != nil {
		log.Es.Errorf("{QueryESData error::%s}", err.Error())
		return
	}

	log.Es.Infof(fmt.Sprintf("{QueryESData esRsp 查询消耗时间 %d ms, 结果总数: %d }", esRsp.TookInMillis, esRsp.TotalHits()))
	if esRsp.TotalHits() <= 0 {
		err = errors.New("查询ES数据为空")
		log.Es.Errorf("QueryESData esRsp is null error::", err.Error())
		return
	}

	totalCount = int(esRsp.TotalHits())
	for _, i := range esRsp.Each(reflect.TypeOf(result)) {
		tmp := i.(interface{})
		rsp = append(rsp, tmp)
	}

	log.Es.Infof("QueryESData end rsp::%s", rsp)
	return
}

// 查询语句SQL化
func QueryESSourceSql(query elastic.Query) (str string, err error) {
	src, err := query.Source()
	if err != nil {
		log.Es.Errorf("QueryESSql Source error::", err.Error())
		return
	}
	data, err := json.Marshal(src)
	if err != nil {
		log.Es.Errorf("QueryESSql marshaling to JSON error::", err.Error())
		return
	}
	str = string(data)
	return
}
