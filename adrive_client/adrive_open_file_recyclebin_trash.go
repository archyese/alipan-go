package alipan

import (
	"fmt"
	"github.com/archyese/alipan-sdk/adrive_client/protos"
	"github.com/archyese/alipan-sdk/common"
)

func (c *AdriveClient) AdriveOpenFileRecyclebinTrash(username string, params *protos.AdriveOpenFileRecyclebinTrashParams) (*protos.AdriveOpenFileRecyclebinTrash, error) {
	apiUrl := fmt.Sprintf("%s/adrive/v1.0/openFile/recyclebin/trash", c.ApiHost)
	return common.DoJsonRequest[protos.AdriveOpenFileRecyclebinTrash](username, c.AccessTokenLoader, c.Agent, apiUrl, params)
}
