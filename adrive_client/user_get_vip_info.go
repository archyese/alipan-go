package alipan

import (
	"fmt"
	"github.com/archyese/alipan-sdk/adrive_client/protos"
	"github.com/archyese/alipan-sdk/common"
	"net/url"
)

func (c *AdriveClient) UserGetVipInfo(username string, _params *protos.UserGetVipInfoParams) (*protos.UserGetVipInfo, error) {
	apiUrl := fmt.Sprintf("%s/oauth/users/info", c.ApiHost)
	params := url.Values{}
	return common.DoGetQuery[protos.UserGetVipInfo](username, c.AccessTokenLoader, c.Agent, apiUrl, params)
}
