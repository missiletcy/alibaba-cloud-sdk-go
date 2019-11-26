package rds

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

// CheckAccountNameAvailable invokes the rds.CheckAccountNameAvailable API synchronously
// api document: https://help.aliyun.com/api/rds/checkaccountnameavailable.html
func (client *Client) CheckAccountNameAvailable(request *CheckAccountNameAvailableRequest) (response *CheckAccountNameAvailableResponse, err error) {
	response = CreateCheckAccountNameAvailableResponse()
	err = client.DoAction(request, response)
	return
}

// CheckAccountNameAvailableWithChan invokes the rds.CheckAccountNameAvailable API asynchronously
// api document: https://help.aliyun.com/api/rds/checkaccountnameavailable.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckAccountNameAvailableWithChan(request *CheckAccountNameAvailableRequest) (<-chan *CheckAccountNameAvailableResponse, <-chan error) {
	responseChan := make(chan *CheckAccountNameAvailableResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckAccountNameAvailable(request)
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

// CheckAccountNameAvailableWithCallback invokes the rds.CheckAccountNameAvailable API asynchronously
// api document: https://help.aliyun.com/api/rds/checkaccountnameavailable.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckAccountNameAvailableWithCallback(request *CheckAccountNameAvailableRequest, callback func(response *CheckAccountNameAvailableResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckAccountNameAvailableResponse
		var err error
		defer close(result)
		response, err = client.CheckAccountNameAvailable(request)
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

// CheckAccountNameAvailableRequest is the request struct for api CheckAccountNameAvailable
type CheckAccountNameAvailableRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	AccountName          string           `position:"Query" name:"AccountName"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// CheckAccountNameAvailableResponse is the response struct for api CheckAccountNameAvailable
type CheckAccountNameAvailableResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCheckAccountNameAvailableRequest creates a request to invoke CheckAccountNameAvailable API
func CreateCheckAccountNameAvailableRequest() (request *CheckAccountNameAvailableRequest) {
	request = &CheckAccountNameAvailableRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "CheckAccountNameAvailable", "", "")
	return
}

// CreateCheckAccountNameAvailableResponse creates a response to parse from CheckAccountNameAvailable response
func CreateCheckAccountNameAvailableResponse() (response *CheckAccountNameAvailableResponse) {
	response = &CheckAccountNameAvailableResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
