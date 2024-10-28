package alipan

import (
	"fmt"
	"github.com/archyese/alipan-sdk/adrive_client/protos"
	"github.com/archyese/alipan-sdk/common"
)

func (c *AdriveClient) AdriveOpenFileGetDownloadUrl(username string, params *protos.AdriveOpenFileGetDownloadUrlParams) (*protos.AdriveOpenFileGetDownloadUrl, error) {
	apiUrl := fmt.Sprintf("%s/adrive/v1.0/openFile/getDownloadUrl", c.ApiHost)
	return common.DoJsonRequest[protos.AdriveOpenFileGetDownloadUrl](username, c.AccessTokenLoader, c.Agent, apiUrl, params)
}
