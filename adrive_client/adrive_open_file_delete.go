package alipan

import (
	"fmt"
	"github.com/archyese/alipan-sdk/adrive_client/protos"
	"github.com/archyese/alipan-sdk/common"
)

func (c *AdriveClient) AdriveOpenFileDelete(username string, params *protos.AdriveOpenFileDeleteParams) (*protos.AdriveOpenFileDelete, error) {
	apiUrl := fmt.Sprintf("%s/adrive/v1.0/openFile/recyclebin/trash", c.ApiHost)
	return common.DoJsonRequest[protos.AdriveOpenFileDelete](username, c.AccessTokenLoader, c.Agent, apiUrl, params)
}
