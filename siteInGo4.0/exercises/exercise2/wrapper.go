package main

import (
    "runtime"
    "os"
    "os/exec"
    "syscall/js"
    "log"
)

func runCLI() {
    cmd := exec.Command("go", "run", "main.go")
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err := cmd.Run(); err != nil {
        log.Fatalf("Errore durante l'esecuzione CLI: %v", err)
    }
}

func runWebAssembly() {
    js.Global().Set("runMain", js.FuncOf(func(this js.Value, args []js.Value) any {
        println("WebAssembly: placeholder per l'interazione con main.go")
        return nil
    }))
    println("WebAssembly pronto. Usa 'runMain()' per interagire.")
    select {}
}

func main() {
    if runtime.GOOS == "js" && runtime.GOARCH == "wasm" {
        runWebAssembly()
    } else {
        runCLI()
    }
}
