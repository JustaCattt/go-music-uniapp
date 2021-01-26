package dataGetter

import (
	"fmt"
	"testing"
)

func TestDataGetter(t *testing.T) {
	var artists = GetArtists()
	for _, artist := range artists {
		var songs = GetAllSongsByArtistId(artist.Id)
		fmt.Printf("\n%s全部的歌(共%d首)：\n", artist.Name, len(songs))
		for index, song := range songs {
			fmt.Printf("《%s》、", song.Name)
			if index%10 == 0 && index != 0 {
				fmt.Println()
			}
		}
		fmt.Println()
	}
}

func TestMap(t *testing.T) {
	m := MapArtistToSongs()
	fmt.Println(m["薛之谦"])
}
