package tools

import (
	"bytes"
	"fmt"
	"os/exec"
)

func GetCommandOutput(command string, arguments ...string) (string, error) {
	status := exec.Command(command, arguments...)
	res, err := status.CombinedOutput()
	
	if err != nil {
		fmt.Printf("La ejecucion del comando ha fallado: %v \n", err)
		return "", err
	}

	return string(res), nil
}


func GetAndFilterCommandOutput(filters string,command string, arguments ...string) (string, error) {
	cmd1 := exec.Command(command, arguments...)
	cmd2 := exec.Command("grep", filters)

	var out bytes.Buffer

	r, err := cmd1.StdoutPipe()

	if err != nil {
		fmt.Println("El comando pasado no sire chcch")
		return "", err
	}

	cmd2.Stdin  = r 
	cmd2.Stdout = &out

	cmd1.Start()
	cmd2.Start()

	cmd1.Wait()
	cmd2.Wait()


	return out.String(), nil
}