package arms

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

// SendCustomIncidents invokes the arms.SendCustomIncidents API synchronously
func (client *Client) SendCustomIncidents(request *SendCustomIncidentsRequest) (response *SendCustomIncidentsResponse, err error) {
	response = CreateSendCustomIncidentsResponse()
	err = client.DoAction(request, response)
	return
}

// SendCustomIncidentsWithChan invokes the arms.SendCustomIncidents API asynchronously
func (client *Client) SendCustomIncidentsWithChan(request *SendCustomIncidentsRequest) (<-chan *SendCustomIncidentsResponse, <-chan error) {
	responseChan := make(chan *SendCustomIncidentsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SendCustomIncidents(request)
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

// SendCustomIncidentsWithCallback invokes the arms.SendCustomIncidents API asynchronously
func (client *Client) SendCustomIncidentsWithCallback(request *SendCustomIncidentsRequest, callback func(response *SendCustomIncidentsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SendCustomIncidentsResponse
		var err error
		defer close(result)
		response, err = client.SendCustomIncidents(request)
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

// SendCustomIncidentsRequest is the request struct for api SendCustomIncidents
type SendCustomIncidentsRequest struct {
	*requests.RpcRequest
	Incidents   string `position:"Query" name:"Incidents"`
	ProxyUserId string `position:"Query" name:"ProxyUserId"`
	ProductType string `position:"Query" name:"ProductType"`
}

// SendCustomIncidentsResponse is the response struct for api SendCustomIncidents
type SendCustomIncidentsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateSendCustomIncidentsRequest creates a request to invoke SendCustomIncidents API
func CreateSendCustomIncidentsRequest() (request *SendCustomIncidentsRequest) {
	request = &SendCustomIncidentsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "SendCustomIncidents", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSendCustomIncidentsResponse creates a response to parse from SendCustomIncidents response
func CreateSendCustomIncidentsResponse() (response *SendCustomIncidentsResponse) {
	response = &SendCustomIncidentsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
