package alipan

import (
	"fmt"
	"github.com/archyese/alipan-sdk/adrive_client/protos"
	"github.com/archyese/alipan-sdk/common"
)

func (c *AdriveClient) AdriveUserGetSpaceInfo(username string, _params *protos.AdriveUserGetSpaceInfoParams) (*protos.AdriveUserGetSpaceInfo, error) {
	apiUrl := fmt.Sprintf("%s/adrive/v1.0/user/getSpaceInfo", c.ApiHost)
	return common.DoPostNoBody[protos.AdriveUserGetSpaceInfo](username, c.AccessTokenLoader, c.Agent, apiUrl)
}
