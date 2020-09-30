package alikafka

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

// GetTopicStatus invokes the alikafka.GetTopicStatus API synchronously
func (client *Client) GetTopicStatus(request *GetTopicStatusRequest) (response *GetTopicStatusResponse, err error) {
	response = CreateGetTopicStatusResponse()
	err = client.DoAction(request, response)
	return
}

// GetTopicStatusWithChan invokes the alikafka.GetTopicStatus API asynchronously
func (client *Client) GetTopicStatusWithChan(request *GetTopicStatusRequest) (<-chan *GetTopicStatusResponse, <-chan error) {
	responseChan := make(chan *GetTopicStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetTopicStatus(request)
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

// GetTopicStatusWithCallback invokes the alikafka.GetTopicStatus API asynchronously
func (client *Client) GetTopicStatusWithCallback(request *GetTopicStatusRequest, callback func(response *GetTopicStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetTopicStatusResponse
		var err error
		defer close(result)
		response, err = client.GetTopicStatus(request)
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

// GetTopicStatusRequest is the request struct for api GetTopicStatus
type GetTopicStatusRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	Topic      string `position:"Query" name:"Topic"`
}

// GetTopicStatusResponse is the response struct for api GetTopicStatus
type GetTopicStatusResponse struct {
	*responses.BaseResponse
	Success     bool        `json:"Success" xml:"Success"`
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	Code        int         `json:"Code" xml:"Code"`
	Message     string      `json:"Message" xml:"Message"`
	TopicStatus TopicStatus `json:"TopicStatus" xml:"TopicStatus"`
}

// CreateGetTopicStatusRequest creates a request to invoke GetTopicStatus API
func CreateGetTopicStatusRequest() (request *GetTopicStatusRequest) {
	request = &GetTopicStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alikafka", "2019-09-16", "GetTopicStatus", "alikafka", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetTopicStatusResponse creates a response to parse from GetTopicStatus response
func CreateGetTopicStatusResponse() (response *GetTopicStatusResponse) {
	response = &GetTopicStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
