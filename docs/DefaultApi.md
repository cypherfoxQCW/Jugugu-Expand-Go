# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FriendLoginPost**](DefaultApi.md#FriendLoginPost) | **Post** /FriendLogin | 4.Cookie登录
[**JuguguGetPhoneCodePost**](DefaultApi.md#JuguguGetPhoneCodePost) | **Post** /Jugugu_GetPhoneCode | 2.获取登录短信验证码
[**JugugugGetReleaseNFT1155PNGCodePost**](DefaultApi.md#JugugugGetReleaseNFT1155PNGCodePost) | **Post** /Jugugug_GetReleaseNFT1155_PNGCode | 1.获得验证码图片
[**TransferNFTPost**](DefaultApi.md#TransferNFTPost) | **Post** /TransferNFT | 转移NFT
[**UserNFTsPost**](DefaultApi.md#UserNFTsPost) | **Post** /UserNFTs | 查询用户NFT



## FriendLoginPost

> InlineResponse2002 FriendLoginPost(ctx, optional)

4.Cookie登录

Cookie无感登录jugugu获取用户信息。使用CookieToken和手机号登录菊咕咕获取信息，安全考虑该方法不会获得操作等级高的token。cookietoken只有有限的操作权限。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FriendLoginPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FriendLoginPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject4** | [**optional.Interface of InlineObject4**](InlineObject4.md)|  | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JuguguGetPhoneCodePost

> InlineResponse2002 JuguguGetPhoneCodePost(ctx, optional)

2.获取登录短信验证码

获取登录短信验证码，验证码重复发送周期为60秒。验证码有效时间为3分钟

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***JuguguGetPhoneCodePostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a JuguguGetPhoneCodePostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject2** | [**optional.Interface of InlineObject2**](InlineObject2.md)|  | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JugugugGetReleaseNFT1155PNGCodePost

> InlineResponse2003 JugugugGetReleaseNFT1155PNGCodePost(ctx, optional)

1.获得验证码图片

获取验证码图片，用于防止机器人。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***JugugugGetReleaseNFT1155PNGCodePostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a JugugugGetReleaseNFT1155PNGCodePostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject3** | [**optional.Interface of InlineObject3**](InlineObject3.md)|  | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransferNFTPost

> InlineResponse2001 TransferNFTPost(ctx, optional)

转移NFT

转移指定ID的NFT，指定数量Amount，至指定用户区块链地址

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TransferNFTPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TransferNFTPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject1** | [**optional.Interface of InlineObject1**](InlineObject1.md)|  | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserNFTsPost

> InlineResponse200 UserNFTsPost(ctx, optional)

查询用户NFT

查询指定合约用户徽章持有ID和对应的数量。

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserNFTsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UserNFTsPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject** | [**optional.Interface of InlineObject**](InlineObject.md)|  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

