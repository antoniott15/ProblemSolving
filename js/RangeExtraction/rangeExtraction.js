const solution = (list) => {
    let prev = '';
    let arr = []
    let len = list.length;
  
    for (let i = 0, j = 1; j <= len;) {
      if (list[j] - list[j - 1] !== 1 || j++ === len) {
        if(j > 1) {
          prev = list[i] + "-"
        }
        if(list[j - 1] - list[i] === 1){
          prev = list[i] + ","
        }
  
        if(list[j - 1] === list[i]){
           prev = '' 
        }
  
        arr.push(prev + list[j-1]);
        prev = list[(i = j++)]
      }
    }
  
    return arr.join`,`;
  }