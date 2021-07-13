package logic

import "github.com/tars-vcms/rbac-server/tars-protocol/vcms"

type AccessPointOperateImpl struct {}

func (a AccessPointOperateImpl) CreateAccessPoint(req *vcms.CreateAccessPointReq, rsp *vcms.AccessPointInfo) (ret int32, err error) {
	panic("implement me")
}

func (a AccessPointOperateImpl) GetAccessPoints(req *vcms.GetAccessPointsReq, rsp *[]vcms.AccessPointInfo) (ret int32, err error) {
	panic("implement me")
}

func (a AccessPointOperateImpl) UpdateAccessPoint(req *vcms.AccessPointInfo, rsp *vcms.AccessPointInfo) (ret int32, err error) {
	panic("implement me")
}

func (a AccessPointOperateImpl) DeleteAccessPoint(req *vcms.AccessPointInfo, rsp *vcms.AccessPointInfo) (ret int32, err error) {
	panic("implement me")
}
