package server

import (
	"bufio"
	"io"
	"log"
	"net"
)

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	log.Println("New Client:", conn.RemoteAddr())

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			if err != io.EOF {
				log.Println("Read error:", err)
			}
			return
		}

		writer.WriteString(line)
		writer.Flush()
	}
}
