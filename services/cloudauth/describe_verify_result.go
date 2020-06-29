package cloudauth

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

// DescribeVerifyResult invokes the cloudauth.DescribeVerifyResult API synchronously
// api document: https://help.aliyun.com/api/cloudauth/describeverifyresult.html
func (client *Client) DescribeVerifyResult(request *DescribeVerifyResultRequest) (response *DescribeVerifyResultResponse, err error) {
	response = CreateDescribeVerifyResultResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVerifyResultWithChan invokes the cloudauth.DescribeVerifyResult API asynchronously
// api document: https://help.aliyun.com/api/cloudauth/describeverifyresult.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeVerifyResultWithChan(request *DescribeVerifyResultRequest) (<-chan *DescribeVerifyResultResponse, <-chan error) {
	responseChan := make(chan *DescribeVerifyResultResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVerifyResult(request)
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

// DescribeVerifyResultWithCallback invokes the cloudauth.DescribeVerifyResult API asynchronously
// api document: https://help.aliyun.com/api/cloudauth/describeverifyresult.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeVerifyResultWithCallback(request *DescribeVerifyResultRequest, callback func(response *DescribeVerifyResultResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVerifyResultResponse
		var err error
		defer close(result)
		response, err = client.DescribeVerifyResult(request)
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

// DescribeVerifyResultRequest is the request struct for api DescribeVerifyResult
type DescribeVerifyResultRequest struct {
	*requests.RpcRequest
	BizType string `position:"Query" name:"BizType"`
	BizId   string `position:"Query" name:"BizId"`
}

// DescribeVerifyResultResponse is the response struct for api DescribeVerifyResult
type DescribeVerifyResultResponse struct {
	*responses.BaseResponse
	RequestId                 string   `json:"RequestId" xml:"RequestId"`
	VerifyStatus              int      `json:"VerifyStatus" xml:"VerifyStatus"`
	AuthorityComparisionScore float64  `json:"AuthorityComparisionScore" xml:"AuthorityComparisionScore"`
	FaceComparisonScore       float64  `json:"FaceComparisonScore" xml:"FaceComparisonScore"`
	IdCardFaceComparisonScore float64  `json:"IdCardFaceComparisonScore" xml:"IdCardFaceComparisonScore"`
	Material                  Material `json:"Material" xml:"Material"`
}

// CreateDescribeVerifyResultRequest creates a request to invoke DescribeVerifyResult API
func CreateDescribeVerifyResultRequest() (request *DescribeVerifyResultRequest) {
	request = &DescribeVerifyResultRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudauth", "2019-03-07", "DescribeVerifyResult", "cloudauth", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeVerifyResultResponse creates a response to parse from DescribeVerifyResult response
func CreateDescribeVerifyResultResponse() (response *DescribeVerifyResultResponse) {
	response = &DescribeVerifyResultResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
