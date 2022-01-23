package dao

import (
	"context"
	"fmt"
	"log"
	"time"

	firestore "cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/iterator"
)

type User struct {
	Id            string
	Name          string
	TwitterHandle string
	PublicKey     string
	Active        bool
}

type Word struct {
	Id     string
	Word   string
	Date   string
	UserId string
}

type WordWrapper struct {
	Id   string
	Word Word
	User User
}

const usersCollectionName = "users"
const wordsCollectionName = "words"

var wordMap = map[string]WordWrapper{}
var userMap = map[string]User{}

func openClient() (context.Context, *firestore.Client, error) {
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		err = fmt.Errorf("error initializing app: %v", err)
		return nil, nil, err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		err = fmt.Errorf("error initializing Firestore: %v", err)
		return nil, nil, err
	}
	return ctx, client, nil
}

func ListUsers() ([]User, error) {
	ctx, client, err := openClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()
	iter := client.Collection(usersCollectionName).Documents(ctx)
	results := []User{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			err = fmt.Errorf("failed to iterate: %v", err)
			return nil, err
		}
		var userObj User
		doc.DataTo(&userObj)
		userObj.Id = doc.Ref.ID
		results = append(results, userObj)
	}
	return results, nil
}

func CreateUser(user User) (string, error) {
	ctx, client, err := openClient()
	if err != nil {
		return "", err
	}
	defer client.Close()

	ref := client.Collection(usersCollectionName).NewDoc()
	user.Id = ref.ID
	_, err = ref.Set(ctx, user)
	if err != nil {
		err = fmt.Errorf("failed to add user: %v", err)
		return "", err
	}
	userMap[user.Id] = user
	return ref.ID, err
}

func MarkUserActive(id string) error {
	ctx, client, err := openClient()
	if err != nil {
		return err
	}
	defer client.Close()

	// Inactivate existing user
	iter := client.Collection(usersCollectionName).Where("Active", "==", true).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to iterate existing active users: %v", err)
		}
		var userObj User
		doc.DataTo(&userObj)
		userObj.Id = doc.Ref.ID
		userObj.Active = false
		_, err = doc.Ref.Set(ctx, userObj)
		if err != nil {
			return fmt.Errorf("failed to inactivate existing active user: %v", err)
		}
		userMap[userObj.Id] = userObj
	}

	// Activate current user
	doc, err := client.Collection(usersCollectionName).Doc(id).Get(ctx)
	if err != nil {
		return fmt.Errorf("failed to get user: %v", err)
	}
	var userObj User
	doc.DataTo(&userObj)
	userObj.Id = doc.Ref.ID
	userObj.Active = true
	_, err = doc.Ref.Set(ctx, userObj)
	if err != nil {
		return fmt.Errorf("failed to activate user: %v", err)
	}
	userMap[userObj.Id] = userObj
	return nil
}

func UpdatePublicKey(id string, publicKey string) error {
	ctx, client, err := openClient()
	if err != nil {
		return err
	}
	defer client.Close()

	doc, err := client.Collection(usersCollectionName).Doc(id).Get(ctx)
	if err != nil {
		return fmt.Errorf("failed to get user: %v", err)
	}
	var userObj User
	doc.DataTo(&userObj)
	userObj.PublicKey = publicKey
	_, err = doc.Ref.Set(ctx, userObj)
	if err != nil {
		return fmt.Errorf("failed to update public key: %v", err)
	}
	userMap[id] = userObj
	return nil
}

func GetUser(id string) (User, error) {
	if user, ok := userMap[id]; ok {
		log.Println("Found user in cache")
		return user, nil
	}
	ctx, client, err := openClient()
	if err != nil {
		return User{}, err
	}
	defer client.Close()

	doc, err := client.Collection(usersCollectionName).Doc(id).Get(ctx)
	if err != nil {
		return User{}, fmt.Errorf("failed to get user: %v", err)
	}
	var userObj User
	doc.DataTo(&userObj)
	userObj.Id = doc.Ref.ID
	userMap[id] = userObj
	return userObj, nil
}

func AddWord(word Word) (string, error) {
	ctx, client, err := openClient()
	if err != nil {
		return "", err
	}
	defer client.Close()

	// Check if word already exists for the day
	w, err := GetWordForTheDay(word.Date)
	if err != nil {
		err = fmt.Errorf("failed to check if word exists for the day: %v", err)
		return "", err
	}
	if w.Id != "" {
		err = fmt.Errorf("word already exists for the day")
		return "", err
	}

	// Disable setting word for t+2
	t := time.Now().Format("2006-01-02")
	t1 := time.Now().Add(time.Hour * 24).Format("2006-01-02")
	t2 := time.Now().Add(time.Hour * 48).Format("2006-01-02")
	if t != word.Date && t1 != word.Date && t2 != word.Date {
		err = fmt.Errorf("word can only be set for today or tomorrow or day after tomorrow")
		return "", err
	}

	ref := client.Collection(wordsCollectionName).NewDoc()
	word.Id = ref.ID
	_, err = ref.Set(ctx, word)
	if err != nil {
		err = fmt.Errorf("failed to add word: %v", err)
		return "", err
	}
	user, _ := GetUser(word.UserId)
	wordWrapper := WordWrapper{
		Id:   word.Id,
		Word: word,
		User: user,
	}
	wordMap[word.Date] = wordWrapper
	return ref.ID, err
}

func GetWordForTheDay(date string) (WordWrapper, error) {
	if wordWrapper, ok := wordMap[date]; ok {
		log.Println("Found word in cache")
		return wordWrapper, nil
	}
	ctx, client, err := openClient()
	if err != nil {
		return WordWrapper{}, err
	}
	defer client.Close()

	iter := client.Collection(wordsCollectionName).Where("Date", "==", date).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return WordWrapper{}, fmt.Errorf("failed to iterate: %v", err)
		}
		var wordObj Word
		doc.DataTo(&wordObj)
		wordObj.Id = doc.Ref.ID
		if wordObj.Word != "" {
			user, _ := GetUser(wordObj.UserId)
			wordWrapper := WordWrapper{
				Id:   wordObj.Id,
				Word: wordObj,
				User: user,
			}
			wordMap[date] = wordWrapper
			return wordWrapper, nil
		}
	}
	return WordWrapper{}, nil
}
