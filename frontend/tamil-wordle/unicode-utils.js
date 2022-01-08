String.prototype.toUnicode = function () {
    var result = "";
    for (var i = 0; i < this.length; i++) {
        // Assumption: all characters are < 0xffff
        result += "\\u{" + ("000" + this[i].charCodeAt(0).toString(16)).substr(-4) + "}";
    }
    return result;
};

const diacritics = {
    "\u0B82": true,
    "\u0BBE": true,
    "\u0BBF": true,
    "\u0BC0": true,
    "\u0BC1": true,
    "\u0BC2": true,
    "\u0BC6": true,
    "\u0BC7": true,
    "\u0BC8": true,
    "\u0BCA": true,
    "\u0BCB": true,
    "\u0BCC": true,
    "\u0BCD": true,
    "\u0BD7": true,
};

String.prototype.forUnicodeEach = function (f) {
    var targetList = [];
    for (var i = 0; i != this.length; i++) {
        var ch = this[i];
        diacritics[ch] ? (targetList[targetList.length - 1] += ch) : targetList.push(ch);
    }
    targetList.forEach(f);
};

String.prototype.unicodeLength = function () {
    var targetList = [];
    for (var i = 0; i != this.length; i++) {
        var ch = this[i];
        diacritics[ch] ? (targetList[targetList.length - 1] += ch) : targetList.push(ch);
    }
    return targetList.length;
};
