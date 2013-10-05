package main

import "io/ioutil";

func check(e error) {
    if e != nil {
        panic(e);
    }
}

func main() {
    speedFile := "/sys/devices/platform/i8042/serio1/serio2/speed";
    sensitivityFile := "/sys/devices/platform/i8042/serio1/serio2/sensitivity";
    err := ioutil.WriteFile(speedFile, []byte("120"), 0644);
    check(err);
    err = ioutil.WriteFile(sensitivityFile, []byte("250"), 0644);
    check(err);
}
