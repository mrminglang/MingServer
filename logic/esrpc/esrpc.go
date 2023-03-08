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

//获取文章列表
//func GetArticleListFromES(req *Bar.GetListReq) comm.ES_Article_List {
//	log.Es.DebInfofugf("{GetArticleList req}|%s", req.Display())
//	reqJson := getArticleReqJson(req)
//	log.ES.Debugf("{articleReqJson json}|%s", reqJson)
//	var esReq FCS.QueryPureReq
//	esReq.Method = "POST"
//	esReq.Path = "/" + articleIndexName + "/_search"
//	esReq.PostBody = reqJson
//	var esRsp FCS.QueryPureRsp
//	esPrx.QueryPure(&esReq, &esRsp)
//	var esRspData comm.ES_Article_List
//	json.Unmarshal([]byte(esRsp.JsonRsp), &esRspData)
//
//	return esRspData
//}
