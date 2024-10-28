package alipan

import (
	"fmt"
	"github.com/archyese/alipan-sdk/adrive_client/protos"
	"github.com/archyese/alipan-sdk/common"
	"net/url"
)

func (c *AdriveClient) OauthUserScopes(username string, _params *protos.OauthUserScopesParams) (*protos.OauthUserScopes, error) {
	apiUrl := fmt.Sprintf("%s/oauth/users/scopes", c.ApiHost)
	params := url.Values{}
	return common.DoGetQuery[protos.OauthUserScopes](username, c.AccessTokenLoader, c.Agent, apiUrl, params)
}
