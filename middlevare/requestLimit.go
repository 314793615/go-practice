package middlevare

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

var rdc *redis.Client


func initRedis() {
	rdc = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Username: "",
		Password: "",
		DB: 0,
	})
}

func init(){
	initRedis()
}


func RequestLimitHandle(f http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		ip := getRequestIp(r)
		val, err := rdc.Get(ctx, ip).Result()
		if err == redis.Nil{
			err := rdc.Set(ctx, ip, 1, time.Minute).Err()
			if err != nil{
				http.Error(w, "Redis error", http.StatusInternalServerError)
				return
			}
		}else if err != nil{
			http.Error(w, "Redis error", http.StatusInternalServerError) 
			return 
		}else{
			countInt,_ := strconv.Atoi(val)
			if countInt >= 10{
				http.Error(w, "Too many requests in one minute", http.StatusTooManyRequests)
				return 
			}
			rdc.Incr(ctx, ip)
		}
		f.ServeHTTP(w, r)
	}
}

func submitHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Request successful")
}


func startServer() {
	http.HandleFunc("/submit", RequestLimitHandle(submitHandler))
	log.Println("server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
	
}



func getRequestIp(r *http.Request) string{
	ip := ""
	xForwardHead := r.Header.Get("X-Forwarded-For")
	if xForwardHead != ""{
		ip = strings.Split(xForwardHead, ",")[0]
	}else{
		remoteAddr := r.RemoteAddr
		ip = strings.Split(remoteAddr, ":")[0]
		
	}
	return ip
}