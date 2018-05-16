package main

import (
  "fmt"
  "log"
  "os"

  proxmox "github.com/FireDrunk/go-proxmox"
  argparse "github.com/akamensky/argparse"
)

func main() {
  fmt.Println("Welcome to the Go Proxmox CLI")

	parser := argparse.NewParser("print", "Prints provided string to stdout")

  c := parser.String("c","create-pool",&argparse.Options {Required: false, Help: "Create pool",})
	n := parser.Flag("n","nodes",&argparse.Options {Required: false, Help: "Show nodes",})
  t := parser.Flag("t","tasks",&argparse.Options {Required: false, Help: "Show tasks",})
  p := parser.Flag("p","pools",&argparse.Options {Required: false, Help: "Show pools",})


	err := parser.Parse(os.Args)

	if err != nil {
		fmt.Print(parser.Usage(err))
	}

  proxmoxClient, error := proxmox.NewProxMox("10.0.2.15:8006", "root", "password")
  if error != nil {
    log.Fatal(error)
  }
  fmt.Printf("Connected to: %s\n", proxmoxClient.Hostname)

  // nodes
  if *n {
    fmt.Println("Nodes:")
    nodes, error := proxmoxClient.Nodes()

    if error != nil {
      fmt.Println(error)
    }

    for _, node := range nodes {
      fmt.Printf("  - %s\n", node.Node)
    }
  // tasks
  } else if *t {
    fmt.Println("Tasks:")
    tasks, error := proxmoxClient.Tasks()

    if error != nil {
      fmt.Println(error)
    }

    for _, task := range tasks {
      fmt.Printf("  - %s (%s - %s)\n", task.UPid, task.StartTime, task.EndTime)
    }
  // pools
  } else if *p {
    fmt.Println("Pools:")
    pools, error := proxmoxClient.Pools()

    if error != nil {
      fmt.Println(error)
    }

    for _, pool := range pools {
      fmt.Printf("  - %s\n", pool.Poolid)
    }
  // create pool
  } else if len(*c) > 0 {
    result, err := proxmoxClient.NewPool(*c, "comment")
    if err != nil {
      fmt.Println("[cli] Error creating pool")
      fmt.Println(err)
    } else {
      fmt.Printf("Pool created: %s", *c)
      fmt.Println(result)
    }
  }
}
