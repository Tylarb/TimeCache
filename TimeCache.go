package timeCache



type TimeCache interface {
  Includes(string) bool
  Add(string)
  Drop(string)
  }
  
