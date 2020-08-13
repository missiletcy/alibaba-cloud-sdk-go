package devops_rdc

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

// GetPipelineInstanceBuildNumberStatus invokes the devops_rdc.GetPipelineInstanceBuildNumberStatus API synchronously
// api document: https://help.aliyun.com/api/devops-rdc/getpipelineinstancebuildnumberstatus.html
func (client *Client) GetPipelineInstanceBuildNumberStatus(request *GetPipelineInstanceBuildNumberStatusRequest) (response *GetPipelineInstanceBuildNumberStatusResponse, err error) {
	response = CreateGetPipelineInstanceBuildNumberStatusResponse()
	err = client.DoAction(request, response)
	return
}

// GetPipelineInstanceBuildNumberStatusWithChan invokes the devops_rdc.GetPipelineInstanceBuildNumberStatus API asynchronously
// api document: https://help.aliyun.com/api/devops-rdc/getpipelineinstancebuildnumberstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetPipelineInstanceBuildNumberStatusWithChan(request *GetPipelineInstanceBuildNumberStatusRequest) (<-chan *GetPipelineInstanceBuildNumberStatusResponse, <-chan error) {
	responseChan := make(chan *GetPipelineInstanceBuildNumberStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetPipelineInstanceBuildNumberStatus(request)
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

// GetPipelineInstanceBuildNumberStatusWithCallback invokes the devops_rdc.GetPipelineInstanceBuildNumberStatus API asynchronously
// api document: https://help.aliyun.com/api/devops-rdc/getpipelineinstancebuildnumberstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetPipelineInstanceBuildNumberStatusWithCallback(request *GetPipelineInstanceBuildNumberStatusRequest, callback func(response *GetPipelineInstanceBuildNumberStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetPipelineInstanceBuildNumberStatusResponse
		var err error
		defer close(result)
		response, err = client.GetPipelineInstanceBuildNumberStatus(request)
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

// GetPipelineInstanceBuildNumberStatusRequest is the request struct for api GetPipelineInstanceBuildNumberStatus
type GetPipelineInstanceBuildNumberStatusRequest struct {
	*requests.RpcRequest
	BuildNum   requests.Integer `position:"Body" name:"BuildNum"`
	UserPk     string           `position:"Body" name:"UserPk"`
	OrgId      string           `position:"Query" name:"OrgId"`
	PipelineId requests.Integer `position:"Query" name:"PipelineId"`
}

// GetPipelineInstanceBuildNumberStatusResponse is the response struct for api GetPipelineInstanceBuildNumberStatus
type GetPipelineInstanceBuildNumberStatusResponse struct {
	*responses.BaseResponse
	Success      bool   `json:"Success" xml:"Success"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Object       Object `json:"Object" xml:"Object"`
}

// CreateGetPipelineInstanceBuildNumberStatusRequest creates a request to invoke GetPipelineInstanceBuildNumberStatus API
func CreateGetPipelineInstanceBuildNumberStatusRequest() (request *GetPipelineInstanceBuildNumberStatusRequest) {
	request = &GetPipelineInstanceBuildNumberStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("devops-rdc", "2020-03-03", "GetPipelineInstanceBuildNumberStatus", "", "")
	request.Method = requests.POST
	return
}

// CreateGetPipelineInstanceBuildNumberStatusResponse creates a response to parse from GetPipelineInstanceBuildNumberStatus response
func CreateGetPipelineInstanceBuildNumberStatusResponse() (response *GetPipelineInstanceBuildNumberStatusResponse) {
	response = &GetPipelineInstanceBuildNumberStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}