package server

import (
	"bufio"
	"log"
	"net"
	"strconv"

	"github.com/Yatin-Malhotra/lumosdb/internal/protocol"
)

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	log.Println("New Client:", conn.RemoteAddr())

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	for {
		cmd, err := protocol.ReadCommand(reader)
		if err != nil {
			log.Println("Protocol error:", err)
			return
		}

		log.Printf("Command: %s %v\n", cmd.Name, cmd.Args)

		switch cmd.Name {
		case "PING":
			writer.WriteString("+PONG\r\n")

		case "SET":
			if len(cmd.Args) != 2 {
				writer.WriteString("-ERR invalid number of arguments provided\r\n")
				break
			}

			key := cmd.Args[0]
			value := cmd.Args[1]

			s.Store.Set(key, value, 0)
			writer.WriteString("+OK\r\n")

		case "GET":
			if len(cmd.Args) != 1 {
				writer.WriteString("-ERR invalid number of arguments provided\r\n")
				break
			}

			key := cmd.Args[0]
			value, ok := s.Store.Get(key)

			if !ok {
				writer.WriteString("$-1\r\n")
			} else {
				writer.WriteString("$" + strconv.Itoa(len(value)) + "\r\n")
				writer.WriteString(value + "\r\n")
			}

		case "DEL":
			if len(cmd.Args) != 1 {
				writer.WriteString("-ERR invalid number of arguments provided\r\n")
				break
			}

			s.Store.Delete(cmd.Args[0])
			writer.WriteString(":1\r\n")

		default:
			writer.WriteString("-ERR unknown command\r\n")
		}

		writer.Flush()
	}
}
