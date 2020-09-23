function bkn(skey) {
  var hash = 5381;
  for (var i = 0; i < skey.length; i++) {
    hash += (hash << 5) + skey[i].charCodeAt(0);
  }
  return hash & 2147483647;
}

console.log(bkn('MP6mJCegHj'));
