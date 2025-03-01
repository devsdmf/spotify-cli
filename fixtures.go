package main

import (
	"fmt"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
)

// NewDebugClient returns Client which can be used in debug mode or in tests,
// It does not communicate with Spotify API.
func NewDebugClient() SpotifyClient {
	return DebugClient{
		player:           &DebugPlayer{},
		searcher:         &DebugSearcher{},
		userAlbumFetcher: &DebugUserAlbumFetcher{},
	}
}

// DebugClient is a dummy struct used when running in debug mode
type DebugClient struct {
	player
	searcher
	userAlbumFetcher
}

type DebugPlayer struct {
}

// Play is a dummy implementation used when running in debug mode
func (dp DebugPlayer) Play() error {
	return nil
}

// PlayOpt is a dummy implementation used when running in debug mode
func (dp DebugPlayer) PlayOpt(opt *spotify.PlayOptions) error {
	return nil
}

type DebugSearcher struct{}

// Search is a dummy implementation used when running in debug mode
func (ds DebugSearcher) Search(query string, t spotify.SearchType) (*spotify.SearchResult, error) {
	return nil, nil
}

type DebugUserAlbumFetcher struct{}

// CurrentUsersAlbumsOpt is a dummy implementation used when running in debug mode
func (debugFetcher DebugUserAlbumFetcher) CurrentUsersAlbumsOpt(options *spotify.Options) (*spotify.SavedAlbumPage, error) {
	return &spotify.SavedAlbumPage{
		Albums: constructNSpotifySavedAlbums(visibleAlbums * 3),
	}, nil
}

func constructNSpotifySavedAlbums(n int) []spotify.SavedAlbum {
	albums := make([]spotify.SavedAlbum, 0)
	for i := 1; i <= n; i++ {
		album := spotify.SavedAlbum{}
		album.Name = fmt.Sprintf("Album Name %d", i)
		album.Artists = []spotify.SimpleArtist{spotify.SimpleArtist{Name: fmt.Sprintf("Artist Name %d", i)}}
		albums = append(albums, album)
	}
	return albums
}

// Previous is a dummy implementation used when running in debug mode
func (fc DebugClient) Previous() error {
	return nil
}

// Pause is a dummy implementation used when running in debug mode
func (fc DebugClient) Pause() error {
	return nil
}

// Next is a dummy implementation used when running in debug mode
func (fc DebugClient) Next() error {
	return nil
}

// PlayerCurrentlyPlaying is a dummy implementation used when running in debug mode
func (fc DebugClient) PlayerCurrentlyPlaying() (*spotify.CurrentlyPlaying, error) {
	return &spotify.CurrentlyPlaying{Item: &spotify.FullTrack{SimpleTrack: spotify.SimpleTrack{
		Name:    "Currently Playing Song",
		Artists: []spotify.SimpleArtist{{Name: "Currently Playing Artist"}}},
		Album: spotify.SimpleAlbum{Name: "Currently Playing Album"}},
	}, nil
}

// PlayerDevices is a dummy implementation used when running in debug mode
func (fc DebugClient) PlayerDevices() ([]spotify.PlayerDevice, error) {

	return []spotify.PlayerDevice{
		{Name: "iPad", Type: "Tablet"},
		{Name: "iPhone", Type: "Smarthphone"},
		{Name: "Mac", Type: "App Player"},
	}, nil
}

// TransferPlayback is a dummy implementation used when running in debug mode
func (fc DebugClient) TransferPlayback(id spotify.ID, play bool) error {
	return nil
}

// CurrentUser is a dummy implementation used when running in debug mode
func (fc DebugClient) CurrentUser() (*spotify.PrivateUser, error) {
	return &spotify.PrivateUser{}, nil
}

// Token is a dummy implementation used when running in debug mode
func (fc DebugClient) Token() (*oauth2.Token, error) {
	return &oauth2.Token{}, nil
}
