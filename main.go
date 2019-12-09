package main

import (
	"crypto/rand"
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/anacrolix/torrent/bencode"
	"github.com/nektro/go-util/alias"
	"github.com/nektro/go-util/util"
	etc "github.com/nektro/go.etc"
	"github.com/spf13/pflag"
)

type TrackerResponse struct {
	FailReason  string        `bencode:"failure reason"`
	WarnMsg     string        `bencode:"warning message"`
	Interval    int           `bencode:"interval"`
	MinInterval int           `bencode:"min interval"`
	TrackerID   string        `bencode:"tracker id"`
	Complete    int           `bencode:"complete"`
	Incomplete  int           `bencode:"incomplete"`
	Peers       []TrackerPeer `bencode:"peers"`
	PeersB      string        `bencode:"peers"`
}

type TrackerPeer struct {
	PeerID string `bencode:"peer id"`
	IP     string `bencode:"ip"`
	Port   int    `bencode:"port"`
}

type Torrent struct {
	Hash     string
	Name     string
	Seeders  int
	Leechers int
}

var client = &http.Client{
	Timeout: time.Second * 5,
}

func main() {

	flagTR := pflag.StringArrayP("tracker", "t", []string{}, "")
	flagMG := pflag.StringArrayP("magnet", "m", []string{}, "")

	pflag.Parse()

	//

	trackers := []string{}

	for i, item := range *flagTR {
		urlO, _ := url.Parse(item)

		if urlO.Scheme == "http" {
			req, err := http.NewRequest(http.MethodGet, item, nil)
			if err != nil {
				util.LogWarn("tracker:", alias.F("[%d/%d]:", i+1, len(*flagTR)), "bad:", item, err.Error())
				continue
			}
			res, err := client.Do(req)
			if err != nil {
				util.LogWarn("tracker:", alias.F("[%d/%d]:", i+1, len(*flagTR)), "bad:", item, err.Error())
				continue
			}
			ioutil.ReadAll(res.Body)
			util.Log("tracker:", alias.F("[%d/%d]:", i+1, len(*flagTR)), "good:", item)
			trackers = append(trackers, item)

		} else {
			util.LogError("tracker:", alias.F("[%d/%d]:", i+1, len(*flagTR)), "unknown scheme:", urlO.Scheme)
		}
	}

	//

	torrents := map[string]*Torrent{}

	//

	for i, item := range *flagMG {
		urlO, _ := url.Parse(item)
		qu := urlO.Query()

		btih := qu["xt"][0][9:]
		name := qu["dn"][0]

		seeders := 0
		leechers := 0

		for _, jtem := range trackers {
			s, l, _ := queryTracker(jtem, btih)
			seeders += s
			leechers += l
		}
		t := &Torrent{btih, name, seeders, leechers}
		torrents[btih] = t
		util.Log(alias.F("[%d/%d]:", i+1, len(*flagMG)), alias.F("%+v", t))
	}

	//

	etc.MFS.Add(http.Dir("www"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		etc.WriteHandlebarsFile(r, w, "/index.hbs", map[string]interface{}{
			"torrents": torrents,
		})
	})

	http.ListenAndServe(":80", nil)
}

func queryTracker(urlS string, btih string) (int, int, error) {
	urlP, _ := url.Parse(urlS)

	if urlP.Scheme == "http" {

		v := url.Values{}
		v.Add("info_hash", string(hashToBin(btih)))
		v.Add("peer_id", random(20))
		v.Add("port", "6882")

		u := (urlS + "?" + v.Encode())
		q, _ := http.NewRequest(http.MethodGet, u, nil)
		s, err := client.Do(q)
		if err != nil {
			// util.LogWarn(err.Error())
			return 0, 0, err
		}
		b, err := ioutil.ReadAll(s.Body)
		util.DieOnError(err)

		tr := TrackerResponse{}
		bencode.Unmarshal(b, &tr)

		return tr.Complete, tr.Incomplete, nil
	}
	return 0, 0, nil
}

func hashToBin(h string) []byte {
	b := make([]byte, hex.DecodedLen(len(h)))
	_, err := hex.Decode(b, []byte(h))
	util.DieOnError(err)
	return b
}

func random(l int) string {
	if l%4 != 0 {
		return ""
	}
	b := make([]byte, l/4)
	rand.Read(b)
	return hex.EncodeToString(b)
}
