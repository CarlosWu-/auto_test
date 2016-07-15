package main


import (

    "net/http"
    "io/ioutil"
    //"encoding/json"

    "fmt"
    //"strings"
    "os"
    "time"
    "log"
    "bufio"
)

type Love struct {
    BoyName string
    GirlName string
}


func main() {

    //num := calc(5,18)
    //Arr := []int{1,2,3,10}
    //fmt.Println(Arr[6])
    //fmt.Println(num)
    //fmt.Println(test(Arr...))
    //fmt.Println(strings.Contains("wujiabing","g"))
    //ch,_ := ReadFile("C:\\Users\\wujiabing\\Desktop\\crash.txt")
    //fmt.Println(string(ch[:]))
    //x := double(3)
    //fmt.Println("\n","dddd",x)



    //var url = "http://testurl/v1/get_stream_list?app_id=201&token=98765"
    //fmt.Print(request(url))


    running := true

    for running{
        fmt.Println("please input somthing:")
        input := bufio.NewReader(os.Stdin)
        data,_,_ := input.ReadLine()

        switch string(data) {

        case "wujiabing": fmt.Println("boy")
        case "liuxiaoning": fmt.Println("girl")
        case "wuyemo": fmt.Println("little girl")
        case "stop":running = false
        default:
            panic(fmt.Sprintf("ivanlid %s",input))


        }

    }



}

func calc(x,y int) int{
    return x+y
}

func request(url string) string {

    resp,err := http.Get(url)
    if err != nil{
        //return err
    }
    if resp.StatusCode != http.StatusOK{
        resp.Body.Close()
        //return  resp.StatusCode
    }
    content,err := ioutil.ReadAll(resp.Body)
    resp.Body.Close()

    return string(content)
}


func test(vars...int) int {

    total := 0
    for _,i := range  vars{

        total += i
    }
    return total
}

func ReadFile(filename string) ([]byte, error) {
    defer trace("ReadFile")()
    time.Sleep(5 * time.Second)
    f, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer f.Close()
    return ioutil.ReadAll(f)
}

/*
@program run time
 */
func trace(msg string) func(){
    start := time.Now()
    log.Printf("enter %s" ,msg)
    return func() {
        log.Printf("%s run time  (%s)",msg,time.Since(start))
    }
}

func double(x int) (result int){

    defer func(){
        result += x
        fmt.Printf("number %d--double:%d",x,result)
    }()
    return x+x
}