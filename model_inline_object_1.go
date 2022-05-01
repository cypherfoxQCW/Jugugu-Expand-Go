/*
 * 扩展接入
 *
 * 接入jugugu区块链全包的全部接口，访问IP需要连续管理员，添加业务服务器IP
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// InlineObject1 struct for InlineObject1
type InlineObject1 struct {
	// 交互APPID
	Appid string `json:"appid"`
	// 目的地址
	To string `json:"to"`
	// NFT的ID
	Id string `json:"id"`
	// 转移的数量
	Amount string `json:"amount"`
}