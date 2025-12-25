package server

import (
	"bufio"
	"log"
	"net"

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

		writer.WriteString("+OK\r\n")
		writer.Flush()
	}
}
