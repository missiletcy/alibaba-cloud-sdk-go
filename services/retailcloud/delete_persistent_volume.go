package retailcloud

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

// DeletePersistentVolume invokes the retailcloud.DeletePersistentVolume API synchronously
// api document: https://help.aliyun.com/api/retailcloud/deletepersistentvolume.html
func (client *Client) DeletePersistentVolume(request *DeletePersistentVolumeRequest) (response *DeletePersistentVolumeResponse, err error) {
	response = CreateDeletePersistentVolumeResponse()
	err = client.DoAction(request, response)
	return
}

// DeletePersistentVolumeWithChan invokes the retailcloud.DeletePersistentVolume API asynchronously
// api document: https://help.aliyun.com/api/retailcloud/deletepersistentvolume.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeletePersistentVolumeWithChan(request *DeletePersistentVolumeRequest) (<-chan *DeletePersistentVolumeResponse, <-chan error) {
	responseChan := make(chan *DeletePersistentVolumeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeletePersistentVolume(request)
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

// DeletePersistentVolumeWithCallback invokes the retailcloud.DeletePersistentVolume API asynchronously
// api document: https://help.aliyun.com/api/retailcloud/deletepersistentvolume.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeletePersistentVolumeWithCallback(request *DeletePersistentVolumeRequest, callback func(response *DeletePersistentVolumeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeletePersistentVolumeResponse
		var err error
		defer close(result)
		response, err = client.DeletePersistentVolume(request)
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

// DeletePersistentVolumeRequest is the request struct for api DeletePersistentVolume
type DeletePersistentVolumeRequest struct {
	*requests.RpcRequest
	PersistentVolumeName string `position:"Body" name:"PersistentVolumeName"`
	ClusterInstanceId    string `position:"Body" name:"ClusterInstanceId"`
}

// DeletePersistentVolumeResponse is the response struct for api DeletePersistentVolume
type DeletePersistentVolumeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      int    `json:"Code" xml:"Code"`
	ErrMsg    string `json:"ErrMsg" xml:"ErrMsg"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateDeletePersistentVolumeRequest creates a request to invoke DeletePersistentVolume API
func CreateDeletePersistentVolumeRequest() (request *DeletePersistentVolumeRequest) {
	request = &DeletePersistentVolumeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailcloud", "2018-03-13", "DeletePersistentVolume", "retailcloud", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeletePersistentVolumeResponse creates a response to parse from DeletePersistentVolume response
func CreateDeletePersistentVolumeResponse() (response *DeletePersistentVolumeResponse) {
	response = &DeletePersistentVolumeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
