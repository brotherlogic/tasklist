package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/brotherlogic/goserver/utils"

	pb "github.com/brotherlogic/tasklist/proto"
)

func main() {
	ctx, cancel := utils.ManualContext("tasklist-cli", time.Minute*5)
	defer cancel()

	conn, err := utils.LFDialServer(ctx, "tasklist")
	if err != nil {
		log.Fatalf("Unable to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewTaskListServiceClient(conn)

	switch os.Args[1] {
	case "file":
		data, err := ioutil.ReadFile(os.Args[2])
		if err != nil {
			log.Fatalf("Load error: %v", err)
		}

		lines := strings.Split(string(data), "\n")
		req := &pb.TaskList{Name: lines[0]}
		for i := 1; i < len(lines); i++ {
			bits := strings.Split(lines[i], "|")
			req.Tasks = append(req.Tasks, &pb.Task{Title: bits[1], Job: bits[0], Index: int32(i)})
		}
		_, err = client.AddTaskList(ctx, &pb.AddTaskListRequest{Add: req})
		if err != nil {
			log.Fatalf("Bad Add: %v", err)
		}
	default:
		fmt.Printf("Unknown command: %v\n", os.Args[1])
	}

}
