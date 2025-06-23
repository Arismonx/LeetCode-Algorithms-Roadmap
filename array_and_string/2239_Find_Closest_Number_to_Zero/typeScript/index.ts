// function findClosestNumber(nums: number[]): number {
//   var num: number = nums[0];
// }

const findClosestNumber = (nums: number[]): number => {
  let closeZero = nums[0];
  let minAbs: number | undefined = abs(closeZero!);

  for (let index = 1; index < nums.length; index++) {
    let curr = nums[index];
    let currAbs = abs(curr!);

    if (currAbs < minAbs) {
      minAbs = currAbs;
      closeZero = curr;
    } else if (currAbs == minAbs) {
      if (curr! > closeZero!) {
        closeZero = curr;
      }
    }
  }
  return closeZero!;
};

const abs = (x: number) => (x < 0 ? -x : x);

let x: number[] = [-4, -2, 1, 4, 8];
let x1: number[] = [-100, -100];
let x2: number[] = [2, -1, 1];

console.log(findClosestNumber(x));
console.log(findClosestNumber(x1));
console.log(findClosestNumber(x2));
