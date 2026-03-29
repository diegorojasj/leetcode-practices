export class ListNode {
    val: number;
    next: ListNode | null;
    constructor(val?: number, next?: ListNode | null) {
        this.val = val ?? 0;
        this.next = next ?? null;
    }
}

export function addTwoNumbers(l1: ListNode | null, l2: ListNode | null): ListNode | null {
    return sumNode(l1 ,l2, 0);
};

const sumNode = (node: ListNode | null, node2: ListNode | null, additionals: number): ListNode | null => {
    let newNode: ListNode | null = null;
    if(node || node2 || additionals > 0) {
        const sum = (node?.val ?? 0) + (node2?.val ?? 0) + additionals;
        const newAdditionals = Math.trunc(sum / 10);
        let num = sum;
        if(newAdditionals > 0) {
            num -= (newAdditionals * 10)
        }
        newNode = new ListNode(num, sumNode(node?.next ?? null, node2?.next ?? null, newAdditionals))
    }
    return newNode;
}

const case1 = new ListNode(2, new ListNode(4, new ListNode(3)));
const case1_2 = new ListNode(5, new ListNode(6, new ListNode(4)));

console.log(addTwoNumbers(case1, case1_2));

const case2 = new ListNode(0);
const case2_2 = new ListNode(0);

console.log(addTwoNumbers(case2, case2_2));

const case3 = new ListNode(9, new ListNode(9, new ListNode(9, new ListNode(9, new ListNode(9, new ListNode(9, new ListNode(9)))))));
const case3_2 = new ListNode(9, new ListNode(9, new ListNode(9, new ListNode(9))));

console.log(addTwoNumbers(case3, case3_2));
