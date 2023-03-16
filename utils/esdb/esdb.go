package esdb

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"gitlab.upchinaproduct.com/taf/tafgo/taf/util/conf"
	"server/logic/esrpc"
	"server/taf-protocol/FCS"
	"server/utils/log"
	"time"
)

var Client *elastic.Client

func Init(conf *conf.Conf) {
	log.Es.Infof("{esdb init start......}")

	// 获取ES集群配置
	esRsp, err := esrpc.GetESClusterList(FCS.GetESClusterListReq{})
	if err != nil {
		log.Es.Errorf("{esdb init GetESClusterList error|%s}", err.Error())
		return
	}

	hosts := make([]string, 0)
	for _, row := range esRsp.EsClusters {
		http := fmt.Sprintf("http://%s:%s", row.Host, row.Port)
		hosts = append(hosts, http)
	}
	log.Es.Infof("{esdb init hosts|%s}", hosts)

	client, err := elastic.NewClient(
		elastic.SetURL(hosts...), // 服务地址
		elastic.SetBasicAuth(conf.GetString("/esConf/<username>"), conf.GetString("/esConf/<password>")),
		elastic.SetHealthcheckInterval(time.Second*5), // 心跳
		elastic.SetMaxRetries(3), // 重试次数
		elastic.SetSniff(false),
	)
	if err != nil {
		log.Es.Errorf("{esdb init error|%s}", err.Error())
		return
	}

	for _, host := range hosts {
		_, _, err := client.Ping(host).Do(context.Background())
		if err != nil {
			log.Es.Errorf("{esdb init Ping error|%s}", err.Error())
			continue
		}

		_, err = client.ElasticsearchVersion(host)
		if err != nil {
			log.Es.Errorf("{esdb init ElasticsearchVersion error|%s}", err.Error())
			continue
		}
	}

	Client = client
	log.Es.Infof("{esdb init success......}")
	return
}
