package milvus

import (
	"context"
	"tiktokrpc/kitex_gen/user"

	"github.com/milvus-io/milvus-sdk-go/v2/entity"
)

func CreateCollection(ctx context.Context, req *user.UploadImagesRequest) error {
	m := MClient.milvus

	exist, err := m.HasCollection(ctx, req.CollectionName)
	if err != nil {
		return err
	}

	if exist {
		return nil
	}

	schema := &entity.Schema{
		CollectionName: req.CollectionName,
		Description:    "search for images",
		Fields: []*entity.Field{
			{
				Name:       "image_id",
				DataType:   entity.FieldTypeInt64,
				PrimaryKey: true,
				AutoID:     true,
			},
			{
				Name:     "image_url",
				DataType: entity.FieldTypeVarChar,
				TypeParams: map[string]string{
					"max_length": "512",
				},
			},
			{
				Name:     "image_intro",
				DataType: entity.FieldTypeFloatVector,
				TypeParams: map[string]string{
					"dim": "1000",
				},
			},
		},
		EnableDynamicField: true,
	}

	if err := m.CreateCollection(ctx, schema, 2); err != nil {
		return err
	}

	return nil
}

func InsertData(ctx context.Context, imagesData []float32, imgUrl, collectionName string) error {
	m := MClient.milvus

	urlColumn := entity.NewColumnVarChar("image_url", []string{imgUrl})
	introColumn := entity.NewColumnFloatVector("image_intro", 1000, [][]float32{imagesData})

	if _, err := m.Insert(
		ctx,
		collectionName,
		"",
		urlColumn,
		introColumn,
	); err != nil {
		return err
	}

	return nil
}
