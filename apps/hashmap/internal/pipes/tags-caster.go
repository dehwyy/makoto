package pipes

import (
	"github.com/dehwyy/makoto/libs/database/models"
	"github.com/dehwyy/makoto/libs/grpc/generated/hashmap"
)

func TagsFromDb2Grpc(tags_input []*models.HashmapTag) (tags []*hashmap.Tag) {
	for _, tag := range tags_input {
		tags = append(tags, &hashmap.Tag{
			TagId:  tag.Id,
			Text:   tag.Text,
			Usages: int32(len(tag.Items)),
		})
	}

	return tags
}
