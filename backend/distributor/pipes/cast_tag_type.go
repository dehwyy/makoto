package pipes

import (
	"github.com/dehwyy/Makoto/backend/distributor/graphql/model"
	dictGrpc "github.com/dehwyy/Makoto/backend/grpc/gen/dict/go/proto"
)

func CastTagsGrpcToGraphQL(tags []*dictGrpc.Tag) []*model.Tag {
	var word_tags []*model.Tag
	for _, tag := range tags {
		word_tags = append(word_tags, &model.Tag{
			TagID: Uint32ToString(tag.TagId),
			Text:  tag.Text,
		})
	}
	return word_tags
}
