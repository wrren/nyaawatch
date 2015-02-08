package main

import (
	"regexp"
	"errors"
	"strconv"
)

func ParseSeries( title string, regexes []string ) (Series,error) {
	var series Series

	for _, r := range regexes {
		regex, err := regexp.Compile( r )

		if err != nil {
			continue
		}

		match  := regex.FindStringSubmatch( title )

		if match == nil {
			continue
		}

		result := make( map[string]string )
		
		for i, name := range regex.SubexpNames() {
			result[name] = match[i]
		}
		
		series.Name 		= result["title"]
		series.Subber		= result["subber"]
		series.Quality 		= result["quality"]
		series.Episode, _	= strconv.Atoi( result["episode"] )

		return series, nil
	}

	return series, errors.New( "Failed to Match Any Regex" )
}