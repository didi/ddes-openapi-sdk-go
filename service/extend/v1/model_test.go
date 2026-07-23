package v1

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/didi/ddes-openapi-sdk-go/core"
)

// --- ExtendInfo Builder 测试 ---

func TestExtendInfoBuilder_FullFields(t *testing.T) {
	child := NewExtendInfoBuilder().
		Id("1125922289295590").
		Status(1).
		Code("EXT_CHILD_001").
		Name("子档案A").
		Build()

	info := NewExtendInfoBuilder().
		Id("1125922289295589").
		Status(1).
		Code("EXT_001").
		Name("北京档案").
		Children([]ExtendInfo{*child}).
		Build()

	if info.Id == nil || *info.Id != "1125922289295589" {
		t.Errorf("Id = %v, want 1125922289295589", info.Id)
	}
	if info.Status == nil || *info.Status != 1 {
		t.Errorf("Status = %v, want 1", info.Status)
	}
	if info.Code == nil || *info.Code != "EXT_001" {
		t.Errorf("Code = %v, want EXT_001", info.Code)
	}
	if info.Name == nil || *info.Name != "北京档案" {
		t.Errorf("Name = %v, want 北京档案", info.Name)
	}
	if len(info.Children) != 1 {
		t.Fatalf("Children len = %d, want 1", len(info.Children))
	}
	if info.Children[0].Id == nil || *info.Children[0].Id != "1125922289295590" {
		t.Errorf("Children[0].Id = %v, want 1125922289295590", info.Children[0].Id)
	}
}

func TestExtendInfoBuilder_PartialFields(t *testing.T) {
	info := NewExtendInfoBuilder().
		Id("1001").
		Build()

	if info.Id == nil || *info.Id != "1001" {
		t.Errorf("Id = %v, want 1001", info.Id)
	}
	if info.Status != nil {
		t.Errorf("Status = %v, want nil", info.Status)
	}
	if info.Code != nil {
		t.Errorf("Code = %v, want nil", info.Code)
	}
	if info.Name != nil {
		t.Errorf("Name = %v, want nil", info.Name)
	}
	if info.Children != nil {
		t.Errorf("Children = %v, want nil", info.Children)
	}
}

func TestExtendInfoBuilder_ZeroIntValue(t *testing.T) {
	info := NewExtendInfoBuilder().
		Status(0).
		Build()

	if info.Status == nil || *info.Status != 0 {
		t.Errorf("Status = %v, want 0", info.Status)
	}
}

// --- ErrorListItem Builder 测试 ---

func TestErrorListItemBuilder_FullFields(t *testing.T) {
	extendData := NewExtendInfoBuilder().
		Id("1125922289295591").
		Code("EXT_ERR_001").
		Name("错误档案").
		Build()

	item := NewErrorListItemBuilder().
		Data(*extendData).
		ErrMsg("子档案code重复").
		Build()

	if item.Data == nil || item.Data.Id == nil || *item.Data.Id != "1125922289295591" {
		t.Errorf("Data.Id = %v, want 1125922289295591", item.Data)
	}
	if item.Data.Code == nil || *item.Data.Code != "EXT_ERR_001" {
		t.Errorf("Data.Code = %v, want EXT_ERR_001", item.Data.Code)
	}
	if item.ErrMsg == nil || *item.ErrMsg != "子档案code重复" {
		t.Errorf("ErrMsg = %v, want 子档案code重复", item.ErrMsg)
	}
}

func TestErrorListItemBuilder_PartialFields(t *testing.T) {
	item := NewErrorListItemBuilder().
		ErrMsg("参数错误").
		Build()

	if item.ErrMsg == nil || *item.ErrMsg != "参数错误" {
		t.Errorf("ErrMsg = %v, want 参数错误", item.ErrMsg)
	}
	if item.Data != nil {
		t.Errorf("Data = %v, want nil", item.Data)
	}
}

// --- CreateExtendBatchRequestBuilder 测试 ---

func TestCreateExtendBatchRequestBuilder_FullParams(t *testing.T) {
	itemListObj := []ExtendInfo{
		*NewExtendInfoBuilder().Id("1125922289295589").Status(1).Code("EXT_001").Name("子档案A").Build(),
	}
	req := NewCreateExtendBatchRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		RootCode("ROOT_001").
		RootName("主档案").
		ItemList("[{\"id\":\"1125922289295589\"}]").
		ItemListObj(itemListObj).
		Build()

	if req.ClientId == nil || *req.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", req.ClientId)
	}
	if req.AccessToken == nil || *req.AccessToken != "test_token" {
		t.Errorf("AccessToken = %v, want test_token", req.AccessToken)
	}
	if req.CompanyId == nil || *req.CompanyId != "test_company" {
		t.Errorf("CompanyId = %v, want test_company", req.CompanyId)
	}
	if req.Timestamp == nil || *req.Timestamp != 1583484681 {
		t.Errorf("Timestamp = %v, want 1583484681", req.Timestamp)
	}
	if req.Sign == nil || *req.Sign != "test_sign" {
		t.Errorf("Sign = %v, want test_sign", req.Sign)
	}
	if req.RootCode == nil || *req.RootCode != "ROOT_001" {
		t.Errorf("RootCode = %v, want ROOT_001", req.RootCode)
	}
	if req.RootName == nil || *req.RootName != "主档案" {
		t.Errorf("RootName = %v, want 主档案", req.RootName)
	}
	if req.ItemList == nil || *req.ItemList != "[{\"id\":\"1125922289295589\"}]" {
		t.Errorf("ItemList = %v, want json string", req.ItemList)
	}
	if len(req.ItemListObj) != 1 {
		t.Errorf("ItemListObj len = %d, want 1", len(req.ItemListObj))
	}
}

func TestCreateExtendBatchRequestBuilder_PartialParams(t *testing.T) {
	req := NewCreateExtendBatchRequestBuilder().
		ClientId("test_client").
		RootCode("ROOT_001").
		Build()

	if req.ClientId == nil || *req.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", req.ClientId)
	}
	if req.RootCode == nil || *req.RootCode != "ROOT_001" {
		t.Errorf("RootCode = %v, want ROOT_001", req.RootCode)
	}
	if req.AccessToken != nil {
		t.Errorf("AccessToken = %v, want nil", req.AccessToken)
	}
	if req.CompanyId != nil {
		t.Errorf("CompanyId = %v, want nil", req.CompanyId)
	}
	if req.Timestamp != nil {
		t.Errorf("Timestamp = %v, want nil", req.Timestamp)
	}
	if req.Sign != nil {
		t.Errorf("Sign = %v, want nil", req.Sign)
	}
	if req.RootName != nil {
		t.Errorf("RootName = %v, want nil", req.RootName)
	}
	if req.ItemList != nil {
		t.Errorf("ItemList = %v, want nil", req.ItemList)
	}
	if req.ItemListObj != nil {
		t.Errorf("ItemListObj = %v, want nil", req.ItemListObj)
	}
}

func TestCreateExtendBatchRequestBuilder_ZeroIntValues(t *testing.T) {
	req := NewCreateExtendBatchRequestBuilder().
		ClientId("test_client").
		Timestamp(0).
		Build()

	if req.Timestamp == nil || *req.Timestamp != 0 {
		t.Errorf("Timestamp = %v, want 0", req.Timestamp)
	}
}

func TestCreateExtendBatchRequestBuilder_OnlyCommonParams(t *testing.T) {
	req := NewCreateExtendBatchRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		Build()

	if req.ClientId == nil || *req.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", req.ClientId)
	}
	if req.AccessToken == nil || *req.AccessToken != "test_token" {
		t.Errorf("AccessToken = %v, want test_token", req.AccessToken)
	}
	if req.CompanyId == nil || *req.CompanyId != "test_company" {
		t.Errorf("CompanyId = %v, want test_company", req.CompanyId)
	}
	if req.Timestamp == nil || *req.Timestamp != 1583484681 {
		t.Errorf("Timestamp = %v, want 1583484681", req.Timestamp)
	}
	if req.Sign == nil || *req.Sign != "test_sign" {
		t.Errorf("Sign = %v, want test_sign", req.Sign)
	}
	if req.RootCode != nil {
		t.Errorf("RootCode = %v, want nil", req.RootCode)
	}
	if req.RootName != nil {
		t.Errorf("RootName = %v, want nil", req.RootName)
	}
	if req.ItemList != nil {
		t.Errorf("ItemList = %v, want nil", req.ItemList)
	}
	if req.ItemListObj != nil {
		t.Errorf("ItemListObj = %v, want nil", req.ItemListObj)
	}
}

// --- CreateExtendBatchApiReply 反序列化测试 ---

func TestCreateExtendBatchApiReply_Deserialization(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": {
			"err_list": [
				{
					"data": {"id": "1125922289295591", "code": "EXT_ERR_001", "name": "错误档案"},
					"err_msg": "子档案code重复"
				}
			]
		},
		"request_id": "test_request_id"
	}`

	var reply CreateExtendBatchApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if reply.Errno == nil || *reply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", reply.Errno)
	}
	if reply.Errmsg == nil || *reply.Errmsg != "SUCCESS" {
		t.Errorf("Errmsg = %v, want SUCCESS", reply.Errmsg)
	}
	if reply.RequestId == nil || *reply.RequestId != "test_request_id" {
		t.Errorf("RequestId = %v, want test_request_id", reply.RequestId)
	}
	if reply.Data == nil {
		t.Fatal("Data is nil")
	}
	if len(reply.Data.ErrList) != 1 {
		t.Fatalf("ErrList len = %d, want 1", len(reply.Data.ErrList))
	}
	if reply.Data.ErrList[0].ErrMsg == nil || *reply.Data.ErrList[0].ErrMsg != "子档案code重复" {
		t.Errorf("ErrList[0].ErrMsg = %v, want 子档案code重复", reply.Data.ErrList[0].ErrMsg)
	}
	if reply.Data.ErrList[0].Data == nil || reply.Data.ErrList[0].Data.Id == nil || *reply.Data.ErrList[0].Data.Id != "1125922289295591" {
		t.Errorf("ErrList[0].Data.Id = %v, want 1125922289295591", reply.Data.ErrList[0].Data)
	}
}

func TestCreateExtendBatchApiReply_ErrorResponse(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"test_request_id"}`

	var reply CreateExtendBatchApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno == nil || *reply.Errno != 10003 {
		t.Errorf("Errno = %v, want 10003", reply.Errno)
	}
	if reply.Data != nil {
		t.Errorf("Data = %v, want nil", reply.Data)
	}
}

func TestCreateExtendBatchApiReply_EmptyErrList(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":{"err_list":[]},"request_id":"req_empty"}`

	var reply CreateExtendBatchApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Data == nil {
		t.Fatal("Data is nil")
	}
	if len(reply.Data.ErrList) != 0 {
		t.Errorf("ErrList len = %d, want 0", len(reply.Data.ErrList))
	}
}

func TestCreateExtendBatchApiReply_MissingDataField(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_nodata"}`

	var reply CreateExtendBatchApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno == nil || *reply.Errno != 10003 {
		t.Errorf("Errno = %v, want 10003", reply.Errno)
	}
	if reply.Data != nil {
		t.Errorf("Data = %v, want nil", reply.Data)
	}
}

func TestCreateExtendBatchApiReply_MultipleItemsPartialMissing(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": {
			"err_list": [
				{"data": {"id": "1001"}, "err_msg": "错误1"},
				{"err_msg": "错误2"},
				{"data": {"id": "1003", "name": "档案C"}}
			]
		},
		"request_id": "req_multi"
	}`

	var reply CreateExtendBatchApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if len(reply.Data.ErrList) != 3 {
		t.Fatalf("ErrList len = %d, want 3", len(reply.Data.ErrList))
	}
	if reply.Data.ErrList[0].ErrMsg == nil || *reply.Data.ErrList[0].ErrMsg != "错误1" {
		t.Errorf("ErrList[0].ErrMsg = %v, want 错误1", reply.Data.ErrList[0].ErrMsg)
	}
	if reply.Data.ErrList[1].Data != nil {
		t.Errorf("ErrList[1].Data = %v, want nil", reply.Data.ErrList[1].Data)
	}
	if reply.Data.ErrList[2].Data == nil || reply.Data.ErrList[2].Data.Name == nil || *reply.Data.ErrList[2].Data.Name != "档案C" {
		t.Errorf("ErrList[2].Data.Name = %v, want 档案C", reply.Data.ErrList[2].Data)
	}
}

// --- ListExtendApiReqBuilder 测试（GET 型）---

func TestListExtendApiReqBuilder_QueryParams(t *testing.T) {
	req := NewListExtendApiReqBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp("1583484681").
		Sign("test_sign").
		RootCode("ROOT_001").
		Build()

	tests := []struct {
		key  string
		want string
	}{
		{"client_id", "test_client"},
		{"access_token", "test_token"},
		{"company_id", "test_company"},
		{"timestamp", "1583484681"},
		{"sign", "test_sign"},
		{"root_code", "ROOT_001"},
	}
	for _, tt := range tests {
		got := req.apiReq.QueryParams.Get(tt.key)
		if got != tt.want {
			t.Errorf("QueryParams[%s] = %q, want %q", tt.key, got, tt.want)
		}
	}
}

func TestListExtendApiReqBuilder_PartialParams(t *testing.T) {
	req := NewListExtendApiReqBuilder().
		ClientId("test_client").
		RootCode("ROOT_001").
		Build()

	if req.apiReq.QueryParams.Get("client_id") != "test_client" {
		t.Errorf("client_id mismatch")
	}
	if req.apiReq.QueryParams.Get("root_code") != "ROOT_001" {
		t.Errorf("root_code mismatch")
	}
	if req.apiReq.QueryParams.Get("access_token") != "" {
		t.Errorf("access_token should be empty, got %q", req.apiReq.QueryParams.Get("access_token"))
	}
	if req.apiReq.QueryParams.Get("company_id") != "" {
		t.Errorf("company_id should be empty, got %q", req.apiReq.QueryParams.Get("company_id"))
	}
	if req.apiReq.QueryParams.Get("sign") != "" {
		t.Errorf("sign should be empty, got %q", req.apiReq.QueryParams.Get("sign"))
	}
}

func TestListExtendApiReqBuilder_OnlyCommonParams(t *testing.T) {
	req := NewListExtendApiReqBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp("1583484681").
		Sign("test_sign").
		Build()

	tests := []struct {
		key  string
		want string
	}{
		{"client_id", "test_client"},
		{"access_token", "test_token"},
		{"company_id", "test_company"},
		{"timestamp", "1583484681"},
		{"sign", "test_sign"},
	}
	for _, tt := range tests {
		got := req.apiReq.QueryParams.Get(tt.key)
		if got != tt.want {
			t.Errorf("QueryParams[%s] = %q, want %q", tt.key, got, tt.want)
		}
	}
	if v := req.apiReq.QueryParams.Get("root_code"); v != "" {
		t.Errorf("QueryParams[root_code] = %q, want empty", v)
	}
}

// --- ListExtendApiReply 反序列化测试 ---

func TestListExtendApiReply_Deserialization(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": {
			"root_id": "1125922289295589",
			"root_code": "ROOT_001",
			"root_name": "北京档案",
			"root_status": 1,
			"children": [
				{
					"id": "1125922289295590",
					"status": 1,
					"code": "EXT_CHILD_001",
					"name": "子档案A",
					"children": [
						{"id": "1125922289295592", "status": 1, "code": "EXT_CHILD_002", "name": "孙档案B"}
					]
				}
			]
		},
		"request_id": "test_request_id"
	}`

	var reply ListExtendApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if reply.Errno == nil || *reply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", reply.Errno)
	}
	if reply.Errmsg == nil || *reply.Errmsg != "SUCCESS" {
		t.Errorf("Errmsg = %v, want SUCCESS", reply.Errmsg)
	}
	if reply.Data == nil {
		t.Fatal("Data is nil")
	}
	if reply.Data.RootId == nil || *reply.Data.RootId != "1125922289295589" {
		t.Errorf("RootId = %v, want 1125922289295589", reply.Data.RootId)
	}
	if reply.Data.RootCode == nil || *reply.Data.RootCode != "ROOT_001" {
		t.Errorf("RootCode = %v, want ROOT_001", reply.Data.RootCode)
	}
	if reply.Data.RootName == nil || *reply.Data.RootName != "北京档案" {
		t.Errorf("RootName = %v, want 北京档案", reply.Data.RootName)
	}
	if reply.Data.RootStatus == nil || *reply.Data.RootStatus != 1 {
		t.Errorf("RootStatus = %v, want 1", reply.Data.RootStatus)
	}
	if len(reply.Data.Children) != 1 {
		t.Fatalf("Children len = %d, want 1", len(reply.Data.Children))
	}
	child := reply.Data.Children[0]
	if child.Id == nil || *child.Id != "1125922289295590" {
		t.Errorf("Children[0].Id = %v, want 1125922289295590", child.Id)
	}
	if child.Status == nil || *child.Status != 1 {
		t.Errorf("Children[0].Status = %v, want 1", child.Status)
	}
	if child.Code == nil || *child.Code != "EXT_CHILD_001" {
		t.Errorf("Children[0].Code = %v, want EXT_CHILD_001", child.Code)
	}
	if child.Name == nil || *child.Name != "子档案A" {
		t.Errorf("Children[0].Name = %v, want 子档案A", child.Name)
	}
	if len(child.Children) != 1 {
		t.Fatalf("Children[0].Children len = %d, want 1", len(child.Children))
	}
	if child.Children[0].Name == nil || *child.Children[0].Name != "孙档案B" {
		t.Errorf("Children[0].Children[0].Name = %v, want 孙档案B", child.Children[0].Name)
	}
}

func TestListExtendApiReply_ErrorResponse(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"test_request_id"}`

	var reply ListExtendApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno == nil || *reply.Errno != 10003 {
		t.Errorf("Errno = %v, want 10003", reply.Errno)
	}
}

func TestListExtendApiReply_EmptyChildren(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":{"root_id":"1125922289295589","root_code":"ROOT_001","children":[]},"request_id":"req_empty"}`

	var reply ListExtendApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if len(reply.Data.Children) != 0 {
		t.Errorf("Children len = %d, want 0", len(reply.Data.Children))
	}
}

func TestListExtendApiReply_MissingDataField(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_nodata"}`

	var reply ListExtendApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno == nil || *reply.Errno != 10003 {
		t.Errorf("Errno = %v, want 10003", reply.Errno)
	}
	if reply.Data != nil {
		t.Errorf("Data = %v, want nil", reply.Data)
	}
}

func TestListExtendApiReply_PartialFieldsMissing(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": {
			"root_id": "1125922289295589",
			"children": [
				{"id": "1001"},
				{"id": "1002", "name": "档案B"}
			]
		},
		"request_id": "req_partial"
	}`

	var reply ListExtendApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Data.RootCode != nil {
		t.Errorf("RootCode = %v, want nil", reply.Data.RootCode)
	}
	if reply.Data.RootStatus != nil {
		t.Errorf("RootStatus = %v, want nil", reply.Data.RootStatus)
	}
	if len(reply.Data.Children) != 2 {
		t.Fatalf("Children len = %d, want 2", len(reply.Data.Children))
	}
	if reply.Data.Children[0].Name != nil {
		t.Errorf("Children[0].Name = %v, want nil", reply.Data.Children[0].Name)
	}
	if reply.Data.Children[1].Name == nil || *reply.Data.Children[1].Name != "档案B" {
		t.Errorf("Children[1].Name = %v, want 档案B", reply.Data.Children[1].Name)
	}
}

// --- UpdateExtendStatusRequestBuilder 测试 ---

func TestUpdateExtendStatusRequestBuilder_FullParams(t *testing.T) {
	itemListObj := []ExtendInfo{
		*NewExtendInfoBuilder().Id("1125922289295589").Status(2).Code("EXT_001").Name("子档案A").Build(),
	}
	req := NewUpdateExtendStatusRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		RootCode("ROOT_001").
		RootStatus(2).
		ItemList("[{\"id\":\"1125922289295589\",\"status\":2}]").
		ItemListObj(itemListObj).
		Build()

	if req.ClientId == nil || *req.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", req.ClientId)
	}
	if req.AccessToken == nil || *req.AccessToken != "test_token" {
		t.Errorf("AccessToken = %v, want test_token", req.AccessToken)
	}
	if req.CompanyId == nil || *req.CompanyId != "test_company" {
		t.Errorf("CompanyId = %v, want test_company", req.CompanyId)
	}
	if req.Timestamp == nil || *req.Timestamp != 1583484681 {
		t.Errorf("Timestamp = %v, want 1583484681", req.Timestamp)
	}
	if req.Sign == nil || *req.Sign != "test_sign" {
		t.Errorf("Sign = %v, want test_sign", req.Sign)
	}
	if req.RootCode == nil || *req.RootCode != "ROOT_001" {
		t.Errorf("RootCode = %v, want ROOT_001", req.RootCode)
	}
	if req.RootStatus == nil || *req.RootStatus != 2 {
		t.Errorf("RootStatus = %v, want 2", req.RootStatus)
	}
	if req.ItemList == nil || *req.ItemList != "[{\"id\":\"1125922289295589\",\"status\":2}]" {
		t.Errorf("ItemList = %v, want json string", req.ItemList)
	}
	if len(req.ItemListObj) != 1 {
		t.Errorf("ItemListObj len = %d, want 1", len(req.ItemListObj))
	}
}

func TestUpdateExtendStatusRequestBuilder_PartialParams(t *testing.T) {
	req := NewUpdateExtendStatusRequestBuilder().
		ClientId("test_client").
		RootCode("ROOT_001").
		RootStatus(3).
		Build()

	if req.ClientId == nil || *req.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", req.ClientId)
	}
	if req.RootCode == nil || *req.RootCode != "ROOT_001" {
		t.Errorf("RootCode = %v, want ROOT_001", req.RootCode)
	}
	if req.RootStatus == nil || *req.RootStatus != 3 {
		t.Errorf("RootStatus = %v, want 3", req.RootStatus)
	}
	if req.AccessToken != nil {
		t.Errorf("AccessToken = %v, want nil", req.AccessToken)
	}
	if req.CompanyId != nil {
		t.Errorf("CompanyId = %v, want nil", req.CompanyId)
	}
	if req.Timestamp != nil {
		t.Errorf("Timestamp = %v, want nil", req.Timestamp)
	}
	if req.Sign != nil {
		t.Errorf("Sign = %v, want nil", req.Sign)
	}
	if req.ItemList != nil {
		t.Errorf("ItemList = %v, want nil", req.ItemList)
	}
	if req.ItemListObj != nil {
		t.Errorf("ItemListObj = %v, want nil", req.ItemListObj)
	}
}

func TestUpdateExtendStatusRequestBuilder_ZeroIntValues(t *testing.T) {
	req := NewUpdateExtendStatusRequestBuilder().
		ClientId("test_client").
		Timestamp(0).
		RootStatus(0).
		Build()

	if req.Timestamp == nil || *req.Timestamp != 0 {
		t.Errorf("Timestamp = %v, want 0", req.Timestamp)
	}
	if req.RootStatus == nil || *req.RootStatus != 0 {
		t.Errorf("RootStatus = %v, want 0", req.RootStatus)
	}
}

func TestUpdateExtendStatusRequestBuilder_OnlyCommonParams(t *testing.T) {
	req := NewUpdateExtendStatusRequestBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp(1583484681).
		Sign("test_sign").
		Build()

	if req.ClientId == nil || *req.ClientId != "test_client" {
		t.Errorf("ClientId = %v, want test_client", req.ClientId)
	}
	if req.AccessToken == nil || *req.AccessToken != "test_token" {
		t.Errorf("AccessToken = %v, want test_token", req.AccessToken)
	}
	if req.CompanyId == nil || *req.CompanyId != "test_company" {
		t.Errorf("CompanyId = %v, want test_company", req.CompanyId)
	}
	if req.Timestamp == nil || *req.Timestamp != 1583484681 {
		t.Errorf("Timestamp = %v, want 1583484681", req.Timestamp)
	}
	if req.Sign == nil || *req.Sign != "test_sign" {
		t.Errorf("Sign = %v, want test_sign", req.Sign)
	}
	if req.RootCode != nil {
		t.Errorf("RootCode = %v, want nil", req.RootCode)
	}
	if req.RootStatus != nil {
		t.Errorf("RootStatus = %v, want nil", req.RootStatus)
	}
	if req.ItemList != nil {
		t.Errorf("ItemList = %v, want nil", req.ItemList)
	}
	if req.ItemListObj != nil {
		t.Errorf("ItemListObj = %v, want nil", req.ItemListObj)
	}
}

// --- UpdateExtendStatusApiReply 反序列化测试 ---

func TestUpdateExtendStatusApiReply_Deserialization(t *testing.T) {
	jsonData := `{
		"errno": 0,
		"errmsg": "SUCCESS",
		"data": {"result": "ok"},
		"request_id": "test_request_id"
	}`

	var reply UpdateExtendStatusApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", reply.Errno)
	}
	if reply.Errmsg != "SUCCESS" {
		t.Errorf("Errmsg = %q, want SUCCESS", reply.Errmsg)
	}
	if reply.RequestId != "test_request_id" {
		t.Errorf("RequestId = %q, want test_request_id", reply.RequestId)
	}
	if reply.Data == nil {
		t.Error("Data should not be nil")
	}
}

func TestUpdateExtendStatusApiReply_ErrorResponse(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"test_request_id"}`

	var reply UpdateExtendStatusApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 10003 {
		t.Errorf("Errno = %d, want 10003", reply.Errno)
	}
	if reply.Errmsg != "param error" {
		t.Errorf("Errmsg = %q, want param error", reply.Errmsg)
	}
}

func TestUpdateExtendStatusApiReply_EmptyData(t *testing.T) {
	jsonData := `{"errno":0,"errmsg":"SUCCESS","data":[],"request_id":"req_empty"}`

	var reply UpdateExtendStatusApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", reply.Errno)
	}
}

func TestUpdateExtendStatusApiReply_MissingDataField(t *testing.T) {
	jsonData := `{"errno":10003,"errmsg":"param error","request_id":"req_nodata"}`

	var reply UpdateExtendStatusApiReply
	if err := json.Unmarshal([]byte(jsonData), &reply); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if reply.Errno != 10003 {
		t.Errorf("Errno = %d, want 10003", reply.Errno)
	}
	if reply.Data != nil {
		t.Errorf("Data = %v, want nil", reply.Data)
	}
}

// --- UpdateExtendStatusReply 测试（空结构体）---

func TestUpdateExtendStatusReply(t *testing.T) {
	reply := UpdateExtendStatusReply{}
	if reply != (UpdateExtendStatusReply{}) {
		t.Errorf("UpdateExtendStatusReply should be zero value struct")
	}
}

// --- CreateExtendBatch 资源方法测试 ---

func newExtendTestOption(serverURL string) *core.Option {
	return &core.Option{
		BaseUrl:        serverURL,
		RequestTimeOut: 5 * time.Second,
		SignMethod:     1,
		Serializer:     &core.DefaultSerializer{},
		Logger:         core.NewLoggerImpl(core.LogLevelInfo, core.NewDefaultLogger(core.LogLevelInfo)),
	}
}

func TestCreateExtendBatch_Success(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/river/ExtendInfo/BatchSync" {
			t.Errorf("expected path /river/ExtendInfo/BatchSync, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"errno": 0,
			"errmsg": "SUCCESS",
			"data": {
				"err_list": [
					{"data": {"id": "1125922289295591", "code": "EXT_ERR_001"}, "err_msg": "子档案code重复"}
				]
			},
			"request_id": "req_001"
		}`))
	}))
	defer testServer.Close()

	option := newExtendTestOption(testServer.URL)
	svc := &extend{option: option}

	req := NewCreateExtendBatchApiReqBuilder().
		CreateExtendBatchRequest(NewCreateExtendBatchRequestBuilder().
			ClientId("test_client").
			AccessToken("test_token").
			CompanyId("test_company").
			Timestamp(1583484681).
			Sign("test_sign").
			RootCode("ROOT_001").
			RootName("主档案").
			Build()).
		Build()

	resp, err := svc.CreateExtendBatch(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("CreateExtendBatch() error = %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusOK)
	}
	if resp.CreateExtendBatchApiReply == nil {
		t.Fatal("CreateExtendBatchApiReply is nil")
	}
	if resp.CreateExtendBatchApiReply.Errno == nil || *resp.CreateExtendBatchApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.CreateExtendBatchApiReply.Errno)
	}
	if resp.CreateExtendBatchApiReply.Data == nil {
		t.Fatal("Data is nil")
	}
	if len(resp.CreateExtendBatchApiReply.Data.ErrList) != 1 {
		t.Fatalf("ErrList len = %d, want 1", len(resp.CreateExtendBatchApiReply.Data.ErrList))
	}
	if resp.CreateExtendBatchApiReply.Data.ErrList[0].ErrMsg == nil || *resp.CreateExtendBatchApiReply.Data.ErrList[0].ErrMsg != "子档案code重复" {
		t.Errorf("ErrList[0].ErrMsg = %v, want 子档案code重复", resp.CreateExtendBatchApiReply.Data.ErrList[0].ErrMsg)
	}
}

func TestCreateExtendBatch_EmptyData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"err_list":[]},"request_id":"req_002"}`))
	}))
	defer testServer.Close()

	option := newExtendTestOption(testServer.URL)
	svc := &extend{option: option}

	req := NewCreateExtendBatchApiReqBuilder().
		CreateExtendBatchRequest(NewCreateExtendBatchRequestBuilder().
			ClientId("test_client").
			RootCode("ROOT_001").
			Build()).
		Build()

	resp, err := svc.CreateExtendBatch(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("CreateExtendBatch() error = %v", err)
	}
	if resp.CreateExtendBatchApiReply == nil {
		t.Fatal("CreateExtendBatchApiReply is nil")
	}
	if len(resp.CreateExtendBatchApiReply.Data.ErrList) != 0 {
		t.Errorf("ErrList len = %d, want 0", len(resp.CreateExtendBatchApiReply.Data.ErrList))
	}
}

func TestCreateExtendBatch_ApiError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":10003,"errmsg":"param error","request_id":"req_003"}`))
	}))
	defer testServer.Close()

	option := newExtendTestOption(testServer.URL)
	svc := &extend{option: option}

	req := NewCreateExtendBatchApiReqBuilder().
		CreateExtendBatchRequest(NewCreateExtendBatchRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := svc.CreateExtendBatch(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("CreateExtendBatch() error = %v", err)
	}
	if resp.CreateExtendBatchApiReply == nil {
		t.Fatal("CreateExtendBatchApiReply is nil")
	}
	if resp.CreateExtendBatchApiReply.Errno == nil || *resp.CreateExtendBatchApiReply.Errno != 10003 {
		t.Errorf("Errno = %v, want 10003", resp.CreateExtendBatchApiReply.Errno)
	}
}

func TestCreateExtendBatch_HttpError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer testServer.Close()

	option := newExtendTestOption(testServer.URL)
	svc := &extend{option: option}

	req := NewCreateExtendBatchApiReqBuilder().
		CreateExtendBatchRequest(NewCreateExtendBatchRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := svc.CreateExtendBatch(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("CreateExtendBatch() error = %v", err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusInternalServerError)
	}
	if resp.CreateExtendBatchApiReply != nil {
		t.Errorf("CreateExtendBatchApiReply should be nil for non-200 response")
	}
}

func TestCreateExtendBatch_WithEncryption_AES128(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":{"err_list":[{"data":{"id":"1125922289295591"},"err_msg":"code重复"}]},"request_id":"req_enc"}`
	key := []byte("16byte-key-12345")
	encrypted, err := core.AESEncryptECB([]byte(plaintext), key)
	if err != nil {
		t.Fatalf("AESEncryptECB() error = %v", err)
	}
	encryptData := base64.StdEncoding.EncodeToString(encrypted)

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"encrypt_data":"` + encryptData + `"}`))
	}))
	defer testServer.Close()

	option := newExtendTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: string(key),
	}
	svc := &extend{option: option}

	req := NewCreateExtendBatchApiReqBuilder().
		CreateExtendBatchRequest(NewCreateExtendBatchRequestBuilder().
			ClientId("test_client").
			RootCode("ROOT_001").
			Build()).
		Build()

	resp, err := svc.CreateExtendBatch(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("CreateExtendBatch() error = %v", err)
	}
	if resp.CreateExtendBatchApiReply == nil {
		t.Fatal("CreateExtendBatchApiReply is nil")
	}
	if resp.CreateExtendBatchApiReply.Errno == nil || *resp.CreateExtendBatchApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.CreateExtendBatchApiReply.Errno)
	}
	if resp.CreateExtendBatchApiReply.RequestId == nil || *resp.CreateExtendBatchApiReply.RequestId != "req_enc" {
		t.Errorf("RequestId = %v, want req_enc", resp.CreateExtendBatchApiReply.RequestId)
	}
	if len(resp.CreateExtendBatchApiReply.Data.ErrList) != 1 {
		t.Fatalf("ErrList len = %d, want 1", len(resp.CreateExtendBatchApiReply.Data.ErrList))
	}
}

func TestCreateExtendBatch_WithEncryption_AES256(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":{"err_list":[]},"request_id":"req_enc256"}`
	key := []byte("32byte-key-1234567890abcdefghijk")
	encrypted, err := core.AESEncryptECB([]byte(plaintext), key)
	if err != nil {
		t.Fatalf("AESEncryptECB() error = %v", err)
	}
	encryptData := base64.URLEncoding.EncodeToString(encrypted)

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"encrypt_data":"` + encryptData + `"}`))
	}))
	defer testServer.Close()

	option := newExtendTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 2,
		Key: string(key),
	}
	svc := &extend{option: option}

	req := NewCreateExtendBatchApiReqBuilder().
		CreateExtendBatchRequest(NewCreateExtendBatchRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := svc.CreateExtendBatch(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("CreateExtendBatch() error = %v", err)
	}
	if resp.CreateExtendBatchApiReply == nil {
		t.Fatal("CreateExtendBatchApiReply is nil")
	}
	if resp.CreateExtendBatchApiReply.RequestId == nil || *resp.CreateExtendBatchApiReply.RequestId != "req_enc256" {
		t.Errorf("RequestId = %v, want req_enc256", resp.CreateExtendBatchApiReply.RequestId)
	}
}

func TestCreateExtendBatch_EncryptionNoEncryptData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"err_list":[]},"request_id":"req_noenc"}`))
	}))
	defer testServer.Close()

	option := newExtendTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: "16byte-key-12345",
	}
	svc := &extend{option: option}

	req := NewCreateExtendBatchApiReqBuilder().
		CreateExtendBatchRequest(NewCreateExtendBatchRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := svc.CreateExtendBatch(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("CreateExtendBatch() error = %v", err)
	}
	if resp.CreateExtendBatchApiReply == nil {
		t.Fatal("CreateExtendBatchApiReply is nil")
	}
	if resp.CreateExtendBatchApiReply.Errno == nil || *resp.CreateExtendBatchApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.CreateExtendBatchApiReply.Errno)
	}
}

func TestCreateExtendBatch_WithReqOption(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if customHeader := r.Header.Get("X-Custom-Header"); customHeader != "custom-value" {
			t.Errorf("expected X-Custom-Header custom-value, got %s", customHeader)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"err_list":[]},"request_id":"req_opt"}`))
	}))
	defer testServer.Close()

	option := newExtendTestOption(testServer.URL)
	svc := &extend{option: option}

	customHeader := http.Header{}
	customHeader.Set("X-Custom-Header", "custom-value")
	reqOption := &core.ReqOption{
		Header: customHeader,
	}

	req := NewCreateExtendBatchApiReqBuilder().
		CreateExtendBatchRequest(NewCreateExtendBatchRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := svc.CreateExtendBatch(context.Background(), req, reqOption)
	if err != nil {
		t.Fatalf("CreateExtendBatch() error = %v", err)
	}
	if resp.CreateExtendBatchApiReply == nil {
		t.Fatal("CreateExtendBatchApiReply is nil")
	}
}

// --- ListExtend 资源方法测试 ---

func TestListExtend_Success(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/river/ExtendInfo/Get" {
			t.Errorf("expected path /river/ExtendInfo/Get, got %s", r.URL.Path)
		}
		if r.URL.Query().Get("client_id") != "test_client" {
			t.Errorf("client_id = %q, want test_client", r.URL.Query().Get("client_id"))
		}
		if r.URL.Query().Get("root_code") != "ROOT_001" {
			t.Errorf("root_code = %q, want ROOT_001", r.URL.Query().Get("root_code"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"errno": 0,
			"errmsg": "SUCCESS",
			"data": {
				"root_id": "1125922289295589",
				"root_code": "ROOT_001",
				"root_name": "北京档案",
				"root_status": 1,
				"children": [
					{"id": "1125922289295590", "status": 1, "code": "EXT_CHILD_001", "name": "子档案A"}
				]
			},
			"request_id": "req_001"
		}`))
	}))
	defer testServer.Close()

	option := newExtendTestOption(testServer.URL)
	svc := &extend{option: option}

	req := NewListExtendApiReqBuilder().
		ClientId("test_client").
		AccessToken("test_token").
		CompanyId("test_company").
		Timestamp("1583484681").
		Sign("test_sign").
		RootCode("ROOT_001").
		Build()

	resp, err := svc.ListExtend(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListExtend() error = %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusOK)
	}
	if resp.ListExtendApiReply == nil {
		t.Fatal("ListExtendApiReply is nil")
	}
	if resp.ListExtendApiReply.Errno == nil || *resp.ListExtendApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.ListExtendApiReply.Errno)
	}
	if resp.ListExtendApiReply.Data == nil {
		t.Fatal("Data is nil")
	}
	if resp.ListExtendApiReply.Data.RootId == nil || *resp.ListExtendApiReply.Data.RootId != "1125922289295589" {
		t.Errorf("RootId = %v, want 1125922289295589", resp.ListExtendApiReply.Data.RootId)
	}
	if len(resp.ListExtendApiReply.Data.Children) != 1 {
		t.Fatalf("Children len = %d, want 1", len(resp.ListExtendApiReply.Data.Children))
	}
	if resp.ListExtendApiReply.Data.Children[0].Name == nil || *resp.ListExtendApiReply.Data.Children[0].Name != "子档案A" {
		t.Errorf("Children[0].Name = %v, want 子档案A", resp.ListExtendApiReply.Data.Children[0].Name)
	}
}

func TestListExtend_EmptyData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"root_id":"1125922289295589","children":[]},"request_id":"req_002"}`))
	}))
	defer testServer.Close()

	option := newExtendTestOption(testServer.URL)
	svc := &extend{option: option}

	req := NewListExtendApiReqBuilder().
		ClientId("test_client").
		RootCode("ROOT_001").
		Build()

	resp, err := svc.ListExtend(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListExtend() error = %v", err)
	}
	if resp.ListExtendApiReply == nil {
		t.Fatal("ListExtendApiReply is nil")
	}
	if len(resp.ListExtendApiReply.Data.Children) != 0 {
		t.Errorf("Children len = %d, want 0", len(resp.ListExtendApiReply.Data.Children))
	}
}

func TestListExtend_ApiError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":10003,"errmsg":"param error","request_id":"req_003"}`))
	}))
	defer testServer.Close()

	option := newExtendTestOption(testServer.URL)
	svc := &extend{option: option}

	req := NewListExtendApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := svc.ListExtend(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListExtend() error = %v", err)
	}
	if resp.ListExtendApiReply == nil {
		t.Fatal("ListExtendApiReply is nil")
	}
	if resp.ListExtendApiReply.Errno == nil || *resp.ListExtendApiReply.Errno != 10003 {
		t.Errorf("Errno = %v, want 10003", resp.ListExtendApiReply.Errno)
	}
}

func TestListExtend_HttpError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer testServer.Close()

	option := newExtendTestOption(testServer.URL)
	svc := &extend{option: option}

	req := NewListExtendApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := svc.ListExtend(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListExtend() error = %v", err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusInternalServerError)
	}
	if resp.ListExtendApiReply != nil {
		t.Errorf("ListExtendApiReply should be nil for non-200 response")
	}
}

func TestListExtend_WithEncryption_AES128(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":{"root_id":"1125922289295589","root_code":"ROOT_001","children":[{"id":"1125922289295590","name":"子档案A"}]},"request_id":"req_enc"}`
	key := []byte("16byte-key-12345")
	encrypted, err := core.AESEncryptECB([]byte(plaintext), key)
	if err != nil {
		t.Fatalf("AESEncryptECB() error = %v", err)
	}
	encryptData := base64.StdEncoding.EncodeToString(encrypted)

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"encrypt_data":"` + encryptData + `"}`))
	}))
	defer testServer.Close()

	option := newExtendTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: string(key),
	}
	svc := &extend{option: option}

	req := NewListExtendApiReqBuilder().
		ClientId("test_client").
		RootCode("ROOT_001").
		Build()

	resp, err := svc.ListExtend(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListExtend() error = %v", err)
	}
	if resp.ListExtendApiReply == nil {
		t.Fatal("ListExtendApiReply is nil")
	}
	if resp.ListExtendApiReply.Errno == nil || *resp.ListExtendApiReply.Errno != 0 {
		t.Errorf("Errno = %v, want 0", resp.ListExtendApiReply.Errno)
	}
	if len(resp.ListExtendApiReply.Data.Children) != 1 {
		t.Fatalf("Children len = %d, want 1", len(resp.ListExtendApiReply.Data.Children))
	}
}

func TestListExtend_WithEncryption_AES256(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":{"root_id":"1125922289295589","children":[{"id":"1125922289295590"}]},"request_id":"req_enc256"}`
	key := []byte("32byte-key-1234567890abcdefghijk")
	encrypted, err := core.AESEncryptECB([]byte(plaintext), key)
	if err != nil {
		t.Fatalf("AESEncryptECB() error = %v", err)
	}
	encryptData := base64.URLEncoding.EncodeToString(encrypted)

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"encrypt_data":"` + encryptData + `"}`))
	}))
	defer testServer.Close()

	option := newExtendTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 2,
		Key: string(key),
	}
	svc := &extend{option: option}

	req := NewListExtendApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := svc.ListExtend(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListExtend() error = %v", err)
	}
	if resp.ListExtendApiReply == nil {
		t.Fatal("ListExtendApiReply is nil")
	}
	if len(resp.ListExtendApiReply.Data.Children) != 1 {
		t.Fatalf("Children len = %d, want 1", len(resp.ListExtendApiReply.Data.Children))
	}
}

func TestListExtend_EncryptionNoEncryptData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"root_id":"1125922289295589","children":[]},"request_id":"req_noenc"}`))
	}))
	defer testServer.Close()

	option := newExtendTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: "16byte-key-12345",
	}
	svc := &extend{option: option}

	req := NewListExtendApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := svc.ListExtend(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("ListExtend() error = %v", err)
	}
	if resp.ListExtendApiReply == nil {
		t.Fatal("ListExtendApiReply is nil")
	}
	if len(resp.ListExtendApiReply.Data.Children) != 0 {
		t.Errorf("Children len = %d, want 0", len(resp.ListExtendApiReply.Data.Children))
	}
}

func TestListExtend_WithReqOption(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if customHeader := r.Header.Get("X-Custom-Header"); customHeader != "custom-value" {
			t.Errorf("expected X-Custom-Header custom-value, got %s", customHeader)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"root_id":"1125922289295589","children":[]},"request_id":"req_opt"}`))
	}))
	defer testServer.Close()

	option := newExtendTestOption(testServer.URL)
	svc := &extend{option: option}

	customHeader := http.Header{}
	customHeader.Set("X-Custom-Header", "custom-value")
	reqOption := &core.ReqOption{
		Header: customHeader,
	}

	req := NewListExtendApiReqBuilder().
		ClientId("test_client").
		Build()

	resp, err := svc.ListExtend(context.Background(), req, reqOption)
	if err != nil {
		t.Fatalf("ListExtend() error = %v", err)
	}
	if resp.ListExtendApiReply == nil {
		t.Fatal("ListExtendApiReply is nil")
	}
}

// --- UpdateExtendStatus 资源方法测试 ---

func TestUpdateExtendStatus_Success(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/river/ExtendInfo/Status" {
			t.Errorf("expected path /river/ExtendInfo/Status, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"errno": 0,
			"errmsg": "SUCCESS",
			"data": {"result": "ok"},
			"request_id": "req_001"
		}`))
	}))
	defer testServer.Close()

	option := newExtendTestOption(testServer.URL)
	svc := &extend{option: option}

	req := NewUpdateExtendStatusApiReqBuilder().
		UpdateExtendStatusRequest(NewUpdateExtendStatusRequestBuilder().
			ClientId("test_client").
			AccessToken("test_token").
			CompanyId("test_company").
			Timestamp(1583484681).
			Sign("test_sign").
			RootCode("ROOT_001").
			RootStatus(2).
			Build()).
		Build()

	resp, err := svc.UpdateExtendStatus(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateExtendStatus() error = %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusOK)
	}
	if resp.UpdateExtendStatusApiReply == nil {
		t.Fatal("UpdateExtendStatusApiReply is nil")
	}
	if resp.UpdateExtendStatusApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.UpdateExtendStatusApiReply.Errno)
	}
	if resp.UpdateExtendStatusApiReply.Errmsg != "SUCCESS" {
		t.Errorf("Errmsg = %q, want SUCCESS", resp.UpdateExtendStatusApiReply.Errmsg)
	}
	if resp.UpdateExtendStatusApiReply.RequestId != "req_001" {
		t.Errorf("RequestId = %q, want req_001", resp.UpdateExtendStatusApiReply.RequestId)
	}
}

func TestUpdateExtendStatus_EmptyData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":[],"request_id":"req_002"}`))
	}))
	defer testServer.Close()

	option := newExtendTestOption(testServer.URL)
	svc := &extend{option: option}

	req := NewUpdateExtendStatusApiReqBuilder().
		UpdateExtendStatusRequest(NewUpdateExtendStatusRequestBuilder().
			ClientId("test_client").
			RootCode("ROOT_001").
			RootStatus(3).
			Build()).
		Build()

	resp, err := svc.UpdateExtendStatus(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateExtendStatus() error = %v", err)
	}
	if resp.UpdateExtendStatusApiReply == nil {
		t.Fatal("UpdateExtendStatusApiReply is nil")
	}
	if resp.UpdateExtendStatusApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.UpdateExtendStatusApiReply.Errno)
	}
}

func TestUpdateExtendStatus_ApiError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":10003,"errmsg":"param error","request_id":"req_003"}`))
	}))
	defer testServer.Close()

	option := newExtendTestOption(testServer.URL)
	svc := &extend{option: option}

	req := NewUpdateExtendStatusApiReqBuilder().
		UpdateExtendStatusRequest(NewUpdateExtendStatusRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := svc.UpdateExtendStatus(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateExtendStatus() error = %v", err)
	}
	if resp.UpdateExtendStatusApiReply == nil {
		t.Fatal("UpdateExtendStatusApiReply is nil")
	}
	if resp.UpdateExtendStatusApiReply.Errno != 10003 {
		t.Errorf("Errno = %d, want 10003", resp.UpdateExtendStatusApiReply.Errno)
	}
	if resp.UpdateExtendStatusApiReply.Errmsg != "param error" {
		t.Errorf("Errmsg = %q, want param error", resp.UpdateExtendStatusApiReply.Errmsg)
	}
}

func TestUpdateExtendStatus_HttpError(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer testServer.Close()

	option := newExtendTestOption(testServer.URL)
	svc := &extend{option: option}

	req := NewUpdateExtendStatusApiReqBuilder().
		UpdateExtendStatusRequest(NewUpdateExtendStatusRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := svc.UpdateExtendStatus(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateExtendStatus() error = %v", err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusInternalServerError)
	}
	if resp.UpdateExtendStatusApiReply != nil {
		t.Errorf("UpdateExtendStatusApiReply should be nil for non-200 response")
	}
}

func TestUpdateExtendStatus_WithEncryption_AES128(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":{"result":"ok"},"request_id":"req_enc"}`
	key := []byte("16byte-key-12345")
	encrypted, err := core.AESEncryptECB([]byte(plaintext), key)
	if err != nil {
		t.Fatalf("AESEncryptECB() error = %v", err)
	}
	encryptData := base64.StdEncoding.EncodeToString(encrypted)

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"encrypt_data":"` + encryptData + `"}`))
	}))
	defer testServer.Close()

	option := newExtendTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: string(key),
	}
	svc := &extend{option: option}

	req := NewUpdateExtendStatusApiReqBuilder().
		UpdateExtendStatusRequest(NewUpdateExtendStatusRequestBuilder().
			ClientId("test_client").
			RootCode("ROOT_001").
			RootStatus(2).
			Build()).
		Build()

	resp, err := svc.UpdateExtendStatus(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateExtendStatus() error = %v", err)
	}
	if resp.UpdateExtendStatusApiReply == nil {
		t.Fatal("UpdateExtendStatusApiReply is nil")
	}
	if resp.UpdateExtendStatusApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.UpdateExtendStatusApiReply.Errno)
	}
	if resp.UpdateExtendStatusApiReply.RequestId != "req_enc" {
		t.Errorf("RequestId = %q, want req_enc", resp.UpdateExtendStatusApiReply.RequestId)
	}
}

func TestUpdateExtendStatus_WithEncryption_AES256(t *testing.T) {
	plaintext := `{"errno":0,"errmsg":"SUCCESS","data":{"result":"ok"},"request_id":"req_enc256"}`
	key := []byte("32byte-key-1234567890abcdefghijk")
	encrypted, err := core.AESEncryptECB([]byte(plaintext), key)
	if err != nil {
		t.Fatalf("AESEncryptECB() error = %v", err)
	}
	encryptData := base64.URLEncoding.EncodeToString(encrypted)

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"encrypt_data":"` + encryptData + `"}`))
	}))
	defer testServer.Close()

	option := newExtendTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 2,
		Key: string(key),
	}
	svc := &extend{option: option}

	req := NewUpdateExtendStatusApiReqBuilder().
		UpdateExtendStatusRequest(NewUpdateExtendStatusRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := svc.UpdateExtendStatus(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateExtendStatus() error = %v", err)
	}
	if resp.UpdateExtendStatusApiReply == nil {
		t.Fatal("UpdateExtendStatusApiReply is nil")
	}
	if resp.UpdateExtendStatusApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.UpdateExtendStatusApiReply.Errno)
	}
	if resp.UpdateExtendStatusApiReply.RequestId != "req_enc256" {
		t.Errorf("RequestId = %q, want req_enc256", resp.UpdateExtendStatusApiReply.RequestId)
	}
}

func TestUpdateExtendStatus_EncryptionNoEncryptData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"result":"ok"},"request_id":"req_noenc"}`))
	}))
	defer testServer.Close()

	option := newExtendTestOption(testServer.URL)
	option.EnableEncryption = true
	option.EncryptionOption = &core.EncryptionOption{
		Ent: 1,
		Key: "16byte-key-12345",
	}
	svc := &extend{option: option}

	req := NewUpdateExtendStatusApiReqBuilder().
		UpdateExtendStatusRequest(NewUpdateExtendStatusRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := svc.UpdateExtendStatus(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("UpdateExtendStatus() error = %v", err)
	}
	if resp.UpdateExtendStatusApiReply == nil {
		t.Fatal("UpdateExtendStatusApiReply is nil")
	}
	if resp.UpdateExtendStatusApiReply.Errno != 0 {
		t.Errorf("Errno = %d, want 0", resp.UpdateExtendStatusApiReply.Errno)
	}
}

func TestUpdateExtendStatus_WithReqOption(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if customHeader := r.Header.Get("X-Custom-Header"); customHeader != "custom-value" {
			t.Errorf("expected X-Custom-Header custom-value, got %s", customHeader)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errno":0,"errmsg":"SUCCESS","data":{"result":"ok"},"request_id":"req_opt"}`))
	}))
	defer testServer.Close()

	option := newExtendTestOption(testServer.URL)
	svc := &extend{option: option}

	customHeader := http.Header{}
	customHeader.Set("X-Custom-Header", "custom-value")
	reqOption := &core.ReqOption{
		Header: customHeader,
	}

	req := NewUpdateExtendStatusApiReqBuilder().
		UpdateExtendStatusRequest(NewUpdateExtendStatusRequestBuilder().
			ClientId("test_client").
			Build()).
		Build()

	resp, err := svc.UpdateExtendStatus(context.Background(), req, reqOption)
	if err != nil {
		t.Fatalf("UpdateExtendStatus() error = %v", err)
	}
	if resp.UpdateExtendStatusApiReply == nil {
		t.Fatal("UpdateExtendStatusApiReply is nil")
	}
}
