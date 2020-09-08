const iqTest = (numbers) => {
    let arr = numbers.split(" ")
    
    let oddNumbers = arr.filter(((val) => { return parseInt(val) % 2 === 1}));
    let evenNumbers = arr.filter(((val) => { return parseInt(val) % 2 === 0}));
    
    if(oddNumbers.length < evenNumbers.length){
      return arr.indexOf(oddNumbers[0]) + 1
    }
      
    return arr.indexOf(evenNumbers[0]) + 1
  }