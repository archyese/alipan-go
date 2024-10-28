package alipan

import (
	"fmt"
	"github.com/archyese/alipan-sdk/adrive_client/protos"
	"github.com/archyese/alipan-sdk/common"
)

func (c *AdriveClient) AdriveOpenFileGet(username string, params *protos.AdriveOpenFileGetParams) (*protos.AdriveOpenFile, error) {
	apiUrl := fmt.Sprintf("%s/adrive/v1.0/openFile/get", c.ApiHost)
	return common.DoJsonRequest[protos.AdriveOpenFile](username, c.AccessTokenLoader, c.Agent, apiUrl, params)
}
