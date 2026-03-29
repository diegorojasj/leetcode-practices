// @ts-ignore
import { expect, test, describe, beforeAll, afterAll } from "bun:test";
import { addTwoNumbers, ListNode } from "./attempt1";

describe("addTwoNumbers", () => {
    beforeAll(() => {
        console.log("Starting addTwoNumbers tests...");
    });

    afterAll(() => {
        console.log("All addTwoNumbers tests finished.");
    });

    test("Case 1", () => {
        const start = performance.now();
        const result = addTwoNumbers(new ListNode(2, new ListNode(4, new ListNode(3))), new ListNode(5, new ListNode(6, new ListNode(4))));
        console.log(`Case 1: ${(performance.now() - start).toFixed(4)}ms`);
        expect(result).toEqual(new ListNode(7, new ListNode(0, new ListNode(8))));
    });

    test("Case 2", () => {
        const start = performance.now();
        const result = addTwoNumbers(new ListNode(0), new ListNode(0));
        console.log(`Case 2: ${(performance.now() - start).toFixed(4)}ms`);
        expect(result).toEqual(new ListNode(0));
    });

    test("Case 3", () => {
        const start = performance.now();
        const result = addTwoNumbers(new ListNode(9, new ListNode(9, new ListNode(9, new ListNode(9, new ListNode(9, new ListNode(9, new ListNode(9))))))), new ListNode(9, new ListNode(9, new ListNode(9, new ListNode(9)))));
        console.log(`Case 3: ${(performance.now() - start).toFixed(4)}ms`);
        expect(result).toEqual(new ListNode(8, new ListNode(9, new ListNode(9, new ListNode(9, new ListNode(0, new ListNode(0, new ListNode(0, new ListNode(1)))))))));
    });
})