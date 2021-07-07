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

//AccessPointOperate struct
type AccessPointOperate struct {
	s m.Servant
}

//CreateAccessPoint is the proxy function for the method defined in the tars file, with the context
func (_obj *AccessPointOperate) CreateAccessPoint(req *CreateAccessPointReq, rsp *string, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = req.WriteBlock(_os, 1)
	if err != nil {
		return ret, err
	}

	err = _os.Write_string((*rsp), 2)
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

	err = _obj.s.Tars_invoke(tarsCtx, 0, "CreateAccessPoint", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}

	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err = _is.Read_string(&(*rsp), 2, true)
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

//CreateAccessPointWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *AccessPointOperate) CreateAccessPointWithContext(tarsCtx context.Context, req *CreateAccessPointReq, rsp *string, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = req.WriteBlock(_os, 1)
	if err != nil {
		return ret, err
	}

	err = _os.Write_string((*rsp), 2)
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

	err = _obj.s.Tars_invoke(tarsCtx, 0, "CreateAccessPoint", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}

	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err = _is.Read_string(&(*rsp), 2, true)
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

//CreateAccessPointOneWayWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *AccessPointOperate) CreateAccessPointOneWayWithContext(tarsCtx context.Context, req *CreateAccessPointReq, rsp *string, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = req.WriteBlock(_os, 1)
	if err != nil {
		return ret, err
	}

	err = _os.Write_string((*rsp), 2)
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

	err = _obj.s.Tars_invoke(tarsCtx, 1, "CreateAccessPoint", _os.ToBytes(), _status, _context, _resp)
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

//GetAccessPoints is the proxy function for the method defined in the tars file, with the context
func (_obj *AccessPointOperate) GetAccessPoints(req *GetAccessPointsReq, rsp *[]AccessPointInfo, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = req.WriteBlock(_os, 1)
	if err != nil {
		return ret, err
	}

	err = _os.WriteHead(codec.LIST, 2)
	if err != nil {
		return ret, err
	}

	err = _os.Write_int32(int32(len((*rsp))), 0)
	if err != nil {
		return ret, err
	}

	for _, v := range *rsp {

		err = v.WriteBlock(_os, 0)
		if err != nil {
			return ret, err
		}

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

	err = _obj.s.Tars_invoke(tarsCtx, 0, "GetAccessPoints", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}

	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err, have, ty = _is.SkipToNoCheck(2, true)
	if err != nil {
		return ret, err
	}

	if ty == codec.LIST {
		err = _is.Read_int32(&length, 0, true)
		if err != nil {
			return ret, err
		}

		(*rsp) = make([]AccessPointInfo, length)
		for i0, e0 := int32(0), length; i0 < e0; i0++ {

			err = (*rsp)[i0].ReadBlock(_is, 0, false)
			if err != nil {
				return ret, err
			}

		}
	} else if ty == codec.SIMPLE_LIST {
		err = fmt.Errorf("not support simple_list type")
		if err != nil {
			return ret, err
		}

	} else {
		err = fmt.Errorf("require vector, but not")
		if err != nil {
			return ret, err
		}

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

//GetAccessPointsWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *AccessPointOperate) GetAccessPointsWithContext(tarsCtx context.Context, req *GetAccessPointsReq, rsp *[]AccessPointInfo, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = req.WriteBlock(_os, 1)
	if err != nil {
		return ret, err
	}

	err = _os.WriteHead(codec.LIST, 2)
	if err != nil {
		return ret, err
	}

	err = _os.Write_int32(int32(len((*rsp))), 0)
	if err != nil {
		return ret, err
	}

	for _, v := range *rsp {

		err = v.WriteBlock(_os, 0)
		if err != nil {
			return ret, err
		}

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

	err = _obj.s.Tars_invoke(tarsCtx, 0, "GetAccessPoints", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}

	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err, have, ty = _is.SkipToNoCheck(2, true)
	if err != nil {
		return ret, err
	}

	if ty == codec.LIST {
		err = _is.Read_int32(&length, 0, true)
		if err != nil {
			return ret, err
		}

		(*rsp) = make([]AccessPointInfo, length)
		for i1, e1 := int32(0), length; i1 < e1; i1++ {

			err = (*rsp)[i1].ReadBlock(_is, 0, false)
			if err != nil {
				return ret, err
			}

		}
	} else if ty == codec.SIMPLE_LIST {
		err = fmt.Errorf("not support simple_list type")
		if err != nil {
			return ret, err
		}

	} else {
		err = fmt.Errorf("require vector, but not")
		if err != nil {
			return ret, err
		}

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

//GetAccessPointsOneWayWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *AccessPointOperate) GetAccessPointsOneWayWithContext(tarsCtx context.Context, req *GetAccessPointsReq, rsp *[]AccessPointInfo, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = req.WriteBlock(_os, 1)
	if err != nil {
		return ret, err
	}

	err = _os.WriteHead(codec.LIST, 2)
	if err != nil {
		return ret, err
	}

	err = _os.Write_int32(int32(len((*rsp))), 0)
	if err != nil {
		return ret, err
	}

	for _, v := range *rsp {

		err = v.WriteBlock(_os, 0)
		if err != nil {
			return ret, err
		}

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

	err = _obj.s.Tars_invoke(tarsCtx, 1, "GetAccessPoints", _os.ToBytes(), _status, _context, _resp)
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

//UpdateAccessPoint is the proxy function for the method defined in the tars file, with the context
func (_obj *AccessPointOperate) UpdateAccessPoint(req *AccessPointInfo, rsp *string, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = req.WriteBlock(_os, 1)
	if err != nil {
		return ret, err
	}

	err = _os.Write_string((*rsp), 2)
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

	err = _obj.s.Tars_invoke(tarsCtx, 0, "UpdateAccessPoint", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}

	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err = _is.Read_string(&(*rsp), 2, true)
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

//UpdateAccessPointWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *AccessPointOperate) UpdateAccessPointWithContext(tarsCtx context.Context, req *AccessPointInfo, rsp *string, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = req.WriteBlock(_os, 1)
	if err != nil {
		return ret, err
	}

	err = _os.Write_string((*rsp), 2)
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

	err = _obj.s.Tars_invoke(tarsCtx, 0, "UpdateAccessPoint", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}

	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err = _is.Read_string(&(*rsp), 2, true)
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

//UpdateAccessPointOneWayWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *AccessPointOperate) UpdateAccessPointOneWayWithContext(tarsCtx context.Context, req *AccessPointInfo, rsp *string, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = req.WriteBlock(_os, 1)
	if err != nil {
		return ret, err
	}

	err = _os.Write_string((*rsp), 2)
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

	err = _obj.s.Tars_invoke(tarsCtx, 1, "UpdateAccessPoint", _os.ToBytes(), _status, _context, _resp)
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

//DeleteAccessPoint is the proxy function for the method defined in the tars file, with the context
func (_obj *AccessPointOperate) DeleteAccessPoint(req *AccessPointInfo, rsp *string, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = req.WriteBlock(_os, 1)
	if err != nil {
		return ret, err
	}

	err = _os.Write_string((*rsp), 2)
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

	err = _obj.s.Tars_invoke(tarsCtx, 0, "DeleteAccessPoint", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}

	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err = _is.Read_string(&(*rsp), 2, true)
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

//DeleteAccessPointWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *AccessPointOperate) DeleteAccessPointWithContext(tarsCtx context.Context, req *AccessPointInfo, rsp *string, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = req.WriteBlock(_os, 1)
	if err != nil {
		return ret, err
	}

	err = _os.Write_string((*rsp), 2)
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

	err = _obj.s.Tars_invoke(tarsCtx, 0, "DeleteAccessPoint", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}

	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err = _is.Read_string(&(*rsp), 2, true)
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

//DeleteAccessPointOneWayWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *AccessPointOperate) DeleteAccessPointOneWayWithContext(tarsCtx context.Context, req *AccessPointInfo, rsp *string, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = req.WriteBlock(_os, 1)
	if err != nil {
		return ret, err
	}

	err = _os.Write_string((*rsp), 2)
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

	err = _obj.s.Tars_invoke(tarsCtx, 1, "DeleteAccessPoint", _os.ToBytes(), _status, _context, _resp)
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
func (_obj *AccessPointOperate) SetServant(s m.Servant) {
	_obj.s = s
}

//TarsSetTimeout sets the timeout for the servant which is in ms.
func (_obj *AccessPointOperate) TarsSetTimeout(t int) {
	_obj.s.TarsSetTimeout(t)
}

//TarsSetProtocol sets the protocol for the servant.
func (_obj *AccessPointOperate) TarsSetProtocol(p m.Protocol) {
	_obj.s.TarsSetProtocol(p)
}

//AddServant adds servant  for the service.
func (_obj *AccessPointOperate) AddServant(imp _impAccessPointOperate, obj string) {
	tars.AddServant(_obj, imp, obj)
}

//AddServant adds servant  for the service with context.
func (_obj *AccessPointOperate) AddServantWithContext(imp _impAccessPointOperateWithContext, obj string) {
	tars.AddServantWithContext(_obj, imp, obj)
}

type _impAccessPointOperate interface {
	CreateAccessPoint(req *CreateAccessPointReq, rsp *string) (ret int32, err error)
	GetAccessPoints(req *GetAccessPointsReq, rsp *[]AccessPointInfo) (ret int32, err error)
	UpdateAccessPoint(req *AccessPointInfo, rsp *string) (ret int32, err error)
	DeleteAccessPoint(req *AccessPointInfo, rsp *string) (ret int32, err error)
}
type _impAccessPointOperateWithContext interface {
	CreateAccessPoint(tarsCtx context.Context, req *CreateAccessPointReq, rsp *string) (ret int32, err error)
	GetAccessPoints(tarsCtx context.Context, req *GetAccessPointsReq, rsp *[]AccessPointInfo) (ret int32, err error)
	UpdateAccessPoint(tarsCtx context.Context, req *AccessPointInfo, rsp *string) (ret int32, err error)
	DeleteAccessPoint(tarsCtx context.Context, req *AccessPointInfo, rsp *string) (ret int32, err error)
}

// Dispatch is used to call the server side implemnet for the method defined in the tars file. _withContext shows using context or not.
func (_obj *AccessPointOperate) Dispatch(tarsCtx context.Context, _val interface{}, tarsReq *requestf.RequestPacket, tarsResp *requestf.ResponsePacket, _withContext bool) (err error) {
	var length int32
	var have bool
	var ty byte
	_is := codec.NewReader(tools.Int8ToByte(tarsReq.SBuffer))
	_os := codec.NewBuffer()
	switch tarsReq.SFuncName {
	case "CreateAccessPoint":
		var req CreateAccessPointReq
		var rsp string

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
			_imp := _val.(_impAccessPointOperate)
			_funRet_, err = _imp.CreateAccessPoint(&req, &rsp)
		} else {
			_imp := _val.(_impAccessPointOperateWithContext)
			_funRet_, err = _imp.CreateAccessPoint(tarsCtx, &req, &rsp)
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

			err = _os.Write_string(rsp, 2)
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
			err = _os.Write_string(rsp, 0)
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
	case "GetAccessPoints":
		var req GetAccessPointsReq
		var rsp []AccessPointInfo

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
			_imp := _val.(_impAccessPointOperate)
			_funRet_, err = _imp.GetAccessPoints(&req, &rsp)
		} else {
			_imp := _val.(_impAccessPointOperateWithContext)
			_funRet_, err = _imp.GetAccessPoints(tarsCtx, &req, &rsp)
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

			err = _os.WriteHead(codec.LIST, 2)
			if err != nil {
				return err
			}

			err = _os.Write_int32(int32(len(rsp)), 0)
			if err != nil {
				return err
			}

			for _, v := range rsp {

				err = v.WriteBlock(_os, 0)
				if err != nil {
					return err
				}

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
			err = _os.WriteHead(codec.LIST, 0)
			if err != nil {
				return err
			}

			err = _os.Write_int32(int32(len(rsp)), 0)
			if err != nil {
				return err
			}

			for _, v := range rsp {

				err = v.WriteBlock(_os, 0)
				if err != nil {
					return err
				}

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
	case "UpdateAccessPoint":
		var req AccessPointInfo
		var rsp string

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
			_imp := _val.(_impAccessPointOperate)
			_funRet_, err = _imp.UpdateAccessPoint(&req, &rsp)
		} else {
			_imp := _val.(_impAccessPointOperateWithContext)
			_funRet_, err = _imp.UpdateAccessPoint(tarsCtx, &req, &rsp)
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

			err = _os.Write_string(rsp, 2)
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
			err = _os.Write_string(rsp, 0)
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
	case "DeleteAccessPoint":
		var req AccessPointInfo
		var rsp string

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
			_imp := _val.(_impAccessPointOperate)
			_funRet_, err = _imp.DeleteAccessPoint(&req, &rsp)
		} else {
			_imp := _val.(_impAccessPointOperateWithContext)
			_funRet_, err = _imp.DeleteAccessPoint(tarsCtx, &req, &rsp)
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

			err = _os.Write_string(rsp, 2)
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
			err = _os.Write_string(rsp, 0)
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
