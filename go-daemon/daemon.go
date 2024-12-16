func startDaemon() error{
	pwd, err := os.Getwd()
	if err != nil{
		return err
	}

	args := []string{"-deamon"}
	args = append(args, os.Args[1:]...)

	cmd := exec.Command(os.Args[0], args...)
	cmd.Dir = pwd
	cmd.Env = os.Environ()
	
	return cmd.Start()
}

func gracefulRestart() error {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil{
		return err
	}

	cmd := exec.Command(os.Args[0], os.Args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.ExtractFiles = []*os.File{listener.(*net.TCPListener).File()}

	errr := cmd.Start()
	if err != nil{
		return err
	}

	waitForConnections()
	return nil

}

type MemoryStats struct{
	Alloc uint64
	TotalAlloc uint64
	Sys uint64
	NumGC uint32
}

func monitorMemory(threshold uint64){
	var stats runtime.MemoryStats
	ticker := time.NewTicker(time.Minute)
	
	for range ticker.C {
		runtime.ReadMemStats(&stats)
		if stats.Alloc > threshold{
			log.Printf("Memory usage exceeds threshold: %d MB", stats.Alloc/1024/1024)
			runtime.GC()
		}
	}
}

