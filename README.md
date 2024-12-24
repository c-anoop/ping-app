# ping-app

This project is a Go console application that sends ICMP packets to a specified IP address using raw sockets.

## Project Structure

```
ping-app
├── src
│   ├── main.go
│   └── ping
│       └── ping.go
├── go.mod
└── README.md
```

## Installation

1. Clone the repository:
   ```
   git clone <repository-url>
   cd ping-app
   ```

2. Install the necessary dependencies:
   ```
   go mod tidy
   ```

## Usage

To run the application, use the following command:

```
sudo go run src/main.go <IP_ADDRESS>
```

Replace `<IP_ADDRESS>` with the target IP address you want to ping.

## Example

```
go run src/main.go 192.168.1.1
```

This command will send 5 ICMP packets of 64 bytes each to the IP address `192.168.1.1` and print the received packet messages.