# SDK 接口对齐 Diff 清单

> 基准：`docs/接口清单.md` 32 个条目（28 个主接口 + 4 个集团账户 v2 备用路径），SDK（ddes-openapi-sdk-go）↔ PHP openapi（服务端老实现）↔ openapi-go（服务端新实现）↔ 官方文档（opendocs.xiaojukeji.com）四方对照。
>
> 生成日期：2026-09-09（v2，按新清单纳入集团账户 v2 路径）。证据标注：`文档:<id>` 官方文档、`PHP:<file>:<line>` 服务端代码、`Go:<file>` openapi-go 实现。

## 0. 总览

| 分类 | 数量 | 处置 |
|------|-----|------|
| P0 真 bug（字段名错误，功能失效） | 2 | 必须修 |
| P1 集团账户参数缺失（老路径，模式 B） | 11 个接口 13 处 | 增量补 |
| P1.5 集团账户 v2 路径（模式 A，4 个域） | 4 个域 | 新增 v2 方法+模型 |
| P2 文档新增字段缺失 | 6 项 | 增量补 |
| P3 服务端返回但文档未公开 | 2 个接口 7 个字段 | 建议补（需决策） |
| P4 功能性缺失 | 1 个接口 | 增量补 |
| 已对齐 / 记录不修 | — | 见 §7/§8 |

## 1. 服务端归属与集团账户模式

```
32 条目归属
├─ PHP openapi 承接（/river/*）21 个主接口
│    权威字段来源: controllers/<域>/<Action>.php + rules/ + models/
├─ openapi-go 承接 7 个主接口
│    ├─ traveler create/update/del（原生路由, OpenapiReq+param_json）
│    ├─ rank create/update/del（gate 转发 → leto 服务）
│    └─ v2/project/outTravelerList（原生路由, idl/traveler.idl）
└─ 集团账户 v2 备用路径 4 个（均由 openapi-go 原生路由承接）
     /open-apis/v2/legalEntity/get    Go idl/legal-entity.idl
     /open-apis/v2/budgetCenter/get   Go idl/budget-center.idl
     /open-apis/v2/member/get         Go idl/member.idl
     /open-apis/v2/regulation/get     Go idl/regulation.idl
```

**集团账户两种模式**（文档模板化表述，与 openapi-go 原生 v2 路由完全对应）：

| 模式 | 表述 | 适用接口 | SDK 对齐方式 |
|------|------|---------|-------------|
| **A：切 v2 路径** | HTTP URL → `/open-apis/v2/<domain>/get`，新增 belong_enterprise_name / taxpayer_no / query_sub_company / next_token；分页改 next_token | LegalEntity/get、Member/get、BudgetCenter/get、Regulation/get | **新增 v2 资源方法 + v2 ApiReq/Reply**（见 §3.5） |
| **B：老路径直加** | 路径不变，新增 belong_enterprise_name + taxpayer_no | 其余全部集团接口 | 现有 builder/请求体增量补参数（见 §3） |

> 架构事实：SDK `core.ReqOption` 不支持 ApiPath 覆盖（路径硬编码于 resource 方法），且 v2 响应结构不同（data 含 next_token/progress/company_total/company_list_has_change），故模式 A 必须新增独立方法与模型，不能在现有方法上换路径。

## 2. P0 — 真 bug（字段名错误，功能失效）

### 2.1 Rank/getRanks 响应 `rank_id` tag 错误

- **位置**：`service/rank/v1/model_rank_info.go:5` — `RankId *string json:"rank_id"`
- **正确值**：`id`
- **证据**：文档 11253 响应字段 `data.records.id`；PHP `models/Rank.php:65` `'id' => strval(...)`
- **影响**：`ListRankApiReply.Records` 反序列化时职级滴滴侧 ID 永远为空
- **修法**：按「禁止复用类型」规则，为 getRanks 响应新建独立 `RankRecord`（`id`/`name`/`out_rank_id`）+ Builder；请求侧 `RankInfo`（create/update/del 用）不动

### 2.2 Member single/edit 常驻地 `redisents_ids` tag 拼写错误

- **位置**：`service/member/v1/model_member_info.go:18` — `RedisentsIds *string json:"redisents_ids"`
- **正确值**：`residents_ids`
- **证据**：文档 11155 data 字段 `residents_ids`；PHP `controllers/Member/Single.php:195` `isset($data['residents_ids'])`
- **影响**：`Member/single`、`Member/edit` 传常驻地 ID 时，服务端读不到该 key，常驻地设置不生效
- **修法**：json tag 改为 `residents_ids`（bug 修复，不算行为变更）

## 3. P1 — 集团账户参数缺失（模式 B：belong_enterprise_name + taxpayer_no）

集团账户客户靠这两个参数定位实际操作的子公司（PHP `fillCompanyIdInParams()`，`controllers/Base.php:620`）。SDK 不支持 = 集团客户无法使用这些接口。

| # | 接口 | 参数落点 | SDK 现状 | 动作 | 证据 |
|---|------|---------|---------|------|------|
| 3.1 | /river/Rank/getRanks | query | ❌ 无 | ListRankApiReqBuilder +2 | 文档 11253；PHP `Rank/getRanks.php:31` |
| 3.2 | /open-apis/v1/rank/create | param_json 内 | ❌ 无 | RankInfo +2 | 文档 11255；Go `api_rank_create.yaml`（paramJson 列表含 belong/taxpayer） |
| 3.3 | /open-apis/v1/rank/update | param_json 内 | ❌ 无 | RankInfo +2 | 文档 11257；Go `api_rank_update.yaml` |
| 3.4 | /open-apis/v1/rank/del | param_json 内 | ❌ 无 | RankInfo +2 | 文档 11259；Go `api_rank_delete.yaml` |
| 3.5 | /river/Member/single | 请求体顶层 | ❌ 无 | CreateMemberRequest +2 | 文档 11155；PHP `Member/Single.php:141`（fill 4 参数，文档公开 2） |
| 3.6 | /river/Member/edit | 请求体顶层 | ❌ 无 | UpdateMemberRequest +2 | PHP `Member/Edit.php:91` |
| 3.7 | /river/Member/del | 请求体顶层 | ❌ 无 | DelMemberRequest +2 | 文档 11159；PHP `Member/Del.php:39` |
| 3.8 | /river/Member/getQuota | query | ❌ 无 | GetMemberQuotaApiReqBuilder +2 | 文档 11165；PHP `Member/GetQuota.php:35` |
| 3.9 | /river/Regulation/detail | query | ❌ 无 | GetRegulationApiReqBuilder +2 | 文档 11315；PHP `Regulation.php:33` |
| 3.10 | /river/Role/get | query | ❌ 无 | ListRoleApiReqBuilder +2 | 文档 11251；PHP `Role.php:35` |
| 3.11 | /river/LegalEntity/add | 请求体（param_json） | ⚠️ 有 TaxpayerNo（业务字段），无 belong | CreateLegalEntityRequest +belong_enterprise_name | PHP `LegalEntity/add.php:29` |
| 3.12 | /river/LegalEntity/edit | 请求体（param_json） | 同上 | UpdateLegalEntityRequest +belong_enterprise_name | PHP `LegalEntity/edit.php:29` |
| 3.13 | /river/LegalEntity/del | 请求体（param_json） | ❌ 无 | DelLegalEntityRequest +belong +taxpayer | PHP `LegalEntity/del.php:30` |

**已对齐无需处理**：BudgetCenter add/edit/del（请求体已含 belong/taxpayer/out_legal_entity_id）、Project detail/updateMember/delMember、Member/detail（builder 已含）、Member/get（SDK 已有，PHP 实际不消费，多余无害）。

## 3.5 P1.5 — 集团账户 v2 路径对齐（模式 A，4 个域）

4 个域均需：**新增 v2 资源方法 + `*V2ApiReq`/`*V2ApiReqBuilder` + `*V2Reply`**（v2 响应外层统一为 `next_token`/`records`/`progress`/`company_total`/`company_list_has_change`）。命名遵循 CLAUDE.md 命名规范，具体方法名 propose 阶段定（建议 `Get<Member|LegalEntity|BudgetCenter|Regulation>Enterprise` 或 `*V2` 后缀风格，与既有 `OutTravelerList` 命名一致性需评审）。

### 各域 v2 请求参数缺口（相对 SDK 现有 get 方法 builder）

| 域 | v2 路径 | v2 请求参数（Go IDL 权威） | SDK 现有 builder 缺失 |
|----|---------|--------------------------|---------------------|
| LegalEntity | /open-apis/v2/legalEntity/get | keyword, legal_entity_id, **out_legal_entity_id**, **query_sub_company**, length, **next_token**, **status**, **taxpayer_no**, **belong_enterprise_name**（无 offset） | out_legal_entity_id、query_sub_company、next_token、status、taxpayer_no、belong_enterprise_name（`Go idl/legal-entity.idl`） |
| BudgetCenter | /open-apis/v2/budgetCenter/get | id, out_budget_id, type, is_exact_name, name, length, **next_token**, **third_out_id**, is_need_limit_rule, **query_sub_company**, **taxpayer_no**, **belong_enterprise_name**（无 offset、**无 is_get_poi / is_get_extend_fields**） | next_token、third_out_id、query_sub_company、taxpayer_no、belong_enterprise_name（`Go idl/budget-center.idl`） |
| Member | /open-apis/v2/member/get | email, employee_number, phone, **query_sub_company**, length, **next_token**, status, **is_last_id**, **taxpayer_no**, **belong_enterprise_name**（无 offset/realname/last_id） | is_last_id（SDK ListMember 已有 next_token/query_sub_company/belong/taxpayer，最接近 v2）（`Go idl/member.idl`） |
| Regulation | /open-apis/v2/regulation/get | **query_sub_company**, **next_token**, **length**, **taxpayer_no**, **belong_enterprise_name**（/river 版无分页，v2 为分页结构） | 全部 5 个（SDK ListRegulation 仅通用 5 参数）（`Go idl/regulation.idl`） |

### 各域 v2 响应字段缺口（相对 SDK 现有 Reply 模型）

| 域 | v2 响应外层 | records 元素 vs SDK 现有模型 | v2 需补/差异 |
|----|-----------|---------------------------|-------------|
| LegalEntity | next_token/records/progress/company_total/company_list_has_change | LegalEntityInfo（14 字段）vs LegalEntityRecord（12 字段） | records 元素 + `out_parent_id`、`belong_enterprise_name` |
| BudgetCenter | 同上 | BudgetCenterInfo（27 字段）vs BudgetCenterRecord（31 字段） | v2 多 `company_id`/`belong_enterprise_name`/`create_time`/`update_time`；v2 **无** poi_list/extend_field/department_id/out_department_id/scope —— 必须独立模型 |
| Member | 同上 | MemberInfo（44 字段）vs MemberRecord（49 字段） | records 元素 + `belong_enterprise_name`（Go IDL 第 44 字段） |
| Regulation | 同上 | RegulationInfo（12 字段）vs SDK RegulationInfo（14 字段） | v2 多 `belong_enterprise_name`；SDK 多 institution_id（/river 版保留）——独立模型 |

> 注意：v2 与 /river 的 records 元素字段大体同构但**结构语义独立演进**（budgetCenter 已出现实质分叉），按「禁止复用类型」规则全部新建。

## 4. P2 — 文档新增字段缺失（版本记录可溯）

| # | 接口 | 方向 | 字段 | 类型 | 说明 | 证据 |
|---|------|------|------|------|------|------|
| 4.1 | /river/LegalEntity/get | 请求 | `out_legal_entity_id` | *string | 外部主体编号精准查询；2025.06.18 版本新增 | 文档 11123 版本记录 |
| 4.2 | /river/LegalEntity/get | 响应 | `out_parent_id` | *string | 上级公司外部编号；2025.06.18 版本新增 | 文档 11123；PHP `LegalEntity.php:166`；Go `legal-entity.idl:13` |
| 4.3 | /river/Member/getQuota | 请求 | `monthly_limit_only_query` | string | 1 仅按自然月（默认）/ 2 查询所有限额 | 文档 11165 |
| 4.4 | /river/Member/single（edit 同） | 请求 data | `residents_adcode` | *string | 常驻地行政区划代码（与 residents_ids 二选一） | 文档 11155；PHP `LoginRequest.php:66`（同构） |
| 4.5 | /river/Member/single（edit 同） | 请求 data | `home_address` | []HomeAddressInfo | 家庭住址（写入模型缺失；读取模型已有） | 文档 11155 |
| 4.6 | /river/Regulation/get、detail | 响应 | `regulation_employee_name`、`regulation_employee_description` | *string | 制度员工侧名称/描述 | PHP `Regulation.php:330,332`；Go `regulation.idl:3,4`（置信高） |

## 5. P3 — 服务端实际返回但文档未公开（建议补，需确认）

原则：用户已确认「SDK 对齐两个服务端」，服务端真实返回即对齐目标；文档滞后不作为排除依据。

| # | 接口 | 字段 | 类型 | 证据 |
|---|------|------|------|------|
| 5.1 | /river/Member/get 响应（MemberRecord） | `country_code` | *string | PHP `MemberV2Model.php:472`（常量 `country_code`，手机号区号） |
| 5.2 | 同上 | `immediate_superior_country_code` | *string | PHP `MemberV2Model.php:500` |
| 5.3 | 同上 | `extend_field` | []MemberExtendField（新结构体） | PHP `MemberV2Model.php:523` `transformMemberExtendFields()` |
| 5.4 | /river/Member/detail 响应（GetMemberDetailReply） | `country_code` | *string | PHP `MemberV2Model.php:637` |
| 5.5 | 同上 | `immediate_superior_country_code` | *string | PHP `MemberV2Model.php:663` |
| 5.6 | 同上 | `third_user_id` | *string | PHP `MemberV2Model.php:641`（list 响应已有，detail 缺） |
| 5.7 | 同上 | `extend_field` | []MemberExtendField | PHP `MemberV2Model.php:686` |

`extend_field` 结构（PHP `transformMemberExtendFields`，两种形态）：

```go
type MemberExtendField struct {
    Id           *string                  `json:"id,omitempty"`
    Name         *string                  `json:"name,omitempty"`
    Type         *string                  `json:"type,omitempty"`           // "text" | "archive"
    Value        *string                  `json:"value,omitempty"`          // text 有值
    RootCode     *string                  `json:"root_code,omitempty"`      // archive 有值
    RootName     *string                  `json:"root_name,omitempty"`
    ArchiveItems []ExtendFieldArchiveItem `json:"archive_items,omitempty"` // archive 有值
}
type ExtendFieldArchiveItem struct {
    ItemCode *string `json:"item_code,omitempty"`
    ItemName *string `json:"item_name,omitempty"`
    ItemType *string `json:"item_type,omitempty"` // "major" | "other"
}
```

## 6. P4 — 功能性缺失

| # | 接口 | 缺失 | 说明 | 证据 |
|---|------|------|------|------|
| 6.1 | /river/Member/del | 请求 `employee_phone`、`employee_email` | 文档公开 4 种删除标识（member_id/employee_number/employee_phone/employee_email），优先级：滴滴侧ID > 手机号 > 工号 > 邮箱；SDK 仅前 2 种 | 文档 11159；PHP `Member/Del.php:39` |

**记录不补**：`third_user_id` 删除标识（PHP 支持：`Del.php:54`，文档未公开，暂不暴露）。

## 7. 已对齐（探索中未发现差异，无需动作）

| 接口/域 | 结论 |
|---------|------|
| /open-apis/v2/project/outTravelerList | 与 Go IDL `traveler.idl` 14 字段完全一致（含 related_employees） |
| /open-apis/v1/traveler/create、update、del | `param_json` 模式与 `OpenapiReq`（idl/document.idl:41）吻合，响应 `TravelerOperateRespData.id` ✅ |
| rank create/update/del 响应 | `CreateRankReply.Id json:"id"` ✅（文档 11255 data.id） |
| /river/Role/get 响应 | `data` 数组 + id/name/alias 与文档 11251 一致 |
| /river/BudgetCenter/get（/river 路径）响应 | BudgetCenterRecord 27 字段完整（含 extend_field/poi_list/out_parent_id/scope）；`is_limit_quota` *int32 与 PHP 特意转 int 吻合（PHP 双路由 get/getV2 对外契约一致） |
| Member/get（/river 路径）请求 | 全参数对齐（phone/employee_number/email/realname/status/offset/length/last_id） |
| Member/single 顶层 | send_message/has_card_info ✅ |
| Project/detail/updateMember/delMember 请求 | 集团参数 ✅；`Lenth` 拼写为文档原样（既有规则） |
| BudgetCenter add/edit/del | 集团参数 + 业务字段 ✅ |

## 8. 记录不修（类型差异 / 历史决策）

| 项 | 现状 | 说明 |
|----|------|------|
| `LegalEntityRecord.LegalEntityId` *int64 | 文档标注 string、示例超 int32（`12345678543222245346`） | 遵循「已有字段不变更」；大 ID 建议值记录在案，新接口才适用 *string 规则（v2 新模型采用 *string） |
| LegalEntity/get（/river）请求 `status` | PHP rules 支持（0/1/2 多选，默认 1），文档未公开 | /river 路径不补；v2 路径 IDL 有 status，随 §3.5 新建 |
| Member/get 响应 `month_quota` 等 SDK 多余字段 | SDK 有、文档无（PHP list 有 month_quota） | omitempty 无害，不动 |
| rank create 请求 `RankInfo.RankId` | 文档 create param_json 无 rank_id（update/del 才有） | 共享结构 omitempty 无害，不动 |
| 两代分页差异 | PHP offset+length / Go v2 next_token+is_last_id | SDK /river 方法已同时兼容 offset/last_id/next_token，v2 方法用 next_token（§3.5 新建） |
| Go IDL `MemberInfo.belong_enterprise_name`（响应） | 仅 v2 路径返回，`/river/Member/get` PHP 不返回 | /river 模型不补；v2 新模型带上（§3.5） |
| `RegulationInfo` 枚举类型（SceneType/RegulationStatus *string） | 文档枚举为数字 | 未取到充分证据，不动；v2 新模型按 IDL 类型（string） |

## 9. 修复批次建议

1. **批次一（P0）**：两个 tag bug —— rank getRanks 响应独立 RankRecord；residents_ids 拼写修正
2. **批次二（P1）**：13 处集团参数（§3 表）—— 模式统一、风险低、直接解锁集团客户（老路径场景）
3. **批次三（P1.5）**：4 个域新增 v2 集团路径方法 + ApiReq/Builder/Reply + 单测 —— 工作量最大，独立 change 推进
4. **批次四（P2+P4）**：文档新增字段 + Member/del 多标识
5. **批次五（P3，需先拍板）**：Member get/detail 响应补 country_code 系列 + extend_field（含新结构体 + Builder + 单测）

每批次均需：单元测试（模型层 + httptest 资源层）、网络请求测试 case（openspec propose 阶段产出）、分支名带 `cq`。
