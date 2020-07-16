package main
 
import (
    "encoding/csv"
    "encoding/json"
    "fmt"
    "os"
    "io/ioutil"
   // "log"
    "net"
    "strconv"
     "bufio"
     "time"
    
)
 
type Virus struct {
     Cumulative int
     CumulativeTestpositive int
     Cumulativetestsperformed int
     Date string 
     Discharged int 
     Expired int 
     Stilladmitted int
     Region string

}
 
func main() {
    arguments := os.Args
      if len(arguments) == 1 {
            fmt.Println("Please provide port number")
            return
      }

      PORT := ":" + arguments[1]
      l, err := net.Listen("tcp", PORT)
      if err != nil {
            fmt.Println(err)
            return
     }
     c, err := l.Accept()
      if err != nil {
           fmt.Println(err)
           return
     }
     //c, err := l.Accept()
     // if err != nil {
           // fmt.Println(err)
           // return
     // }
      defer l.Close()
    csvFile, err := os.Open("./covid_final_data.csv")
    if err != nil {
        fmt.Println(err)
    }
    defer csvFile.Close()
 
    reader := csv.NewReader(csvFile)
    reader.FieldsPerRecord = -1
 
    csvData, err := reader.ReadAll()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
 
    var virus1 Virus
    var viruses []Virus
 
    for _, each := range csvData {
        virus1.Cumulative, _ = strconv.Atoi(each[1])
        virus1.CumulativeTestpositive,_=strconv.Atoi(each[2])
        virus1.Cumulativetestsperformed,_ = strconv.Atoi(each[3])
        virus1.Date=each[4]
        virus1.Discharged,_=strconv.Atoi(each[5])
        virus1.Expired,_=strconv.Atoi(each[6])
        virus1.Region=each[9]
        virus1.Stilladmitted,_=strconv.Atoi(each[10])
        viruses = append(viruses, virus1)
    }
 
    // Convert to JSON
    jsonData, err := json.Marshal(viruses)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
 
    fmt.Println(string(jsonData))
 
    jsonFile, err := os.Create("./data.json")
    if err != nil {
        fmt.Println(err)
    }
    defer jsonFile.Close()
 
    jsonFile.Write(jsonData)
    jsonFile.Close()
    byteValue, _ := ioutil.ReadAll(jsonFile)
    var users Virus
    json.Unmarshal(byteValue, &users)
    //byte[] data = File.ReadAllText(data);
     //var dat []map[string]int

 //   if err := json.Unmarshal(data, &dat); err != nil {
    //    panic(err)
   // }
    //fmt.Printf("CumulativeTestpositive: %.2f\n", virus1.CumulativeTestpositive);
    //fmt.Printf("EUR : %.2f\n", obj.EUR);
   // fmt.Printf("GBP : %.2f\n", obj.GBP);


   // for id := range dat {

       // if dat[id]["Region"] =="Punjab "{
          //  fmt.Println(" ", virus1.CumulativeTestpositive,virus1., each[4], strconv.Atoi(each[5]),strconv.Atoi(each[6]),each[9],strconv.Atoi(each[10]))
       // }
    //}

    //jq := gojsonq.New().File("./data.json")
      //res := jq.From("Virus").Where("Region", "=", "Punjab").Get()
      //fmt.Println(res)
      for {

             netData, err := bufio.NewReader(c).ReadString('\n')
             if err != nil {
                  fmt.Println(err)
                  return
             }
             fmt.Print("-> ", string(netData))
            t := time.Now()
             myTime := t.Format(time.RFC3339) + "\n"
             c.Write([]byte(myTime))
}
}