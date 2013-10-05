package main

import "io/ioutil";
import "flag";
import "fmt";
import "strconv";

func check(e error) {
    if e != nil {
        panic(e);
    }
}

func main() {
    var speed = flag.Int("speed", 120, "Pointer speed");
    var sensitivity = flag.Int("sensitivity", 250, "Pointer sensitivity");
    flag.Parse()
    fmt.Println("Speed:", *speed);
    fmt.Println("Sensitivity:", *sensitivity);
    speedFile := "/sys/devices/platform/i8042/serio1/serio2/speed";
    sensitivityFile := "/sys/devices/platform/i8042/serio1/serio2/sensitivity";
    err := ioutil.WriteFile(speedFile, []byte(strconv.Itoa(*speed)), 0644);
    check(err);
    err = ioutil.WriteFile(sensitivityFile, []byte(strconv.Itoa(*sensitivity)), 0644);
    check(err);
}
