package pipes

import (
	"fmt"

	"github.com/dehwyy/Makoto/backend/distributor/graphql/model"
	dictGrpc "github.com/dehwyy/Makoto/backend/grpc/gen/dict/go/proto"
)

// As Go compiler says if your convert uint32 to string it would be rune not a string of number
func Uint32ToString(v uint32) string {
	return fmt.Sprint(v)
}

func CastWordsGrpcToGraphQL(words []*dictGrpc.WordResponse) []*model.WordWithID {
	var res []*model.WordWithID

	// pipe each Word
	for _, word := range words {

		// pipe tags
		var word_tags []*model.Tag
		for _, tag := range word.Tags {
			word_tags = append(word_tags, &model.Tag{
				TagID: Uint32ToString(tag.TagId),
				Text:  tag.Text,
			})
		}

		res = append(res, &model.WordWithID{
			WordID: Uint32ToString(word.Id),
			Word:   word.Word,
			Value:  word.Value,
			Extra:  word.Extra,
			Tags:   word_tags,
		})
	}

	return res
}
