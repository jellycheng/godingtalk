package godingtalk

// 企业配置
type DingCorpConfig struct {
	CorpId   string
	APIToken string
}

// 应用配置
type DingAgentConfig struct {
	AgentId   string
	AppKey    string
	AppSecret string
	AesKey    string
	Token     string
}
