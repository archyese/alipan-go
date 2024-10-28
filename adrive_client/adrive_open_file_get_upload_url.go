package alipan

import (
	"fmt"
	"github.com/archyese/alipan-sdk/adrive_client/protos"
	"github.com/archyese/alipan-sdk/common"
)

func (c *AdriveClient) AdriveOpenFileGetUploadUrl(username string, params *protos.AdriveOpenFileGetUploadUrlParams) (*protos.AdriveOpenFileGetUploadUrl, error) {
	apiUrl := fmt.Sprintf("%s/adrive/v1.0/openFile/getUploadUrl", c.ApiHost)
	return common.DoJsonRequest[protos.AdriveOpenFileGetUploadUrl](username, c.AccessTokenLoader, c.Agent, apiUrl, params)
}
