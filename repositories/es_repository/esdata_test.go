package es_repository_test

import (
	"github.com/mrminglang/tools/dumps"
	"github.com/olivere/elastic/v7"
	"github.com/stretchr/testify/assert"
	"server/boot"
	"server/common"
	"server/repositories/es_repository"
	"strconv"
	"testing"
	"time"
)

var serverName = boot.RootPath() + "/MingServer"

func TestMain(m *testing.M) {
	_ = boot.Boot([]string{serverName}, serverName)
	m.Run()
}

type Person struct {
	Name    string              `json:"name"`
	Age     int                 `json:"age"`
	Birth   string              `json:"birth"`
	Bool    bool                `json:"bool"`
	Address address             `json:"address"`
	Cars    []car               `json:"cars"`
	Class   map[string][]string `json:"class"`
}

type address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

type car struct {
	Brand  string `json:"brand"`
	Colour string `json:"colour"`
	Model  string `json:"model"`
}

func TestClusterHealth(t *testing.T) {
	health, err := es_repository.ClusterHealth()
	if err != nil {
		assert.Error(t, err)
	}
	dumps.Dump(health)
}

func TestCatIndices(t *testing.T) {
	index := "se"
	indices, err := es_repository.CatIndices(index)
	if err != nil {
		assert.Error(t, err)
	}
	dumps.Dump(indices)
}

func TestCatAllocation(t *testing.T) {
	nodes, err := es_repository.CatAllocation()
	if err != nil {
		assert.Error(t, err)
	}
	dumps.Dump(nodes)
}

func TestCatMaster(t *testing.T) {
	node, err := es_repository.CatMaster()
	if err != nil {
		assert.Error(t, err)
	}
	dumps.Dump(node)
}

func TestIndexExists(t *testing.T) {
	indices := []string{"userss"}
	ok, err := es_repository.IndexExists(indices)
	if err != nil {
		assert.Error(t, err)
	}
	assert.True(t, ok)
}

func TestCreateIndex(t *testing.T) {
	name := "person_index_trans2"
	body := `
{
  "aliases": {
    "person2": {}
  },
  "mappings": {
    "properties": {
      "name": {"type": "text"},
      "age": {"type": "integer"},
      "birth": {
        "type": "date",
        "format": "yyyy-MM-dd HH:mm:ss"
      },
      "bool": {"type": "boolean"},
      "address": {
        "type": "object",
        "properties": {
          "city": {"type": "keyword"},
          "state": {"type": "keyword"}
        }
      },
      "cars": {
        "type": "nested",
        "properties": {
          "brand": {"type": "keyword"},
          "colour": {"type": "keyword"},
		  "model": {"type": "keyword"}
        }
      },
      "class": {
        "type": "nested",
        "properties": {
          "key": { "type": "keyword" },
          "values": { "type": "keyword" }
        }
      }
    }
  }
}`
	rsp, err := es_repository.CreateIndex(name, body)
	if err != nil {
		assert.Error(t, err)
	}
	dumps.Dump(rsp)

	boot.Destroy()
}

func TestDeleteIndex(t *testing.T) {
	indices := []string{"userss3"}

	rsp, err := es_repository.DeleteIndex(indices)
	if err != nil {
		assert.Error(t, err)
	}
	dumps.Dump(rsp)
}

func TestSetESData(t *testing.T) {
	m1 := map[string][]string{
		"A": {"A1", "A2"},
		"B": {"B1"},
	}
	p1 := Person{Name: "zhangsan2", Age: 25, Birth: time.Now().Format(common.TimeFormat), Bool: true, Address: address{City: "武汉", State: "鄂"}, Class: m1}
	c1 := car{
		Brand:  "BYD",
		Colour: "红色",
		Model:  "唐",
	}
	cc1 := car{
		Brand:  "BYD",
		Colour: "蓝色",
		Model:  "汉",
	}
	p1.Cars = append(p1.Cars, c1, cc1)
	ret1, err1 := es_repository.SetESData("person", "_doc", strconv.Itoa(1), p1)
	if err1 != nil {
		assert.Error(t, err1)
	}
	dumps.Dump(ret1)

	m2 := map[string][]string{
		"A": {"000001"},
		"B": {"B1"},
		"C": {"C1"},
	}
	p2 := Person{Name: "lisi2", Age: 30, Birth: time.Now().Format(common.TimeFormat), Bool: false, Address: address{City: "北京", State: "京"}, Class: m2}
	c2 := car{
		Brand:  "长安",
		Colour: "灰色",
		Model:  "sl03",
	}
	cc2 := car{
		Brand:  "吉利",
		Colour: "黑色",
		Model:  "GS",
	}
	p2.Cars = append(p2.Cars, c2, cc2)
	ret2, err2 := es_repository.SetESData("person", "_doc", strconv.Itoa(2), p2)
	if err2 != nil {
		assert.Error(t, err2)
	}
	dumps.Dump(ret2)

}

func TestGetESDataById(t *testing.T) {
	ret, err := es_repository.GetESDataById("userss", "person", strconv.Itoa(234))
	if err != nil {
		assert.Error(t, err)
	}
	dumps.Dump(ret)
}

func TestQueryESData(t *testing.T) {
	index := "person"

	// 分页
	from := 0  // 偏移数
	size := 10 // 偏移量

	// 字段查询
	//tq1 := elastic.NewTermQuery("name", "lisi")
	//totalCount1, rsp1, err1 := es_repository.QueryESData(index, tq1, from, size, Person{})
	//if err1 != nil {
	//	assert.Error(t, err1)
	//	return
	//}
	//dumps.Dump(totalCount1)
	//dumps.Dump(rsp1)

	// 对象查询
	//bq := elastic.NewBoolQuery()
	//mq1 := elastic.NewMatchQuery("address.city", "北京")
	//mq2 := elastic.NewMatchQuery("address.state", "京")
	//
	//bq = bq.Must(mq1, mq2)
	//totalCount, rsp, err := es_repository.QueryESData(index, bq, from, size, Person{})
	//if err != nil {
	//	assert.Error(t, err)
	//	return
	//}
	//dumps.Dump(totalCount)
	//dumps.Dump(rsp)

	// 数组查询
	bq := elastic.NewBoolQuery()
	mq := elastic.NewMatchQuery("cars.brand", "BYD")
	tq := elastic.NewTermsQuery("cars.colour", "红色", "蓝色")

	// where in 条件
	//values := make([]interface{}, 0)
	//values = append(values, 25, 26, 30)
	//tsq := elastic.NewTermsQuery("age", values...)

	bq = bq.Must(mq, tq)

	//ih := elastic.NewInnerHit().Sort("cars.brand", true)
	nq := elastic.NewNestedQuery("cars", bq)

	query := nq
	//query := nq.InnerHit(ih)

	// 多字段排序：text类型不支持
	//sorter1 := elastic.NewFieldSort("birth").Asc()
	//sorter2 := elastic.NewFieldSort("age").Desc()

	totalCount2, rsp2, err2 := es_repository.QueryESData(index, query, from, size, Person{})
	if err2 != nil {
		assert.Error(t, err2)
		return
	}
	dumps.Dump(totalCount2)
	dumps.Dump(rsp2)

	boot.Destroy()
}

// map[string][]string 类查询
func TestQueryESData2(t *testing.T) {
	index := "person"

	// 分页
	from := 0  // 偏移数
	size := 10 // 偏移量

	bq := elastic.NewBoolQuery()
	mq1 := elastic.NewMatchQuery("class.A", "A1")
	mq2 := elastic.NewMatchQuery("class.A", "A2")
	//mq3 := elastic.NewMatchQuery("class.B", "B2")

	bq = bq.Must(mq1, mq2)

	//ih := elastic.NewInnerHit().Sort("class.key", true)
	nq := elastic.NewNestedQuery("class", bq)

	//query := nq.InnerHit(ih)

	totalCount, rsp, err := es_repository.QueryESData(index, nq, from, size, Person{})
	if err != nil {
		assert.Error(t, err)
		return
	}
	dumps.Dump(totalCount)
	dumps.Dump(rsp)

	boot.Destroy()
}
