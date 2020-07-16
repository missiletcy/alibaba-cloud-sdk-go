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

// UpdateSummaries invokes the opensearch.UpdateSummaries API synchronously
// api document: https://help.aliyun.com/api/opensearch/updatesummaries.html
func (client *Client) UpdateSummaries(request *UpdateSummariesRequest) (response *UpdateSummariesResponse, err error) {
	response = CreateUpdateSummariesResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateSummariesWithChan invokes the opensearch.UpdateSummaries API asynchronously
// api document: https://help.aliyun.com/api/opensearch/updatesummaries.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateSummariesWithChan(request *UpdateSummariesRequest) (<-chan *UpdateSummariesResponse, <-chan error) {
	responseChan := make(chan *UpdateSummariesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateSummaries(request)
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

// UpdateSummariesWithCallback invokes the opensearch.UpdateSummaries API asynchronously
// api document: https://help.aliyun.com/api/opensearch/updatesummaries.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateSummariesWithCallback(request *UpdateSummariesRequest, callback func(response *UpdateSummariesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateSummariesResponse
		var err error
		defer close(result)
		response, err = client.UpdateSummaries(request)
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

// UpdateSummariesRequest is the request struct for api UpdateSummaries
type UpdateSummariesRequest struct {
	*requests.RoaRequest
	DryRun           requests.Boolean `position:"Query" name:"dryRun"`
	AppId            requests.Integer `position:"Path" name:"appId"`
	AppGroupIdentity string           `position:"Path" name:"appGroupIdentity"`
}

// UpdateSummariesResponse is the response struct for api UpdateSummaries
type UpdateSummariesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"requestId" xml:"requestId"`
	Result    bool   `json:"result" xml:"result"`
}

// CreateUpdateSummariesRequest creates a request to invoke UpdateSummaries API
func CreateUpdateSummariesRequest() (request *UpdateSummariesRequest) {
	request = &UpdateSummariesRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "UpdateSummaries", "/v4/openapi/app-groups/[appGroupIdentity]/apps/[appId]/summaries", "opensearch", "openAPI")
	request.Method = requests.PUT
	return
}

// CreateUpdateSummariesResponse creates a response to parse from UpdateSummaries response
func CreateUpdateSummariesResponse() (response *UpdateSummariesResponse) {
	response = &UpdateSummariesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
