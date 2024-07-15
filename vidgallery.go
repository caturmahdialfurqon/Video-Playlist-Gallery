package main
import (
	"bufio"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
)

type Video struct {
	Title  string
	Source string
}

func main() {
	clearScreen()

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Directory Path Here => ")
	directory, _ := reader.ReadString('\n')
	directory = strings.TrimSpace(directory)

	fmt.Print("Enter Title Bar for Your HTML file => ")
	titleDisplay, _ := reader.ReadString('\n')
	titleDisplay = strings.TrimSpace(titleDisplay)

	if _, err := os.Stat(directory); os.IsNotExist(err) {
		fmt.Printf("Error: Directory '%s' does not exist.\n", directory)
		return
	}

	files, err := ioutil.ReadDir(directory)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return
	}

	var videoList []Video

	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(strings.ToLower(file.Name()), ".mp4") {
			continue
		}

		title := strings.TrimSuffix(file.Name(), ".mp4")
		video := Video{
			Title:  title,
			Source: filepath.Join(directory, file.Name()),
		}

		found := false
		for _, item := range videoList {
			if strings.EqualFold(item.Title, title) {
				found = true
				break
			}
		}

		if !found {
			videoList = append(videoList, video)
		}
	}

	sort.Slice(videoList, func(i, j int) bool {
		return strings.ToLower(videoList[i].Title) < strings.ToLower(videoList[j].Title)
	})

	tmpl := template.Must(template.New("videoPlayer").Parse(htmlTemplate))

	data := struct {
		TitleDisplay string
		VideoList    []Video
	}{
		TitleDisplay: titleDisplay,
		VideoList:    videoList,
	}

	fmt.Print("Enter your file name *without .html* => ")
	fileName, _ := reader.ReadString('\n')
	fileName = strings.TrimSpace(fileName) + ".html"
	filePath := filepath.Join(directory, fileName)

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	err = tmpl.Execute(file, data)
	if err != nil {
		fmt.Printf("Error executing template: %v\n", err)
		return
	}

	fmt.Printf("Generated HTML file saved to: %s\n", filePath)
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

const htmlTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.TitleDisplay}}</title>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
            text-transform: capitalize;
            font-family: sans-serif;
            font-weight: normal;
        }
        body {
            background: #000033;
            color: rgba(255, 255, 255, 0.95);
        }
        .heading {
            color: #eee;
            font-size: 40px;
            text-align: center;
            padding: 10px;
        }
        .container {
            display: grid;
            grid-template-columns: 2fr 1fr;
            gap: 15px;
            align-items: flex-start;
            padding: 5px 5%;
        }
        .container .main-video {
            background: #152057;
            border-radius: 5px;
            padding: 10px;
        }
        .container .main-video video {
            width: 100%;
            border-radius: 5px;
            aspect-ratio: 16 / 9;
        }
        .container .main-video .title {
            color: rgba(255, 255, 255, 0.95);
            font-size: 23px;
            padding-top: 15px;
            padding-bottom: 15px;
        }
        .container .video-list {
            background: #0000;
            border-radius: 5px;
            height: 720px;
            overflow-y: auto;
        }
        .container .video-list::-webkit-scrollbar {
            width: 7px;
        }
        .container .video-list::-webkit-scrollbar-track {
            background: #ccc;
            border-radius: 50px;
        }
        .container .video-list::-webkit-scrollbar-thumb {
            background: #666;
            border-radius: 50px;
        }
        .container .video-list .vid video {
            width: 100px;
            border-radius: 5px;
            aspect-ratio: 16 / 9;
        }
        .container .video-list .title {
            color: rgba(255, 255, 255, 0.95);
            font-size: 10px;
        }
        .container .video-list .vid {
            display: flex;
            align-items: center;
            gap: 15px;
            background: #152057;
            border-radius: 5px;
            margin: 10px;
            padding: 10px;
            border: 1px solid rgba(0,0,0,.1);
            cursor: pointer;
        }
        .container .video-list .vid:hover {
            background: #1C63B2;
        }
        .container .video-list .vid.active {
            background: #FF0000;
        }
        @media (max-width: 991px) {
            .container {
                grid-template-columns: 2fr 1fr;
                padding: 10px;
            }
        }
        @media (max-width: 768px) {
            .container {
                grid-template-columns: 1fr;
                padding: 10px;
            }
        }
    </style>
</head>
<body>
    <h3 class="heading">{{.TitleDisplay}}</h3>
    <div class="container">
        <div class="main-video">
            <div class="video">
                <video src="" controls></video>
                <h3 class="title"></h3> 
            </div>
        </div>
        <div class="video-list">
            {{range .VideoList}}
            <div class="vid">
                <video>
                    <source src="{{.Source}}" type="video/mp4">
                    Your browser does not support the video tag.
                </video>
                <h3 class="title">{{.Title}}</h3>
            </div>
            {{end}}
        </div>
    </div>
    <script>
        let listVideo = document.querySelectorAll('.video-list .vid');
        let mainVideo = document.querySelector('.main-video video');
        let title = document.querySelector('.main-video .title');

        listVideo.forEach(video => {
            video.onclick = () => {
                listVideo.forEach(vid => vid.classList.remove('active'));
                video.classList.add('active');
                if (video.children[0].children[0].getAttribute('src')) {
                    mainVideo.src = video.children[0].children[0].getAttribute('src');
                }
                title.innerHTML = video.children[1].innerHTML;
            }
        });
    </script>
</body>
</html>`

