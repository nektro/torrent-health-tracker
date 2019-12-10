# Torrent Health Tracker
![loc](https://tokei.rs/b1/github/nektro/torrent-health-tracker)
[![license](https://img.shields.io/github/license/nektro/torrent-health-tracker.svg)](https://github.com/nektro/torrent-health-tracker/blob/master/LICENSE)
[![astheno discord](https://img.shields.io/discord/551971034593755159.svg)](https://discord.gg/P6Y4zQC)
[![the-eye discord](https://img.shields.io/discord/302796547656253441.svg)](https://discord.gg/the-eye)
[![paypal](https://img.shields.io/badge/donate-paypal-009cdf)](https://paypal.me/nektro)
[![goreportcard](https://goreportcard.com/badge/github.com/nektro/torrent-health-tracker)](https://goreportcard.com/report/github.com/nektro/torrent-health-tracker)

**Torrent Health Tracker** is a way to provide a web interface into the statistics on the health of a collection of torrents. This is accomplished by looking at assembling tracker stats and periodically re-pinging the swarm. Simplely provide `torrent-health-tracker` with a list of trackers and magnet links and it'll do its thing!

> Note: currently only supports HTTP trackers.

## Development

### Prerequisites
- The Go Language 1.12+ (https://golang.org/dl/)

### Installing
Run
```
$ go get -u -v github.com/nektro/torrent-health-tracker
```
and then make your way to `$GOPATH/src/github.com/nektro/torrent-health-tracker/`.

Once there, add your `--tracker` and `--magnet` URLs and run:
```
$ ./start.sh
```

## Deployment
Coming soon!

## Built With
Output of `go list -f '{{ join .Imports "\n" }}'` sans stdlib packages.
- github.com/anacrolix/torrent/bencode
- github.com/nektro/go-util/alias
- github.com/nektro/go-util/util
- github.com/nektro/go.etc
- github.com/nektro/torrent-health-tracker/statik
- github.com/rakyll/statik/fs
- github.com/spf13/pflag

## Contributing
[![issues](https://img.shields.io/github/issues/nektro/torrent-health-tracker.svg)](https://github.com/nektro/torrent-health-tracker/issues)
[![pulls](https://img.shields.io/github/issues-pr/nektro/torrent-health-tracker.svg)](https://github.com/nektro/torrent-health-tracker/pulls)

We listen to issues all the time right here on GitHub. Labels are extensively to show the progress through the fixing process. Question issues are okay but make sure to close the issue when it has been answered! Off-topic and '+1' comments will be deleted. Please use post/comment reactions for this purpose.

When making a pull request, please have it be associated with an issue and make a comment on the issue saying that you're working on it so everyone else knows what's going on :D

## Contact
- hello@nektro.net
- Meghan#2032 on discordapp.com
- https://twitter.com/nektro

## License
MIT
