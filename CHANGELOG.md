# 变更日志

本文记录当前分支相对 `master` 的变更。

## [1.2.0] - 2026-08-03

### 破坏性变更

以下为相对 `master` 的公开类型变更，共涉及 6 个模型、11 个字段。已有调用方如果直接赋值或调用 Builder，需要同步调整代码：

| 接口/模型 | 字段类型变更 | Builder 入参变更 |
| --- | --- | --- |
| `/river/Approval/getOrder` — `ApprovalOrderRecord` | `CarLevel`、`OrderStatus`、`PayType`、`IsInvoice`：`*int32` → `*string` | 对应参数：`int32` → `string` |
| `/river/Bill/get` — `BillListItemOfWangYC` | `MemberId`：`*int64` → `*string`；`PersonalRealPay`：`*float32` → `*string` | `MemberId`：`int64` → `string`；`PersonalRealPay`：`float32` → `string` |
| `/river/Bill/getNotGeneratedBillDetail` — `NotGenBDOfWangYCItem` | `CompanyRealPay`：`*float64` → `*string`；`IsSensitive`：`*int32` → `*string` | 对应参数：`float64/int32` → `string` |
| `/river/Bill/detail` — `GetBillDetailOfWangYCReply` | `LastId`：`*int64` → `*string` | — |
| `/river/Member/del` — `DelMemberApiReply` | `Data`：`[]int64` → `[]string` | — |
| 订单查询/详情 — `OrderRecord` | `RegulationId`：`*int64` → `*string` | `OrderRecordBuilder.RegulationId`：`int64` → `string` |

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
