# Video-Playlist-Gallery

## Create A Responsive Video Playlist Gallery Using HTML - CSS - Javascript

This script written with php can Instantly Create a single `HTML` file including `CSS` and `JAVASCRIPT`.
For Your Video Gallery, also you can watch it Offline on your browser.

Visit my Github Page For `Source code` [Video-Playlist-Galery](https://caturmahdialfurqon.github.io/posts/Video-Playlist-Galery/)

## Here The Screenshot

![video playlist galery](/Asset/lewiscapaldi.gif)

## UPDATE AND UPGRADE
`Fri Jul 19 14:44:28 WIB 2024`

### Adding Features

- adding Prev and Next Button also Search Tittle Videos
  - just edit the `<style>` and `<script>` on both `GO` or `PHP` source code .

### CSS
```css
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
            text-transform: capitalize;
            font-family: 'Arial', sans-serif;
            font-weight: normal;
        }
        body {
            background: rgba(0, 0, 51, 1);
            color: rgba(255, 255, 255, 0.95);
        }
        .heading {
            color: #fff;
            font-size: 40px;
            text-align: center;
            padding: 20px;
            background: rgba(0, 72, 131, 0.25);
            text-shadow: 2px 2px 4px rgba(255, 255, 255, 1.25);
        }
        .container {
            display: grid;
            grid-template-columns: 2fr 1fr;
            gap: 20px;
            align-items: flex-start;
            padding: 20px 5%;
        }
        .container .main-video {
            background: #00194D;
            border-radius: 10px;
            padding: 15px;
            box-shadow: 0 4px 10px rgba(255, 255, 255, 1);
        }
        .container .main-video video {
            width: 100%;
            border-radius: 9px;
            aspect-ratio: 16 / 9;
            box-shadow: 0 2px 10px rgba(0, 127, 255, 0.5);
        }
        .container .main-video .title {
            color: rgba(255, 255, 255, 0.95);
            font-size: 24px;
            padding: 15px 0;
            text-align: center;
            font-weight: bold;
        }
        .container .video-list {
            background: #00194D;
            border-radius: 10px;
            height: 720px;
            overflow-y: auto;
            padding: 10px;
            box-shadow: 0 4px 10px rgba(0, 0, 0, 0.5);
        }
        .container .video-list::-webkit-scrollbar {
            width: 8px;
        }
        .container .video-list::-webkit-scrollbar-track {
            background: rgba(0, 127, 255, 0.35);
            border-radius: 50px;
        }
        .container .video-list::-webkit-scrollbar-thumb {
            background: rgba(229, 9, 20, 0.5);
            border-radius: 50px;
        }
        .container .video-list .vid video {
            width: 100px;
            border-radius: 5px;
            aspect-ratio: 16 / 9;
        }
        .container .video-list .title {
            color: rgba(255, 255, 255, 0.95);
            font-size: 12px;
            text-align: center;
            margin-top: 1em;
        }
        .container .video-list .vid {
            display: flex;
            align-items: center;
            gap: 15px;
            background: rgba(0, 0, 51, 0.5);
            border-radius: 5px;
            margin: 10px 0;
            padding: 15px;
            border: 1px solid rgba(255, 255, 255, 0.2);
            cursor: pointer;
            transition: background 0.3s;
            box-shadow: 0 2px 5px rgba(0, 127, 255, 1);
        }
        .container .video-list .vid:hover {
            background: rgba(0, 127, 255, 0.25);
        }
        .container .video-list .vid.active {
            background: rgba(229, 9, 20, 0.35);
        }

        button {
          background-color: rgba(0, 0, 51, 1);
          color: white;
          padding: 0.5em 1em;
          margin-top: 10px;
          border: none;
          border-radius: 5px;
          cursor: pointer;
          transition: background 0.3s;
          box-shadow: 0 1px 2px rgba(255, 255, 255, 1.25);
        }

        button:hover {
          background-color: rgba(229, 9, 20, 0.5);
        }

        .search-container {
          margin-bottom: 10px;
        }

        .search-container input {
          width: 100%;
          padding: 10px;
          border: none;
          border-radius: 5px;
          background-color: #010000;
          color: rgba(255, 255, 255, 0.95);
          box-shadow: 0 2px 5px rgba(255, 255, 255, 1.15);
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
```

### HTML

```html
</head>
<body>
    <h3 class=\"heading\">$titledisplay</h3>
    <div class=\"container\">
        <div class=\"main-video\">
            <div class=\"video\">
                <video src=\"\" controls></video>
              </div>
              <div class=\"navigation-buttons\">
                <button id=\"prevButton\">⏮ Prev </button>
                <button id=\"nextButton\">Next ⏭</button>
              </div>
              <h3 class=\"title\"></h3>
            </div>
        <div class=\"video-list\">
            <div class=\"search-container\">
            <input type=\"text\" id=\"searchInput\" placeholder=\"Search video title...\">";

foreach ($videoList as $video) {
    $html .= "
            <div class=\"vid\">
                <video>
                    <source src=\"{$video["source"]}\" type=\"video/mp4\">
                    Your browser does not support the video tag.
                </video>
                <h3 class=\"title\">{$video["title"]}</h3>
            </div>";
}

$html .= "
        </div>
    </div>
```


### Javascript

```javascript
      <script>
        let listVideo = document.querySelectorAll('.video-list .vid');
        let mainVideo = document.querySelector('.main-video video');
        let title = document.querySelector('.main-video .title');
        let currentIndex = 0;
        let searchInput = document.getElementById('searchInput');

        function playVideo(index) {
          if (index >= 0 && index < listVideo.length) {
            listVideo.forEach(vid => vid.classList.remove('active'));
            listVideo[index].classList.add('active');
            if (listVideo[index].children[0].children[0].getAttribute('src')) {
              mainVideo.src = listVideo[index].children[0].children[0].getAttribute('src');
            }
            title.innerHTML = listVideo[index].children[1].innerHTML;
            currentIndex = index;
          }
        }

        listVideo.forEach((video, index) => {
          video.onclick = () => {
            playVideo(index);
          }
        });

        document.getElementById('prevButton').addEventListener('click', () => {
          playVideo(currentIndex - 1);
        });

        document.getElementById('nextButton').addEventListener('click', () => {
          playVideo(currentIndex + 1);
        });

        searchInput.addEventListener('input', function() {
          const searchTerm = this.value.toLowerCase();
          listVideo.forEach((video, index) => {
            const videoTitle = video.querySelector('.title').textContent.toLowerCase();
            if (videoTitle.includes(searchTerm)) {
              video.style.display = 'flex';
            } else {
              video.style.display = 'none';
            }
          });
        });

        // Play the first video initially
        playVideo(0);
      </script>
```
