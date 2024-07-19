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

```css
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
            text-align: center; /* Center title below video */
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
            text-align: center; /* Center title below video */
            margin-top: 1em; /* Add space between video and title */
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
        button {
          background-color: #000033; /* Semi-transparent black background */
          color: white; /* White text for legibility */
          padding: 0.5em 1em; /* Set button padding */
          border: none; /* Remove default border */
          cursor: pointer; /* Indicate clickable button */
          transition: all 0.2s ease-in-out; /* Smooth hover effect */
        }

        button:hover {
          background-color: #FF0000; /* Darken background on hover */
        }

        /* Button positioning and spacing */
        #prevButton {
          margin-right: 0em; /* Add spacing between buttons */
          margin-left: 0em;
        }
        .search-container {
          margin-bottom: 10px;
        }

        .search-container input {
          width: 100%;
          padding: 10px;
          border: none;
          border-radius: 5px;
          background-color: #152057;
          color: white;
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
```

```javascript
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
```
