package ehpc

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

// GetInstanceReport invokes the ehpc.GetInstanceReport API synchronously
// api document: https://help.aliyun.com/api/ehpc/getinstancereport.html
func (client *Client) GetInstanceReport(request *GetInstanceReportRequest) (response *GetInstanceReportResponse, err error) {
	response = CreateGetInstanceReportResponse()
	err = client.DoAction(request, response)
	return
}

// GetInstanceReportWithChan invokes the ehpc.GetInstanceReport API asynchronously
// api document: https://help.aliyun.com/api/ehpc/getinstancereport.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetInstanceReportWithChan(request *GetInstanceReportRequest) (<-chan *GetInstanceReportResponse, <-chan error) {
	responseChan := make(chan *GetInstanceReportResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetInstanceReport(request)
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

// GetInstanceReportWithCallback invokes the ehpc.GetInstanceReport API asynchronously
// api document: https://help.aliyun.com/api/ehpc/getinstancereport.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetInstanceReportWithCallback(request *GetInstanceReportRequest, callback func(response *GetInstanceReportResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetInstanceReportResponse
		var err error
		defer close(result)
		response, err = client.GetInstanceReport(request)
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

// GetInstanceReportRequest is the request struct for api GetInstanceReport
type GetInstanceReportRequest struct {
	*requests.RpcRequest
	EndTime     requests.Integer `position:"Query" name:"EndTime"`
	FilterValue string           `position:"Query" name:"FilterValue"`
	ClusterId   string           `position:"Query" name:"ClusterId"`
	StartTime   requests.Integer `position:"Query" name:"StartTime"`
}

// GetInstanceReportResponse is the response struct for api GetInstanceReport
type GetInstanceReportResponse struct {
	*responses.BaseResponse
	RequestId string                  `json:"RequestId" xml:"RequestId"`
	Metrics   string                  `json:"Metrics" xml:"Metrics"`
	Data      DataInGetInstanceReport `json:"Data" xml:"Data"`
}

// CreateGetInstanceReportRequest creates a request to invoke GetInstanceReport API
func CreateGetInstanceReportRequest() (request *GetInstanceReportRequest) {
	request = &GetInstanceReportRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "GetInstanceReport", "ehs", "openAPI")
	return
}

// CreateGetInstanceReportResponse creates a response to parse from GetInstanceReport response
func CreateGetInstanceReportResponse() (response *GetInstanceReportResponse) {
	response = &GetInstanceReportResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
