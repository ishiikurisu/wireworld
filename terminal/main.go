package main

import (
    "os"
    "fmt"
    "io/ioutil"
    "time"
    "os/exec"
    "runtime"
    "github.com/ishiikurisu/wireworld"
)

var clear map[string]func() //create a map for storing clear funcs

func init() {
    clear = make(map[string]func()) //Initialize it
    clear["linux"] = func() {
        cmd := exec.Command("clear") //Linux example, its tested
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
    clear["windows"] = func() {
        cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
}

func CallClear() {
    value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
    if ok { //if we defined a clear func for that platform:
        value()  //we execute it
    } else { //unsupported platform
        panic("Your platform is unsupported! I can't clear terminal screen :(")
    }
}

func toString(w wireworld.Wireworld) string {
    s := ""
    for y := 0; y < w.Dy; y++ {
        for x := 0; x < w.Dx; x++ {
            switch w.World[y][x].Kind {
            case wireworld.POSITIVE:
                s += "#"
                break
            case wireworld.NEGATIVE:
                s += "*"
                break
            case wireworld.CONDUCTOR:
                s += "_"
                break
            default:
                s += " "
                break
            }
        }
        s += "\n"
    }
    return s
}

func main() {
    // setup
    filename := os.Args[1]
    rawCsv, oops := ioutil.ReadFile(filename)
    if oops != nil {
        fmt.Println("invalid file")
        return
    }
    csv := string(rawCsv)
    w, oops := wireworld.LoadCsv(csv)
    if oops != nil {
        fmt.Println("invalid csv")
        return
    }

    // loop
    for {
        s := toString(w)
        fmt.Println(s)
        time.Sleep(time.Second)
        w = wireworld.Update(w)
    }
}
