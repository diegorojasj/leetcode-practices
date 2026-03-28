export function twoSum(nums: number[], target: number): number[] {
    const secondArr = [...nums.entries()];
    for(const [inx, num] of nums.entries()) {
        secondArr.shift();
        for(const [inx2, num2] of secondArr) {
            if(num+num2 === target) return [inx, inx2];
        }
    }
    return [];
};

const case1 = [2, 7, 11, 15];
const target1 = 9;

console.log(twoSum(case1, target1));

const case2 = [3, 2, 4];
const target2 = 6;

console.log(twoSum(case2, target2));

const case3 = [3, 3];
const target3 = 6;

console.log(twoSum(case3, target3));

// Extra cases

const case4 = [3, 2 , 3];
const target4 = 6;

console.log(twoSum(case4, target4));