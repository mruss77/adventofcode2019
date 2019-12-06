package main

import (
   "fmt"
   "os"
   "bufio"
   "strconv"
   "strings"
)

func main() {
   if len(os.Args) != 2 {
      fmt.Println("Need one argument.")
      os.Exit(0)
   }
   inp, err := strconv.Atoi(os.Args[1])
   if err != nil { panic(err) }
   arg := int32(inp)
   file, err := os.Open("input")
   if (err != nil) { panic(err) }
   defer file.Close()

   reader := bufio.NewReader(file)
   data := getData(reader)

   keepGoing := true;
   ptr := 0
   for keepGoing {
      //fmt.Println(data, ptr)
      data, ptr, keepGoing = runOp(data, ptr, arg)
   }
}

func runOp(data []int32, ptr int, arg int32) ([]int32, int, bool) {
   //Runs the opcode at ptr and returns the data and new ptr value
   opstr := data[ptr]
   opcode := opstr % 100
   pMode := opstr / 100
   nParms := 0
   keepGoing := true

   //fmt.Println("opcode = ", opcode)
   var parms []*int32
   switch opcode {
      case 1, 2, 7, 8:
         nParms = 3
      case 3, 4:
         nParms = 1
      case 99:
         nParms = 0
      case 5, 6:
         nParms = 2
      default:
         panic("bad opcode")
   }
   //fmt.Println("nParms = ", nParms)
   for i:=0; i<nParms; i++ {
      ptr++
      if pMode / int32(intPow(10,i)) % 10 == 0 { //position mode
         parms = append(parms, &data[data[ptr]])
      } else {                              //direct mode
         parms = append(parms, &data[ptr])
      }
   }
   //fmt.Println(parms)
   switch opcode {
      case 1:
         *parms[2] = *parms[0] + *parms[1]
         ptr++
      case 2:
         *parms[2] = *parms[0] * *parms[1]
         ptr++
      case 3:
         *parms[0] = arg
         ptr++
      case 4:
         fmt.Println("Output: ", *parms[0])
         ptr++
      case 5:
         if *parms[0] != 0 {
            ptr = int(*parms[1])
         } else {
            ptr++
         }
      case 6:
         if *parms[0] == 0 {
            ptr = int(*parms[1])
         } else {
            ptr++
         }
      case 7:
         if *parms[0] < *parms[1] {
            *parms[2] = int32(1)
         } else {
            *parms[2] = int32(0)
         }
         ptr++
      case 8:
         if *parms[0] == *parms[1] {
            *parms[2] = int32(1)
         } else {
            *parms[2] = int32(0)
         }
         ptr++
      case 99:
         keepGoing = false
         ptr++
   }
   return data, ptr, keepGoing
}

func getData(r *bufio.Reader) []int32  {
   var err error = nil
   var dat []int32
   for err == nil {
      str := ""
      str, err = r.ReadString(',')
      if str != "" {
         dat = append(dat, getInt(str))
      }
   }
   return dat
}

func getInt(s string) int32 {
   //remove commas & newlines
   s = strings.Trim(s, ",\n")
   n, err := strconv.Atoi(s)
   if err != nil {
      panic(err)
   }
   return int32(n)
}

func intPow(x int, y int) int {
   ans := 1
   for i:=0; i < y; i++ {
      ans = ans * x
   }
   return ans
}

