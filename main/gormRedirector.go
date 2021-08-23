package main

type Redirection struct {
	Path string
	URL string
}

func ListToMapRedir (list []Redirection) map[string]string {
	if list == nil || len(list) == 0 {
		return nil
	}
	redirMap := make(map[string]string)
	for _, redir := range list {
		redirMap[redir.Path] = redir.URL
	}
	return redirMap
}