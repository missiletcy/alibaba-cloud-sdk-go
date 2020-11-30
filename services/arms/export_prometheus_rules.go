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

// ExportPrometheusRules invokes the arms.ExportPrometheusRules API synchronously
func (client *Client) ExportPrometheusRules(request *ExportPrometheusRulesRequest) (response *ExportPrometheusRulesResponse, err error) {
	response = CreateExportPrometheusRulesResponse()
	err = client.DoAction(request, response)
	return
}

// ExportPrometheusRulesWithChan invokes the arms.ExportPrometheusRules API asynchronously
func (client *Client) ExportPrometheusRulesWithChan(request *ExportPrometheusRulesRequest) (<-chan *ExportPrometheusRulesResponse, <-chan error) {
	responseChan := make(chan *ExportPrometheusRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ExportPrometheusRules(request)
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

// ExportPrometheusRulesWithCallback invokes the arms.ExportPrometheusRules API asynchronously
func (client *Client) ExportPrometheusRulesWithCallback(request *ExportPrometheusRulesRequest, callback func(response *ExportPrometheusRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ExportPrometheusRulesResponse
		var err error
		defer close(result)
		response, err = client.ExportPrometheusRules(request)
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

// ExportPrometheusRulesRequest is the request struct for api ExportPrometheusRules
type ExportPrometheusRulesRequest struct {
	*requests.RpcRequest
	NameSpace string `position:"Query" name:"NameSpace"`
	Name      string `position:"Query" name:"Name"`
	ClusterId string `position:"Query" name:"ClusterId"`
}

// ExportPrometheusRulesResponse is the response struct for api ExportPrometheusRules
type ExportPrometheusRulesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateExportPrometheusRulesRequest creates a request to invoke ExportPrometheusRules API
func CreateExportPrometheusRulesRequest() (request *ExportPrometheusRulesRequest) {
	request = &ExportPrometheusRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "ExportPrometheusRules", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateExportPrometheusRulesResponse creates a response to parse from ExportPrometheusRules response
func CreateExportPrometheusRulesResponse() (response *ExportPrometheusRulesResponse) {
	response = &ExportPrometheusRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}