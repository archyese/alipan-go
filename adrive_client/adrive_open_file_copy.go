package alipan

import (
	"fmt"
	"github.com/archyese/alipan-sdk/adrive_client/protos"
	"github.com/archyese/alipan-sdk/common"
)

func (c *AdriveClient) AdriveOpenFileCopy(username string, params *protos.AdriveOpenFileCopyParams) (*protos.AdriveOpenFileCopy, error) {
	apiUrl := fmt.Sprintf("%s/adrive/v1.0/openFile/copy", c.ApiHost)
	return common.DoJsonRequest[protos.AdriveOpenFileCopy](username, c.AccessTokenLoader, c.Agent, apiUrl, params)
}
