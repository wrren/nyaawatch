package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"io"
	"net/http"
	"strings"
	"path/filepath"
	"time"
	"strconv"
	"sync"
)

var Downloads []Series

func alreadyDownloaded( episode Series ) bool {
	for _, download := range Downloads {
		if 	strings.EqualFold( download.Name, episode.Name ) 	&& 
			strings.EqualFold( download.Subber, episode.Subber)	&&
			strings.EqualFold( download.Quality, episode.Quality ) 	&&
			download.Episode == episode.Episode {
			return true;
		}
	}

	return false
}

func downloaded( episode Series ) {
	if !alreadyDownloaded( episode ) {
		Downloads = append( Downloads, episode )
	}
}

func refresh( config WatchConfig, wg *sync.WaitGroup ) {
	resp, err := http.Get( config.URL )

	if err != nil {
		fmt.Println( "Could Not Retrieve RSS Feed: ", err )
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll( resp.Body )

	if err != nil {
		fmt.Println( "Could Not Read RSS Feed Body: ", err )
		return
	}

	rss, err := ReadRSS( body )

	if err != nil {
		fmt.Println( "Could Not Parse RSS Feed Body: ", err )
		return
	}

	for _, e := range rss.Items.ItemList {
		series, err := ParseSeries( e.Title, config.Regexes )

		if err != nil || alreadyDownloaded( series ) {
			continue
		}

		for _, s := range config.Series {
			if strings.EqualFold( s.Name, series.Name ) && strings.EqualFold( s.Subber, series.Subber ) && strings.EqualFold( s.Quality, series.Quality ){
				path := filepath.Join( config.Directory, s.Subber + "." + s.Name + "." +  strconv.Itoa( series.Episode ) + "." + s.Quality + ".torrent" )

				out, err := os.Create( path )
				defer out.Close()

				if err != nil {
					fmt.Println( "Failed to create file at ", path, ": ", err )
					continue
				}

				resp, err := http.Get( e.Link )
				defer resp.Body.Close()

				_, err = io.Copy( out, resp.Body )

				if err != nil {
					fmt.Println( "Failed to write file at ", path, ": ", err )
					continue
				}

				downloaded( series )

				fmt.Println( "Downloaded Torrent for ", s.Name, " to ", path )
			}

		}
	}

	time.Sleep( time.Second * time.Duration( config.Refresh ) )
}

func main() {
	Downloads = make( []Series, 0, 100 )
	args := os.Args[1:]

	if len( args ) == 0 {
		fmt.Println( "usage:" )
		fmt.Println( "\tnyaawatch config_file..." )

		return
	}

	var wg sync.WaitGroup

	for _, e := range args {
		fmt.Println( "Reading Configuration from ", e )

		b, err := ioutil.ReadFile( e )

		if err != nil {
			fmt.Println( "Could Not Read Config File: ", err )
			return
		}

		config, err := ReadConfig( b )

		if err != nil {
			fmt.Println( "Could Not Parse Config File: ", err )
			return
		}

		wg.Add( 1 )
		go refresh( config, &wg )
	}

	wg.Wait()
}