// socket-server project main.go
package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
)

func main() {
	fmt.Println("Server Running...")
	server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer server.Close()
	fmt.Println("Listening on " + SERVER_HOST + ":" + SERVER_PORT)
	fmt.Println("Waiting for client...")
	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		fmt.Println("client connected")
		go processClient(connection)
	}
}
func processClient(connection net.Conn) {
	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Println("Received: ", string(buffer[:mLen]))
	_, err = connection.Write([]byte("Hola  cliente."))
	//var comando string
	/*fmt.Println("escriba el comando")
	if _, err := fmt.Scanln(&comando); err != nil {
		fmt.Println("Error:", err)
	} */
	/* reader := bufio.NewReader(os.Stdin)
	fmt.Println("escribe el comando a ejecutar")
	comand, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("el comando es ", comand)

	if strings.TrimRight(comand, "\r\n") == "exit" {
		fmt.Println("entre a exit")
		_, err = connection.Write([]byte("exit"))
		//break
	}
	if strings.TrimRight(comand, "\r\n") == "dir" {
		fmt.Println("entre a dir")
		buffer := make([]byte, 1024)
		_, err = connection.Write([]byte("dir"))
		mLen, err := connection.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}
		fmt.Println(string(buffer[:mLen]))

	} */

	for {
		fmt.Println("escribe el comando a ejecutar")
		//reader := bufio.NewReader(os.Stdin)
		reader := bufio.NewScanner(os.Stdin)
		reader.Scan()
		if err := reader.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}

		/* comand, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error:", err)
		} */
		fmt.Println("el comando es ", reader.Text())
		//strings.TrimRight(comand, "\r\n")
		if reader.Text() == "exit" {
			fmt.Println("entre a exit")
			_, err = connection.Write([]byte("exit"))
			mLen, err := connection.Read(buffer)
			if err != nil {
				fmt.Println("Error reading:", err.Error())
			}
			fmt.Println("Received: ", string(buffer[:mLen]))
			break
		}
		if reader.Text() == "dir" {
			fmt.Println("entre a dir")
			buffer := make([]byte, 1024)
			_, err = connection.Write([]byte("dir"))
			mLen, err := connection.Read(buffer)
			if err != nil {
				fmt.Println("Error reading:", err.Error())
			}
			fmt.Println(string(buffer[:mLen]))

		}

		if reader.Text()[:2] == "cd" {
			_, err = connection.Write([]byte(reader.Text()))
			mLen, err := connection.Read(buffer)
			if err != nil {
				fmt.Println("Error reading:", err.Error())
			}
			fmt.Println(string(buffer[:mLen]))

		}
		if reader.Text() == "screenshot" {
			_, err = connection.Write([]byte(reader.Text()))
			if err != nil {
				fmt.Println("Error reading:", err.Error())
			}
			fo, err := os.Create("./sss2.png")
			if err != nil {
				panic(err)
			}
			defer fo.Close()
			defer connection.Close()

			_, err = io.Copy(fo, connection)
			if err != nil {
				fmt.Print("error recibir imagen putoo", err)
			}
			fmt.Println("Serrvido Termino")

			//mLen, err := connection.Read(buffer)

			//fmt.Println(string(buffer[:mLen]))
		}

	}
	connection.Close()
}
