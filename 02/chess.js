for (var i = 1; i <= 8; i++) {
  var string = '';
  var sharp = "#";
  var blank = " ";
  for (var j = 1; j <= 8; j++) {
    if (j % 2 != 0) {
      string += (i % 2 != 0) ? sharp : blank;
    } else {
      string += (i % 2 == 0) ? sharp : blank;
    }
  }
  console.log(string);
}