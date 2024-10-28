package alipan

import (
	"fmt"
	"github.com/archyese/alipan-sdk/adrive_client/protos"
	"github.com/archyese/alipan-sdk/common"
)

func (c *AdriveClient) AdriveOpenFileUpdate(username string, params *protos.AdriveOpenFileUpdateParams) (*protos.AdriveOpenFile, error) {
	apiUrl := fmt.Sprintf("%s/adrive/v1.0/openFile/update", c.ApiHost)
	return common.DoJsonRequest[protos.AdriveOpenFile](username, c.AccessTokenLoader, c.Agent, apiUrl, params)
}
