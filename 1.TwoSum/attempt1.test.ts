// @ts-ignore
import { expect, test, describe, beforeAll, afterAll } from "bun:test";
import { twoSum } from "./attempt1";

describe("twoSum", () => {
    beforeAll(() => {
        console.log("Starting twoSum tests...");
    });

    afterAll(() => {
        console.log("All twoSum tests finished.");
    });

    test("Case 1", () => {
        const start = performance.now();
        const result = twoSum([2, 7, 11, 15], 9);
        console.log(`Case 1: ${(performance.now() - start).toFixed(4)}ms`);
        expect(result).toEqual([0, 1]);
    });

    test("Case 2", () => {
        const start = performance.now();
        const result = twoSum([3, 2, 4], 6);
        console.log(`Case 2: ${(performance.now() - start).toFixed(4)}ms`);
        expect(result).toEqual([1, 2]);
    });

    test("Case 3", () => {
        const start = performance.now();
        const result = twoSum([3, 3], 6);
        console.log(`Case 3: ${(performance.now() - start).toFixed(4)}ms`);
        expect(result).toEqual([0, 1]);
    });

    test("Case 4", () => {
        const start = performance.now();
        const result = twoSum([3, 2 , 3], 6);
        console.log(`Case 4: ${(performance.now() - start).toFixed(4)}ms`);
        expect(result).toEqual([0, 2]);
    });
});