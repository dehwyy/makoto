package pipes

import (
	rpc "github.com/dehwyy/Makoto/backend/grpc/gen/dict/go/proto"
	"github.com/dehwyy/makoto/backend/dict/db/models"
)

func CastWordsToGrpcOutput(words []*models.Word) []*rpc.WordResponse {
	var res []*rpc.WordResponse

	// pipe each Word
	for _, word := range words {

		// pipe tags
		var word_tags []*rpc.Tag
		for _, tag := range word.Tags {
			word_tags = append(word_tags, &rpc.Tag{
				TagId: tag.Id,
				Text:  tag.Text,
			})
		}

		res = append(res, &rpc.WordResponse{
			Id:    uint32(word.Id),
			Word:  word.Word,
			Value: word.Value,
			Extra: word.Extra,
			Tags:  word_tags,
		})
	}

	return res
}
