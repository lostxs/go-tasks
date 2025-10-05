package api

import (
	"3-struct/config"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const baseURL = "https://api.jsonbin.io/v3"

type Api struct {
	MasterKey string
}

func NewApi(cfg *config.Config) *Api {
	return &Api{
		MasterKey: cfg.MasterKey,
	}
}

type CreateBinResponse struct {
	Record   map[string]any `json:"record"`
	Metadata struct {
		ID        string `json:"id"`
		CreatedAt string `json:"createdAt"`
		Private   bool   `json:"private"`
	} `json:"metadata"`
}

func (a *Api) CreateBin(filename, name string) (*CreateBinResponse, error) {
	postData, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/b", baseURL), bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", a.MasterKey)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var response CreateBinResponse
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

type UpdateBinResponse struct {
	Record   map[string]any `json:"record"`
	Metadata struct {
		ParentID string `json:"parentId"`
		Private  bool   `json:"private"`
	} `json:"metadata"`
}

func (a *Api) UpdateBin(filename, id string) (*UpdateBinResponse, error) {
	postData, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/b/%s", baseURL, id), bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", a.MasterKey)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var response UpdateBinResponse
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

type DeleteBinResponse struct {
	Metadata struct {
		ID string `json:"id"`
	} `json:"metadata"`
}

func (a *Api) DeleteBin(id string) (*DeleteBinResponse, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/b/%s", baseURL, id), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Master-Key", a.MasterKey)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var response DeleteBinResponse
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

type GetBinResponse struct {
	Record   map[string]any `json:"record"`
	Metadata struct {
		ID      string `json:"id"`
		Private bool   `json:"private"`
	} `json:"metadata"`
}

func (a *Api) GetBin(id string) (*GetBinResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/b/%s", baseURL, id), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Master-Key", a.MasterKey)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var response GetBinResponse
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
