package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/zmb3/spotify"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
)

var (
	auth = spotify.NewAuthenticator(
		redirectURI,
		spotify.ScopeUserReadPrivate,
		spotify.ScopeUserReadCurrentlyPlaying,
		spotify.ScopeUserReadPlaybackState,
		spotify.ScopeUserModifyPlaybackState,
		spotify.ScopeUserLibraryRead,
		// Used for Web Playback SDK
		"streaming",
		spotify.ScopeUserReadBirthdate,
		spotify.ScopeUserReadEmail,
	)
	ch           = make(chan *spotify.Client)
	state        = uuid.New().String()
	redirectURI  = "http://localhost:8888/spotify-cli"
	clientId     = os.Getenv("SPOTIFY_CLIENT_ID")
	clientSecret = os.Getenv("SPOTIFY_SECRET")
)

// authenticate authenticate user with Sotify API
func authenticate() SpotifyClient {
	http.HandleFunc("/spotify-cli", authCallback)
	go http.ListenAndServe(":8888", nil)

	auth.SetAuthInfo(clientId, clientSecret)
	url := auth.AuthURL(state)

	err := openBroswerWith(url)
	if err != nil {
		log.Fatal(err)
	}

	client := <-ch

	token, _ := client.Token()
	insertTokenToTemplate(token.AccessToken)

	return client
}

// openBrowserWith open browsers with given url
func openBroswerWith(url string) error {
	var err error
	switch runtime.GOOS {
	case "darwin":
		err = exec.Command("open", url).Start()
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	default:
		err = fmt.Errorf("Sorry, %v OS is not supported", runtime.GOOS)
	}
	return err
}

// authCallback is a function to by Spotify upon successful
// user login at their site
func authCallback(w http.ResponseWriter, r *http.Request) {
	token, err := auth.Token(state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusNotFound)
		return
	}
	client := auth.NewClient(token)
	ch <- &client

	user, err := client.CurrentUser()
	fmt.Fprintf(w, "<h1>Logged into spotify cli as:</h1>\n<p>%v</p>", user.DisplayName)
}

func insertTokenToTemplate(token string) {
	t := template.New("web player")
	t, _ = template.ParseFiles("index_tmpl.html")
	tokenToInsert := struct {
		Token string
	}{
		Token: token,
	}
	f, err := os.Create("index.html")
	check(err)
	defer f.Close()

	t.Execute(f, tokenToInsert)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
