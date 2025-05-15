package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: mrmint {init|start|stop}")
		return
	}

	switch os.Args[1] {
	case "init":
		fmt.Println("✅ Initializing mrmint...")
	case "start":
		fmt.Println("🚀 Starting mrmint...")
	case "stop":
		fmt.Println("🛑 Stopping mrmint...")
	default:
		fmt.Println("❌ Unknown command:", os.Args[1])
	}
}
