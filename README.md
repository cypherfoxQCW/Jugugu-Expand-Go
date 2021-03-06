# Go API client for openapi

接入jugugu区块链全包的全部接口，访问IP需要连续管理员，添加业务服务器IP

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:

```golang
import "./openapi"
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**FriendLoginPost**](docs/DefaultApi.md#friendloginpost) | **Post** /FriendLogin | 4.Cookie登录
*DefaultApi* | [**JuguguGetPhoneCodePost**](docs/DefaultApi.md#jugugugetphonecodepost) | **Post** /Jugugu_GetPhoneCode | 2.获取登录短信验证码
*DefaultApi* | [**JugugugGetReleaseNFT1155PNGCodePost**](docs/DefaultApi.md#juguguggetreleasenft1155pngcodepost) | **Post** /Jugugug_GetReleaseNFT1155_PNGCode | 1.获得验证码图片
*DefaultApi* | [**TransferNFTPost**](docs/DefaultApi.md#transfernftpost) | **Post** /TransferNFT | 转移NFT
*DefaultApi* | [**UserNFTsPost**](docs/DefaultApi.md#usernftspost) | **Post** /UserNFTs | 查询用户NFT


## Documentation For Models

 - [InlineObject](docs/InlineObject.md)
 - [InlineObject1](docs/InlineObject1.md)
 - [InlineObject2](docs/InlineObject2.md)
 - [InlineObject3](docs/InlineObject3.md)
 - [InlineObject4](docs/InlineObject4.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [InlineResponse2001](docs/InlineResponse2001.md)
 - [InlineResponse2002](docs/InlineResponse2002.md)
 - [InlineResponse2003](docs/InlineResponse2003.md)
 - [InlineResponse200Nfts](docs/InlineResponse200Nfts.md)


## Documentation For Authorization

 Endpoints do not require authorization.



## Author



