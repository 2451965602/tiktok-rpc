package milvus

import (
	"context"
	"github.com/milvus-io/milvus-sdk-go/v2/client"
	"github.com/milvus-io/milvus-sdk-go/v2/entity"
	"tiktokrpc/kitex_gen/user"
)

func Search(vector []float32, req *user.SearchImagesRequest) (string, error) {
	m := MClient.milvus
	ctx := context.Background()

	idx, err := entity.NewIndexIvfFlat( // NewIndex func
		entity.L2, // metricType
		1024,      // ConstructParams
	)
	if err != nil {
		return "", err
	}

	if err = m.CreateIndex(
		ctx,                // ctx
		req.CollectionName, // CollectionName
		"image_intro",      // fieldName
		idx,                // entity.Index
		false,              // async
	); err != nil {
		return "", err
	}

	err = m.LoadCollection(ctx, req.CollectionName, false)
	if err != nil {
		return "", err
	}

	sp, _ := entity.NewIndexIvfFlatSearchParam( // NewIndex*SearchParam func
		10, // searchParam
	)

	opt := client.SearchQueryOptionFunc(func(option *client.SearchQueryOption) {
		option.Limit = 3
		option.Offset = 0
		option.ConsistencyLevel = entity.ClStrong //
		option.IgnoreGrowing = false
	})

	floatv := vector

	searchResult, err := m.Search(
		ctx,                   // ctx
		req.CollectionName,    // CollectionName
		[]string{},            // partitionNames
		"",                    // expr
		[]string{"image_url"}, // outputFields
		[]entity.Vector{entity.FloatVector(floatv)}, // vectors
		"image_intro", // vectorField
		entity.L2,     // metricType
		10,            // topK
		sp,            // searchParams
		opt,
	)
	if err != nil {
		return "", err
	}

	// TODO: return result
	//bookIDList := make([]int64, 0, req.ResultCount)
	var imageURLs []string
	for _, sr := range searchResult {
		for _, fieldData := range sr.Fields {
			if fieldData.Name() == "image_url" {
				urls := fieldData.(*entity.ColumnVarChar).Data()
				imageURLs = append(imageURLs, urls...)
			}
		}
	}

	err = m.ReleaseCollection(ctx, req.CollectionName)
	if err != nil {
		return "", err
	}

	// 返回第一个 image_url
	if len(imageURLs) > 0 {
		return imageURLs[0], nil
	}

	return "", nil
}
