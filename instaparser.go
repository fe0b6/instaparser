package instaparser

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

var (
	dataReg *regexp.Regexp
)

func init() {
	dataReg = regexp.MustCompile(`>window._sharedData = (.*);</script>`)
}

func GetLikes(lnk string) (pi PostInfo, err error) {
	// Читаем страницу
	resp, err := http.Get(lnk)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		log.Println("[error]", err)
		return
	}

	// Что-то пошло не так
	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		log.Println("[error]", resp.Status, resp.StatusCode)
		return
	}

	// Читаем ответ
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("[error]", err)
		return
	}

	arr := dataReg.FindStringSubmatch(string(content))
	if len(arr) == 0 {
		err = errors.New("can't find json data")
		return
	}

	var pd PostData
	err = json.Unmarshal([]byte(arr[1]), &pd)
	if err != nil {
		log.Println("[error]", err)
		return
	}

	if len(pd.EntryData.PostPage) == 0 {
		err = errors.New("bad json data")
		return
	}

	pi.Likes = pd.EntryData.PostPage[0].Graphql.ShortcodeMedia.EdgeMediaPreviewLike.Count
	pi.Comments = pd.EntryData.PostPage[0].Graphql.ShortcodeMedia.EdgeMediaToComment.Count
	pi.Id = pd.EntryData.PostPage[0].Graphql.ShortcodeMedia.Id
	pi.Published = pd.EntryData.PostPage[0].Graphql.ShortcodeMedia.TakenAtTimestamp

	if len(pd.EntryData.PostPage[0].Graphql.ShortcodeMedia.EdgeMediaToCaption.Edges) > 0 {
		pi.Text = pd.EntryData.PostPage[0].Graphql.ShortcodeMedia.EdgeMediaToCaption.Edges[0].Node.Text
	}

	pi.Owner = pd.EntryData.PostPage[0].Graphql.ShortcodeMedia.Owner

	return
}

func GetProfile(lnk string) (pi ProfileInfo, err error) {
	// Читаем страницу
	resp, err := http.Get(lnk)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		log.Println("[error]", err)
		return
	}

	// Что-то пошло не так
	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		log.Println("[error]", resp.Status, resp.StatusCode)
		return
	}

	// Читаем ответ
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("[error]", err)
		return
	}

	arr := dataReg.FindStringSubmatch(string(content))
	if len(arr) == 0 {
		err = errors.New("can't find json data")
		return
	}

	var pd ProfileData
	err = json.Unmarshal([]byte(arr[1]), &pd)
	if err != nil {
		log.Println("[error]", err)
		return
	}

	if len(pd.EntryData.ProfilePage) == 0 {
		err = errors.New("bad json data")
		return
	}

	pi.FollowedBy = pd.EntryData.ProfilePage[0].Graphql.User.FollowedBy.Count

	return
}
