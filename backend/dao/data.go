package dao

import (
	"context"
	"fmt"
	"log"
	"time"
	"unicode"

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
	Id   string
	Word string
	Date string
	User User
}

type WordWrapper struct {
	Word       Word
	Letters    []string
	LettersMap map[string]struct{}
}

const usersCollectionName = "users"
const wordsCollectionName = "words"

var wordMap = map[string]WordWrapper{}
var userMap = map[string]User{}

var empty struct{}
var isDiacritic = map[rune]struct{}{'\u0B82': empty,
	'\u0BBE': empty,
	'\u0BBF': empty,
	'\u0BC0': empty,
	'\u0BC1': empty,
	'\u0BC2': empty,
	'\u0BC6': empty,
	'\u0BC7': empty,
	'\u0BC8': empty,
	'\u0BCA': empty,
	'\u0BCB': empty,
	'\u0BCC': empty,
	'\u0BCD': empty,
	'\u0BD7': empty,
}

func splitWordGetLetters(word string) ([]string, error) {
	var letters []string

	for _, r := range word {
		if !unicode.Is(unicode.Tamil, r) {
			return nil, fmt.Errorf("Non-Tamil word")
		}

		if _, yes := isDiacritic[r]; yes {
			if len(letters) == 0 {
				return nil, fmt.Errorf("invalid diacritic position")
			}
			letters[len(letters)-1] += string(r)
		} else {
			letters = append(letters, string(r))
		}
	}

	return letters, nil
}

func getWordLettersMap(wantLetters []string) map[string]struct{} {
	wantLettersMap := make(map[string]struct{})
	for _, letter := range wantLetters {
		wantLettersMap[letter] = empty
	}
	return wantLettersMap
}

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
	/*if user, ok := userMap[id]; ok {
		log.Println("Found user in cache")
		return user, nil
	}*/
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

	// Check if its a splittable word
	letters, err := splitWordGetLetters(word.Word)
	if err != nil {
		return "", err
	}

	// Check if word already exists for the day
	w, err := GetWordForTheDay(word.Date)
	if err != nil {
		err = fmt.Errorf("failed to check if word exists for the day: %v", err)
		return "", err
	}
	if w.Word.Id != "" {
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

	var wordWrapper WordWrapper = WordWrapper{
		Word:       word,
		Letters:    letters,
		LettersMap: getWordLettersMap(letters),
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
		var word Word
		doc.DataTo(&word)
		word.Id = doc.Ref.ID
		if word.Word != "" {
			letters, err := splitWordGetLetters(word.Word)
			if err != nil {
				return WordWrapper{}, err
			}
			var wordWrapper WordWrapper = WordWrapper{
				Word:       word,
				Letters:    letters,
				LettersMap: getWordLettersMap(letters),
			}
			wordMap[word.Date] = wordWrapper
			return wordWrapper, nil
		}
	}
	return WordWrapper{}, nil
}
