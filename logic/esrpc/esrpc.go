package esrpc

import (
	"gitlab.upchinaproduct.com/upgo/utils/server_utils/log"
	"server/taf-protocol/FCS"
	"server/utils/trpc"
)

// RPC 请求es
func QueryPure(esReq FCS.QueryPureReq) (esRsp FCS.QueryPureRsp, err error) {
	log.Es.Infof("{ES RPC QueryPure req::%s}", esReq.Display())
	_, err = trpc.ESPrx.QueryPure(&esReq, &esRsp)
	if err != nil {
		log.Es.Errorf("{ES RPC QueryPure error::%s}", err.Error())
		return
	}
	log.Es.Infof("{ES RPC QueryPure esRsp::%s}", esRsp.Display())
	return
}

// RPC 批量请求es
func QueryPureBatch(esReq FCS.QueryPureBatchReq) (esRsp FCS.QueryPureBatchRsp, err error) {
	log.Es.Infof("{ES RPC QueryPureBatch req::%s}", esReq.Display())
	_, err = trpc.ESPrx.QueryPureBatch(&esReq, &esRsp)
	if err != nil {
		log.Es.Errorf("{ES RPC QueryPureBatch error::%s}", err.Error())
		return
	}
	log.Es.Infof("{ES RPC QueryPureBatch esRsp::%s}", esRsp.Display())
	return
}

// RPC 获取ES集群配置列表
func GetESClusterList(esReq FCS.GetESClusterListReq) (esRsp FCS.GetESClusterListRsp, err error) {
	log.Es.Infof("{ES RPC GetESClusterList req::%s}", esReq.Display())
	_, err = trpc.ESPrx.GetESClusterList(&esReq, &esRsp)
	if err != nil {
		log.Es.Errorf("{ES RPC GetESClusterList error::%s}", err.Error())
		return
	}
	log.Es.Infof("{ES RPC GetESClusterList esRsp::%s}", esRsp.Display())
	return
}
