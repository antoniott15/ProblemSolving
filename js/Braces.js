const word = "[()]{}"

const openBraces = new Map()
openBraces.set("{", false)
openBraces.set("(", false)
openBraces.set("[", false)

let closeBraces = new Map()
closeBraces.set("}", "{")
closeBraces.set(")", "(")
closeBraces.set("]", "[")

// O(n) solution
const isAcepted = (w) => {
    if(w.length % 2 !== 0) return false
    w = w.split("")

    const qWord = []

    w.forEach((val) => {
        if(openBraces.has(val)) qWord.push(val)
        
        if(closeBraces.has(val)){
            if(qWord[qWord.length - 1] === closeBraces.get(val)) {
                qWord.pop()
            }else {
                return false
            }
        }
    })

    return qWord.length !== 0 ? false : true
}

console.log(isAcepted(word))