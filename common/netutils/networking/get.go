package networking

func RequestGet(url string) (*ExternalInfo, error) {
	c := NewReqClient()
	resp, err := c.Client.Get(url)
	if err != nil {
		return nil, err
	}
	info, err := parseBody(resp)
	if err != nil {
		return nil, err
	}
	return info, nil
}
