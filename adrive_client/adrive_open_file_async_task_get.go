package alipan

import (
	"fmt"
	"github.com/archyese/alipan-sdk/adrive_client/protos"
	"github.com/archyese/alipan-sdk/common"
)

func (c *AdriveClient) AdriveOpenFileAsyncTaskGet(username string, params *protos.AdriveOpenFileAsyncTaskGetParams) (*protos.AdriveOpenFileAsyncTaskGet, error) {
	apiUrl := fmt.Sprintf("%s/adrive/v1.0/openFile/async_task/get", c.ApiHost)
	return common.DoJsonRequest[protos.AdriveOpenFileAsyncTaskGet](username, c.AccessTokenLoader, c.Agent, apiUrl, params)
}
