package bssopenapi

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

// ApplyInvoice invokes the bssopenapi.ApplyInvoice API synchronously
func (client *Client) ApplyInvoice(request *ApplyInvoiceRequest) (response *ApplyInvoiceResponse, err error) {
	response = CreateApplyInvoiceResponse()
	err = client.DoAction(request, response)
	return
}

// ApplyInvoiceWithChan invokes the bssopenapi.ApplyInvoice API asynchronously
func (client *Client) ApplyInvoiceWithChan(request *ApplyInvoiceRequest) (<-chan *ApplyInvoiceResponse, <-chan error) {
	responseChan := make(chan *ApplyInvoiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ApplyInvoice(request)
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

// ApplyInvoiceWithCallback invokes the bssopenapi.ApplyInvoice API asynchronously
func (client *Client) ApplyInvoiceWithCallback(request *ApplyInvoiceRequest, callback func(response *ApplyInvoiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ApplyInvoiceResponse
		var err error
		defer close(result)
		response, err = client.ApplyInvoice(request)
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

// ApplyInvoiceRequest is the request struct for api ApplyInvoice
type ApplyInvoiceRequest struct {
	*requests.RpcRequest
	InvoicingType   requests.Integer `position:"Query" name:"InvoicingType"`
	ProcessWay      requests.Integer `position:"Query" name:"ProcessWay"`
	OwnerId         requests.Integer `position:"Query" name:"OwnerId"`
	InvoiceAmount   requests.Integer `position:"Query" name:"InvoiceAmount"`
	AddressId       requests.Integer `position:"Query" name:"AddressId"`
	ApplyUserNick   string           `position:"Query" name:"ApplyUserNick"`
	InvoiceByAmount requests.Boolean `position:"Query" name:"InvoiceByAmount"`
	CustomerId      requests.Integer `position:"Query" name:"CustomerId"`
	SelectedIds     *[]string        `position:"Query" name:"SelectedIds"  type:"Repeated"`
	UserRemark      string           `position:"Query" name:"UserRemark"`
}

// ApplyInvoiceResponse is the response struct for api ApplyInvoice
type ApplyInvoiceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateApplyInvoiceRequest creates a request to invoke ApplyInvoice API
func CreateApplyInvoiceRequest() (request *ApplyInvoiceRequest) {
	request = &ApplyInvoiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "ApplyInvoice", "", "")
	request.Method = requests.POST
	return
}

// CreateApplyInvoiceResponse creates a response to parse from ApplyInvoice response
func CreateApplyInvoiceResponse() (response *ApplyInvoiceResponse) {
	response = &ApplyInvoiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
