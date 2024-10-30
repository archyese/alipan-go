package alipan

import (
	"fmt"
	"github.com/archyese/alipan-sdk/adrive_client/protos"
	"github.com/archyese/alipan-sdk/common"
)

func (c *AdriveClient) AdriveOpenFileListUploadedParts(username string, params *protos.AdriveOpenFileListUploadedPartsParams) (*protos.AdriveOpenFileListUploadedParts, error) {
	apiUrl := fmt.Sprintf("%s/adrive/v1.0/openFile/listUploadedParts", c.ApiHost)
	return common.DoJsonRequest[protos.AdriveOpenFileListUploadedParts](username, c.AccessTokenLoader, c.Agent, apiUrl, params)
}
