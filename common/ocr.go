package common

import (
	"encoding/json"
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	ocr_api20210707 "github.com/alibabacloud-go/ocr-api-20210707/v2/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"os"
	"strings"
)

type ApiResponse struct {
	Body struct {
		Data struct {
			Content string `json:"content"`
		} `json:"data"`
	} `json:"body"`
}

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *ocr_api20210707.Client, _err error) {
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// Endpoint 请参考 https://api.aliyun.com/product/ocr-api
	config.Endpoint = tea.String("ocr-api.cn-hangzhou.aliyuncs.com")
	_result = &ocr_api20210707.Client{}
	_result, _err = ocr_api20210707.NewClient(config)
	return _result, _err
}

func Ocr(path string) (string, error) {
	// 请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_ID 和 ALIBABA_CLOUD_ACCESS_KEY_SECRET。
	// 工程代码泄露可能会导致 AccessKey 泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考，建议使用更安全的 STS 方式，更多鉴权访问方式请参见：https://help.aliyun.com/document_detail/378661.html
	client, _err := CreateClient(tea.String(os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_ID")), tea.String(os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET")))
	if _err != nil {
		return "", _err
	}

	// 需要安装额外的依赖库，直接点击下载完整工程即可看到所有依赖。
	bodyStream, _ := os.Open(path)
	recognizeAllTextRequest := &ocr_api20210707.RecognizeAllTextRequest{
		Type:         tea.String("Advanced"),
		OutputFigure: tea.Bool(false),
		Body:         bodyStream,
	}
	runtime := &util.RuntimeOptions{}
	resp, tryErr := func() (_resp string, _e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		resp, _err := client.RecognizeAllTextWithOptions(recognizeAllTextRequest, runtime)
		if _err != nil {
			return "", _err
		}

		//console.Log(util.ToJSONString(resp))

		// 解析JSON字符串
		var apiResp ApiResponse
		jsonstr := *util.ToJSONString(resp)
		err := json.Unmarshal([]byte(jsonstr), &apiResp)
		if err != nil {
			return "", fmt.Errorf("JSON unmarshal error: %v", err)
		}

		//fmt.Println(apiResp.Body.Data.Content)
		_resp = apiResp.Body.Data.Content

		return _resp, nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		// 错误 message
		fmt.Println(tea.StringValue(error.Message))
		// 诊断地址
		var data interface{}
		d := json.NewDecoder(strings.NewReader(tea.StringValue(error.Data)))
		d.Decode(&data)
		if m, ok := data.(map[string]interface{}); ok {
			recommend, _ := m["Recommend"]
			fmt.Println(recommend)
		}
		_, _err = util.AssertAsString(error.Message)
		if _err != nil {
			return "", _err
		}
	}
	return resp, _err
}
