const decodeMorse = (morseCode) => {
    let result  = morseCode.split('   ').reduce((word, code, i, arr) => {
      return word + code.split(' ').reduce((wordFrom, codeFrom) => {
        return wordFrom + (MORSE_CODE[codeFrom] || '');
      }, '') + (i < arr.length - 1 ? ' ' : '');
    }, '').trim();
    
    return result;
}