package alipan

import (
	"fmt"
	"github.com/archyese/alipan-sdk/adrive_client/protos"
	"github.com/archyese/alipan-sdk/common"
)

func (c *AdriveClient) AdriveOpenFileList(username string, params *protos.AdriveOpenFileListParams) (*protos.AdriveOpenFileList, error) {
	apiUrl := fmt.Sprintf("%s/adrive/v1.0/openFile/list", c.ApiHost)
	return common.DoJsonRequest[protos.AdriveOpenFileList](username, c.AccessTokenLoader, c.Agent, apiUrl, params)
}
