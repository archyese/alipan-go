package alipan

import (
	"fmt"
	"github.com/archyese/alipan-sdk/adrive_client/protos"
	"github.com/archyese/alipan-sdk/common"
)

func (c *AdriveClient) AdriveOpenFileCreate(username string, params *protos.AdriveOpenFileCreateParams) (*protos.AdriveOpenFileCreate, error) {
	apiUrl := fmt.Sprintf("%s/adrive/v1.0/openFile/create", c.ApiHost)
	return common.DoJsonRequest[protos.AdriveOpenFileCreate](username, c.AccessTokenLoader, c.Agent, apiUrl, params)
}
