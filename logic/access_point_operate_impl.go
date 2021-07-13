package logic

import (
	"fmt"
	"github.com/TarsCloud/TarsGo/tars/protocol/res/basef"
	"github.com/tars-vcms/rbac-server/repo/mysql"
	"github.com/tars-vcms/rbac-server/tars-protocol/vcms"
)

type AccessPointOperateImpl struct {}

func (a AccessPointOperateImpl) CreateAccessPoint(req *vcms.CreateAccessPointReq, rsp *vcms.AccessPointInfo) (ret int32, err error) {
	thisRole, rowCount, err := mysql.AccessPointSqlImpl.Create(req.AccessIdent, req.Comment)
	if rowCount != 1{
		rsp.AccessId = int32(thisRole.ID)
		rsp.AccessIdent = thisRole.Ident
		rsp.Comment = thisRole.Comment
		return basef.TARSSERVERSUCCESS, nil
	}else{
		return basef.TARSSERVERSUCCESS, err
	}
}

func (a AccessPointOperateImpl) GetAccessPoints(req *vcms.GetAccessPointsReq, rsp *[]vcms.AccessPointInfo) (ret int32, err error) {
	whereConditionString := fmt.Sprintf("`id` = %d AND `ident` LIKE %%%s%% AND `comment` LIKE %%%s%%", req.WhereCondition.AccessId, req.WhereCondition.AccessIdent, req.WhereCondition.Comment)
	thisRoles, _, err := mysql.AccessPointSqlImpl.Read(whereConditionString, int(req.Limit.PageSize), int(req.Limit.Offset))
	if err != nil {
		for _, v := range thisRoles {
			*rsp = append(*rsp, vcms.AccessPointInfo{
				AccessId:    int32(v.ID),
				AccessIdent: v.Ident,
				Comment:   v.Comment,
			})
		}
		return basef.TARSSERVERSUCCESS, nil
	}else{
		return basef.TARSSERVERSUCCESS, err
	}
}

func (a AccessPointOperateImpl) UpdateAccessPoint(req *vcms.AccessPointInfo, rsp *vcms.AccessPointInfo) (ret int32, err error) {
	thisRole, rowCount, err := mysql.AccessPointSqlImpl.Update(int(req.AccessId), req.AccessIdent, req.Comment)
	if rowCount != 1{
		rsp.AccessId = int32(thisRole.ID)
		rsp.AccessIdent = thisRole.Ident
		rsp.Comment = thisRole.Comment
		return basef.TARSSERVERSUCCESS, nil
	}else{
		return basef.TARSSERVERSUCCESS, err
	}
}

func (a AccessPointOperateImpl) DeleteAccessPoint(req *vcms.AccessPointInfo, rsp *vcms.AccessPointInfo) (ret int32, err error) {
	thisRole, rowCount, err := mysql.AccessPointSqlImpl.Delete(int(req.AccessId))
	if rowCount != 1{
		rsp.AccessId = int32(thisRole.ID)
		rsp.AccessIdent = thisRole.Ident
		rsp.Comment = thisRole.Comment
		return basef.TARSSERVERSUCCESS, nil
	}else{
		return basef.TARSSERVERSUCCESS, err
	}
}
