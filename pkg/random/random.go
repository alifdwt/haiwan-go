package random

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/gosimple/slug"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gopkg.in/loremipsum.v1"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomInt(min, max int) int {
	return int(rand.Intn(max-min) + min)
}

func RandomMoney() int {
	return int(RandomInt(100, 1000))
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwner() string {
	str := RandomString(6) + " " + RandomString(6)
	return cases.Title(language.English).String(str)
}

func RandomUsername(owner string) string {
	return slug.Make(owner)
}

func RandomEmail(username string) string {
	return username + "@example.com"
}

func RandomBirthDate(age int) time.Time {
	return time.Now().AddDate(-age, 0, 0)
}

// func RandomProfilePictureUrl() string {
// 	var gender string
// 	randomNumber := rand.Intn(2)
// 	if randomNumber == 0 {
// 		gender = "women"
// 	} else {
// 		gender = "men"
// 	}

// 	return fmt.Sprintf("https://randomuser.me/api/portraits/%s/%d.jpg", gender, RandomInt(0, 99))
// }

func RandomProductPhotoUrl() string {
	return fmt.Sprintf("https://picsum.photos/id/%d/720", RandomInt(0, 100))
}

func RandomFaviconUrl() string {
	return fmt.Sprintf("https://picsum.photos/id/%d/32/32", RandomInt(0, 100))
}

func RandomSliderPhotoUrl() string {
	return fmt.Sprintf("https://picsum.photos/id/%d/1208/302", RandomInt(0, 100))
}

func RandomWords(wordCount int) string {
	return loremipsum.New().Words(wordCount)
}

func RandomPassword() string {
	return RandomString(12)
}
