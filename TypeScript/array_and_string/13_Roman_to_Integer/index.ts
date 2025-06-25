function romanToInt(s: string): number {
  let mymap = new Map<string, number>();
  mymap.set("I", 1);
  mymap.set("V", 5);
  mymap.set("X", 10);
  mymap.set("L", 50);
  mymap.set("C", 100);
  mymap.set("D", 500);
  mymap.set("M", 1000);
  let result: number = 0;
  for (let i = 0; i < s.length; i++) {
    if (
      i + 1 < s.length &&
      mymap.get(s[i] as string)! < mymap.get(s[i + 1] as string)!
    ) {
      result -= mymap.get(s[i] as string)!;
    } else {
      result += mymap.get(s[i] as string)!;
    }
  }
  return result;
}

// console.log(romanToInt("III"));
console.log(romanToInt("MCMXCIV"));
