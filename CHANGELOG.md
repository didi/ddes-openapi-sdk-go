# 变更日志

本文记录当前分支相对 `master` 的变更。

## [1.3.0] - 2026-09-09

### 破坏性变更

以下为相对 `master` 仍保留的公开类型变更，共 3 处（2 个响应模型字段类型 + 1 组字段/Builder 方法更名）。已有调用方如果直接引用，需要同步调整代码：

| 接口/模型 | 变更 | 说明 |
| --- | --- | --- |
| `/river/Rank/getRanks` — `ListRankReply` | `Records`：`[]RankInfo` → `[]RankRecord` | 修复性变更：服务端响应字段为 `id`，原复用请求侧 `RankInfo`（`rank_id`）导致职级 ID 永远解析为空；`RankRecord` 字段为 `Id`/`Name`/`OutRankId`。请求侧 `RankInfo`（create/update/del）不受影响 |
| `/river/Member/edit` — `UpdateMemberApiReply` | `Data`：`interface{}` → `*UpdateMemberReply` | 修复性变更：原 `UpdateMemberReply` 为空结构体、Data 无类型语义；现补齐 `MemberId`/`Phone`/`Status`（员工信息更新失败场景返回），对 Data 做类型断言的调用方需同步调整 |
| `/river/Member/single`、`/river/Member/edit` 请求 — `MemberInfo` | 字段及 Builder 方法更名：`RedisentsIds` → `ResidentsIds`，json key `redisents_ids` → `residents_ids` | 修复性变更：原拼写 `redisents_ids` 服务端不识别，该参数此前实际不生效；引用旧字段 `RedisentsIds` 或调用 `MemberInfoBuilder.RedisentsIds()` 的代码需改为 `ResidentsIds`/`.ResidentsIds()`，序列化 key 同步变更 |

### 新增

- 集团账户参数支持（`belong_enterprise_name`/`taxpayer_no`，任选其一），覆盖 11 个接口、13 处落点：
  - query 参数：`/river/Rank/getRanks`、`/river/Member/getQuota`、`/river/Regulation/detail`、`/river/Role/get`；
  - 请求体顶层：`/river/Member/single`、`/river/Member/edit`、`/river/Member/del`；
  - param_json 内：`/open-apis/v1/rank/create|update|del`、`/river/LegalEntity/add|edit|del`。
- `/river/LegalEntity/get`：请求新增 `out_legal_entity_id` 精准查询；响应新增 `out_parent_id`（上级公司外部编号，2025.06.18 上线）。
- `/river/Member/getQuota`：请求新增 `monthly_limit_only_query`（1 仅按自然月（默认）/ 2 查询所有限额）。
- `/river/Member/single`、`/river/Member/edit`：data 新增 `residents_adcode`（常驻地行政区划代码）、`home_address`（家庭住址）。
- `/river/Regulation/get`、`/river/Regulation/detail`：响应新增 `regulation_employee_name`、`regulation_employee_description`。
- `/open-apis/v1/traveler/create|update|del`：`TravelerInfo`（param_json 内）新增 `type`（外部出行人类型：0-公司维护（默认）、3-项目维度）及集团账户参数 `belong_enterprise_name`/`taxpayer_no`（与 Java SDK 落点一致），修复集团账户场景无法传参与项目维度出行人无法创建的问题。
- `/river/Regulation/detail`：响应新增 `car_rule`（用车规则，`CarRuleInfo{call_type, use_scene}`，文档 11315）。
- `/river/Member/single`：响应新增 `member_id`、`status`（添加员工失败-员工已存在场景返回；status：1-正常、4-离职、6-未绑定手机号）。
- `/river/Member/edit`：响应补齐 `member_id`/`phone`/`status`（员工信息更新失败-员工已离职或被删除场景返回），`UpdateMemberApiReply.Data` 类型见破坏性变更表。
- 上述 `member_id` 字段真实流量可能为 number 或 string，已按 SDK 惯例实现 `UnmarshalJSON` 容错（number/string 统一转 `*string`）。
- `/river/Member/del`：新增 `employee_phone`、`employee_email` 删除标识（与 `member_id`/`employee_number` 四选一，优先级：滴滴侧ID > 手机号 > 工号 > 邮箱）。
- 新增集团账户（上级公司统一对接）v2 查询接口支持，4 个域各新增一个 `*V2` 方法（GET `/open-apis/v2/<domain>/get`），现有 `/river/*` 方法不受影响：
  - `MemberService.V1.Member.ListMemberV2`：员工列表，支持 `is_last_id` 游标分页；
  - `LegalEntityService.V1.LegalEntity.GetLegalEntityV2`：公司主体查询，支持 `keyword`/`legal_entity_id`/`out_legal_entity_id`/`status` 精准筛选；
  - `BudgetCenterService.V1.BudgetCenter.GetBudgetCenterV2`：部门/项目查询，支持 `third_out_id`（v2 契约无 `is_get_poi`/`is_get_extend_fields`）；
  - `RegulationService.V1.Regulation.ListRegulationV2`：制度列表（v2 为分页结构，`/river` 版无分页）。
- 四个 V2 方法均支持集团参数 `belong_enterprise_name`/`taxpayer_no`（定位子公司）、`query_sub_company`（聚合下级公司）、`next_token` 跨公司分页；响应外层含 `next_token`/`records`/`progress`/`company_total`/`company_list_has_change`，records 元素按 openapi-go IDL 契约新建独立模型（`MemberV2Record`、`LegalEntityV2Record`（`legal_entity_id` 为 string，不丢大数精度）、`BudgetCenterV2Record`（新增 `company_id`/`create_time`/`update_time`）、`RegulationV2Record`）。
- `sample/member/list_member_v2.go` 新增 next_token 自动翻页循环示例。

### 修复

- 修复 `/river/Rank/getRanks` 响应职级 ID 解析为空的问题（见破坏性变更）。
- 修复 `/river/Member/single`、`/river/Member/edit` 常驻地 ID 不生效的问题：`MemberInfo` 字段 `RedisentsIds` 更名为 `ResidentsIds`，json key 由 `redisents_ids` 修正为服务端实际读取的 `residents_ids`（见破坏性变更）。

### 测试与工程化

- 为上述全部新增字段与修复补齐单元测试（模型序列化、query 写入、响应反序列化、Builder、多标识删除场景）。
- 四域 V2 全量单元测试（builder 参数矩阵、分页响应解析、大 ID 精度、多余字段忽略、httptest 资源层含 AES128/AES256 加密分支）。
- 新增 `docs/接口对齐diff清单.md`（SDK ↔ PHP openapi ↔ openapi-go ↔ 官方文档四方对照证据链）。
- SDK 版本号由 `1.2.0` 调整为 `1.3.0`。

## [1.2.0] - 2026-08-03

### 破坏性变更

以下为相对 `master` 仍保留的公开类型变更，共涉及 2 个模型、2 个字段。已有调用方如果直接赋值，需要同步调整代码：

| 接口/模型 | 字段类型变更 | Builder 入参变更 |
| --- | --- | --- |
| `/river/Bill/detail` — `GetBillDetailOfWangYCReply` | `LastId`：`*int64` → `*string` | — |
| `/river/Member/del` — `DelMemberApiReply` | `Data`：`[]int64` → `[]string` | — |

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
- 恢复审批订单、网约车账单及敏感订单字段的历史公开数据类型，同时保留 number/string 混合响应的容错反序列化能力。

### 测试与工程化

- 为多个服务补充模型、请求构造、JSON 序列化和响应反序列化单元测试。
- 新增基于真实响应样本的回放测试能力，并忽略本地生成的、可能包含敏感信息的 fixture 和构建产物。
- SDK 版本号由 `1.0.0` 调整为 `1.2.0`。
