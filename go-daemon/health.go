type HealthCheck struct{
    Status string `json:"status"`
    Uptime float64 `json:"uptime"`
    MemoryUse uint64 `json:"memory_use"`
    NumGoroutine int `json:"num_goroutine"`
    LastError string `json:last_error,omitempty`
}

func setupHealthCheck(){
    http.HandleFunc("/health", func(w http.ResponseWriter, r * http.Request){
        stats := &HealthCheck{
            Status: "running",
            Uptime: time.Since(startTime).Seconds(),
            Memory: getMemoryUsage(),
            NumGoroutine: runtime.NumGoroutine(),
        }
        json.NewEncoder(w).Encode(stats)

    })

    go http.ListenAndServer(":8081", nil)
}