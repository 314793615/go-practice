type AsyncLogger struct{
    file *os.File
    maxSize int64
    mu sync.Mutex
    logChan chan []byte
    filename string
}


func NewAsyncLogger(filename string, maxsize int64) *AsyncLogger{
    logger := &AsyncLogger{
        filename: filename,
        maxSize: maxSize,
        logChan: make(chan []byte, 10000),
    }

    go logger.writeLoop()
    return logger

}

func (l *AsyncLogger) writeLoop(){
    for msg := range l.logChan{
        l.mu.Lock()
        if l.shouldRotate(){
            l.rotate()
        }
        l.file.Write(msg)
        l.mu.Unlock()
    }
}

func (l *AsyncLogger) Write(p []byte)(n int, err error){
    l.logChan <- append([]byte{}, p...)
    return len(p), nil
}