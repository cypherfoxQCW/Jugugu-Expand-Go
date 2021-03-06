/*
 * 扩展接入
 *
 * 接入jugugu区块链全包的全部接口，访问IP需要连续管理员，添加业务服务器IP
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// InlineResponse2002 struct for InlineResponse2002
type InlineResponse2002 struct {
	// -1代表错误 1代表成功 0代表提示
	State string `json:"state"`
	// 状态码对应的详细信息
	Msg string `json:"msg"`
	// 手机号
	Phone string `json:"phone"`
	// 验证码图片ID
	Virifycodeid string `json:"virifycodeid"`
	// 验证码图片数据data:image/png;base64,格式
	Virifyimage string `json:"virifyimage"`
	// 树图区块链地址
	Confluxaddress string `json:"confluxaddress"`
	// 以太坊及其侧链地址
	Ethaddress string `json:"ethaddress"`
	// 登录成功后获得的令牌，用于jugugu其他各类操作
	Token string `json:"token"`
	// cookie无感登录的令牌
	Cookietoken string `json:"cookietoken"`
}
