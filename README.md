# RapidAPI GetVideo

GetVideo is a simple wrapper for RapidAPI GetVideo apis written in go.

## Usage
```go
package main

import (
	"fmt"

	getvideo "gitlab.com/emaele/rapidapi-getvideo"
)

func main() {
	cl := getvideo.NewClient("apikey")

	video, err := cl.NewRequest("")
	if err != nil {
		// Error handling
	}

	fmt.Printf("Title: %s", video.Title)
}
```