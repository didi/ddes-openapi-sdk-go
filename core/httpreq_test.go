package core

import "testing"

// 测试完整地址拼接函数
func TestAssemblyFullUrl(t *testing.T) {
	baseUrl := "http://127.0.0.1:8080"
	apiPath := "/river/auth/{name}/authorize/{age}"
	pathParams := make(map[string]string)    // 路径参数
	queryParams := make(map[string][]string) // 查询参数

	pathParams["name"] = "test"
	pathParams["age"] = "100"
	queryParams["city"] = []string{"chengdu", "chongqing"}
	queryParams["street"] = []string{"tianfuStreet"}
	// 期望拼接结果：http://127.0.0.1:8080/river/auth/test/authorize/100?city=chengdu&city=chongqing&street=tianfuStreet
	fullUrl := AssemblyFullUrl(baseUrl, apiPath, pathParams, queryParams)
	t.Log(fullUrl)

	//type args struct {
	//	baseUrl     string
	//	apiPath     string
	//	pathParams  map[string]string
	//	queryParams map[string][]string
	//}
	//tests := []struct {
	//	name string
	//	args args
	//	want string
	//}{
	//	{
	//		name: "",
	//		args: args{
	//			baseUrl:     "",
	//			apiPath:     "",
	//			pathParams:  nil,
	//			queryParams: nil,
	//		},
	//		want: "",
	//	},
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		if got := AssemblyFullUrl(tt.args.baseUrl, tt.args.apiPath, tt.args.pathParams, tt.args.queryParams); got != tt.want {
	//			t.Errorf("AssemblyFullUrl() = %v, want %v", got, tt.want)
	//		}
	//	})
	//}
}
