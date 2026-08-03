# 变更日志

本文记录当前分支相对 `master` 的变更。

## [1.2.0] - 2026-08-03

### 破坏性变更

以下公开模型字段及对应 Builder 方法的参数类型相对 `master` 发生变化。已有调用方如果直接赋值或调用 Builder，需要同步调整代码：

- `service/approval/v1`：`ApprovalOrderRecord.CarLevel`、`OrderStatus`、`PayType`、`IsInvoice` 由 `*int32` 改为 `*string`；对应 Builder 参数由 `int32` 改为 `string`。
- `service/approval/v1`：`TravelCity.Id` 由 `*int32` 改为 `*string`；`TravelCityBuilder.Id` 参数由 `int32` 改为 `string`。
- `service/bill/v1`：`BillListItemOfWangYC.IsSensitive` 由 `*float32` 改为 `*string`，`MemberId` 由 `*int64` 改为 `*string`，`PersonalRealPay` 由 `*float32` 改为 `*string`；对应 Builder 参数同步改为 `string`。
- `service/bill/v1`：`NotGenBDOfWangYCItem.CompanyRealPay` 由 `*float64` 改为 `*string`，`IsSensitive` 由 `*int32` 改为 `*string`；对应 Builder 参数同步改为 `string`。
- `service/bill/v1`：`GetBillDetailOfWangYCReply.LastId` 由 `*int64` 改为 `*string`。
- `service/city/v1`：`HotelCityInfo.CityId`、`TrainCityInfo.CityId` 由 `*string` 改为 `*int64`；对应 Builder 参数同步改为 `int64`。
- `service/city/v1`：`ListAirportCityReply.CityId`、`CountryId` 以及 `ListCountryReply.CountryId` 由 `*int32` 改为 `*string`。
- `service/member/v1`：`DelMemberApiReply.Data` 由 `[]int64` 改为 `[]string`。
- `service/order/v1`：`OrderRecord.RegulationId` 由 `*int64` 改为 `*string`；对应 Builder 参数由 `int64` 改为 `string`。

### 新增

- 新增项目服务 `ProjectService`，支持：
  - 查询项目成员；
  - 查询项目外部出行人信息（会议场景）；
  - 绑定项目与人员关系；
  - 解绑项目与人员关系。
- 扩展员工列表（批量查询）和员工明细接口模型，补充家庭住址、限额规则、常住地等字段。
- 扩展部门/项目查询、新增/编辑、删除接口模型，补充扩展字段、限额规则、POI、外部出行人和关联员工等数据结构。
- 新增 `core.RawMessageToString` 和 `core.SmartDecode`，用于兼容真实流量中数字、字符串、空值混用的字段类型。

### 修复

- 修复审批、账单、预算中心、城市、员工、订单、规章等接口模型的数据类型不匹配问题，提升响应反序列化兼容性。
- 修复部分 ID、订单号、枚举及金额相关字段在 number/string 混合返回时的解析失败问题。

### 测试与工程化

- 为多个服务补充模型、请求构造、JSON 序列化和响应反序列化单元测试。
- 新增基于真实响应样本的回放测试能力，并忽略本地生成的、可能包含敏感信息的 fixture 和构建产物。
- SDK 版本号由 `1.0.0` 调整为 `1.2.0`。
