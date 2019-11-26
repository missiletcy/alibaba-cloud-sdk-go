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

// ModifyDBInstanceAutoUpgradeMinorVersion invokes the rds.ModifyDBInstanceAutoUpgradeMinorVersion API synchronously
// api document: https://help.aliyun.com/api/rds/modifydbinstanceautoupgrademinorversion.html
func (client *Client) ModifyDBInstanceAutoUpgradeMinorVersion(request *ModifyDBInstanceAutoUpgradeMinorVersionRequest) (response *ModifyDBInstanceAutoUpgradeMinorVersionResponse, err error) {
	response = CreateModifyDBInstanceAutoUpgradeMinorVersionResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDBInstanceAutoUpgradeMinorVersionWithChan invokes the rds.ModifyDBInstanceAutoUpgradeMinorVersion API asynchronously
// api document: https://help.aliyun.com/api/rds/modifydbinstanceautoupgrademinorversion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDBInstanceAutoUpgradeMinorVersionWithChan(request *ModifyDBInstanceAutoUpgradeMinorVersionRequest) (<-chan *ModifyDBInstanceAutoUpgradeMinorVersionResponse, <-chan error) {
	responseChan := make(chan *ModifyDBInstanceAutoUpgradeMinorVersionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDBInstanceAutoUpgradeMinorVersion(request)
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

// ModifyDBInstanceAutoUpgradeMinorVersionWithCallback invokes the rds.ModifyDBInstanceAutoUpgradeMinorVersion API asynchronously
// api document: https://help.aliyun.com/api/rds/modifydbinstanceautoupgrademinorversion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDBInstanceAutoUpgradeMinorVersionWithCallback(request *ModifyDBInstanceAutoUpgradeMinorVersionRequest, callback func(response *ModifyDBInstanceAutoUpgradeMinorVersionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDBInstanceAutoUpgradeMinorVersionResponse
		var err error
		defer close(result)
		response, err = client.ModifyDBInstanceAutoUpgradeMinorVersion(request)
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

// ModifyDBInstanceAutoUpgradeMinorVersionRequest is the request struct for api ModifyDBInstanceAutoUpgradeMinorVersion
type ModifyDBInstanceAutoUpgradeMinorVersionRequest struct {
	*requests.RpcRequest
	ResourceOwnerId         requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount    string           `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken             string           `position:"Query" name:"ClientToken"`
	OwnerId                 requests.Integer `position:"Query" name:"OwnerId"`
	AutoUpgradeMinorVersion string           `position:"Query" name:"AutoUpgradeMinorVersion"`
	DBInstanceId            string           `position:"Query" name:"DBInstanceId"`
}

// ModifyDBInstanceAutoUpgradeMinorVersionResponse is the response struct for api ModifyDBInstanceAutoUpgradeMinorVersion
type ModifyDBInstanceAutoUpgradeMinorVersionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDBInstanceAutoUpgradeMinorVersionRequest creates a request to invoke ModifyDBInstanceAutoUpgradeMinorVersion API
func CreateModifyDBInstanceAutoUpgradeMinorVersionRequest() (request *ModifyDBInstanceAutoUpgradeMinorVersionRequest) {
	request = &ModifyDBInstanceAutoUpgradeMinorVersionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "ModifyDBInstanceAutoUpgradeMinorVersion", "", "")
	return
}

// CreateModifyDBInstanceAutoUpgradeMinorVersionResponse creates a response to parse from ModifyDBInstanceAutoUpgradeMinorVersion response
func CreateModifyDBInstanceAutoUpgradeMinorVersionResponse() (response *ModifyDBInstanceAutoUpgradeMinorVersionResponse) {
	response = &ModifyDBInstanceAutoUpgradeMinorVersionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
