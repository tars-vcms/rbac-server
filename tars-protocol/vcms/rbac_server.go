// Package vcms comment
// This file was generated by tars2go 1.1.4
// Generated from rbac_server.tars
package vcms

import (
	"fmt"

	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = codec.FromInt8

// AccessAbilityRsp struct implement
type AccessAbilityRsp struct {
	Ident      string `json:"Ident"`
	Name       string `json:"Name"`
	AccessList string `json:"AccessList"`
	Comment    string `json:"Comment"`
}

func (st *AccessAbilityRsp) ResetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *AccessAbilityRsp) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.ResetDefault()

	err = _is.Read_string(&st.Ident, 0, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Name, 1, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.AccessList, 2, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Comment, 3, true)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *AccessAbilityRsp) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.ResetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require AccessAbilityRsp, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(_is)
	if err != nil {
		return err
	}

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *AccessAbilityRsp) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_string(st.Ident, 0)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Name, 1)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.AccessList, 2)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Comment, 3)
	if err != nil {
		return err
	}

	_ = err

	return nil
}

//WriteBlock encode struct
func (st *AccessAbilityRsp) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(_os)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}

// IsValidAccessReq struct implement
type IsValidAccessReq struct {
	AccessIdent string `json:"AccessIdent"`
	RoleIdent   string `json:"RoleIdent"`
}

func (st *IsValidAccessReq) ResetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *IsValidAccessReq) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.ResetDefault()

	err = _is.Read_string(&st.AccessIdent, 0, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.RoleIdent, 1, true)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *IsValidAccessReq) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.ResetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require IsValidAccessReq, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(_is)
	if err != nil {
		return err
	}

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *IsValidAccessReq) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_string(st.AccessIdent, 0)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.RoleIdent, 1)
	if err != nil {
		return err
	}

	_ = err

	return nil
}

//WriteBlock encode struct
func (st *IsValidAccessReq) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(_os)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}

// AccessPointInfo struct implement
type AccessPointInfo struct {
	AccessId    int32  `json:"AccessId"`
	AccessIdent string `json:"AccessIdent"`
	Comment     string `json:"Comment"`
}

func (st *AccessPointInfo) ResetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *AccessPointInfo) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.ResetDefault()

	err = _is.Read_int32(&st.AccessId, 0, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.AccessIdent, 1, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Comment, 2, true)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *AccessPointInfo) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.ResetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require AccessPointInfo, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(_is)
	if err != nil {
		return err
	}

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *AccessPointInfo) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_int32(st.AccessId, 0)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.AccessIdent, 1)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Comment, 2)
	if err != nil {
		return err
	}

	_ = err

	return nil
}

//WriteBlock encode struct
func (st *AccessPointInfo) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(_os)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}

// CreateAccessPointReq struct implement
type CreateAccessPointReq struct {
	AccessIdent string `json:"AccessIdent"`
	Comment     string `json:"Comment"`
}

func (st *CreateAccessPointReq) ResetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *CreateAccessPointReq) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.ResetDefault()

	err = _is.Read_string(&st.AccessIdent, 0, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Comment, 1, true)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *CreateAccessPointReq) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.ResetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require CreateAccessPointReq, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(_is)
	if err != nil {
		return err
	}

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *CreateAccessPointReq) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_string(st.AccessIdent, 0)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Comment, 1)
	if err != nil {
		return err
	}

	_ = err

	return nil
}

//WriteBlock encode struct
func (st *CreateAccessPointReq) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(_os)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}

// PageLimiter struct implement
type PageLimiter struct {
	PageSize int32 `json:"PageSize"`
	Offset   int32 `json:"Offset"`
}

func (st *PageLimiter) ResetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *PageLimiter) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.ResetDefault()

	err = _is.Read_int32(&st.PageSize, 0, true)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&st.Offset, 1, true)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *PageLimiter) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.ResetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require PageLimiter, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(_is)
	if err != nil {
		return err
	}

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *PageLimiter) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_int32(st.PageSize, 0)
	if err != nil {
		return err
	}

	err = _os.Write_int32(st.Offset, 1)
	if err != nil {
		return err
	}

	_ = err

	return nil
}

//WriteBlock encode struct
func (st *PageLimiter) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(_os)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}

// AccessPointWhereCondition struct implement
type AccessPointWhereCondition struct {
	AccessId    int32  `json:"AccessId"`
	AccessIdent string `json:"AccessIdent"`
	Comment     string `json:"Comment"`
}

func (st *AccessPointWhereCondition) ResetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *AccessPointWhereCondition) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.ResetDefault()

	err = _is.Read_int32(&st.AccessId, 0, false)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.AccessIdent, 1, false)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Comment, 2, true)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *AccessPointWhereCondition) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.ResetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require AccessPointWhereCondition, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(_is)
	if err != nil {
		return err
	}

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *AccessPointWhereCondition) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_int32(st.AccessId, 0)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.AccessIdent, 1)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Comment, 2)
	if err != nil {
		return err
	}

	_ = err

	return nil
}

//WriteBlock encode struct
func (st *AccessPointWhereCondition) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(_os)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}

// GetAccessPointsReq struct implement
type GetAccessPointsReq struct {
	WhereCondition AccessPointWhereCondition `json:"WhereCondition"`
	Limit          PageLimiter               `json:"Limit"`
}

func (st *GetAccessPointsReq) ResetDefault() {
	st.WhereCondition.ResetDefault()
	st.Limit.ResetDefault()
}

//ReadFrom reads  from _is and put into struct.
func (st *GetAccessPointsReq) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.ResetDefault()

	err = st.WhereCondition.ReadBlock(_is, 0, true)
	if err != nil {
		return err
	}

	err = st.Limit.ReadBlock(_is, 1, true)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *GetAccessPointsReq) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.ResetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require GetAccessPointsReq, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(_is)
	if err != nil {
		return err
	}

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *GetAccessPointsReq) WriteTo(_os *codec.Buffer) error {
	var err error

	err = st.WhereCondition.WriteBlock(_os, 0)
	if err != nil {
		return err
	}

	err = st.Limit.WriteBlock(_os, 1)
	if err != nil {
		return err
	}

	_ = err

	return nil
}

//WriteBlock encode struct
func (st *GetAccessPointsReq) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(_os)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}

// CreateAccessRoleReq struct implement
type CreateAccessRoleReq struct {
	RoleIdent string `json:"RoleIdent"`
	RoleName  string `json:"RoleName"`
	Access    string `json:"Access"`
	Comment   string `json:"Comment"`
}

func (st *CreateAccessRoleReq) ResetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *CreateAccessRoleReq) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.ResetDefault()

	err = _is.Read_string(&st.RoleIdent, 0, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.RoleName, 1, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Access, 2, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Comment, 3, true)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *CreateAccessRoleReq) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.ResetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require CreateAccessRoleReq, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(_is)
	if err != nil {
		return err
	}

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *CreateAccessRoleReq) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_string(st.RoleIdent, 0)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.RoleName, 1)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Access, 2)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Comment, 3)
	if err != nil {
		return err
	}

	_ = err

	return nil
}

//WriteBlock encode struct
func (st *CreateAccessRoleReq) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(_os)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}

// AccessRoleInfo struct implement
type AccessRoleInfo struct {
	RoleId    int32  `json:"RoleId"`
	RoleIdent string `json:"RoleIdent"`
	RoleName  string `json:"RoleName"`
	Access    string `json:"Access"`
	Comment   string `json:"Comment"`
}

func (st *AccessRoleInfo) ResetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *AccessRoleInfo) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.ResetDefault()

	err = _is.Read_int32(&st.RoleId, 0, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.RoleIdent, 1, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.RoleName, 2, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Access, 3, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Comment, 4, true)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *AccessRoleInfo) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.ResetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require AccessRoleInfo, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(_is)
	if err != nil {
		return err
	}

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *AccessRoleInfo) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_int32(st.RoleId, 0)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.RoleIdent, 1)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.RoleName, 2)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Access, 3)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Comment, 4)
	if err != nil {
		return err
	}

	_ = err

	return nil
}

//WriteBlock encode struct
func (st *AccessRoleInfo) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(_os)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}

// AccessRoleWhereCondition struct implement
type AccessRoleWhereCondition struct {
	RoleId    int32  `json:"RoleId"`
	RoleIdent string `json:"RoleIdent"`
	RoleName  string `json:"RoleName"`
	Access    string `json:"Access"`
	Comment   string `json:"Comment"`
}

func (st *AccessRoleWhereCondition) ResetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *AccessRoleWhereCondition) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.ResetDefault()

	err = _is.Read_int32(&st.RoleId, 0, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.RoleIdent, 1, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.RoleName, 2, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Access, 3, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Comment, 4, true)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *AccessRoleWhereCondition) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.ResetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require AccessRoleWhereCondition, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(_is)
	if err != nil {
		return err
	}

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *AccessRoleWhereCondition) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_int32(st.RoleId, 0)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.RoleIdent, 1)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.RoleName, 2)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Access, 3)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Comment, 4)
	if err != nil {
		return err
	}

	_ = err

	return nil
}

//WriteBlock encode struct
func (st *AccessRoleWhereCondition) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(_os)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}

// GetAccessRoleReq struct implement
type GetAccessRoleReq struct {
	WhereCondition AccessRoleWhereCondition `json:"WhereCondition"`
	Limit          PageLimiter              `json:"Limit"`
}

func (st *GetAccessRoleReq) ResetDefault() {
	st.WhereCondition.ResetDefault()
	st.Limit.ResetDefault()
}

//ReadFrom reads  from _is and put into struct.
func (st *GetAccessRoleReq) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.ResetDefault()

	err = st.WhereCondition.ReadBlock(_is, 0, true)
	if err != nil {
		return err
	}

	err = st.Limit.ReadBlock(_is, 1, true)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *GetAccessRoleReq) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.ResetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require GetAccessRoleReq, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(_is)
	if err != nil {
		return err
	}

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *GetAccessRoleReq) WriteTo(_os *codec.Buffer) error {
	var err error

	err = st.WhereCondition.WriteBlock(_os, 0)
	if err != nil {
		return err
	}

	err = st.Limit.WriteBlock(_os, 1)
	if err != nil {
		return err
	}

	_ = err

	return nil
}

//WriteBlock encode struct
func (st *GetAccessRoleReq) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(_os)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}
