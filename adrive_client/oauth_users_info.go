package alipan

import (
	"fmt"
	"github.com/archyese/alipan-sdk/adrive_client/protos"
	"github.com/archyese/alipan-sdk/common"
	"net/url"
)

func (c *AdriveClient) OauthUserInfo(username string, _params *protos.OauthUserInfoParams) (*protos.OauthUserInfo, error) {
	apiUrl := fmt.Sprintf("%s/oauth/users/info", c.ApiHost)
	params := url.Values{}
	return common.DoGetQuery[protos.OauthUserInfo](username, c.AccessTokenLoader, c.Agent, apiUrl, params)
}
