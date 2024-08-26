package speller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
)

// Структура ответа спеллера
type SpellerResponse []struct {
	Code int      `json:"code"`
	Pos  int      `json:"pos"`
	Row  int      `json:"row"`
	Col  int      `json:"col"`
	Len  int      `json:"len"`
	Word string   `json:"word"`
	Sugg []string `json:"s"`
}

// Исправляет орфографию
func Spell(text string) (string, error) {

	params := url.Values{}
	params.Add("text", text)

	apiURL := "https://speller.yandex.net/services/spellservice.json/checkText?" + params.Encode()

	resp, err := http.Get(apiURL)
	if err != nil {
		log.Println("Failed to request Speller: ", err)
		return text, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Failed to read body: ", err)
		return text, err
	}

	var results SpellerResponse
	err = json.Unmarshal(body, &results)
	if err != nil {
		log.Println("Failed to unmarshal body: ", err)
		return text, err
	}

	if len(results) == 0 {
		log.Println("No spelling errors")
		return text, nil
	}

	// Исправляем ошибочные слова
	runes := []rune(text)
	ans := []rune{}
	offset := 0
	for _, result := range results {
		pos := result.Pos
		len := result.Len
		word := result.Sugg[0]
		ans = append(ans, runes[offset:pos]...)
		ans = append(ans, []rune(word)...)
		offset = pos + len
	}
	ans = append(ans, runes[offset:]...)

	log.Println("Spelling errors corrected successfully")
	return string(ans), nil
}
