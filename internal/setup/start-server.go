package setup

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type client struct {
	conn *websocket.Conn
	logs chan string
}

func ExecuteCommand(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Error while upgrading to Websocket")
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	clnt := &client{
		conn: conn,
		logs: make(chan string),
	}

	go clnt.readMessage()
	go clnt.writeMessage()
}

func RunCommand(command string, clnt *client) {
	parts := strings.Split(command, " ")
	workingDirPath := "/base_nextjs_page_router_js"
	cmd := exec.Command(parts[0], parts[1:]...)
	cmd.Dir = workingDirPath

	// capture stdout
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Println("Error while reading stdout " + err.Error())
		return
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Println("Error while reading stderr " + err.Error())
		return
	}

	// start the process
	if err := cmd.Start(); err != nil {
		log.Println("Error running given command " + err.Error())
		return
	}

	go StreamLogs(stdout, "STDOUT", clnt)
	go StreamLogs(stderr, "STDERR", clnt)

	// wait for the process to finish
	if err := cmd.Wait(); err != nil {
		log.Println("Process exited with error " + err.Error())
	}

	log.Println("Closing logs channel")
	close(clnt.logs)
}

// read message
func (c *client) readMessage() {
	defer func() {
		log.Println("Closing connection....")
		c.conn.Close()
		close(c.logs)
	}()
	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			log.Println("Error while reading message " + err.Error())
			break
		}
		log.Println("Message = " + string(msg))
		RunCommand(string(msg), c)
	}
}

// write message
func (c *client) writeMessage() {
	defer func() {
		log.Println("Closing connection.......")
		c.conn.Close()
	}()
	for msg := range c.logs {
		err := c.conn.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			log.Println("Error while writing message " + err.Error() + " closing connection...")
			break
		}
	}
}

// streamlogs read from the pipe and print logs
func StreamLogs(pipe io.ReadCloser, label string, clnt *client) {
	scanner := bufio.NewScanner(pipe)

	for scanner.Scan() {
		clnt.logs <- scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Println("Error reading logs " + err.Error())
		return
	}
}
