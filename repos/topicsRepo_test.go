package repos_test

import (
	"encoding/json"
	"github.com/nu7hatch/gouuid"
	"github.com/x1m3/Tertulia/model"
	"github.com/x1m3/Tertulia/repos"
	"os"
	"testing"
	"time"

	"errors"
	"fmt"
	"github.com/asdine/storm"
)

type csvDTO struct {
	Title     string
	Body      string
	Author    string
	CreatedOn string
}

func TestTopicsRepo(t *testing.T) {

	db, err := storm.Open("test.db")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		db.Close()
		os.Remove("test.db")
	}()

	repo := repos.NewTopicsStorm(db)

	topics, err := loadTopicsFromCSV("testdata/random_topics.json")
	if err != nil {
		t.Error(err)
	}

	// Adding all topics to repo
	for _, topic := range topics.GetByCreatedDateDesc(0, 0) {
		repo.Add(topic)
	}

	// Let's ckeck if all entries exists
	for _, topic := range topics.GetByCreatedDateDesc(0, 0) {
		dto, err := repo.Get(topic.ID())
		if err != nil {
			t.Error(err)
		}
		if topic.ID().String() != dto.Id.String() {
			t.Errorf("Topic ID differ. Expecting <%v>, got <%v>", topic.ID().String(), dto.Id.String())
		}
		if topic.Author().ID().String() != dto.Author.String() {
			t.Errorf("Author ID differ. Expecting <%v>, got <%v>", topic.Author().ID().String(), dto.Author.String())
		}
		if topic.Body() != dto.Body {
			t.Errorf("Topic Body differ. Expecting <%v>, got <%v>", topic.Body(), dto.Body)
		}
		if topic.CreationDate() != dto.CreationDate {
			t.Errorf("Topic creation date differ. Expecting <%v>, got <%v>", topic.CreationDate(), dto.CreationDate)
		}
		if topic.ModDate() != dto.ModDate {
			t.Errorf("Topic modification date differ. Expecting <%v>, got <%v>", topic.ModDate(), dto.ModDate)
		}
	}

	// Let's check for a non existing topic
	randUUID, _ := uuid.NewV4()
	dto, err := repo.Get(randUUID)
	if dto != nil || err != model.ErrNotFound {
		t.Error("expecting nil value and err= model.ErrNotFound for non existing item")
	}

	// Let's update a non existing topic
	randUUID, _ = uuid.NewV4()
	rt := model.NewTopic(randUUID)
	rt.SetAuthor(model.NewPerson("lala"))
	err = repo.Update(rt)
	if err != model.ErrNotFound {
		t.Error("expecting err= model.ErrNotFound updating a non existing item")
	}

	// Let's remove a non existing topic
	err = repo.Delete(rt)
	if err != model.ErrNotFound {
		t.Error("expecting err= model.ErrNotFound removing a non existing item")
	}

	// Let's update a topic
	list := topics.GetByCreatedDateDesc(0, 1)
	atopic := &model.Topic{}
	*atopic = *list[0]
	atopic.SetTitle("lala")

	err = repo.Update(atopic)
	if err != nil {
		t.Error(err)
	}

	storedTopic, err := repo.Get(list[0].ID())
	if err != nil {
		t.Error(err)
	}
	if storedTopic.Title != "lala" {
		t.Error("Fail updating topic title. Expecting <lala>, got <%s>", storedTopic.Title)
	}

	// Let's remove all topics
	for _, topic := range topics.GetByCreatedDateDesc(0, 0) {
		repo.Delete(topic)
	}
	for _, topic := range topics.GetByCreatedDateDesc(0, 0) {
		dto, err := repo.Get(topic.ID())
		if dto != nil || err != model.ErrNotFound {
			t.Error("expecting nil value and err= model.ErrNotFound for non existing item")
		}
	}
}

func loadTopicsFromCSV(filename string) (*model.TopicsMemory, error) {

	topicsRepo := model.NewTopicsMemory()

	csvFile, err := os.Open("testdata/random_topics.json")
	if err != nil {
		panic("Cannot open test file")
	}
	defer csvFile.Close()

	encoder := json.NewDecoder(csvFile)

	fileItems := make([]csvDTO, 0)
	encoder.Decode(&fileItems)

	for _, item := range fileItems {
		id, _ := uuid.NewV4()
		topic := model.NewTopic(id)
		topic.SetTitle(item.Title)
		topic.SetBody(item.Body)
		createdOn, err := time.Parse("2006-01-02 15:04:05", item.CreatedOn)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Bad date format in testdata <%s>", item.CreatedOn))
		}
		topic.SetAuthor(model.NewPerson(item.Author))
		topic.SetCreationDate(createdOn)
		topic.SetModDate(createdOn)

		err = topicsRepo.Add(topic)
		if err != nil {
			return nil, err
		}
	}
	return topicsRepo, nil
}
