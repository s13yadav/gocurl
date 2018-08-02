package main

import (
        "bufio"
        "io"
        "os"
        "strings"
)

// ReadConf reads a "=" separated key value pair file and return a map[string]string
func ReadConf(filename string) (conf map[string]string, err error) {
        //log.Println("INFO  : reading ", filename)
        f, err := os.Open(filename)
        if err != nil {
                //fmt.Println("error opening file ", err)
                return nil, err
        }
        defer f.Close()

        conf = make(map[string]string)

        r := bufio.NewReader(f)
        var param string
        nLine := 0

        // anonymous function
        set_param := func() {
                param = strings.TrimSpace(param)
                if len(param) < 2 {
                        return
                }
                if strings.HasPrefix(param, "#") {
                        return
                }
                split_items := strings.Split(param, "=")
                if len(split_items) != 2 {
                        return
                }
                k := strings.TrimSpace(split_items[0])
                v := strings.TrimSpace(split_items[1])
                if len(k) == 0 {
                        Logger.Printf("warning: line:%d - missing key, attempting to continue", nLine)
                        return
                }
                conf[k] = v
        }

        for {
                nLine = nLine + 1
                param, err = r.ReadString('\n')
                if err == io.EOF {
                        err = nil
                        set_param()
                        break
                } else if err != nil {
                        Logger.Printf("%s - error while parsing\n", filename)
                        return // if you return error
                }

                set_param()
        }
        return
}
