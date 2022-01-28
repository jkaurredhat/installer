package edas

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

// GetWebContainerConfig invokes the edas.GetWebContainerConfig API synchronously
func (client *Client) GetWebContainerConfig(request *GetWebContainerConfigRequest) (response *GetWebContainerConfigResponse, err error) {
	response = CreateGetWebContainerConfigResponse()
	err = client.DoAction(request, response)
	return
}

// GetWebContainerConfigWithChan invokes the edas.GetWebContainerConfig API asynchronously
func (client *Client) GetWebContainerConfigWithChan(request *GetWebContainerConfigRequest) (<-chan *GetWebContainerConfigResponse, <-chan error) {
	responseChan := make(chan *GetWebContainerConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetWebContainerConfig(request)
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

// GetWebContainerConfigWithCallback invokes the edas.GetWebContainerConfig API asynchronously
func (client *Client) GetWebContainerConfigWithCallback(request *GetWebContainerConfigRequest, callback func(response *GetWebContainerConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetWebContainerConfigResponse
		var err error
		defer close(result)
		response, err = client.GetWebContainerConfig(request)
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

// GetWebContainerConfigRequest is the request struct for api GetWebContainerConfig
type GetWebContainerConfigRequest struct {
	*requests.RoaRequest
	AppId string `position:"Query" name:"AppId"`
}

// GetWebContainerConfigResponse is the response struct for api GetWebContainerConfig
type GetWebContainerConfigResponse struct {
	*responses.BaseResponse
	RequestId          string             `json:"RequestId" xml:"RequestId"`
	Message            string             `json:"Message" xml:"Message"`
	Code               int                `json:"Code" xml:"Code"`
	WebContainerConfig WebContainerConfig `json:"WebContainerConfig" xml:"WebContainerConfig"`
}

// CreateGetWebContainerConfigRequest creates a request to invoke GetWebContainerConfig API
func CreateGetWebContainerConfigRequest() (request *GetWebContainerConfigRequest) {
	request = &GetWebContainerConfigRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "GetWebContainerConfig", "/pop/v5/oam/web_container_config", "Edas", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetWebContainerConfigResponse creates a response to parse from GetWebContainerConfig response
func CreateGetWebContainerConfigResponse() (response *GetWebContainerConfigResponse) {
	response = &GetWebContainerConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
