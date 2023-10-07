package pipes

import (
	rpc "github.com/dehwyy/Makoto/backend/grpc/gen/dict/go/proto"
	"github.com/dehwyy/makoto/backend/dict/db/models"
)

func CastTagsToGrpcOutput(tags []models.Tag) []*rpc.Tag {
	var word_tags []*rpc.Tag
	for _, tag := range tags {
		word_tags = append(word_tags, &rpc.Tag{
			TagId: tag.Id,
			Text:  tag.Text,
		})
	}

	return word_tags
}
