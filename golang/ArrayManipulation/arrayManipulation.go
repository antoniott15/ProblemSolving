package main

func findMax(newArray []int64, n int64) (max int64) {
	var newValue, maxim int64 = 0, 0
	for i := 1; i <= int(n); i++ {
		newValue += newArray[i]
		if maxim < newValue {
			maxim = newValue
		}
	}
	return maxim
}

// Complete the arrayManipulation function below.
func arrayManipulation(n int64, queries [][]int64) int64 {
	var newArray = make([]int64, n+1)
	/* var position = make([] int64, n)
	   var value int64
	   for _,elements := range queries{
	       for j, valuesofQuery := range elements{
	           position[j] = valuesofQuery - 1
	           if(j==2){
	               value = valuesofQuery
	           }
	       }
	       for i:=0;i<len(position);i++{
	           if position[1]-position[0]==1 || position[1]-position[0]==0 {
	               if i==0||i==1 {
	                   newArray[position[i]] += value;
	               }
	           }else if position[1]-position[0] > 1 {
	              for newJ:=position[0];newJ<=position[1];newJ++{
	                  if(i==0){
	                      newArray[newJ] += value;
	                  }
	               }
	           }
	       }
	   }
	   maxValue := findMax(newArray)
	   fmt.Println(maxValue)
	   return int64(maxValue)*/

	for _, element := range queries {
		firsElement := int64(element[0])
		secondElement := int64(element[1])
		value := int64(element[2])

		newArray[firsElement] += value

		if (secondElement + 1) <= n {
			newArray[secondElement+1] -= value
		}
	}

	maxValue := findMax(newArray, n)
	return maxValue
}
