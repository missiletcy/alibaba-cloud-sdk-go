package ehpc

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

// ListClustersMeta invokes the ehpc.ListClustersMeta API synchronously
// api document: https://help.aliyun.com/api/ehpc/listclustersmeta.html
func (client *Client) ListClustersMeta(request *ListClustersMetaRequest) (response *ListClustersMetaResponse, err error) {
	response = CreateListClustersMetaResponse()
	err = client.DoAction(request, response)
	return
}

// ListClustersMetaWithChan invokes the ehpc.ListClustersMeta API asynchronously
// api document: https://help.aliyun.com/api/ehpc/listclustersmeta.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListClustersMetaWithChan(request *ListClustersMetaRequest) (<-chan *ListClustersMetaResponse, <-chan error) {
	responseChan := make(chan *ListClustersMetaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListClustersMeta(request)
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

// ListClustersMetaWithCallback invokes the ehpc.ListClustersMeta API asynchronously
// api document: https://help.aliyun.com/api/ehpc/listclustersmeta.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListClustersMetaWithCallback(request *ListClustersMetaRequest, callback func(response *ListClustersMetaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListClustersMetaResponse
		var err error
		defer close(result)
		response, err = client.ListClustersMeta(request)
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

// ListClustersMetaRequest is the request struct for api ListClustersMeta
type ListClustersMetaRequest struct {
	*requests.RpcRequest
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
}

// ListClustersMetaResponse is the response struct for api ListClustersMeta
type ListClustersMetaResponse struct {
	*responses.BaseResponse
	RequestId  string                     `json:"RequestId" xml:"RequestId"`
	TotalCount int                        `json:"TotalCount" xml:"TotalCount"`
	PageNumber int                        `json:"PageNumber" xml:"PageNumber"`
	PageSize   int                        `json:"PageSize" xml:"PageSize"`
	Clusters   ClustersInListClustersMeta `json:"Clusters" xml:"Clusters"`
}

// CreateListClustersMetaRequest creates a request to invoke ListClustersMeta API
func CreateListClustersMetaRequest() (request *ListClustersMetaRequest) {
	request = &ListClustersMetaRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "ListClustersMeta", "ehs", "openAPI")
	return
}

// CreateListClustersMetaResponse creates a response to parse from ListClustersMeta response
func CreateListClustersMetaResponse() (response *ListClustersMetaResponse) {
	response = &ListClustersMetaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
