package service

func SortName(urls []string) []string {
	_len := len(urls)
	for i := 0; i < _len; i++ {
		for j := i; j < _len; j++ {
			if urls[i] > urls[j] {
				tmp_url := urls[i]
				urls[i] = urls[j]
				urls[j] = tmp_url
			}
		}
	}

	return urls
}
