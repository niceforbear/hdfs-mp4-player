package web_hdfs

import (
	"encoding/json"
	"github.com/niceforbear/hdfs-mp4-player/helpers"
)

type ListDirFileStatusDetail struct {
	Group				string		`json:"group"`
	Length				int			`json:"length"`
	ModificationTime	int64		`json:"modificationTime"`
	Owner				string		`json:"owner"`
	PathSuffix			string		`json:"pathSuffix"`
	Permission			string		`json:"permission"`
	Type				string		`json:"type"`
}

type ListDirFileStatus struct {
	FileStatus	[]ListDirFileStatusDetail	`json:"FileStatus"`
}

type ListDirResponse struct {
	FileStatuses ListDirFileStatus	`json:"FileStatuses"`
}

func (this *WebHdfsApi) ListDir() ([]ListDirFileStatusDetail, error) {
	address, _ := WrapperListDirUrl(this.RequestUrl, this.FilePath)
	logger.Printf("ListDir addr: %v", address)

	reqBody := map[string]string{
		"op": this.Operation,
	}

	respBody, err := helpers.HttpWrapperGet(address, reqBody, nil)
	if err != nil {
		logger.Printf("[ERROR] WebHdfsApi.ListDir error: %+v", err)
		return nil, err
	}

	listDirResp := &ListDirResponse{}
	err = json.Unmarshal(respBody, &listDirResp)
	if err != nil {
		logger.Printf("[ERROR] json unmarshal list dir resp error: %v", err)
		return nil, err
	}

	return listDirResp.FileStatuses.FileStatus, nil
}
