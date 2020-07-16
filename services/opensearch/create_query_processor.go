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

// CreateQueryProcessor invokes the opensearch.CreateQueryProcessor API synchronously
// api document: https://help.aliyun.com/api/opensearch/createqueryprocessor.html
func (client *Client) CreateQueryProcessor(request *CreateQueryProcessorRequest) (response *CreateQueryProcessorResponse, err error) {
	response = CreateCreateQueryProcessorResponse()
	err = client.DoAction(request, response)
	return
}

// CreateQueryProcessorWithChan invokes the opensearch.CreateQueryProcessor API asynchronously
// api document: https://help.aliyun.com/api/opensearch/createqueryprocessor.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateQueryProcessorWithChan(request *CreateQueryProcessorRequest) (<-chan *CreateQueryProcessorResponse, <-chan error) {
	responseChan := make(chan *CreateQueryProcessorResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateQueryProcessor(request)
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

// CreateQueryProcessorWithCallback invokes the opensearch.CreateQueryProcessor API asynchronously
// api document: https://help.aliyun.com/api/opensearch/createqueryprocessor.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateQueryProcessorWithCallback(request *CreateQueryProcessorRequest, callback func(response *CreateQueryProcessorResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateQueryProcessorResponse
		var err error
		defer close(result)
		response, err = client.CreateQueryProcessor(request)
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

// CreateQueryProcessorRequest is the request struct for api CreateQueryProcessor
type CreateQueryProcessorRequest struct {
	*requests.RoaRequest
	DryRun           requests.Boolean `position:"Query" name:"dryRun"`
	AppId            requests.Integer `position:"Path" name:"appId"`
	AppGroupIdentity string           `position:"Path" name:"appGroupIdentity"`
}

// CreateQueryProcessorResponse is the response struct for api CreateQueryProcessor
type CreateQueryProcessorResponse struct {
	*responses.BaseResponse
	RequestId string `json:"requestId" xml:"requestId"`
	Result    Result `json:"result" xml:"result"`
}

// CreateCreateQueryProcessorRequest creates a request to invoke CreateQueryProcessor API
func CreateCreateQueryProcessorRequest() (request *CreateQueryProcessorRequest) {
	request = &CreateQueryProcessorRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "CreateQueryProcessor", "/v4/openapi/app-groups/[appGroupIdentity]/apps/[appId]/query-processors", "opensearch", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateQueryProcessorResponse creates a response to parse from CreateQueryProcessor response
func CreateCreateQueryProcessorResponse() (response *CreateQueryProcessorResponse) {
	response = &CreateQueryProcessorResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
