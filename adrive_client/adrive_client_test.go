package alipan

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/archyese/alipan-sdk/adrive_client/local"
	"github.com/archyese/alipan-sdk/adrive_client/protos"
	"github.com/archyese/alipan-sdk/oauth_client"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
)

func oauthClient(t *testing.T) *oauth_client.OauthClient {
	buff, err := os.ReadFile("../tests_data/app_info.toml")
	if err != nil {
		t.Fatal(err)
	}
	appInfo := map[string]string{}
	err = toml.Unmarshal(buff, &appInfo)
	if err != nil {
		t.Fatal(err)
	}
	client := oauth_client.NewOauthClient(appInfo["client_id"], appInfo["client_secret"])
	return client
}

type TomlAccessTokenStore struct{}

func (t *TomlAccessTokenStore) SaveAccessToken(token *local.AccessToken) error {
	buff, err := toml.Marshal(token)
	if err != nil {
		return err
	}
	err = os.WriteFile("../tests_data/access_token.toml", buff, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (t *TomlAccessTokenStore) LoadAccessToken() (*local.AccessToken, error) {
	buff, err := os.ReadFile("../tests_data/access_token.toml")
	if err != nil {
		return nil, err
	}
	token := &local.AccessToken{}
	err = toml.Unmarshal(buff, token)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func adriveClient(t *testing.T) *AdriveClient {
	buff, err := os.ReadFile("../tests_data/app_info.toml")
	if err != nil {
		t.Fatal(err)
	}
	appInfo := map[string]string{}
	err = toml.Unmarshal(buff, &appInfo)
	if err != nil {
		t.Fatal(err)
	}
	client := NewAdriveClient(appInfo["client_id"], &local.OauthClientTokenLoader{
		AccessTokenStore: &TomlAccessTokenStore{},
		OauthClient:      oauthClient(t),
	})
	return client
}

func TestAdriveClient_OauthUserInfo(t *testing.T) {
	client := adriveClient(t)
	info, err := client.OauthUserInfo("yese", &protos.OauthUserInfoParams{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(info)
}

func TestAdriveClient_OauthUserScopes(t *testing.T) {
	client := adriveClient(t)
	info, err := client.OauthUserScopes("yese", &protos.OauthUserScopesParams{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(info)
}

func TestAdriveClient_AdriveUserGetSpaceInfo(t *testing.T) {
	client := adriveClient(t)
	info, err := client.AdriveUserGetSpaceInfo("yese", &protos.AdriveUserGetSpaceInfoParams{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(info)
}

func TestAdriveClient_AdriveUserGetDriveInfo(t *testing.T) {
	client := adriveClient(t)
	info, err := client.AdriveUserGetDriveInfo("yese", &protos.AdriveUserGetDriveInfoParams{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(info)
}

func TestAdriveClient_AdriveOpenFileList(t *testing.T) {
	client := adriveClient(t)
	info, err := client.AdriveOpenFileList("yese", &protos.AdriveOpenFileListParams{
		DriveId:      "drive_id",
		ParentFileId: "root",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(info)
}

func TestAdriveClient_AdriveOpenFileGet(t *testing.T) {
	client := adriveClient(t)
	info, err := client.AdriveOpenFileGet("yese", &protos.AdriveOpenFileGetParams{
		DriveId: "drive_id",
		FileId:  "file_id",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(info)
}

func TestAdriveClient_AdriveOpenFileGetByPath(t *testing.T) {
	client := adriveClient(t)
	info, err := client.AdriveOpenFileGetByPath("yese", &protos.AdriveOpenFileGetByPathParams{
		DriveId:  "drive_id",
		FilePath: "/file_path",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(info)
}

func TestAdriveClient_AdriveOpenFileGetDownloadUrl(t *testing.T) {
	client := adriveClient(t)
	info, err := client.AdriveOpenFileGetDownloadUrl("yese", &protos.AdriveOpenFileGetDownloadUrlParams{
		DriveId: "drive_id",
		FileId:  "file_id",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(info)
}

func TestAdriveClient_AdriveOpenFileRecyclebinTrash(t *testing.T) {
	client := adriveClient(t)
	info, err := client.AdriveOpenFileRecyclebinTrash("yese", &protos.AdriveOpenFileRecyclebinTrashParams{
		DriveId: "drive_id",
		FileId:  "file_id",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(info)
}

func TestAdriveClient_AdriveOpenFileDelete(t *testing.T) {
	client := adriveClient(t)
	info, err := client.AdriveOpenFileDelete("yese", &protos.AdriveOpenFileDeleteParams{
		DriveId: "drive_id",
		FileId:  "file_id",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(info)
}

func TestAdriveClient_AdriveOpenFileCopy(t *testing.T) {
	client := adriveClient(t)
	info, err := client.AdriveOpenFileCopy("yese", &protos.AdriveOpenFileCopyParams{
		DriveId:        "drive_id",
		FileId:         "file_id",
		ToDriveId:      "to_drive_id",
		ToParentFileId: "to_parent_file_id",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(info)
}

func TestAdriveClient_AdriveOpenFileMove(t *testing.T) {
	client := adriveClient(t)
	info, err := client.AdriveOpenFileMove("yese", &protos.AdriveOpenFileMoveParams{
		DriveId:        "drive_id",
		FileId:         "file_id",
		ToParentFileId: "to_parent_file_id",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(info)
}

func TestAdriveClient_AdriveOpenFileUpdate(t *testing.T) {
	client := adriveClient(t)
	info, err := client.AdriveOpenFileUpdate("yese", &protos.AdriveOpenFileUpdateParams{
		DriveId: "drive_id",
		FileId:  "file_id",
		Name:    "new_name",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(info)
}

// 秒传

const FILE = "HELLO.TXT"
const TEXT = "Hello, World!"

func TestAdriveClient_AdriveOpenFileCreate_RapidUpload(t *testing.T) {
	client := adriveClient(t)
	info, err := client.AdriveOpenFileCreate("yese", &protos.AdriveOpenFileCreateParams{
		DriveId:         "drive_id",
		ParentFileId:    "root",
		Type:            "file",
		Name:            FILE,
		Size:            int64(len(TEXT)),
		ContentHashName: "sha1",
		ContentHash:     sha1ForText(TEXT),
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(info)
}

func sha1ForText(text string) string {
	hash := sha1.New()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}

// 非秒传

func TestAdriveClient_AdriveOpenFileCreate_NormalUpload(t *testing.T) {
	client := adriveClient(t)
	// 创建文件
	create, err := client.AdriveOpenFileCreate("yese", &protos.AdriveOpenFileCreateParams{
		DriveId:      "drive_id",
		ParentFileId: "root",
		Type:         "file",
		Name:         FILE,
		Size:         int64(len(TEXT)),
		// 避免秒传
		// ContentHashName: "sha1",
		// ContentHash:     sha1ForText(TEXT),
		PartInfoList: []*protos.AdriveOpenFilePartInfo{
			&protos.AdriveOpenFilePartInfo{
				PartNumber: 1,
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(create)
	//
	if create.Exist {
		log.Fatal("exist")
	}
	if create.RapidUpload {
		log.Printf("rapid upload")
		return
	}
	// 获取上传链接
	uploadUrlGet, err := client.AdriveOpenFileGetUploadUrl("yese", &protos.AdriveOpenFileGetUploadUrlParams{
		DriveId:  create.DriveId,
		FileId:   create.FileId,
		UploadId: create.UploadId,
		PartInfoList: []*protos.AdriveOpenFilePartInfo{
			{
				PartNumber: 1,
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	uploadUrl := uploadUrlGet.PartInfoList[0].UploadUrl
	// 上传文件
	req, err := http.NewRequest("PUT", uploadUrl, strings.NewReader(TEXT))
	if err != nil {
		t.Fatal(err)
	}
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.Fatal(err)
		}
	}(rsp.Body)
	body, err := io.ReadAll(rsp.Body)
	if err != nil {
		t.Fatal(err)
	}
	responseCode := rsp.StatusCode
	responseText := string(body)
	if responseCode/100 != 2 {
		t.Fatal(fmt.Sprintf("responseCode: %d, responseText: %s", responseCode, responseText))
	}
	t.Log(fmt.Sprintf("responseCode: %d, responseText: %s", responseCode, responseText))
	// 完成
	complete, err := client.AdriveOpenFileComplete("yese", &protos.AdriveOpenFileCompleteParams{
		DriveId:  create.DriveId,
		FileId:   create.FileId,
		UploadId: create.UploadId,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(complete)
}
