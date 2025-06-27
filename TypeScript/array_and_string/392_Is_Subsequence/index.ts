function isSubsequence(s: string, t: string): boolean {
  if (s.length > t.length) return false;
  let [countS, countT] = [0, 0];

  while (countS < s.length && countT < t.length) {
    if (s[countS] == t[countT]) {
      countS++;
    }
    countT++;
  }
  return countS == s.length;
}

console.log(isSubsequence("abc", "ahbgdc"));
console.log(isSubsequence("axc", "ahbgdc"));
