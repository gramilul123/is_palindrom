package word

import (
	"isPalindrom/db"

	_ "github.com/lib/pq"
)

// Word
type Word struct {
	ID          int    `db:"id"`
	Word        string `db:"word"`
	IsPalindrom bool   `db:"is_palindrom"`
}

// Save
func (word Word) Save() error {

	dbpg := db.GetDBConnect()

	if dbpg.Error != nil {
		return dbpg.Error
	}

	_, err := dbpg.Connect.NamedExec(`INSERT INTO public.words (word,is_palindrom) VALUES (:word,:is_palindrom)`, word)

	if err != nil {
		return err
	}

	return nil

}

// GetWordsList
func GetWordsList() ([]Word, error) {

	dbpg := db.GetDBConnect()
	words := []Word{}
	err := dbpg.Connect.Select(&words, "SELECT * FROM words ORDER BY id")

	return words, err
}

// IsPalindrom
func IsPalindrom(word string) bool {
	result := true
	lenWord := len(word)
	for i := 0; i < lenWord/2; i++ {
		if word[i] != word[lenWord-1-i] {
			result = false
			break
		}
	}

	return result
}
