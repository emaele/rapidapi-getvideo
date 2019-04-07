package getvideo

// NewClient returns a getvideoClient capable of requests
func NewClient(key string) Client {
	var getvideoclient Client
	getvideoclient.apikey = key

	return getvideoclient
}
