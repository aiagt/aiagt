package main

import (
	"context"
	"encoding/json"
	"github.com/aiagt/aiagt/apps/model/conf"
	"github.com/aiagt/aiagt/apps/model/model"
	"github.com/aiagt/aiagt/common/confutil"
	"github.com/aiagt/aiagt/pkg/closer"
	"github.com/aiagt/aiagt/pkg/hash/hmap"
	"github.com/aiagt/aiagt/pkg/hash/hset"
	"github.com/aiagt/aiagt/pkg/snowflake"
	ktdb "github.com/aiagt/kitextool/option/server/db"
	"github.com/shopspring/decimal"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

type getVModelResp struct {
	Data []struct {
		Id          int                     `json:"id"`
		ModelName   string                  `json:"model_name"`
		Supplier    string                  `json:"supplier"`
		MaxTokens   int64                   `json:"max_tokens"`
		Description string                  `json:"description"`
		CreatedTime int                     `json:"created_time"`
		Tags        []*getVModelRespDataTag `json:"tags"`
	} `json:"data"`
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type getVModelRespDataTag struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	IsDisplay bool   `json:"is_display"`
}

type getPricingResp struct {
	Data       []*getPricingRespData `json:"data"`
	GroupRatio struct {
		Default int `json:"default"`
		Gf      int `json:"gf"`
		Vip     int `json:"vip"`
	} `json:"group_ratio"`
	Success bool `json:"success"`
}

type getPricingRespData struct {
	ModelName       string          `json:"model_name"`
	QuotaType       int             `json:"quota_type"`
	ModelRatio      decimal.Decimal `json:"model_ratio"`
	ModelPrice      decimal.Decimal `json:"model_price"`
	OwnerBy         string          `json:"owner_by"`
	CompletionRatio decimal.Decimal `json:"completion_ratio"`
	EnableGroups    []string        `json:"enable_groups"`
}

func (g *getPricingRespData) UnmarshalJSON(data []byte) error {
	var aux struct {
		ModelName       string   `json:"model_name"`
		QuotaType       int      `json:"quota_type"`
		ModelRatio      float64  `json:"model_ratio"`
		ModelPrice      float64  `json:"model_price"`
		OwnerBy         string   `json:"owner_by"`
		CompletionRatio float64  `json:"completion_ratio"`
		EnableGroups    []string `json:"enable_groups"`
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	g.ModelName = aux.ModelName
	g.QuotaType = aux.QuotaType
	g.ModelRatio = decimal.NewFromFloat(aux.ModelRatio)
	g.ModelPrice = decimal.NewFromFloat(aux.ModelPrice)
	g.OwnerBy = aux.OwnerBy
	g.CompletionRatio = decimal.NewFromFloat(aux.CompletionRatio)
	g.EnableGroups = aux.EnableGroups

	return nil
}

func httpGET[T any](url string) (*T, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer closer.Close(resp.Body)

	var result T
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func save(models []*model.Models) error {
	confutil.LoadConf(conf.Conf(), filepath.Join("..", "conf"))
	ktdb.WithDB(ktdb.NewMySQLDial()).Apply(nil, conf.Conf().GetServerConf())

	err := ktdb.DB().AutoMigrate(new(model.Models))
	if err != nil {
		return err
	}

	for _, m := range models {
		m.ID = snowflake.Generate().Int64()
	}

	return ktdb.DBCtx(context.TODO()).Model(new(model.Models)).CreateInBatches(models, 100).Error
}

func main() {
	vModelResp, err := httpGET[getVModelResp]("https://api.gpt.ge/api/vmodel/")
	if err != nil {
		log.Fatal(err)
	}
	if !vModelResp.Success {
		log.Fatalf("get vmodel failed: %s", vModelResp.Message)
	}

	pricingResp, err := httpGET[getPricingResp]("https://api.gpt.ge/api/pricing")
	if err != nil {
		log.Fatal(err)
	}
	if !pricingResp.Success {
		log.Fatal("get pricing failed")
	}

	modelPricingMap := hmap.FromSliceEntries(pricingResp.Data, func(t *getPricingRespData) (string, *getPricingRespData, bool) {
		return t.ModelName, t, true
	})

	var models []*model.Models

	for _, item := range vModelResp.Data {
		if item.Supplier != "OpenAI" {
			continue
		}

		tags := hset.FromSlice(item.Tags, func(t *getVModelRespDataTag) string { return t.Name })
		if !tags.Has("文本") || tags.Has("弃用") {
			continue
		}

		// tags filter
		for tag := range tags {
			if len(tag) < 2 {
				tags.Remove(tag)
			}
		}

		pricing, ok := modelPricingMap[item.ModelName]
		if !ok {
			continue
		}

		var datum = decimal.NewFromFloat(0.002)

		log.Println(item.ModelName)
		inputPrice := pricing.ModelRatio.Mul(datum)
		outputPrice := pricing.CompletionRatio.Mul(inputPrice)

		models = append(models, &model.Models{
			Name:        strings.ReplaceAll(strings.ReplaceAll(item.ModelName, "chatgpt", "GPT"), "gpt", "GPT"),
			Description: item.Description,
			Source:      "OPENAI",
			ModelKey:    item.ModelName,
			Logo:        "app_logo/5070bbab-b27d-4ef4-ae7d-1df994ba913a.svg",
			MaxToken:    item.MaxTokens * 1000,
			Tags:        tags.List(),
			InputPrice:  inputPrice,
			OutputPrice: outputPrice,
		})
	}

	err = save(models)
	if err != nil {
		log.Fatal(err)
	}
}
