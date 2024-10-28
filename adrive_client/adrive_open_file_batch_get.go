package alipan

import (
	"fmt"
	"github.com/archyese/alipan-sdk/adrive_client/protos"
	"github.com/archyese/alipan-sdk/common"
)

func (c *AdriveClient) AdriveOpenFileBatchGet(username string, params *protos.AdriveOpenFileBatchGetParams) (*protos.AdriveOpenFileBatchGet, error) {
	apiUrl := fmt.Sprintf("%s/adrive/v1.0/openFile/get", c.ApiHost)
	return common.DoJsonRequest[protos.AdriveOpenFileBatchGet](username, c.AccessTokenLoader, c.Agent, apiUrl, params)
}
