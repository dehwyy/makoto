package pipes

import (
	"github.com/dehwyy/makoto/libs/database/models"
	"github.com/dehwyy/makoto/libs/grpc/generated/hashmap"
)

func ItemsDb2Grpc(items []*models.HashmapItem) (response []*hashmap.ItemResponse) {
	// pipe each Item
	for _, item := range items {

		item_tags := TagsFromDb2Grpc(item.Tags)

		response = append(response, &hashmap.ItemResponse{
			Id:    item.Id,
			Key:   item.Key,
			Value: item.Value,
			Extra: item.Extra,
			Tags:  item_tags,
		})
	}

	return response
}
