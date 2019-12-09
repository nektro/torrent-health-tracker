package main

var client = &http.Client{
	Timeout: time.Second * 5,
}

func main() {

	flagTR := pflag.StringArrayP("tracker", "t", []string{}, "")

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
}
