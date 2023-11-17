package pipes

import (
	"net/http"
	"testing"

	"github.com/dehwyy/makoto/libs/logger"
)

var (
	log = logger.New()
)

func Test_FromHttpBody2Map_Parallel(t *testing.T) {
	tests := []struct {
		name   string
		url    string
		result map[string]interface{}
	}{
		{
			name: "thunder-client-req",
			url:  "https://www.thunderclient.com/welcome",
			result: map[string]interface{}{
				"message": "Welcome to Thunder Client",
				"about":   "Lightweight Rest API Client for VSCode",
			},
		},
		{
			name: "json-placeholder",
			url:  "https://jsonplaceholder.typicode.com/posts/1",
			result: map[string]interface{}{
				"userId": 1., // JSON parses number into float64, not int
				"id":     1.,
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			res, err := http.Get(tc.url)
			if err != nil {
				log.Errorf("request with err: %v", err)
				t.Fail()
			}

			body_map, err := FromHttpBody2Map(res.Body)
			if err != nil {
				log.Errorf("cannot parse body: %v", err)
				t.Fail()
			}

			for k, v := range tc.result {
				found, ok := body_map[k]
				if !ok {
					log.Errorf("map doesn't have key %s ", k)
					t.Fail()
				}

				if found != v {
					log.Errorf("value from map is not valid: %v != %v", v, found)
					t.Fail()
				}
			}

			log.Infof("Test succeded: %v!", tc.name)
		})
	}
}

func Test_Body2Struct(t *testing.T) {
	type TestFakeAPI struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	tests := []struct {
		name   string
		url    string
		store  TestFakeAPI
		result TestFakeAPI
	}{
		{
			name:  "json-placeholdere-posts-1",
			url:   "https://jsonplaceholder.typicode.com/posts/1",
			store: TestFakeAPI{},
			result: TestFakeAPI{
				Title: "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
				Body:  "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto",
			},
		},
		{
			name:  "json-placeholder-posts-2",
			url:   "https://jsonplaceholder.typicode.com/posts/2",
			store: TestFakeAPI{},
			result: TestFakeAPI{
				Title: "qui est esse",
				Body:  "est rerum tempore vitae\nsequi sint nihil reprehenderit dolor beatae ea dolores neque\nfugiat blanditiis voluptate porro vel nihil molestiae ut reiciendis\nqui aperiam non debitis possimus qui neque nisi nulla",
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			res, err := http.Get(tc.url)
			if err != nil {
				log.Errorf("request with err: %v", err)
				t.Fail()
			}

			err = Body2Struct(res.Body, &tc.store)
			if err != nil {
				log.Errorf("cannot parse body: %v", err)
				t.Fail()
			}

			if tc.store.Body != tc.result.Body || tc.store.Title != tc.result.Title {
				log.Errorf("value from map is not valid: %v != %v", tc.result, tc.store)
				t.Fail()
			}

		})
	}
}
