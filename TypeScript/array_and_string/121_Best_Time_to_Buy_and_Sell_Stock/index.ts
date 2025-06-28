function maxProfit(prices: number[]): number {
  let minPre: number | undefined = prices[0];
  let result: number = 0;

  for (let i = 1; i < prices.length; i++) {
    minPre = min(prices[i]!, minPre!);
    result = max(result, prices[i]! - minPre);
  }
  return result;
}

const min = (x: number, y: number) => (x < y ? x : y);
const max = (x: number, y: number) => (x < y ? y : x);

let nums1: number[] = [7, 1, 5, 3, 6, 4];
let nums2: number[] = [7, 6, 4, 3, 1];
console.log(maxProfit(nums1));
console.log(maxProfit(nums2));
