package main


func burbuja(Lista []int64){
 var auxiliar int64
 for i := 0; i < len(Lista); i++ {
  for j := 0; j < len(Lista); j++ {
   if Lista[i] < Lista[j] {
    auxiliar = Lista[i]
    Lista[i] = Lista[j]
    Lista[j] = auxiliar
    }
  }
}
}

func main() {
  
}