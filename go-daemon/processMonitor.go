type ProcessMonitor struct{
	pid int
    restarts int
    lastRestart time.Time
    maxRestarts int
}


func (pm *ProcessMonitor) Monitor(){
    for {

        if pm.restarts >= p.maxRestarts {
            log.Fatal("Too many restarts, giving up")
        }
        if err := checkProcess(pm.pid); err != nil{
            pm.restartProcess()
        }

        time.Sleep(time.Second * 5)
    }


}

func setupCrashRecovery(){

    defer func(){
        if r := recover(); r != nil{
            log.Printf("Recovered from panic: %v", r)
            debug.PrintStack()
        }

        logCrash(r)
        restartService()
    }()

}


type WorkPool struct{

    workerCount int
    jobQueue chan jobQueue
    workers []*WorkerCount
}

func NewWorkerPool(count int) *WorkerPool {
    pool := &WorkerPool{
        workerCount: count,
        jobQueue: make(chan Job, 100),
        workers: make([]*Worker, count),

    }
    pool.Start()
    return pool
}

type ResourceLimiter struct{
    maxCPU float64
    MaxMemory uint64
    interval time.Duration
}

func (rl *ResourceLimiter) Monitor(){
    ticker := time.NewTicker(rl.interval)
    for range ticker.C{
        if rl.checkResourceUsage(){
            rl.applyLimits()
        }
        
    }
}