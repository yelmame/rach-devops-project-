package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os/exec"
    "runtime"
    "time"
)

type HealthResponse struct {
    Status     string    `json:"status"`
    GoVersion  string    `json:"go_version"`
    SolcVersion string   `json:"solc_version"`
    Timestamp  time.Time `json:"timestamp"`
    PID        int       `json:"pid"`
}

func main() {
    http.HandleFunc("/health", healthHandler)
    http.HandleFunc("/solc", solcHandler)
    http.HandleFunc("/", rootHandler)
    
    log.Println("=== Go SDK App Started on :8080 ===")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
    solcOut, _ := exec.Command("solc", "--version").CombinedOutput()
    hc := HealthResponse{
        Status:     "healthy",
        GoVersion:  runtime.Version(),
        SolcVersion: string(solcOut),
        Timestamp:  time.Now(),
        PID:        os.Getpid(),
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(hc)
}

func solcHandler(w http.ResponseWriter, r *http.Request) {
    cmd := exec.Command("solc", "--version")
    out, err := cmd.CombinedOutput()
    if err != nil {
        http.Error(w, fmt.Sprintf("solc error: %v", err), 500)
        return
    }
    w.Header().Set("Content-Type", "text/plain")
    w.Write(out)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, `Autheo Go SDK Pipeline App v1.0
===================
Endpoints:
/health    - JSON health + versions  
/solc      - Raw solc --version
/          - This page

Ready for ECS/EKS via Jenkins pipeline!
`)
}
