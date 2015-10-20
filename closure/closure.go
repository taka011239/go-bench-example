package closure

type Request struct {
	Url string
}

func Closure1(reqs []*Request) {
	for _, req := range reqs {
		go func(req *Request) {
			req.Url = "http://example.org"
		}(req)
	}
}

func Closure2(reqs []*Request) {
	for _, req := range reqs {
		req := req
		go func() {
			req.Url = "http://example.org"
		}()
	}
}
