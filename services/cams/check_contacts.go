package cams

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

// CheckContacts invokes the cams.CheckContacts API synchronously
func (client *Client) CheckContacts(request *CheckContactsRequest) (response *CheckContactsResponse, err error) {
	response = CreateCheckContactsResponse()
	err = client.DoAction(request, response)
	return
}

// CheckContactsWithChan invokes the cams.CheckContacts API asynchronously
func (client *Client) CheckContactsWithChan(request *CheckContactsRequest) (<-chan *CheckContactsResponse, <-chan error) {
	responseChan := make(chan *CheckContactsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckContacts(request)
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

// CheckContactsWithCallback invokes the cams.CheckContacts API asynchronously
func (client *Client) CheckContactsWithCallback(request *CheckContactsRequest, callback func(response *CheckContactsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckContactsResponse
		var err error
		defer close(result)
		response, err = client.CheckContacts(request)
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

// CheckContactsRequest is the request struct for api CheckContacts
type CheckContactsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ChannelType          string           `position:"Body" name:"ChannelType"`
	From                 string           `position:"Body" name:"From"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Contacts             string           `position:"Body" name:"Contacts"`
}

// CheckContactsResponse is the response struct for api CheckContacts
type CheckContactsResponse struct {
	*responses.BaseResponse
	RequestId     string          `json:"RequestId" xml:"RequestId"`
	ResultCode    string          `json:"ResultCode" xml:"ResultCode"`
	ResultMessage string          `json:"ResultMessage" xml:"ResultMessage"`
	Contacts      []ContactStatus `json:"Contacts" xml:"Contacts"`
}

// CreateCheckContactsRequest creates a request to invoke CheckContacts API
func CreateCheckContactsRequest() (request *CheckContactsRequest) {
	request = &CheckContactsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cams", "2020-06-06", "CheckContacts", "cams", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCheckContactsResponse creates a response to parse from CheckContacts response
func CreateCheckContactsResponse() (response *CheckContactsResponse) {
	response = &CheckContactsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
