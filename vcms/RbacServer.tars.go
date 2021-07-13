// Package vcms comment
// This file was generated by tars2go 1.1.4
// Generated from rbac_server.tars
package vcms

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/TarsCloud/TarsGo/tars"
	m "github.com/TarsCloud/TarsGo/tars/model"
	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
	"github.com/TarsCloud/TarsGo/tars/protocol/res/basef"
	"github.com/TarsCloud/TarsGo/tars/protocol/res/requestf"
	"github.com/TarsCloud/TarsGo/tars/protocol/tup"
	"github.com/TarsCloud/TarsGo/tars/util/current"
	"github.com/TarsCloud/TarsGo/tars/util/tools"
	"unsafe"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = codec.FromInt8
var _ = unsafe.Pointer(nil)
var _ = bytes.ErrTooLarge

//RbacServer struct
type RbacServer struct {
	s m.Servant
}

//GetAccessAbilityByIdent is the proxy function for the method defined in the tars file, with the context
func (_obj *RbacServer) GetAccessAbilityByIdent(req string, rsp *AccessAbilityRsp, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_string(req, 1)
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
	_resp := new(requestf.ResponsePacket)
	tarsCtx := context.Background()

	err = _obj.s.Tars_invoke(tarsCtx, 0, "GetAccessAbilityByIdent", _os.ToBytes(), _status, _context, _resp)
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

//GetAccessAbilityByIdentWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *RbacServer) GetAccessAbilityByIdentWithContext(tarsCtx context.Context, req string, rsp *AccessAbilityRsp, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_string(req, 1)
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
	_resp := new(requestf.ResponsePacket)

	err = _obj.s.Tars_invoke(tarsCtx, 0, "GetAccessAbilityByIdent", _os.ToBytes(), _status, _context, _resp)
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

//GetAccessAbilityByIdentOneWayWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *RbacServer) GetAccessAbilityByIdentOneWayWithContext(tarsCtx context.Context, req string, rsp *AccessAbilityRsp, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_string(req, 1)
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
	_resp := new(requestf.ResponsePacket)

	err = _obj.s.Tars_invoke(tarsCtx, 1, "GetAccessAbilityByIdent", _os.ToBytes(), _status, _context, _resp)
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

//GetAccessPointInfo is the proxy function for the method defined in the tars file, with the context
func (_obj *RbacServer) GetAccessPointInfo(reqAccessPointId int32, rsp *AccessPointInfo, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_int32(reqAccessPointId, 1)
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
	_resp := new(requestf.ResponsePacket)
	tarsCtx := context.Background()

	err = _obj.s.Tars_invoke(tarsCtx, 0, "GetAccessPointInfo", _os.ToBytes(), _status, _context, _resp)
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

//GetAccessPointInfoWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *RbacServer) GetAccessPointInfoWithContext(tarsCtx context.Context, reqAccessPointId int32, rsp *AccessPointInfo, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_int32(reqAccessPointId, 1)
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
	_resp := new(requestf.ResponsePacket)

	err = _obj.s.Tars_invoke(tarsCtx, 0, "GetAccessPointInfo", _os.ToBytes(), _status, _context, _resp)
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

//GetAccessPointInfoOneWayWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *RbacServer) GetAccessPointInfoOneWayWithContext(tarsCtx context.Context, reqAccessPointId int32, rsp *AccessPointInfo, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_int32(reqAccessPointId, 1)
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
	_resp := new(requestf.ResponsePacket)

	err = _obj.s.Tars_invoke(tarsCtx, 1, "GetAccessPointInfo", _os.ToBytes(), _status, _context, _resp)
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

//IsValidAccess is the proxy function for the method defined in the tars file, with the context
func (_obj *RbacServer) IsValidAccess(req *IsValidAccessReq, rsp *bool, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = req.WriteBlock(_os, 1)
	if err != nil {
		return ret, err
	}

	err = _os.Write_bool((*rsp), 2)
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
	_resp := new(requestf.ResponsePacket)
	tarsCtx := context.Background()

	err = _obj.s.Tars_invoke(tarsCtx, 0, "IsValidAccess", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}

	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err = _is.Read_bool(&(*rsp), 2, true)
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

//IsValidAccessWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *RbacServer) IsValidAccessWithContext(tarsCtx context.Context, req *IsValidAccessReq, rsp *bool, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = req.WriteBlock(_os, 1)
	if err != nil {
		return ret, err
	}

	err = _os.Write_bool((*rsp), 2)
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
	_resp := new(requestf.ResponsePacket)

	err = _obj.s.Tars_invoke(tarsCtx, 0, "IsValidAccess", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}

	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err = _is.Read_bool(&(*rsp), 2, true)
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

//IsValidAccessOneWayWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *RbacServer) IsValidAccessOneWayWithContext(tarsCtx context.Context, req *IsValidAccessReq, rsp *bool, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = req.WriteBlock(_os, 1)
	if err != nil {
		return ret, err
	}

	err = _os.Write_bool((*rsp), 2)
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
	_resp := new(requestf.ResponsePacket)

	err = _obj.s.Tars_invoke(tarsCtx, 1, "IsValidAccess", _os.ToBytes(), _status, _context, _resp)
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
func (_obj *RbacServer) SetServant(s m.Servant) {
	_obj.s = s
}

//TarsSetTimeout sets the timeout for the servant which is in ms.
func (_obj *RbacServer) TarsSetTimeout(t int) {
	_obj.s.TarsSetTimeout(t)
}

//TarsSetProtocol sets the protocol for the servant.
func (_obj *RbacServer) TarsSetProtocol(p m.Protocol) {
	_obj.s.TarsSetProtocol(p)
}

//AddServant adds servant  for the service.
func (_obj *RbacServer) AddServant(imp _impRbacServer, obj string) {
	tars.AddServant(_obj, imp, obj)
}

//AddServant adds servant  for the service with context.
func (_obj *RbacServer) AddServantWithContext(imp _impRbacServerWithContext, obj string) {
	tars.AddServantWithContext(_obj, imp, obj)
}

type _impRbacServer interface {
	GetAccessAbilityByIdent(req string, rsp *AccessAbilityRsp) (ret int32, err error)
	GetAccessPointInfo(reqAccessPointId int32, rsp *AccessPointInfo) (ret int32, err error)
	IsValidAccess(req *IsValidAccessReq, rsp *bool) (ret int32, err error)
}
type _impRbacServerWithContext interface {
	GetAccessAbilityByIdent(tarsCtx context.Context, req string, rsp *AccessAbilityRsp) (ret int32, err error)
	GetAccessPointInfo(tarsCtx context.Context, reqAccessPointId int32, rsp *AccessPointInfo) (ret int32, err error)
	IsValidAccess(tarsCtx context.Context, req *IsValidAccessReq, rsp *bool) (ret int32, err error)
}

// Dispatch is used to call the server side implemnet for the method defined in the tars file. _withContext shows using context or not.
func (_obj *RbacServer) Dispatch(tarsCtx context.Context, _val interface{}, tarsReq *requestf.RequestPacket, tarsResp *requestf.ResponsePacket, _withContext bool) (err error) {
	var length int32
	var have bool
	var ty byte
	_is := codec.NewReader(tools.Int8ToByte(tarsReq.SBuffer))
	_os := codec.NewBuffer()
	switch tarsReq.SFuncName {
	case "GetAccessAbilityByIdent":
		var req string
		var rsp AccessAbilityRsp

		if tarsReq.IVersion == basef.TARSVERSION {

			err = _is.Read_string(&req, 1, true)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.TUPVERSION {
			_reqTup_ := tup.NewUniAttribute()
			_reqTup_.Decode(_is)

			var _tupBuffer_ []byte

			_reqTup_.GetBuffer("req", &_tupBuffer_)
			_is.Reset(_tupBuffer_)
			err = _is.Read_string(&req, 0, true)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.JSONVERSION {
			var _jsonDat_ map[string]interface{}
			_decoder_ := json.NewDecoder(bytes.NewReader(_is.ToBytes()))
			_decoder_.UseNumber()
			err = _decoder_.Decode(&_jsonDat_)
			if err != nil {
				return fmt.Errorf("Decode reqpacket failed, error: %+v", err)
			}
			{
				_jsonStr_, _ := json.Marshal(_jsonDat_["req"])
				if err = json.Unmarshal([]byte(_jsonStr_), &req); err != nil {
					return err
				}
			}

		} else {
			err = fmt.Errorf("Decode reqpacket fail, error version: %d", tarsReq.IVersion)
			return err
		}

		var _funRet_ int32
		if _withContext == false {
			_imp := _val.(_impRbacServer)
			_funRet_, err = _imp.GetAccessAbilityByIdent(req, &rsp)
		} else {
			_imp := _val.(_impRbacServerWithContext)
			_funRet_, err = _imp.GetAccessAbilityByIdent(tarsCtx, req, &rsp)
		}

		if err != nil {
			return err
		}

		if tarsReq.IVersion == basef.TARSVERSION {
			_os.Reset()

			err = _os.Write_int32(_funRet_, 0)
			if err != nil {
				return err
			}

			err = rsp.WriteBlock(_os, 2)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.TUPVERSION {
			_tupRsp_ := tup.NewUniAttribute()

			err = _os.Write_int32(_funRet_, 0)
			if err != nil {
				return err
			}

			_tupRsp_.PutBuffer("", _os.ToBytes())
			_tupRsp_.PutBuffer("tars_ret", _os.ToBytes())

			_os.Reset()
			err = rsp.WriteBlock(_os, 0)
			if err != nil {
				return err
			}

			_tupRsp_.PutBuffer("rsp", _os.ToBytes())

			_os.Reset()
			err = _tupRsp_.Encode(_os)
			if err != nil {
				return err
			}
		} else if tarsReq.IVersion == basef.JSONVERSION {
			_rspJson_ := map[string]interface{}{}
			_rspJson_["tars_ret"] = _funRet_
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
	case "GetAccessPointInfo":
		var reqAccessPointId int32
		var rsp AccessPointInfo

		if tarsReq.IVersion == basef.TARSVERSION {

			err = _is.Read_int32(&reqAccessPointId, 1, true)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.TUPVERSION {
			_reqTup_ := tup.NewUniAttribute()
			_reqTup_.Decode(_is)

			var _tupBuffer_ []byte

			_reqTup_.GetBuffer("reqAccessPointId", &_tupBuffer_)
			_is.Reset(_tupBuffer_)
			err = _is.Read_int32(&reqAccessPointId, 0, true)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.JSONVERSION {
			var _jsonDat_ map[string]interface{}
			_decoder_ := json.NewDecoder(bytes.NewReader(_is.ToBytes()))
			_decoder_.UseNumber()
			err = _decoder_.Decode(&_jsonDat_)
			if err != nil {
				return fmt.Errorf("Decode reqpacket failed, error: %+v", err)
			}
			{
				_jsonStr_, _ := json.Marshal(_jsonDat_["reqAccessPointId"])
				if err = json.Unmarshal([]byte(_jsonStr_), &reqAccessPointId); err != nil {
					return err
				}
			}

		} else {
			err = fmt.Errorf("Decode reqpacket fail, error version: %d", tarsReq.IVersion)
			return err
		}

		var _funRet_ int32
		if _withContext == false {
			_imp := _val.(_impRbacServer)
			_funRet_, err = _imp.GetAccessPointInfo(reqAccessPointId, &rsp)
		} else {
			_imp := _val.(_impRbacServerWithContext)
			_funRet_, err = _imp.GetAccessPointInfo(tarsCtx, reqAccessPointId, &rsp)
		}

		if err != nil {
			return err
		}

		if tarsReq.IVersion == basef.TARSVERSION {
			_os.Reset()

			err = _os.Write_int32(_funRet_, 0)
			if err != nil {
				return err
			}

			err = rsp.WriteBlock(_os, 2)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.TUPVERSION {
			_tupRsp_ := tup.NewUniAttribute()

			err = _os.Write_int32(_funRet_, 0)
			if err != nil {
				return err
			}

			_tupRsp_.PutBuffer("", _os.ToBytes())
			_tupRsp_.PutBuffer("tars_ret", _os.ToBytes())

			_os.Reset()
			err = rsp.WriteBlock(_os, 0)
			if err != nil {
				return err
			}

			_tupRsp_.PutBuffer("rsp", _os.ToBytes())

			_os.Reset()
			err = _tupRsp_.Encode(_os)
			if err != nil {
				return err
			}
		} else if tarsReq.IVersion == basef.JSONVERSION {
			_rspJson_ := map[string]interface{}{}
			_rspJson_["tars_ret"] = _funRet_
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
	case "IsValidAccess":
		var req IsValidAccessReq
		var rsp bool

		if tarsReq.IVersion == basef.TARSVERSION {

			err = req.ReadBlock(_is, 1, true)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.TUPVERSION {
			_reqTup_ := tup.NewUniAttribute()
			_reqTup_.Decode(_is)

			var _tupBuffer_ []byte

			_reqTup_.GetBuffer("req", &_tupBuffer_)
			_is.Reset(_tupBuffer_)
			err = req.ReadBlock(_is, 0, true)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.JSONVERSION {
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
			err = fmt.Errorf("Decode reqpacket fail, error version: %d", tarsReq.IVersion)
			return err
		}

		var _funRet_ int32
		if _withContext == false {
			_imp := _val.(_impRbacServer)
			_funRet_, err = _imp.IsValidAccess(&req, &rsp)
		} else {
			_imp := _val.(_impRbacServerWithContext)
			_funRet_, err = _imp.IsValidAccess(tarsCtx, &req, &rsp)
		}

		if err != nil {
			return err
		}

		if tarsReq.IVersion == basef.TARSVERSION {
			_os.Reset()

			err = _os.Write_int32(_funRet_, 0)
			if err != nil {
				return err
			}

			err = _os.Write_bool(rsp, 2)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.TUPVERSION {
			_tupRsp_ := tup.NewUniAttribute()

			err = _os.Write_int32(_funRet_, 0)
			if err != nil {
				return err
			}

			_tupRsp_.PutBuffer("", _os.ToBytes())
			_tupRsp_.PutBuffer("tars_ret", _os.ToBytes())

			_os.Reset()
			err = _os.Write_bool(rsp, 0)
			if err != nil {
				return err
			}

			_tupRsp_.PutBuffer("rsp", _os.ToBytes())

			_os.Reset()
			err = _tupRsp_.Encode(_os)
			if err != nil {
				return err
			}
		} else if tarsReq.IVersion == basef.JSONVERSION {
			_rspJson_ := map[string]interface{}{}
			_rspJson_["tars_ret"] = _funRet_
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

	default:
		return fmt.Errorf("func mismatch")
	}
	var _status map[string]string
	s, ok := current.GetResponseStatus(tarsCtx)
	if ok && s != nil {
		_status = s
	}
	var _context map[string]string
	c, ok := current.GetResponseContext(tarsCtx)
	if ok && c != nil {
		_context = c
	}
	*tarsResp = requestf.ResponsePacket{
		IVersion:     tarsReq.IVersion,
		CPacketType:  0,
		IRequestId:   tarsReq.IRequestId,
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