const duplicateCount = (text) => {
    if(text.length === 0) {
      return 0;
    }
    let counter = 0;
    let myMap = new Map()
    for(let i = 0; i < text.length; i++){
      let temp = text[i];
      if(typeof temp === "string"){
        temp = temp.toUpperCase();
      }
      
      if(myMap.has(temp)){
        myMap.set(temp, myMap.get(temp) + 1)
      }else {
        myMap.set(temp,1)
      }
    }
    
    myMap.forEach((val) => {
      if(val > 1) {
        counter++;
      }
    })
 
    return counter;
}