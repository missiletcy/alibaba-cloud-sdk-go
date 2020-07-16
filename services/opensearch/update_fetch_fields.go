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

// UpdateFetchFields invokes the opensearch.UpdateFetchFields API synchronously
// api document: https://help.aliyun.com/api/opensearch/updatefetchfields.html
func (client *Client) UpdateFetchFields(request *UpdateFetchFieldsRequest) (response *UpdateFetchFieldsResponse, err error) {
	response = CreateUpdateFetchFieldsResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateFetchFieldsWithChan invokes the opensearch.UpdateFetchFields API asynchronously
// api document: https://help.aliyun.com/api/opensearch/updatefetchfields.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateFetchFieldsWithChan(request *UpdateFetchFieldsRequest) (<-chan *UpdateFetchFieldsResponse, <-chan error) {
	responseChan := make(chan *UpdateFetchFieldsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateFetchFields(request)
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

// UpdateFetchFieldsWithCallback invokes the opensearch.UpdateFetchFields API asynchronously
// api document: https://help.aliyun.com/api/opensearch/updatefetchfields.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateFetchFieldsWithCallback(request *UpdateFetchFieldsRequest, callback func(response *UpdateFetchFieldsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateFetchFieldsResponse
		var err error
		defer close(result)
		response, err = client.UpdateFetchFields(request)
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

// UpdateFetchFieldsRequest is the request struct for api UpdateFetchFields
type UpdateFetchFieldsRequest struct {
	*requests.RoaRequest
	DryRun           requests.Boolean `position:"Query" name:"dryRun"`
	AppId            requests.Integer `position:"Path" name:"appId"`
	AppGroupIdentity string           `position:"Path" name:"appGroupIdentity"`
}

// UpdateFetchFieldsResponse is the response struct for api UpdateFetchFields
type UpdateFetchFieldsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"requestId" xml:"requestId"`
	Result    bool   `json:"result" xml:"result"`
}

// CreateUpdateFetchFieldsRequest creates a request to invoke UpdateFetchFields API
func CreateUpdateFetchFieldsRequest() (request *UpdateFetchFieldsRequest) {
	request = &UpdateFetchFieldsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "UpdateFetchFields", "/v4/openapi/app-groups/[appGroupIdentity]/apps/[appId]/fetch-fields", "opensearch", "openAPI")
	request.Method = requests.PUT
	return
}

// CreateUpdateFetchFieldsResponse creates a response to parse from UpdateFetchFields response
func CreateUpdateFetchFieldsResponse() (response *UpdateFetchFieldsResponse) {
	response = &UpdateFetchFieldsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
