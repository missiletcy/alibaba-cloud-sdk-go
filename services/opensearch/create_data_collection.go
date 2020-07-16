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

// CreateDataCollection invokes the opensearch.CreateDataCollection API synchronously
// api document: https://help.aliyun.com/api/opensearch/createdatacollection.html
func (client *Client) CreateDataCollection(request *CreateDataCollectionRequest) (response *CreateDataCollectionResponse, err error) {
	response = CreateCreateDataCollectionResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDataCollectionWithChan invokes the opensearch.CreateDataCollection API asynchronously
// api document: https://help.aliyun.com/api/opensearch/createdatacollection.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateDataCollectionWithChan(request *CreateDataCollectionRequest) (<-chan *CreateDataCollectionResponse, <-chan error) {
	responseChan := make(chan *CreateDataCollectionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDataCollection(request)
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

// CreateDataCollectionWithCallback invokes the opensearch.CreateDataCollection API asynchronously
// api document: https://help.aliyun.com/api/opensearch/createdatacollection.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateDataCollectionWithCallback(request *CreateDataCollectionRequest, callback func(response *CreateDataCollectionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDataCollectionResponse
		var err error
		defer close(result)
		response, err = client.CreateDataCollection(request)
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

// CreateDataCollectionRequest is the request struct for api CreateDataCollection
type CreateDataCollectionRequest struct {
	*requests.RoaRequest
	AppGroupIdentity string `position:"Path" name:"appGroupIdentity"`
}

// CreateDataCollectionResponse is the response struct for api CreateDataCollection
type CreateDataCollectionResponse struct {
	*responses.BaseResponse
	RequestId string                       `json:"requestId" xml:"requestId"`
	Result    ResultInCreateDataCollection `json:"result" xml:"result"`
}

// CreateCreateDataCollectionRequest creates a request to invoke CreateDataCollection API
func CreateCreateDataCollectionRequest() (request *CreateDataCollectionRequest) {
	request = &CreateDataCollectionRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "CreateDataCollection", "/v4/openapi/app-groups/[appGroupIdentity]/data-collections", "opensearch", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateDataCollectionResponse creates a response to parse from CreateDataCollection response
func CreateCreateDataCollectionResponse() (response *CreateDataCollectionResponse) {
	response = &CreateDataCollectionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
