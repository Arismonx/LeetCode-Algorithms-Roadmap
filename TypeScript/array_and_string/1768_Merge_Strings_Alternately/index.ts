function mergeAlternately(word1: string, word2: string): string {
  let result: string[] = [];
  let i = 0;
  let j = 0;

  while (i < word1.length && j < word2.length) {
    result.push(word1[i]!);
    result.push(word2[j]!);
    i++;
    j++;
  }

  while (i < word1.length) {
    result.push(word1[i]!);
    i++;
  }

  while (j < word2.length) {
    result.push(word2[j]!);
    j++;
  }
  return result.join("");
}

// let str1 = "abc";
// let str2 = "pqr";
// console.log(mergeAlternately(str1, str2));

let str1 = "abcd";
let str2 = "pq";
console.log(mergeAlternately(str1, str2)); //a p b q c   d
