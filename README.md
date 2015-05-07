# nyaawatch
Watches nyaa.se for releases matching a filter based on subber, series and quality level and downloads the torrent file to a designated directory

# Build

  Install [golang](https://golang.org/doc/install)
  Navigate to the directory where you checked out nyaawatch
  Call `go build`to build the application 
  
# Run

Once you've built the nyaawatch binary, run it with a single argument specifying the path to the configuration file it should use. An example configuration file is included under the `config` directory.

# Configuration

The first argument to the program should be the path to the watch configuration file. The file should be laid out like example/config.json. 

`url` points to the RSS feed URL to be read. 

`refresh` specifies the interval between RSS reads in seconds. Don't set this too low or you'll hammer the server.

`series` is an array of series objects that we would like to retrieve, with title, quality and subber specified.

`directory` refers to the torrent file download directory; when a watched series is found, the torrent file for the episode will be downloaded to this directory. If you're using rtorrent, uTorrent or a client capable of watching a directory for new torrent files then this will allow the torrent to be started automatically. 

`regexes` contains a list of regular expression patterns to be used when parsing the RSS feed. If the series you're watching
for cannot be parsed, formulate a regular expression that can match the RSS item title with groupings for subber, title, episode and quality, then add it
to the list.
