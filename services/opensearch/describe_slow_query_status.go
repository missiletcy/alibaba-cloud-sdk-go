package opensearch

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

// DescribeSlowQueryStatus invokes the opensearch.DescribeSlowQueryStatus API synchronously
// api document: https://help.aliyun.com/api/opensearch/describeslowquerystatus.html
func (client *Client) DescribeSlowQueryStatus(request *DescribeSlowQueryStatusRequest) (response *DescribeSlowQueryStatusResponse, err error) {
	response = CreateDescribeSlowQueryStatusResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSlowQueryStatusWithChan invokes the opensearch.DescribeSlowQueryStatus API asynchronously
// api document: https://help.aliyun.com/api/opensearch/describeslowquerystatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSlowQueryStatusWithChan(request *DescribeSlowQueryStatusRequest) (<-chan *DescribeSlowQueryStatusResponse, <-chan error) {
	responseChan := make(chan *DescribeSlowQueryStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSlowQueryStatus(request)
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

// DescribeSlowQueryStatusWithCallback invokes the opensearch.DescribeSlowQueryStatus API asynchronously
// api document: https://help.aliyun.com/api/opensearch/describeslowquerystatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSlowQueryStatusWithCallback(request *DescribeSlowQueryStatusRequest, callback func(response *DescribeSlowQueryStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSlowQueryStatusResponse
		var err error
		defer close(result)
		response, err = client.DescribeSlowQueryStatus(request)
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

// DescribeSlowQueryStatusRequest is the request struct for api DescribeSlowQueryStatus
type DescribeSlowQueryStatusRequest struct {
	*requests.RoaRequest
	AppGroupIdentity string `position:"Path" name:"appGroupIdentity"`
}

// DescribeSlowQueryStatusResponse is the response struct for api DescribeSlowQueryStatus
type DescribeSlowQueryStatusResponse struct {
	*responses.BaseResponse
	RequestId string `json:"requestId" xml:"requestId"`
	Result    Result `json:"result" xml:"result"`
}

// CreateDescribeSlowQueryStatusRequest creates a request to invoke DescribeSlowQueryStatus API
func CreateDescribeSlowQueryStatusRequest() (request *DescribeSlowQueryStatusRequest) {
	request = &DescribeSlowQueryStatusRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "DescribeSlowQueryStatus", "/v4/openapi/app-groups/[appGroupIdentity]/optimizers/slow-query", "opensearch", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeSlowQueryStatusResponse creates a response to parse from DescribeSlowQueryStatus response
func CreateDescribeSlowQueryStatusResponse() (response *DescribeSlowQueryStatusResponse) {
	response = &DescribeSlowQueryStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
