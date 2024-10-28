package alipan

import (
	"fmt"
	"github.com/archyese/alipan-sdk/adrive_client/protos"
	"github.com/archyese/alipan-sdk/common"
)

func (c *AdriveClient) AdriveOpenFileMove(username string, params *protos.AdriveOpenFileMoveParams) (*protos.AdriveOpenFileMove, error) {
	apiUrl := fmt.Sprintf("%s/adrive/v1.0/openFile/move", c.ApiHost)
	return common.DoJsonRequest[protos.AdriveOpenFileMove](username, c.AccessTokenLoader, c.Agent, apiUrl, params)
}
