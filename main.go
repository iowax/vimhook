package main

import (
	"fmt"

	"./vimeo"
)

func main() {
	fmt.Println("Search for videos related to:")
	//reader := bufio.NewReader(os.Stdin)
	//searchtext, _ := reader.ReadString('\n')

	//Client instance for connecting to the vimeo apis
	vimeoClient, err := vimeo.NewVimeoClient()

	if err != nil {
		fmt.Println(err)
	}
	//vimeoClient.SearchAllVideos(searchtext)
	vimeoClient.DownloadVideo("20732587", "/tmp/video.mp4")
}
