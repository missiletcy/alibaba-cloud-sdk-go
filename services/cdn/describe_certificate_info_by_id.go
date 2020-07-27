package cdn

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

// DescribeCertificateInfoByID invokes the cdn.DescribeCertificateInfoByID API synchronously
// api document: https://help.aliyun.com/api/cdn/describecertificateinfobyid.html
func (client *Client) DescribeCertificateInfoByID(request *DescribeCertificateInfoByIDRequest) (response *DescribeCertificateInfoByIDResponse, err error) {
	response = CreateDescribeCertificateInfoByIDResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCertificateInfoByIDWithChan invokes the cdn.DescribeCertificateInfoByID API asynchronously
// api document: https://help.aliyun.com/api/cdn/describecertificateinfobyid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCertificateInfoByIDWithChan(request *DescribeCertificateInfoByIDRequest) (<-chan *DescribeCertificateInfoByIDResponse, <-chan error) {
	responseChan := make(chan *DescribeCertificateInfoByIDResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCertificateInfoByID(request)
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

// DescribeCertificateInfoByIDWithCallback invokes the cdn.DescribeCertificateInfoByID API asynchronously
// api document: https://help.aliyun.com/api/cdn/describecertificateinfobyid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCertificateInfoByIDWithCallback(request *DescribeCertificateInfoByIDRequest, callback func(response *DescribeCertificateInfoByIDResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCertificateInfoByIDResponse
		var err error
		defer close(result)
		response, err = client.DescribeCertificateInfoByID(request)
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

// DescribeCertificateInfoByIDRequest is the request struct for api DescribeCertificateInfoByID
type DescribeCertificateInfoByIDRequest struct {
	*requests.RpcRequest
	CertId  string           `position:"Query" name:"CertId"`
	OwnerId requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeCertificateInfoByIDResponse is the response struct for api DescribeCertificateInfoByID
type DescribeCertificateInfoByIDResponse struct {
	*responses.BaseResponse
	RequestId string                                 `json:"RequestId" xml:"RequestId"`
	CertInfos CertInfosInDescribeCertificateInfoByID `json:"CertInfos" xml:"CertInfos"`
}

// CreateDescribeCertificateInfoByIDRequest creates a request to invoke DescribeCertificateInfoByID API
func CreateDescribeCertificateInfoByIDRequest() (request *DescribeCertificateInfoByIDRequest) {
	request = &DescribeCertificateInfoByIDRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DescribeCertificateInfoByID", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribeCertificateInfoByIDResponse creates a response to parse from DescribeCertificateInfoByID response
func CreateDescribeCertificateInfoByIDResponse() (response *DescribeCertificateInfoByIDResponse) {
	response = &DescribeCertificateInfoByIDResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
