# go-teleparty

go-teleparty is a Go application that allows two or more users to watch a video together while chatting in real-time. The application uses WebSockets for real-time communication and synchronization of video playback.

## Features

Watch a video together with friends
Chat in real-time during the video playback
Synchronize video playback (play, pause, rewind, and forward)
Secure communication over HTTPS

## Prerequisites

Go (version 1.20 or later)
A web browser (Google Chrome recommended)
ngrok (for exposing the local server to the internet)

## Getting Started

## Clone the repository:

```git clone https://github.com/vinay1668/go-teleparty.git```



## Navigate to the project directory:

```cd go-teleparty```

## Generate a self-signed SSL certificate:

```openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -nodes```

This will generate cert.pem and key.pem files in the project directory.

Place your video file (e.g., fire-drill.mp4) in the static directory.

## Start the Go server:

```go run main.go```

The server will start on https://localhost:8080.

In a separate terminal, start ngrok to expose the local server to the internet:

```ngrok http https://localhost:8080```

ngrok will provide a public URL that you can share with your friends.

Open Google Chrome and navigate to the ngrok URL. Since the server uses a self-signed certificate, you may need to bypass the security warning.

```open -na "Google Chrome" --args --ignore-certificate-errors```

This command will open Google Chrome and ignore the certificate errors.

Share the ngrok URL with your friends, and they can join the video session by opening the URL in their browsers.

### Usage
Click the "Play" button to start the video playback.
Use the "Rewind" and "Forward" buttons to navigate through the video.
Type your messages in the chat input box and press "Enter" or click the "Send" button to send them.
All participants will see the synchronized video playback and chat messages in real-time.
Contributing
Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License.



