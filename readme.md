# RBAC-Server
本服务用于实施对访问点以及访问角色的管理以及访问鉴权。

## 表设计
### 接入点表 (AccessPoint)
| 键名 | 类型 | 说明 |
| :---: | :---: | :---: |
| id  | int | PK、自增 |
| ident | varchar | 接入点名 |
| comment | varchar | 备注 | 

### 用户组表 (AccessRole)
| 键名 | 类型 | 说明 |
| :---: | :---: | :---: |
| id  | int | PK、自增 |
| ident | varchar | 用户组标识，唯一 |
| name | varchar | 用户组名，建议唯一 |
| access | text | 接入点id 逗分 |
| comment | varchar | 备注 | 
