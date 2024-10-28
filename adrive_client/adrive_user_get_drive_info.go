package alipan

import (
	"fmt"
	"github.com/archyese/alipan-sdk/adrive_client/protos"
	"github.com/archyese/alipan-sdk/common"
)

func (c *AdriveClient) AdriveUserGetDriveInfo(username string, _params *protos.AdriveUserGetDriveInfoParams) (*protos.AdriveUserGetDriveInfo, error) {
	apiUrl := fmt.Sprintf("%s/adrive/v1.0/user/getDriveInfo", c.ApiHost)
	return common.DoPostNoBody[protos.AdriveUserGetDriveInfo](username, c.AccessTokenLoader, c.Agent, apiUrl)
}
