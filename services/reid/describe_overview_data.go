package reid

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

// DescribeOverviewData invokes the reid.DescribeOverviewData API synchronously
// api document: https://help.aliyun.com/api/reid/describeoverviewdata.html
func (client *Client) DescribeOverviewData(request *DescribeOverviewDataRequest) (response *DescribeOverviewDataResponse, err error) {
	response = CreateDescribeOverviewDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeOverviewDataWithChan invokes the reid.DescribeOverviewData API asynchronously
// api document: https://help.aliyun.com/api/reid/describeoverviewdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeOverviewDataWithChan(request *DescribeOverviewDataRequest) (<-chan *DescribeOverviewDataResponse, <-chan error) {
	responseChan := make(chan *DescribeOverviewDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeOverviewData(request)
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

// DescribeOverviewDataWithCallback invokes the reid.DescribeOverviewData API asynchronously
// api document: https://help.aliyun.com/api/reid/describeoverviewdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeOverviewDataWithCallback(request *DescribeOverviewDataRequest, callback func(response *DescribeOverviewDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeOverviewDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeOverviewData(request)
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

// DescribeOverviewDataRequest is the request struct for api DescribeOverviewData
type DescribeOverviewDataRequest struct {
	*requests.RpcRequest
	Date     string `position:"Body" name:"Date"`
	StoreIds string `position:"Body" name:"StoreIds"`
}

// DescribeOverviewDataResponse is the response struct for api DescribeOverviewData
type DescribeOverviewDataResponse struct {
	*responses.BaseResponse
	ErrorCode              string                 `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage           string                 `json:"ErrorMessage" xml:"ErrorMessage"`
	Message                string                 `json:"Message" xml:"Message"`
	Code                   string                 `json:"Code" xml:"Code"`
	DynamicCode            string                 `json:"DynamicCode" xml:"DynamicCode"`
	RequestId              string                 `json:"RequestId" xml:"RequestId"`
	Success                bool                   `json:"Success" xml:"Success"`
	DynamicMessage         string                 `json:"DynamicMessage" xml:"DynamicMessage"`
	OverviewDetail         OverviewDetail         `json:"OverviewDetail" xml:"OverviewDetail"`
	AccurateOverviewDetail AccurateOverviewDetail `json:"AccurateOverviewDetail" xml:"AccurateOverviewDetail"`
}

// CreateDescribeOverviewDataRequest creates a request to invoke DescribeOverviewData API
func CreateDescribeOverviewDataRequest() (request *DescribeOverviewDataRequest) {
	request = &DescribeOverviewDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("reid", "2019-09-28", "DescribeOverviewData", "", "")
	return
}

// CreateDescribeOverviewDataResponse creates a response to parse from DescribeOverviewData response
func CreateDescribeOverviewDataResponse() (response *DescribeOverviewDataResponse) {
	response = &DescribeOverviewDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
