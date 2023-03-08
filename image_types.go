package openai

/*
Request body
prompt
string
Required
A text description of the desired image(s). The maximum length is 1000 characters.

n
integer
Optional
Defaults to 1
The number of images to generate. Must be between 1 and 10.

size
string
Optional
Defaults to 1024x1024
The size of the generated images. Must be one of 256x256, 512x512, or 1024x1024.

response_format
string
Optional
Defaults to url
The format in which the generated images are returned. Must be one of url or b64_json.

user
string
Optional
A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse. Learn more.
除了 prompt 之外的参数都是可选的，但是如果你想要生成多张图片，那么你必须指定 n 参数，否则只会生成一张图片。参数都为指针
*/

type CreateImageRequest struct {
	Prompt         string `json:"prompt"`
	N              *int   `json:"n,omitempty"`
	Size           *int   `json:"size,omitempty"`
	ResponseFormat *int   `json:"response_format,omitempty"`
	User           *int   `json:"user,omitempty"`
}

/*
{
  "created": 1589478378,
  "data": [
    {
      "url": "https://..."
    },
    {
      "url": "https://..."
    }
  ]
}

*/

type ImageResponse struct {
	Created int `json:"created"`
	Data    []struct {
		URL string `json:"url"`
	} `json:"data"`
}

/*
Request body
image
string
Required
The image to edit. Must be a valid PNG file, less than 4MB, and square. If mask is not provided, image must have transparency, which will be used as the mask.

mask
string
Optional
An additional image whose fully transparent areas (e.g. where alpha is zero) indicate where image should be edited. Must be a valid PNG file, less than 4MB, and have the same dimensions as image.

prompt
string
Required
A text description of the desired image(s). The maximum length is 1000 characters.

n
integer
Optional
Defaults to 1
The number of images to generate. Must be between 1 and 10.

size
string
Optional
Defaults to 1024x1024
The size of the generated images. Must be one of 256x256, 512x512, or 1024x1024.

response_format
string
Optional
Defaults to url
The format in which the generated images are returned. Must be one of url or b64_json.

user
string
Optional
A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse. Learn more.
参数都为指针
*/

type EditImageRequest struct {
	Image          string  `json:"image,omitempty"`
	Mask           *string `json:"mask,omitempty"`
	Prompt         string  `json:"prompt,omitempty"`
	N              *int    `json:"n,omitempty"`
	Size           *int    `json:"size,omitempty"`
	ResponseFormat *int    `json:"response_format,omitempty"`
	User           *int    `json:"user,omitempty"`
}

/*
Request body
image
string
Required
The image to use as the basis for the variation(s). Must be a valid PNG file, less than 4MB, and square.

n
integer
Optional
Defaults to 1
The number of images to generate. Must be between 1 and 10.

size
string
Optional
Defaults to 1024x1024
The size of the generated images. Must be one of 256x256, 512x512, or 1024x1024.

response_format
string
Optional
Defaults to url
The format in which the generated images are returned. Must be one of url or b64_json.

user
string
Optional
A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse. Learn more.
*/

type VariationImageRequest struct {
	Image          string  `json:"image,omitempty"`
	N              *int    `json:"n,omitempty"`
	Size           *string `json:"size,omitempty"`
	ResponseFormat *string `json:"response_format,omitempty"`
	User           *string `json:"user,omitempty"`
}
