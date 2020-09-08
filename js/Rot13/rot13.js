function rot13(message) {
    let newWord = [];
    for (let i = 0; i < message.length; i++) {
        let upper = false;
        let lower = false;
        let ascii = message[i].charCodeAt()
        if (ascii >= 65 && ascii <= 90) {
            upper = true;
        } else if (ascii >= 97 && ascii <= 122) {
            lower = true;
        }

        if (upper) {
            ascii = ascii + 13
            if (ascii > 90) {
                ascii = ascii - 90 + 64
            }
        } else if (lower) {
            ascii = ascii + 13
            if (ascii > 122) {
                ascii = ascii - 122 + 96
            }
        }

        newWord.push(ascii)
    }


    let result = "";
    for (let i = 0; i < newWord.length; i++) {
        result += String.fromCharCode(newWord[i])
    }
    return result
}