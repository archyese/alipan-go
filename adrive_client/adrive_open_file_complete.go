package alipan

import (
	"fmt"
	"github.com/archyese/alipan-sdk/adrive_client/protos"
	"github.com/archyese/alipan-sdk/common"
)

func (c *AdriveClient) AdriveOpenFileComplete(username string, params *protos.AdriveOpenFileCompleteParams) (*protos.AdriveOpenFile, error) {
	apiUrl := fmt.Sprintf("%s/adrive/v1.0/openFile/complete", c.ApiHost)
	return common.DoJsonRequest[protos.AdriveOpenFile](username, c.AccessTokenLoader, c.Agent, apiUrl, params)
}
