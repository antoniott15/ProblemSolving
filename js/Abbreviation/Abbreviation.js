const abr = (a,b) => {
    a = a.split("");
    b = b.split("");
    let mapB = new Map();

    b.forEach((val) => {
        mapB.set(val.toUpperCase(), val)
    })

    let helper = []
    a.forEach((val) => {
        if(!mapB.has(val.toUpperCase())) {
            helper.push(0)
        }else {
            helper.push(1)
        }
    })

    let wordResult = []
    helper.forEach((val, i) => {
        if(val === 1) {
            wordResult.push(a[i])
        }
    })

    const result = JSON.stringify(wordResult).toUpperCase() === JSON.stringify(b).toUpperCase()

    return result ? "YES" : "NO"
}



console.log(abr("AFDE", "ABDE"))