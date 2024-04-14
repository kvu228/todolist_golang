package rgpc_http

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"time"
	"to_do_list/module/post/usecase"
)

type rpcGetUsersByIds struct {
	url string
}

func NewRpcGetUsersByIds(url string) *rpcGetUsersByIds {
	return &rpcGetUsersByIds{url: url}
}

func (r *rpcGetUsersByIds) FindWithIds(ctx context.Context, ids []uuid.UUID) (owners []usecase.OwnerDTO, err error) {
	url := r.url
	method := "POST"

	// create a data variable as an empty array to contain ownerIds
	data := struct {
		Ids []uuid.UUID `json:"ids"`
	}{
		ids,
	}

	dataByte, _ := json.Marshal(data)
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest(method, url, bytes.NewReader(dataByte))

	if err != nil {
		log.Println(err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var responseData struct {
		Data []usecase.OwnerDTO `json:"data"`
	}

	if err := json.Unmarshal(body, &responseData); err != nil {
		return nil, err
	}

	return responseData.Data, nil

}
