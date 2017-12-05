package model

import (
	"testing"
	"os"
	"encoding/json"
	"github.com/nu7hatch/gouuid"
	"time"
)

type csvDTO struct {
	Title     string
	Body      string
	Author    string
	CreatedOn string
}

// TODO.. Clean this test breaking in subtests

func TestTopicRepo(t *testing.T) {


	topicsRepo := NewTopicsMemory()

	csvFile, err := os.Open("testdata/random_topics.json")
	if err != nil {
		panic("Cannot open test file")
	}
	defer csvFile.Close()

	encoder := json.NewDecoder(csvFile)

	fileItems := make([]csvDTO, 0)
	encoder.Decode(&fileItems)

	// Load all items in memory
	generatedIds := make([]*uuid.UUID, 0, len(fileItems))
	for _, item := range (fileItems) {
		id, _ := uuid.NewV4()
		topic := NewTopic(id)
		topic.SetTitle(item.Title)
		topic.SetBody(item.Body)
		createdOn, err := time.Parse("2006-01-02 15:04:05", item.CreatedOn, )
		if err != nil {
			t.Fatalf("Bad date format in testdata <%s>", item.CreatedOn)
		}
		topic.SetCreationDate(createdOn)
		topic.SetModDate(createdOn)

		err = topicsRepo.Add(topic)
		if err != nil {
			t.Error(err)
		}
		generatedIds = append(generatedIds, topic.ID())

	}

	// Let's see if we got all the items loaded
	for _, id := range generatedIds {
		topic, err := topicsRepo.Get(id)
		if err != nil {
			t.Error(err)
		}
		if topic.ID().String() != id.String() {
			t.Errorf("Topic id differ <%v>, <%v>", topic.ID(), id)
		}
	}

	// Let's see if all are ordered by time
	last := time.Now().Add(100 * 12 * 30 * 24 * time.Hour)
	for _, topic := range topicsRepo.GetByCreatedDateDesc(0, 0) {
		if last.Before(topic.CreationDate()) {
			t.Errorf("Error getting ordered topics by date. Expecting a date smaller than <%v>, got <%v>", last, topic.CreationDate())
		}
		last = topic.CreationDate()
	}

	// Let's test from and limit
	topics := topicsRepo.GetByCreatedDateDesc(0, 30)
	if len(topics) != 30 {
		t.Errorf("Expecting 30 elements. Got <%d>", len(topics))
	}

	allTopics := topicsRepo.GetByCreatedDateDesc(0, 0)
	allTopicsByStep := make([]*Topic, 0, len(allTopics))
	for i := 0; i <= len(allTopics); i += 30 {
		allTopicsByStep = append(allTopicsByStep, topicsRepo.GetByCreatedDateDesc(i, 30)...)
	}

	for i := 0; i < len(allTopics); i++ {
		if allTopics[i].ID().String() != allTopicsByStep[i].ID().String() {
			t.Errorf("Topics differ <%v>, <%v>", allTopics[i].ID(), allTopicsByStep[i].ID())
		}
	}
	csvFile.Close()

}
