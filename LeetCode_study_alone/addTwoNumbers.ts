
 //* Definition for singly-linked list.
class ListNode {
      val: number
      next: ListNode | null
      constructor(val?: number, next?: ListNode | null) {
          this.val = (val===undefined ? 0 : val)
         this.next = (next===undefined ? null : next)
      }
  }
 
function addTwoNumbers(l1:ListNode|null, l2:ListNode|null) : ListNode|null{
    let ans = new ListNode(0); // create ListNode which will save the answer nodes.
    let current = ans //Pointer moves on the ListNode 'ans'
    let carry = 0

    while(l1 != null || l2 != null){
        const val1 = l1?.val ?? 0; // if the value is not exist use 0 instead. 
        const val2 = l2?.val ?? 0;
        const sum = val1 + val2 + carry 

        current.next = new ListNode(sum % 10) // connect new value to the next node. 
        carry = Math.floor(sum / 10) //get the floor value only. 
        current = current.next 

        if (l1 != null) l1 = l1.next // move to next list node. if don't have next node, error occurs-> check the node exists. 
        if (l2 != null) l2 = l2.next;
    }
    if (carry > 0) current.next = new ListNode(carry) // when the lists are end and last carry value remained. 


    return ans.next
}

 
  
