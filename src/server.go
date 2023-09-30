package main

import (
	"bufio"
	// "time"

	// "encoding/hex"
	// "encoding/hex"
	"encoding/json"
	"fmt"
	"io"

	// "io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"

	"golang.org/x/net/websocket"
)

var lineNum int = 0

func getRoot(w http.ResponseWriter, r *http.Request) {
	// fmt.Printf("got / request\n")
	http.ServeFile(w, r, "./static/index.html")
	// http.FileServer(http.Dir("../static/index1.html"))
}
func getCSS(w http.ResponseWriter, r *http.Request) {
	// fmt.Printf("got / request\n")
	http.ServeFile(w, r, "./static/style.css")
	// http.FileServer(http.Dir("../static/index1.html"))
}

func getHello(w http.ResponseWriter, _ *http.Request) {
	// fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

func listTask(w http.ResponseWriter, _ *http.Request) {
	tastList := getTODO()
	io.WriteString(w, tastList)
}

func removeTask(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	asciiString := string(body)
	decodedValue, err := url.QueryUnescape(asciiString)
	t := strings.Replace(decodedValue, "task=", "", -1)
	// fmt.Println(t)
	deleteLineFromFile("database.txt", t)
}

func deletea(task string) {
	f, err := os.Open("database.txt")

	// fmt.Println(task)

	if err != nil {
		// return 0, err
		fmt.Println("Error :", err)
	}
	defer f.Close()

	tmpFile, err := os.Open("tempfile")
	if err != nil {
		fmt.Println("Error :", err)
	}
	defer tmpFile.Close()

	// Splits on newlines by default.
	scanner := bufio.NewScanner(f)

	lineNo := 1

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(scanner.Text())
		if strings.Contains(line, task) {
			fmt.Println("found on line", lineNo, line)
			continue
		}
		lineNo++
	}

	f.Close()

	// Remove the original file
	err = os.Remove("tempfile")
	if err != nil {
		fmt.Println("Error :", err)

	}

	// Rename the temporary file to the original file name
	err = os.Rename("tempfile", "database.txt")
	if err != nil {
		fmt.Println("Error :", err)

	}

	// readFile, err := os.Open("database.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fileScanner := bufio.NewScanner(readFile)

	// fileScanner.Split(bufio.ScanLines)
	// lineNum1 := 0
	// for fileScanner.Scan() {
	// 	task += fileScanner.Text()
	// }
	// readFile.Close()

}

func deleteLineFromFile(filePath string, lineToDelete string) error {
	// Open the file for reading and writing
	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a temporary file to store the modified content
	tempFilePath := filePath + ".temp"
	tempFile, err := os.Create(tempFilePath)
	if err != nil {
		return err
	}
	defer tempFile.Close()

	// Create a scanner to read the original file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Check if the line matches the line to be deleted
		if strings.TrimSpace(line) == lineToDelete {
			continue // Skip this line
		}

		// Write the line to the temporary file
		fmt.Fprintln(tempFile, line)
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		return err
	}

	// Close the original file
	file.Close()

	// Remove the original file
	if err := os.Remove(filePath); err != nil {
		return err
	}

	// Rename the temporary file to the original file name
	if err := os.Rename(tempFilePath, filePath); err != nil {
		return err
	}

	return nil
}

func gets(w http.ResponseWriter, r *http.Request) {
	// fmt.Printf("got /save request\n")
	r.ParseForm()
	storeTODO(r.FormValue("task"))
	fmt.Println(r.FormValue("task"))

	lineNum++
	tastList := `<div hx-boost="true" class="new-element"> <a href="task"` + strconv.Itoa(
		lineNum,
	) + `\>` + r.FormValue(
		"task",
	) + `</a></div>`
	io.WriteString(w, tastList)
}

func storeTODO(task string) {
	readFile, err := os.OpenFile("database.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	_, err = fmt.Fprintln(readFile, task)
	if err != nil {
		fmt.Println(err)
	}
	readFile.Close()
}

func getTODO() string {
	readFile, err := os.Open("database.txt")
	task := ""
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	lineNum1 := 0
	for fileScanner.Scan() {
		lineNum1++
		lineNum = lineNum1
		// fmt.Println(fileScanner.Text())
		task += "<div hx-boost=\"true\" class=\"new-element\"> <a href=\"/task/" + strconv.Itoa(
			lineNum1,
		) + "\">" + fileScanner.Text() + "</a></div>"
	}
	// fmt.Println("New string 2: ", task)
	readFile.Close()
	return task
}

func getTaskById(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	// fmt.Fprintf("Dynamic Value: %s", path)
	segments := strings.Split(path, "/")

	var result []string
	for _, s := range segments {
		if s != "" {
			result = append(result, s)
		}
	}
	// fmt.Print(result[1])

	// fmt.Print(len(result))
	// fmt.Fprintf(w, "Dynamic Value: %s \n %s \n %d", path, segments,len(segments))
	if len(result) == 2 && len(result) > 0 {
		dynamicValue := result[1] // Assuming the dynamic value is the third segment
		// Use the dynamicValue as needed
		// fmt.Fprintf(w, "Dynamic Value: %s", dynamicValue)
		file, err := os.Open("database.txt")
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		currentLine := 0

		for scanner.Scan() {
			currentLine++
			lineno, _ := strconv.ParseInt(dynamicValue, 10, 64)
			if currentLine == int(lineno) {
				io.WriteString(w, scanner.Text())
			}
		}

		if err := scanner.Err(); err != nil {
		}

	} else {
		// Handle cases where there are not enough segments or the dynamic value is missing
		http.NotFound(w, r)
	}
}
func getTaskPage(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	fmt.Println(path)
	segments := strings.Split(path, "/")

	var result []string
	for _, s := range segments {
		if s != "" {
			result = append(result, s)
		}
	}
	// if len(result) <= 2 && len(result) > 0 {
	fmt.Println(len(result))
	fmt.Println(result)

	if len(result) > 0 {
		// fmt.Print(result[1])
		fmt.Printf(result[1])
		lineNum, _ = strconv.Atoi(result[1])

		fmt.Printf("new value of line %d\n", lineNum)

		http.ServeFile(w, r, "./static/task.html")

	}

	// getTaskById(w, r)
}

func setid(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<div hx-get=\"/id/"+strconv.Itoa(lineNum)+"\"  hx-trigger=\"load\"></div>")
}

func CORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().
			Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		if r.Method == "OPTIONS" {
			http.Error(w, "No Content", http.StatusNoContent)
			return
		}
		next(w, r)
	}
}

type Headers struct {
	HXRequest     string `json:"HX-Request"`
	HXTrigger     string `json:"HX-Trigger"`
	HXTriggerName string `json:"HX-Trigger-Name"`
	HXTarget      string `json:"HX-Target"`
	HXCurrentURL  string `json:"HX-Current-URL"`
}

type Data struct {
	Task    string  `json:"task"`
	Headers Headers `json:"HEADERS"`
}

type Server struct {
	conns map[*websocket.Conn]bool
	mutex sync.Mutex
}

func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
		mutex: sync.Mutex{},
	}
}

func (s *Server) websocketHandler(ws *websocket.Conn) {
	defer ws.Close() // Ensure the WebSocket connection is closed when the function exits
	fmt.Println("WebSocket connection established.")
	s.conns[ws] = true
	s.readLoop(ws)

	// for {
	// 	var message string
	// 	err := websocket.Message.Receive(ws, &message)
	// 	if err != nil {
	// 		fmt.Println("Error receiving message:", err)
	// 		break
	// 	}
	// 	jsonBytes := []byte(message)
	// 	var data Data
	// 	if err := json.Unmarshal(jsonBytes, &data); err != nil {
	// 		fmt.Println("Error unmarshaling JSON:", err)
	// 		return
	// 	}

	// 	fmt.Println("Received message:", data.Task)
	// 	storeTODO(data.Task)
	// 	// Echo back the received message
	// 	// broadcast(data.Task)
	// 	tastList := "<div hx-boost=\"true\"> <a href=\"/task\">" + data.Task + " </a></div>"
	// 	err = websocket.Message.Send(ws, tastList)
	// 	if err != nil {
	// 		fmt.Println("Error sending message:", err)
	// 		break
	// 	}
	// }
}

func (s *Server) readLoop(ws *websocket.Conn) {
	defer ws.Close()
	fmt.Println("WebSocket connection established.")
	s.conns[ws] = true

	for {
		// n, err := ws.Read()

		var message string
		err := websocket.Message.Receive(ws, &message)
		if err != nil {
			fmt.Println("Error receiving message:", err)
			delete(s.conns, ws)
			// return
			break
		}
		jsonBytes := []byte(message)
		var data Data
		if err := json.Unmarshal(jsonBytes, &data); err != nil {
			fmt.Println("Error unmarshaling JSON:", err)
			return
		}

		fmt.Println("Received message:", data.Task)
		storeTODO(data.Task)
		// Echo back the received message
		// broadcast(data.Task)
		// tastList := "<div hx-boost=\"true\"> <a href=\"/task\">" + data.Task + " </a></div>"
		lineNum++
		tastList := "<div hx-swap-oob=\"beforeend:#todo-list\"><div hx-boost=\"true\"> <a href=\"/task/" + strconv.Itoa(
			lineNum,
		) + "\">" + data.Task + "</a></div></div>"
		s.broadcast(tastList)
		// err = websocket.Message.Send(ws, data.Task)
		// if err != nil {
		// 	fmt.Println("Error sending message:", err)
		// 	break
		// }
	}
}

func (s *Server) broadcast(text string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for ws := range s.conns {
		go func(ws *websocket.Conn) {
			if _, err := ws.Write([]byte(text)); err != nil {
				// fmt.Println("err a", err)
				fmt.Println("Error sending message:", err)
				// Handle the error or remove the closed connection from the map
				ws.Close()
				delete(s.conns, ws)
			}
		}(ws)
	}
}

// func handleWebSocket(w http.ResponseWriter, r *http.Request) {
//     // Upgrade the HTTP connection to a WebSocket connection
//     conn, err := upgrader.Upgrade(w, r, nil)
//     if err != nil {
//         fmt.Println(err)
//         return
//     }
//     defer conn.Close()

//     fmt.Println("WebSocket connection established.")

//     // Handle WebSocket communication
//     for {
//         messageType, p, err := conn.ReadMessage()
//         if err != nil {
//             if err == io.EOF {
//                 fmt.Println("WebSocket connection closed by client.")
//             } else {
//                 fmt.Println("Error reading message:", err)
//             }
//             return
//         }

//         // Handle different message types (text, binary, ping, pong, etc.) if needed.
//         // For now, we'll simply echo back received messages.

//         if err := conn.WriteMessage(messageType, p); err != nil {
//             fmt.Println("Error writing message:", err)
//             return
//         }
//     }
// }

func ani(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `<div id="list" class="fadein">ssss</div>`)
}

func main() {
	server := NewServer()
	// conn()

	http.HandleFunc("/", CORS(getRoot))
	http.HandleFunc("/style.css", CORS(getCSS))
	http.HandleFunc("/taskList", CORS(listTask))
	http.HandleFunc("/hello", CORS(getHello))
	http.HandleFunc("/save", CORS(gets))
	http.HandleFunc("/delete", (removeTask))
	http.HandleFunc("/task/", CORS(getTaskPage))
	http.HandleFunc("/id/", CORS(getTaskById))
	http.HandleFunc("/ids", CORS(setid))
	http.HandleFunc("/ani", CORS(ani))

	http.Handle("/ws", websocket.Handler(server.websocketHandler))
	http.ListenAndServe(":3333", nil)
}
