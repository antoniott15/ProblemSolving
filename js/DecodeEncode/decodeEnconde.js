const encode = (message) => {
    let word = "";
    let map = new Map();
    map.set("a", 1);
    map.set("e", 2);
    map.set("i", 3);
    map.set("o", 4);
    map.set("u", 5);

    [...message].forEach((val) => {
        if (map.has(val)) {
            word += map.get(val)
        } else {
            word += val
        }
    })

    return word
}

const decode = (message) => {
    let word = "";
    let map = new Map();

    map.set("1", "a");
    map.set("2", "e");
    map.set("3", "i");
    map.set("4", "o");
    map.set("5", "u");
    [...message].forEach((val) => {
        if (map.has(val)) {
            word += map.get(val)
        } else {
            word += val
        }
    })

    return word
}