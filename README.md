# Secure Chat Server 💬

A secure TCP-based chat server implementation in Go that supports multiple clients with encrypted message communication. 🔒

## Features ✨

- Multi-client support with concurrent connections 👥
- Real-time message broadcasting 📡
- Secure message encryption using AES-CFB 🔐
- Username-based client identification 👤
- Clean disconnection handling 🔌
- Thread-safe client management 🛡️

## Technical Stack 🛠️

- Go 1.24.0 🚀
- Standard Go libraries:
  - `net` for TCP networking 🌐
  - `crypto/aes` for message encryption 🔑
  - `sync` for thread-safe operations 🔄

## Project Structure 📁

```
chat-app/
├── main.go         # Server initialization and main entry point
├── client.go       # Client handling and message broadcasting
├── encryption.go   # Message encryption/decryption utilities
├── server_test.go  # Server testing
├── go.mod         # Go module definition
└── Makefile       # Build and run automation
```

## Getting Started 🚀

### Prerequisites 📋

- Go 1.24.0 or higher

### Installation 💻

1. Clone the repository:
```bash
git clone https://github.com/Amul-Thantharate/chat-app.git
cd chat-app
```

2. Build the server:
```bash
make build
```

### Running the Server ▶️

Start the server using:
```bash
make run
```
Or directly run the binary:
```bash
./chat-server
```

The server will start listening on port 8080.

## Architecture 🏗️

### Server Components 🔧

1. **Main Server (main.go)** 🎯
   - Initializes TCP listener
   - Handles client acceptance

2. **Client Handler (client.go)** 👥
   - Manages client connections
   - Handles message broadcasting
   - Maintains thread-safe client list

3. **Encryption Module (encryption.go)** 🔒
   - Provides message encryption
   - Implements AES-CFB encryption/decryption

## Security 🛡️

- All messages are encrypted using AES-CFB encryption 🔐
- Secure key management for message encryption 🔑
- Protection against concurrent access issues using mutex locks 🔒

## Testing 🧪

Run the tests using:
```bash
make test
```

## Building 🏗️

The project includes a Makefile with the following commands:
- `make build`: Build the server 🔨
- `make run`: Run the server ▶️
- `make test`: Run tests ✅
- `make clean`: Clean build artifacts 🧹

## Usage Examples 📝

### Connecting to the Server using Netcat 🖥️

1. Start the server:
```bash
./chat-server
```

2. Connect as a client using netcat (open multiple terminals for multiple clients):
```bash
nc localhost 8080
```

3. When connected, you'll be prompted to enter your username:
```
Enter your username: Alice
```

4. Start chatting! Here's an example conversation:

```
[Server] Welcome Alice!
[Server] Type your message and press Enter to send
Hello everyone!                           # Alice types
[Alice] Hello everyone!                   # All clients see this
[Bob] Hi Alice!                          # Bob responds
[Charlie] Welcome to the chat Alice!     # Charlie responds
```

### Example Chat Session 💭

Here's a typical chat session with multiple users:

```
[Server] Welcome to the chat server!
[Server] Enter your username: Alice
[Server] Welcome Alice!
[Server] Bob has joined the chat
Hello everyone!
[Alice] Hello everyone!
[Bob] Hi Alice! How are you?
[Server] Charlie has joined the chat
[Charlie] Hey folks!
/users
[Server] Connected users: Alice, Bob, Charlie
/msg Bob Hey, can you help me with something?
[Alice -> Bob] Hey, can you help me with something?
[Bob -> Alice] Sure, what do you need?
/quit
[Server] Alice has left the chat

## Contributing 🤝

We welcome contributions to improve the chat server! Here's how you can contribute:

### Development Setup 🛠️

1. Fork the repository
2. Clone your fork:
```bash
git clone https://github.com/YOUR_USERNAME/chat-app.git
cd chat-app
```
3. Create a new branch:
```bash
git checkout -b feature/your-feature-name
```

### Coding Guidelines 📚

- Follow Go best practices and coding conventions ✨
- Use meaningful variable and function names 📝
- Add comments for complex logic 💡
- Write unit tests for new features ✅
- Keep functions small and focused 🎯
- Handle errors appropriately ⚠️

### Pull Request Process 📤

1. Update the documentation to reflect any changes
2. Add or update tests as needed
3. Ensure all tests pass locally:
```bash
make test
```
4. Push to your fork and submit a pull request
5. Wait for review and address any feedback

### Reporting Issues 🐛

When reporting issues, please include:

- A clear description of the problem 📝
- Steps to reproduce the issue 🔄
- Expected vs actual behavior ⚖️
- Version information ℹ️
- Any relevant error messages or logs 📋

## Author ✍️

**Amul Thantharate**

- GitHub: [@Amul-Thantharate](https://github.com/Amul-Thantharate) 🐙
- LinkedIn: [Amul Thantharate](https://linkedin.com/in/amul-thantharate) 💼

## Acknowledgments 🙏

- Thanks to all contributors who have helped improve this project ⭐
- Special thanks to the Go community for their excellent networking libraries 🌟
- Inspired by various open-source chat implementations 💫

## Support 💪

If you need help or have questions:

1. Check the [Issues](https://github.com/Amul-Thantharate/chat-app/issues) page 📋
2. Create a new issue if your problem isn't already listed ➕
3. Reach out to the author directly through GitHub 📧
