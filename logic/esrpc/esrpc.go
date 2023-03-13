package esrpc

import (
	"server/taf-protocol/FCS"
	"server/utils/log"
	"server/utils/taftool"
)

var esPrx FCS.ESDriver

func Init(esObj string) {
	//初始化esrpc链接
	log.Es.Infof("init esrpc server start....")
	taftool.InitTafPrx(&esPrx, esObj)
	log.Es.Infof("init esrpc server success....")
}

// RPC 请求es
func QueryPure(esReq FCS.QueryPureReq) (esRsp FCS.QueryPureRsp, err error) {
	log.Es.Infof("{ES RPC QueryPure req::%s}", esReq.Display())
	_, err = esPrx.QueryPure(&esReq, &esRsp)
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
	_, err = esPrx.QueryPureBatch(&esReq, &esRsp)
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
	_, err = esPrx.GetESClusterList(&esReq, &esRsp)
	if err != nil {
		log.Es.Errorf("{ES RPC GetESClusterList error::%s}", err.Error())
		return
	}
	log.Es.Infof("{ES RPC GetESClusterList esRsp::%s}", esRsp.Display())
	return
}
