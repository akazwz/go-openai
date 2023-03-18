package openai

const (
	CreateImageURL    = "/v1/images/generations"
	EditImageURL      = "/v1/images/edits"
	VariationImageURL = "/v1/images/variations"
)

func (c *Client) CreateImage(request *CreateImageRequest) (*ImageResponse, error) {
	res, err := c.DoCommonPostReq(request, c.apiURL+CreateImageURL)
	if err != nil {
		return nil, err
	}
	typedRes, ok := res.(*ImageResponse)
	if !ok {
		return nil, err
	}
	return typedRes, nil
}

func (c *Client) EditImage(request *EditImageRequest) (*ImageResponse, error) {
	res, err := c.DoCommonPostReq(request, c.apiURL+EditImageURL)
	if err != nil {
		return nil, err
	}
	typedRes, ok := res.(*ImageResponse)
	if !ok {
		return nil, err
	}
	return typedRes, nil
}

func (c *Client) VariationImage(request *VariationImageRequest) (*ImageResponse, error) {
	res, err := c.DoCommonPostReq(request, c.apiURL+VariationImageURL)
	if err != nil {
		return nil, err
	}
	typedRes, ok := res.(*ImageResponse)
	if !ok {
		return nil, err
	}
	return typedRes, nil
}
