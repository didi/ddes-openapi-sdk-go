package v1

// CarRuleInfo struct for CarRuleInfo 用车规则
type CarRuleInfo struct {
	CallType []int32 `json:"call_type,omitempty"` // 叫车方式，枚举值数字：1-自用，2-代叫
	UseScene []int32 `json:"use_scene,omitempty"` // 使用场景，枚举值数字：1-市内用车，2-接送机，3、4-接送其他
}

type CarRuleInfoBuilder struct {
	callType    []int32 // 叫车方式，枚举值数字：1-自用，2-代叫
	callTypeSet bool
	useScene    []int32 // 使用场景，枚举值数字：1-市内用车，2-接送机，3、4-接送其他
	useSceneSet bool
}

func NewCarRuleInfoBuilder() *CarRuleInfoBuilder {
	return &CarRuleInfoBuilder{}
}
func (builder *CarRuleInfoBuilder) CallType(callType []int32) *CarRuleInfoBuilder {
	builder.callType = callType
	builder.callTypeSet = true
	return builder
}
func (builder *CarRuleInfoBuilder) UseScene(useScene []int32) *CarRuleInfoBuilder {
	builder.useScene = useScene
	builder.useSceneSet = true
	return builder
}

func (builder *CarRuleInfoBuilder) Build() *CarRuleInfo {
	data := &CarRuleInfo{}
	if builder.callTypeSet {
		data.CallType = builder.callType
	}
	if builder.useSceneSet {
		data.UseScene = builder.useScene
	}
	return data
}
