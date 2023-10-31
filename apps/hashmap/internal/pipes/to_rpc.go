package pipes

import (
	"github.com/dehwyy/makoto/libs/database/models"
	"github.com/dehwyy/makoto/libs/grpc/generated/hashmap"
)

func ToRpcItems(items []*models.HashmapItem) (response []*hashmap.Item) {
	// pipe each Item
	for _, item := range items {

		item_tags := ToRpcTags(item.Tags)

		response = append(response, &hashmap.Item{
			Id:    item.Id,
			Key:   item.Key,
			Value: item.Value,
			Extra: item.Extra,
			Tags:  item_tags,
		})
	}

	return response
}

func ToRpcTags(tags_input []*models.HashmapTag) (tags []*hashmap.Tag) {
	for _, tag := range tags_input {
		tags = append(tags, &hashmap.Tag{
			Text:   tag.Text,
			Usages: uint32(len(tag.Items)),
		})
	}

	return tags
}
