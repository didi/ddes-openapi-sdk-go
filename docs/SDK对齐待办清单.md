# SDK 对齐待办清单

> 生成于 2026-09-11，依据：接口参数清单复核审计（7 域 32 接口）+ 口径冲突字段级明细。
> 处理方式：自上而下逐项推进；文档类直接改清单，代码类走 openspec propose→apply，验证类用 hulu-request-skill 真机打 osim。
> 状态标记：☐ 待办 / 🔄 进行中 / ✅ 完成 / ⏸ 搁置（注明原因）。

## 总览

| 批次 | 内容 | 性质 | 项数 |
|---|---|---|---|
| 一 | 清单文档修正（A 组补录、E/F 注记、C 组确认后删） | 文档 | T1-T4 |
| 二 | 真机验证（D 组 regulation v2 层级） | 验证 | T5 |
| 三 | SDK 代码修复（小项） | 代码 | T6-T9 |
| 四 | core 层专项（加密+ParamJsonObj） | 代码（大项） | T10 |
| 五 | 遗留拍板项 | 决策 | T11 |

---

## 批次一：清单文档修正

### T1 ☐ A 组：v2 四域 SDK 已实现字段补入清单（28 条）

依据：服务端 Go IDL（openapi-go `idl/*.idl`），文档未公开但服务端真实支持；SDK 均已实现且 omitempty 可选。逐条补入清单对应 v2 小节，备注统一标注"依据：服务端 IDL，文档未公开"。

| URI | 方向 | 字段 |
|---|---|---|
| /open-apis/v2/legalEntity/get | 请求 | `status`（int32，0 停用/1 启用/2 过期）、`length`（注：与 next_token 并存关系待联调） |
| /open-apis/v2/legalEntity/get | 响应 data | `progress`（float64 查询进度）、`company_total`（int32 涉及公司总数）、`company_list_has_change`（bool 翻页期间列表是否变化） |
| /open-apis/v2/legalEntity/get | 响应 records | `status`、`belong_enterprise_name` |
| /open-apis/v2/budgetCenter/get | 请求 | `length`、`third_out_id` |
| /open-apis/v2/budgetCenter/get | 响应 data | `progress`、`company_total`、`company_list_has_change` |
| /open-apis/v2/budgetCenter/get | 响应 records | `third_out_id`、`company_id`、`belong_enterprise_name`、`create_time`、`update_time`（清单已有注记，转正式字段） |
| /open-apis/v2/member/get | 请求 | `length`、`is_last_id`（注：IDL 版游标开关，与 next_token 关系待联调） |
| /open-apis/v2/member/get | 响应 data | `progress`、`company_total`（query_sub_company=true 聚合时）、`company_list_has_change` |
| /open-apis/v2/member/get | 响应 records | `third_user_id`、`department`（老参数将废弃）、`branch_name`（老参数将废弃）、`guest_car_right`、`belong_enterprise_name`、`month_quota` |
| /open-apis/v2/regulation/get | 请求 | `length` |
| /open-apis/v2/regulation/get | 响应 data | `progress`、`company_total`、`company_list_has_change` |
| /open-apis/v2/regulation/get | 响应 records | `regulation_employee_name`、`regulation_employee_description`、`belong_enterprise_name` |

**验收**：清单四个 v2 小节含上述全部字段；`docs/接口对齐diff清单.md` 同步更新 A 组结论。
**SDK 定义参考**：model_get_legal_entity_v2{,_reply}.go、model_legal_entity_v2_record.go、model_get_budget_center_v2{,_reply}.go、model_budget_center_v2_record.go、model_list_member_v2{,_reply}.go、model_member_v2_record.go、model_list_regulation_v2{,_reply}.go、model_regulation_v2_record.go。

### T2 ☐ E 组：query_sub_company 类型注记（4 处）

清单四个 v2 小节的 `query_sub_company` 备注追加："SDK 以 string 类型传 'true'/'false'，HTTP query 传输等价"。
**验收**：4 处注记到位；SDK 代码不动。

### T3 ☐ F 组：member 域类型互反字段在清单加注（F1-F10 记录项）

清单对应接口小节加注"SDK 现用类型与文档标注相反，服务端实际返回形态待联调验证"：status（int32↔string）、source（string↔int）、is_remark（int32/string 混用）、card_type（string↔int）、member_id（int64 vs 域内 string）、monthly_limit_only_query（int32 等价）、project_codes_detail（JSON 字符串形态）。
**验收**：清单 member 域相关小节含类型注记。

### T4 ☐ C 组：member v2 的 realname/home_address 查证后处置

清单 v2 member/get 小节现列有 `realname`（请求）与 `home_address`（响应 records），但 SDK 无实现、Go IDL（idl/member.idl）无此二字段，疑为整理清单时从主接口 #15 全量表复制带入。
**动作**：先查 openapi-go 网关源码（controller/handlers + application 层）确认 v2 路径是否消费 realname、是否返回 home_address；不支持则从清单删除该 2 行（同 B 组处置），支持则补 SDK。
**验收**：查证结论记录在案，清单/SDK 按结论对齐。

---

## 批次二：真机验证

### T5 ☐ D 组：regulation v2 响应层级真机裁决

**冲突**：SDK `Data *ListRegulationV2Reply`（object：records/next_token 嵌 data 内，依据 IDL+文档请求备注）vs 清单 data=array+顶层 next_token（依据文档响应变化表）。若服务端按顶层返回，SDK 反序列化直接失败。
**动作**：openspec propose 阶段产出网络测试 case，用 hulu-request-skill 对 osim 打一次 `/open-apis/v2/regulation/get`，hulu-log-skill 查日志确认真实 JSON 结构；按结果改 SDK 或改清单（legalEntity/budgetCenter/member 三域 v2 文档未写 data 变 array，不受影响）。
**验收**：真机响应样例 + trace_id 记录在案；败方完成修改。

---

## 批次三：SDK 代码修复（每项独立 openspec change）

### T6 ☐ LegalEntity add/edit 请求补 `out_parent_id`

`CreateLegalEntityRequest` / `UpdateLegalEntityRequest` 增 `OutParentId *string json:"out_parent_id,omitempty"` + Builder 方法（含"与 parent_id 仅需传一组""不支持通过 parent_id 修改 out_parent_id"注释）+ 单测（CLAUDE.md 完整矩阵）。
**证据**：文档 11125/11127；PHP LegalEntity.php:41/69 实际消费。
**验收**：编译通过、单测全绿、清单无需改（清单本就列出该字段）。

### T7 ☐ Member/detail 响应补 `other_companies_list` + `belong_enterprise_name`

一人多企业场景字段（清单 1503 行附近已详细展开）。`GetMemberDetailReply` 增 `BelongEnterpriseName *string` 与 `OtherCompaniesList []<明细对象>`（元素=完整明细字段+belong_enterprise_name；注意"禁止复用类型"——元素需独立类型或经评审确认复用形态）+ Builder + 单测。
**证据**：php Member/Detail.php:330；Java SDK 仅 List<Object>。
**验收**：编译+单测；模型含触发条件注释（集团账户+phone+未传集团参数）。

### T8 ☐ getQuota 响应 `member_id` 加容错（F11）

`GetMemberQuotaReply.MemberId` 现 `*int64`（model_get_member_quota_reply.go:5），有 residents_list.id 前科（服务端返字符串致反序列化失败），集测未覆盖 Data 解析。按项目规则"有问题的数据类型优先 string+UnmarshalJSON"加容错（保留对外兼容）。
**验收**：单测覆盖数字与字符串两种返回形态。

### T9 ☐ budgetcenter v2 嵌套类型独立化 + 小修打包

1. `BudgetCenterV2Record` 的 `LeaderItem`/`LimitRuleItem` 复用 /river 版（model_budget_center_v2_record.go:21,32），违反"禁止复用类型"——新建 v2 专属嵌套类型（v2 limit_rule_list 无 limit_management_scope 字段，正好按 IDL 收敛）。
2. rank getRanks 3 个加密测试 mock 改 `id` tag + 补 Id 值断言（model_test.go:1574,1619,1663）。
3. 注释修正：traveler card `expire_date` 格式注释改 YYYYMMDD（model_travel_card_info.go:7）；budgetcenter add/edit `out_travelers` 上限注释"最多500个"改"最多 200 人"。
**验收**：编译+四域相关单测全绿。

---

## 批次四：core 层专项

### T10 ☐ core：加密+ParamJsonObj 组合失效修复（存量 15 个月，波及 5 域)

**缺陷**：translate() 中 Body 加密（req_translator.go:44）先于签名阶段 SignPostBody 内的 `__obj__` 替换（sign.go:93-107）；加密模式下 ParamJsonObj 的 map 值走 `fmt.Sprintf("%v")` 拼成 Go map 文本进密文，`__obj__` 替换空转，服务端解出非法 param_json。
**引入**：奠基提交 45428d2（2025-06-16）生来如此，非回归。
**修复方向**：`__obj__` 替换逻辑从 SignPostBody 提出为独立步骤，挪到 AssemblyBodyMapToEncryptionDataMap 之前。
**影响面**：traveler/order/city/workplace/rank 全部使用 ParamJsonObj 的接口 × EnableEncryption 用户；需补"加密+ParamJsonObj"组合测试（现有矩阵零覆盖）。
**验收**：新增组合单测（AES128/AES256 × ParamJsonObj）证明密文解出合法 JSON；全量回归。

---

## 批次五：遗留拍板项

### T11 ⏸ member get/detail 响应补 country_code 系列 + extend_field（旧审计 §5，维持待拍板）

`country_code`、`immediate_superior_country_code`（get+detail）、`third_user_id`（仅 detail）、`extend_field []MemberExtendField`（get+detail，结构见旧审计 §5）。证据：PHP MemberV2Model.php:472-686。与滴滴侧确认公开口径后再动。

---

## 已解决记录

| 项 | 结论 | 时间 |
|---|---|---|
| 冲突 B：v2 budgetCenter 清单多列 is_get_poi/is_get_extend_fields 及响应 5 字段+2 嵌套表 | 已从清单删除，v2 小节加差异注记 | 2026-09-11 |
| 历史修复验证（fa0beda/9dc93de/1211b64/6f2945e） | 全部真实落地，无虚报 | 2026-09-11 审计 |
| project/role/rank/traveler 域 | 生产代码对齐，零功能问题 | 2026-09-11 审计 |
