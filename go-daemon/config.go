type Config struct {
    LogLevel string `json:"log_level"`
    MaxMemory int64 `json:"max_memory"`
    WorkerCount int `json:"worker_count"`
    mu sync.RWMutex
}

func (c *Config) Reload() error{
    c.mu.Lock()
    defer c.mu.Unlock()

    data, err := os.ReadFile("/pwd")

    if err != nil{
        return err
    }

    return json.Unmarshall(data, c)

}

