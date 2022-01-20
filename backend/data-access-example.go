package main

import (
	"log"

	dao "example.com/tamil-wordle/dao"
)

func main() {
	/*results, err := dao.ListUsers()
	if err != nil {
		log.Fatalf("Failed to list users: %v", err)
	}
	log.Printf("Users: %v", results)*/

	/*id, err := dao.CreateUser(dao.User{
		Name:          "TestUser2",
		TwitterHandle: "testuser2",
		PublicKey:     "public-key",
		Active:        true,
	})
	log.Printf("Created user with id: %s", id)*/

	/*err = dao.MarkUserActive("FKC1u8Us13YLdYnj3Wko")
	if err != nil {
		log.Fatalf("Failed to mark user active: %v", err)
	}
	log.Printf("Marked user active")*/

	/*err = dao.UpdatePublicKey("FKC1u8Us13YLdYnj3Wko", "new-public-key")
	if err != nil {
		log.Fatalf("Failed to update public key: %v", err)
	}
	log.Printf("Updated public key")*/

	word, err := dao.GetWordForTheDay("2022-01-21")
	if err != nil {
		log.Fatalf("Failed to get word for the day: %v", err)
	}
	log.Printf("Word for the day: %v", word.User.Name)

	// should read from cache
	word, err = dao.GetWordForTheDay("2022-01-22")
	if err != nil {
		log.Fatalf("Failed to get word for the day: %v", err)
	}
	log.Printf("Word for the day: %v", word.User.Name)

	/*id, err := dao.AddWord(dao.Word{
		Word:   "தமிழ்",
		Date:   "2022-01-22",
		UserId: "FKC1u8Us13YLdYnj3Wko",
	})
	if err != nil {
		log.Fatalf("Failed to add word: %v", err)
	}
	log.Printf("Added word with id: %s", id)*/
}
