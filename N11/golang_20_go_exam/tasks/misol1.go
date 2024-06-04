package tasks

import "errors"

func SumOflice(slc []int) (int, error) {
 if len(slc) == 0 {
  return 0, errors.New("slice uzunligi 0 ga teng")
 }
 sum := 0
 for _, v := range slc {
  sum += v
 }
 if sum == 0 {
  return 0, errors.New("yig'indisi 0 ga teng")
 }
 return sum, nil
}