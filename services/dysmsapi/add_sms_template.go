package dysmsapi

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// AddSmsTemplate invokes the dysmsapi.AddSmsTemplate API synchronously
// api document: https://help.aliyun.com/api/dysmsapi/addsmstemplate.html
func (client *Client) AddSmsTemplate(request *AddSmsTemplateRequest) (response *AddSmsTemplateResponse, err error) {
	response = CreateAddSmsTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// AddSmsTemplateWithChan invokes the dysmsapi.AddSmsTemplate API asynchronously
// api document: https://help.aliyun.com/api/dysmsapi/addsmstemplate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddSmsTemplateWithChan(request *AddSmsTemplateRequest) (<-chan *AddSmsTemplateResponse, <-chan error) {
	responseChan := make(chan *AddSmsTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddSmsTemplate(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// AddSmsTemplateWithCallback invokes the dysmsapi.AddSmsTemplate API asynchronously
// api document: https://help.aliyun.com/api/dysmsapi/addsmstemplate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddSmsTemplateWithCallback(request *AddSmsTemplateRequest, callback func(response *AddSmsTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddSmsTemplateResponse
		var err error
		defer close(result)
		response, err = client.AddSmsTemplate(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// AddSmsTemplateRequest is the request struct for api AddSmsTemplate
type AddSmsTemplateRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Remark               string           `position:"Query" name:"Remark"`
	TemplateType         requests.Integer `position:"Query" name:"TemplateType"`
	TemplateName         string           `position:"Query" name:"TemplateName"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	TemplateContent      string           `position:"Query" name:"TemplateContent"`
}

// AddSmsTemplateResponse is the response struct for api AddSmsTemplate
type AddSmsTemplateResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	TemplateCode string `json:"TemplateCode" xml:"TemplateCode"`
	Code         string `json:"Code" xml:"Code"`
	Message      string `json:"Message" xml:"Message"`
}

// CreateAddSmsTemplateRequest creates a request to invoke AddSmsTemplate API
func CreateAddSmsTemplateRequest() (request *AddSmsTemplateRequest) {
	request = &AddSmsTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dysmsapi", "2017-05-25", "AddSmsTemplate", "dysmsapi", "openAPI")
	return
}

// CreateAddSmsTemplateResponse creates a response to parse from AddSmsTemplate response
func CreateAddSmsTemplateResponse() (response *AddSmsTemplateResponse) {
	response = &AddSmsTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
