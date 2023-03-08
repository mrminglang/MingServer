// Package FCS comment
// This file was generated by jce2go 2.0.0
// Generated from ESDriver.jce
package FCS

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"gitlab.upchinaproduct.com/taf/tafgo/taf"
	m "gitlab.upchinaproduct.com/taf/tafgo/taf/model"
	"gitlab.upchinaproduct.com/taf/tafgo/taf/protocol/codec"
	"gitlab.upchinaproduct.com/taf/tafgo/taf/protocol/res/basef"
	"gitlab.upchinaproduct.com/taf/tafgo/taf/protocol/res/requestf"
	"gitlab.upchinaproduct.com/taf/tafgo/taf/protocol/wup"
	"gitlab.upchinaproduct.com/taf/tafgo/taf/util/current"
	"gitlab.upchinaproduct.com/taf/tafgo/taf/util/tools"
	"gitlab.upchinaproduct.com/taf/tafgo/taf/util/trace"
	"unsafe"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = codec.FromInt8
var _ = unsafe.Pointer(nil)
var _ = bytes.ErrTooLarge

//ESDriver struct
type ESDriver struct {
	s m.Servant
}

//QueryPure is the proxy function for the method defined in the jce file, with the context
func (_obj *ESDriver) QueryPure(req *QueryPureReq, rsp *QueryPureRsp, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = req.WriteBlock(_os, 1)
	if err != nil {
		return ret, err
	}

	err = (*rsp).WriteBlock(_os, 2)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.BasePacket)
	tafCtx := context.Background()

	err = _obj.s.Taf_invoke(tafCtx, 0, "queryPure", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}

	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err = (*rsp).ReadBlock(_is, 2, true)
	if err != nil {
		return ret, err
	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range _resp.Status {
			_status[k] = v
		}

	}
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//QueryPureWithContext is the proxy function for the method defined in the jce file, with the context
func (_obj *ESDriver) QueryPureWithContext(tafCtx context.Context, req *QueryPureReq, rsp *QueryPureRsp, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = req.WriteBlock(_os, 1)
	if err != nil {
		return ret, err
	}

	err = (*rsp).WriteBlock(_os, 2)
	if err != nil {
		return ret, err
	}

	traceData, ok := current.GetTraceData(tafCtx)
	if ok && traceData.TraceCall {
		traceData.NewSpan()
		var traceParam string
		traceParamFlag := traceData.NeedTraceParam(trace.EstCS, uint(_os.Len()))
		if traceParamFlag == trace.EnpNormal {
			value := map[string]interface{}{}
			value["req"] = req
			p, _ := json.Marshal(value)
			traceParam = string(p)
		} else if traceParamFlag == trace.EnpOverMaxLen {
			traceParam = "{\"trace_param_over_max_len\":true}"
		}
		taf.Trace(traceData.GetTraceKey(trace.EstCS), trace.TraceAnnotationCS, taf.GetClientConfig().ModuleName, _obj.s.Name(), "QueryPure", 0, traceParam, "")
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.BasePacket)

	err = _obj.s.Taf_invoke(tafCtx, 0, "queryPure", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}

	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err = (*rsp).ReadBlock(_is, 2, true)
	if err != nil {
		return ret, err
	}

	if ok && traceData.TraceCall {
		var traceParam string
		traceParamFlag := traceData.NeedTraceParam(trace.EstCR, uint(_is.Len()))
		if traceParamFlag == trace.EnpNormal {
			value := map[string]interface{}{}
			value[""] = ret
			value["taf_ret"] = ret
			value["rsp"] = *rsp
			p, _ := json.Marshal(value)
			traceParam = string(p)
		} else if traceParamFlag == trace.EnpOverMaxLen {
			traceParam = "{\"trace_param_over_max_len\":true}"
		}
		taf.Trace(traceData.GetTraceKey(trace.EstCR), trace.TraceAnnotationCR, taf.GetClientConfig().ModuleName, _obj.s.Name(), "QueryPure", int(_resp.IRet), traceParam, "")
	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range _resp.Status {
			_status[k] = v
		}

	}
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//QueryPureOneWayWithContext is the proxy function for the method defined in the jce file, with the context
func (_obj *ESDriver) QueryPureOneWayWithContext(tafCtx context.Context, req *QueryPureReq, rsp *QueryPureRsp, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = req.WriteBlock(_os, 1)
	if err != nil {
		return ret, err
	}

	err = (*rsp).WriteBlock(_os, 2)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.BasePacket)

	err = _obj.s.Taf_invoke(tafCtx, 1, "queryPure", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range _resp.Status {
			_status[k] = v
		}

	}
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//QueryPureBatch is the proxy function for the method defined in the jce file, with the context
func (_obj *ESDriver) QueryPureBatch(req *QueryPureBatchReq, rsp *QueryPureBatchRsp, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = req.WriteBlock(_os, 1)
	if err != nil {
		return ret, err
	}

	err = (*rsp).WriteBlock(_os, 2)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.BasePacket)
	tafCtx := context.Background()

	err = _obj.s.Taf_invoke(tafCtx, 0, "queryPureBatch", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}

	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err = (*rsp).ReadBlock(_is, 2, true)
	if err != nil {
		return ret, err
	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range _resp.Status {
			_status[k] = v
		}

	}
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//QueryPureBatchWithContext is the proxy function for the method defined in the jce file, with the context
func (_obj *ESDriver) QueryPureBatchWithContext(tafCtx context.Context, req *QueryPureBatchReq, rsp *QueryPureBatchRsp, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = req.WriteBlock(_os, 1)
	if err != nil {
		return ret, err
	}

	err = (*rsp).WriteBlock(_os, 2)
	if err != nil {
		return ret, err
	}

	traceData, ok := current.GetTraceData(tafCtx)
	if ok && traceData.TraceCall {
		traceData.NewSpan()
		var traceParam string
		traceParamFlag := traceData.NeedTraceParam(trace.EstCS, uint(_os.Len()))
		if traceParamFlag == trace.EnpNormal {
			value := map[string]interface{}{}
			value["req"] = req
			p, _ := json.Marshal(value)
			traceParam = string(p)
		} else if traceParamFlag == trace.EnpOverMaxLen {
			traceParam = "{\"trace_param_over_max_len\":true}"
		}
		taf.Trace(traceData.GetTraceKey(trace.EstCS), trace.TraceAnnotationCS, taf.GetClientConfig().ModuleName, _obj.s.Name(), "QueryPureBatch", 0, traceParam, "")
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.BasePacket)

	err = _obj.s.Taf_invoke(tafCtx, 0, "queryPureBatch", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}

	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err = (*rsp).ReadBlock(_is, 2, true)
	if err != nil {
		return ret, err
	}

	if ok && traceData.TraceCall {
		var traceParam string
		traceParamFlag := traceData.NeedTraceParam(trace.EstCR, uint(_is.Len()))
		if traceParamFlag == trace.EnpNormal {
			value := map[string]interface{}{}
			value[""] = ret
			value["taf_ret"] = ret
			value["rsp"] = *rsp
			p, _ := json.Marshal(value)
			traceParam = string(p)
		} else if traceParamFlag == trace.EnpOverMaxLen {
			traceParam = "{\"trace_param_over_max_len\":true}"
		}
		taf.Trace(traceData.GetTraceKey(trace.EstCR), trace.TraceAnnotationCR, taf.GetClientConfig().ModuleName, _obj.s.Name(), "QueryPureBatch", int(_resp.IRet), traceParam, "")
	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range _resp.Status {
			_status[k] = v
		}

	}
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//QueryPureBatchOneWayWithContext is the proxy function for the method defined in the jce file, with the context
func (_obj *ESDriver) QueryPureBatchOneWayWithContext(tafCtx context.Context, req *QueryPureBatchReq, rsp *QueryPureBatchRsp, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = req.WriteBlock(_os, 1)
	if err != nil {
		return ret, err
	}

	err = (*rsp).WriteBlock(_os, 2)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.BasePacket)

	err = _obj.s.Taf_invoke(tafCtx, 1, "queryPureBatch", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range _resp.Status {
			_status[k] = v
		}

	}
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//SetServant sets servant for the service.
func (_obj *ESDriver) SetServant(s m.Servant) {
	_obj.s = s
}

//TafSetTimeout sets the timeout for the servant which is in ms.
func (_obj *ESDriver) TafSetTimeout(t int) {
	_obj.s.TafSetTimeout(t)
}

//TafSetProtocol sets the protocol for the servant.
func (_obj *ESDriver) TafSetProtocol(p m.Protocol) {
	_obj.s.TafSetProtocol(p)
}

//AddServant adds servant  for the service.
func (_obj *ESDriver) AddServant(imp _impESDriver, obj string) {
	taf.AddServant(_obj, imp, obj)
}

//AddServantWithContext adds servant  for the service with context.
func (_obj *ESDriver) AddServantWithContext(imp _impESDriverWithContext, obj string) {
	taf.AddServantWithContext(_obj, imp, obj)
}

type _impESDriver interface {
	QueryPure(req *QueryPureReq, rsp *QueryPureRsp) (ret int32, err error)
	QueryPureBatch(req *QueryPureBatchReq, rsp *QueryPureBatchRsp) (ret int32, err error)
}
type _impESDriverWithContext interface {
	QueryPure(tafCtx context.Context, req *QueryPureReq, rsp *QueryPureRsp) (ret int32, err error)
	QueryPureBatch(tafCtx context.Context, req *QueryPureBatchReq, rsp *QueryPureBatchRsp) (ret int32, err error)
}

// Dispatch is used to call the server side implemnet for the method defined in the jce file. _withContext shows using context or not.
func (_obj *ESDriver) Dispatch(tafCtx context.Context, _val interface{}, tafReq *requestf.BasePacket, tafResp *requestf.BasePacket, _withContext bool) (err error) {
	var length int32
	var have bool
	var ty byte
	_is := codec.NewReader(tools.Int8ToByte(tafReq.SBuffer))
	_os := codec.NewBuffer()
	switch tafReq.SFuncName {
	case "queryPure":
		var req QueryPureReq
		var rsp QueryPureRsp

		if tafReq.IVersion == basef.JCEVERSION {

			err = req.ReadBlock(_is, 1, true)
			if err != nil {
				return err
			}

		} else if tafReq.IVersion == basef.WUPVERSION {
			_reqWup_ := wup.NewUniAttribute()
			_reqWup_.Decode(_is)

			var _wupBuffer_ []byte

			_reqWup_.GetBuffer("req", &_wupBuffer_)
			_is.Reset(_wupBuffer_)
			err = req.ReadBlock(_is, 0, true)
			if err != nil {
				return err
			}

		} else if tafReq.IVersion == basef.JSONVERSION {
			var _jsonDat_ map[string]interface{}
			_decoder_ := json.NewDecoder(bytes.NewReader(_is.ToBytes()))
			_decoder_.UseNumber()
			err = _decoder_.Decode(&_jsonDat_)
			if err != nil {
				return fmt.Errorf("Decode reqpacket failed, error: %+v", err)
			}
			{
				_jsonStr_, _ := json.Marshal(_jsonDat_["req"])
				req.ResetDefault()
				if err = json.Unmarshal([]byte(_jsonStr_), &req); err != nil {
					return err
				}
			}

		} else {
			err = fmt.Errorf("Decode reqpacket fail, error version: %d", tafReq.IVersion)
			return err
		}

		traceData, ok := current.GetTraceData(tafCtx)
		if ok && traceData.TraceCall {
			var traceParam string
			traceParamFlag := traceData.NeedTraceParam(trace.EstSR, uint(_is.Len()))
			if traceParamFlag == trace.EnpNormal {
				value := map[string]interface{}{}
				value["req"] = req
				p, _ := json.Marshal(value)
				traceParam = string(p)
			} else if traceParamFlag == trace.EnpOverMaxLen {
				traceParam = "{\"trace_param_over_max_len\":true}"
			}
			taf.Trace(traceData.GetTraceKey(trace.EstSR), trace.TraceAnnotationSR, taf.GetClientConfig().ModuleName, tafReq.SServantName, "queryPure", 0, traceParam, "")
		}

		var _funRet_ int32
		if _withContext == false {
			_imp := _val.(_impESDriver)
			_funRet_, err = _imp.QueryPure(&req, &rsp)
		} else {
			_imp := _val.(_impESDriverWithContext)
			_funRet_, err = _imp.QueryPure(tafCtx, &req, &rsp)
		}

		if err != nil {
			return err
		}

		if tafReq.IVersion == basef.JCEVERSION {
			_os.Reset()

			err = _os.Write_int32(_funRet_, 0)
			if err != nil {
				return err
			}

			err = rsp.WriteBlock(_os, 2)
			if err != nil {
				return err
			}

		} else if tafReq.IVersion == basef.WUPVERSION {
			_wupRsp_ := wup.NewUniAttribute()

			err = _os.Write_int32(_funRet_, 0)
			if err != nil {
				return err
			}

			_wupRsp_.PutBuffer("", _os.ToBytes())
			_wupRsp_.PutBuffer("taf_ret", _os.ToBytes())

			_os.Reset()
			err = rsp.WriteBlock(_os, 0)
			if err != nil {
				return err
			}

			_wupRsp_.PutBuffer("rsp", _os.ToBytes())

			_os.Reset()
			err = _wupRsp_.Encode(_os)
			if err != nil {
				return err
			}
		} else if tafReq.IVersion == basef.JSONVERSION {
			_rspJson_ := map[string]interface{}{}
			_rspJson_["taf_ret"] = _funRet_
			_rspJson_[""] = _funRet_
			_rspJson_["rsp"] = rsp

			var _rspByte_ []byte
			if _rspByte_, err = json.Marshal(_rspJson_); err != nil {
				return err
			}

			_os.Reset()
			err = _os.Write_slice_uint8(_rspByte_)
			if err != nil {
				return err
			}
		}

		if ok && traceData.TraceCall {
			var traceParam string
			traceParamFlag := traceData.NeedTraceParam(trace.EstSS, uint(_os.Len()))
			if traceParamFlag == trace.EnpNormal {
				value := map[string]interface{}{}
				value[""] = _funRet_
				value["taf_ret"] = _funRet_
				value["rsp"] = rsp
				p, _ := json.Marshal(value)
				traceParam = string(p)
			} else if traceParamFlag == trace.EnpOverMaxLen {
				traceParam = "{\"trace_param_over_max_len\":true}"
			}
			taf.Trace(traceData.GetTraceKey(trace.EstSS), trace.TraceAnnotationSS, taf.GetClientConfig().ModuleName, tafReq.SServantName, "queryPure", 0, traceParam, "")
		}

	case "queryPureBatch":
		var req QueryPureBatchReq
		var rsp QueryPureBatchRsp

		if tafReq.IVersion == basef.JCEVERSION {

			err = req.ReadBlock(_is, 1, true)
			if err != nil {
				return err
			}

		} else if tafReq.IVersion == basef.WUPVERSION {
			_reqWup_ := wup.NewUniAttribute()
			_reqWup_.Decode(_is)

			var _wupBuffer_ []byte

			_reqWup_.GetBuffer("req", &_wupBuffer_)
			_is.Reset(_wupBuffer_)
			err = req.ReadBlock(_is, 0, true)
			if err != nil {
				return err
			}

		} else if tafReq.IVersion == basef.JSONVERSION {
			var _jsonDat_ map[string]interface{}
			_decoder_ := json.NewDecoder(bytes.NewReader(_is.ToBytes()))
			_decoder_.UseNumber()
			err = _decoder_.Decode(&_jsonDat_)
			if err != nil {
				return fmt.Errorf("Decode reqpacket failed, error: %+v", err)
			}
			{
				_jsonStr_, _ := json.Marshal(_jsonDat_["req"])
				req.ResetDefault()
				if err = json.Unmarshal([]byte(_jsonStr_), &req); err != nil {
					return err
				}
			}

		} else {
			err = fmt.Errorf("Decode reqpacket fail, error version: %d", tafReq.IVersion)
			return err
		}

		traceData, ok := current.GetTraceData(tafCtx)
		if ok && traceData.TraceCall {
			var traceParam string
			traceParamFlag := traceData.NeedTraceParam(trace.EstSR, uint(_is.Len()))
			if traceParamFlag == trace.EnpNormal {
				value := map[string]interface{}{}
				value["req"] = req
				p, _ := json.Marshal(value)
				traceParam = string(p)
			} else if traceParamFlag == trace.EnpOverMaxLen {
				traceParam = "{\"trace_param_over_max_len\":true}"
			}
			taf.Trace(traceData.GetTraceKey(trace.EstSR), trace.TraceAnnotationSR, taf.GetClientConfig().ModuleName, tafReq.SServantName, "queryPureBatch", 0, traceParam, "")
		}

		var _funRet_ int32
		if _withContext == false {
			_imp := _val.(_impESDriver)
			_funRet_, err = _imp.QueryPureBatch(&req, &rsp)
		} else {
			_imp := _val.(_impESDriverWithContext)
			_funRet_, err = _imp.QueryPureBatch(tafCtx, &req, &rsp)
		}

		if err != nil {
			return err
		}

		if tafReq.IVersion == basef.JCEVERSION {
			_os.Reset()

			err = _os.Write_int32(_funRet_, 0)
			if err != nil {
				return err
			}

			err = rsp.WriteBlock(_os, 2)
			if err != nil {
				return err
			}

		} else if tafReq.IVersion == basef.WUPVERSION {
			_wupRsp_ := wup.NewUniAttribute()

			err = _os.Write_int32(_funRet_, 0)
			if err != nil {
				return err
			}

			_wupRsp_.PutBuffer("", _os.ToBytes())
			_wupRsp_.PutBuffer("taf_ret", _os.ToBytes())

			_os.Reset()
			err = rsp.WriteBlock(_os, 0)
			if err != nil {
				return err
			}

			_wupRsp_.PutBuffer("rsp", _os.ToBytes())

			_os.Reset()
			err = _wupRsp_.Encode(_os)
			if err != nil {
				return err
			}
		} else if tafReq.IVersion == basef.JSONVERSION {
			_rspJson_ := map[string]interface{}{}
			_rspJson_["taf_ret"] = _funRet_
			_rspJson_[""] = _funRet_
			_rspJson_["rsp"] = rsp

			var _rspByte_ []byte
			if _rspByte_, err = json.Marshal(_rspJson_); err != nil {
				return err
			}

			_os.Reset()
			err = _os.Write_slice_uint8(_rspByte_)
			if err != nil {
				return err
			}
		}

		if ok && traceData.TraceCall {
			var traceParam string
			traceParamFlag := traceData.NeedTraceParam(trace.EstSS, uint(_os.Len()))
			if traceParamFlag == trace.EnpNormal {
				value := map[string]interface{}{}
				value[""] = _funRet_
				value["taf_ret"] = _funRet_
				value["rsp"] = rsp
				p, _ := json.Marshal(value)
				traceParam = string(p)
			} else if traceParamFlag == trace.EnpOverMaxLen {
				traceParam = "{\"trace_param_over_max_len\":true}"
			}
			taf.Trace(traceData.GetTraceKey(trace.EstSS), trace.TraceAnnotationSS, taf.GetClientConfig().ModuleName, tafReq.SServantName, "queryPureBatch", 0, traceParam, "")
		}

	default:
		return fmt.Errorf("func mismatch")
	}
	var _status map[string]string
	s, ok := current.GetResponseStatus(tafCtx)
	if ok && s != nil {
		_status = s
	}
	var _context map[string]string
	c, ok := current.GetResponseContext(tafCtx)
	if ok && c != nil {
		_context = c
	}
	*tafResp = requestf.BasePacket{
		IVersion:     tafReq.IVersion,
		CPacketType:  0,
		IRequestId:   tafReq.IRequestId,
		IMessageType: 0,
		IRet:         0,
		SBuffer:      tools.ByteToInt8(_os.ToBytes()),
		Status:       _status,
		SResultDesc:  "",
		Context:      _context,
	}

	_ = _is
	_ = _os
	_ = length
	_ = have
	_ = ty
	return nil
}